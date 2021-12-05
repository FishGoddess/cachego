// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/03/14 16:28:56

package cachego

import (
	"context"
	"time"

	"github.com/FishGoddess/cachego/internal/config"
)

const (
	// defaultMapSize determines the initialized map size of one segment.
	defaultMapSize = 256

	// defaultSegmentSize determines the size of segments.
	// This value will affect the performance of concurrency.
	defaultSegmentSize = 256
)

// Cache is a struct of cache.
type Cache struct {
	// mapSize is the size of map inside.
	mapSize int

	// segmentSize is the size of segments.
	// This value will affect the performance of concurrency.
	// It should be the pow of 2 (such as 64) or the segments may be uneven.
	segmentSize int

	// segments is a slice stores the real data.
	segments []*segment
}

// NewCache returns a new Cache holder for use.
func NewCache(opts ...Option) *Cache {
	cache := applyOptions(&Cache{
		mapSize:     defaultMapSize,
		segmentSize: defaultSegmentSize,
	}, opts...)

	cache.segments = newSegments(cache.mapSize, cache.segmentSize)
	return cache
}

// newSegments returns a slice of initialized segments.
func newSegments(mapSize int, segmentSize int) []*segment {
	segments := make([]*segment, segmentSize)

	for i := 0; i < segmentSize; i++ {
		segments[i] = newSegment(mapSize)
	}

	return segments
}

// indexOf returns a position in segments of this key.
func (c *Cache) indexOf(key string) int {
	index := 1469598103934665603
	keyBytes := []byte(key)

	for _, b := range keyBytes {
		index = (index << 5) - index + int(b&0xff)
		index *= 1099511628211
	}

	return index
}

// segmentOf returns the segment of this key.
func (c *Cache) segmentOf(key string) *segment {
	return c.segments[c.indexOf(key)&(c.segmentSize-1)]
}

// Set sets key and value to Cache.
// The key will not expire.
func (c *Cache) Set(key string, value interface{}, opts ...SetOption) {
	conf := applySetOptions(config.NewDefaultSetConfig(), opts...)
	c.segmentOf(key).set(key, value, conf.TTL)
}

// AutoSet starts a goroutine to execute Set() at fixed duration.
// It returns a channel which can be used to stop this goroutine.
func (c *Cache) AutoSet(key string, loadFunc func(ctx context.Context) (interface{}, error), opts ...AutoSetOption) chan<- struct{} {
	conf := applyAutoSetOptions(config.NewDefaultAutoSetConfig(), opts...)

	quitChan := make(chan struct{}, 1)

	go func() {
		ticker := time.NewTicker(conf.Gap)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if data, err := loadFunc(conf.Ctx); err == nil {
					c.Set(key, data, WithSetTTL(conf.TTL))
				}
			case <-quitChan:
				return
			}
		}
	}()

	return quitChan
}

// Get fetches value of key from c first, and returns it if ok.
func (c *Cache) Get(key string, opts ...GetOption) (interface{}, error) {
	v, ok := c.segmentOf(key).get(key)
	if ok {
		return v, nil
	}

	conf := applyGetOptions(config.NewDefaultGetConfig(), opts...)
	if conf.OnMissed != nil {
		data, err := conf.OnMissed(conf.Ctx)
		if err != nil {
			return nil, err
		}

		if conf.OnMissedSet {
			c.Set(key, data, WithSetTTL(conf.OnMissedSetTTL))
		}
		return data, nil
	}

	return nil, newNotFoundErr(key)
}

// Delete removes the value of key.
// If this key is not existed, nothing will happen.
func (c *Cache) Delete(key string) {
	c.segmentOf(key).delete(key)
}

// DeleteAll removes all keys in Cache.
// Notice that this method is weak-consistency.
func (c *Cache) DeleteAll() {
	for _, segment := range c.segments {
		segment.deleteAll()
	}
}

// Size returns the size of Cache.
// Notice that this method is weak-consistency.
func (c *Cache) Size() int {
	size := 0

	for _, segment := range c.segments {
		size += segment.size()
	}

	return size
}

// Gc removes dead entries in Cache.
// Notice that this method is weak-consistency and
// it doesn't guarantee 100% removed.
func (c *Cache) Gc() {
	for _, segment := range c.segments {
		segment.gc()
	}
}

// AutoGc starts a goroutine to execute Gc() at fixed duration.
// It returns a channel which can be used to stop this goroutine.
func (c *Cache) AutoGc(duration time.Duration) chan<- struct{} {
	quitChan := make(chan struct{})

	go func() {
		ticker := time.NewTicker(duration)
		for {
			select {
			case <-ticker.C:
				c.Gc()
			case <-quitChan:
				ticker.Stop()
				return
			}
		}
	}()

	return quitChan
}
