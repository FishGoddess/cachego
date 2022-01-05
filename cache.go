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
	"errors"
	"time"

	"github.com/FishGoddess/cachego/internal/config"
	"github.com/FishGoddess/cachego/pkg/singleflight"
)

var (
	// errNotFound is the error of key not found.
	errNotFound = errors.New("cachego: key not found")
)

// IsNotFound returns if this error is key not found.
func IsNotFound(err error) bool {
	return err != nil && err == errNotFound
}

// Cache is a struct of cache.
type Cache struct {
	// conf is the config of cache.
	conf config.Config

	// segments is a slice stores the real data.
	segments []*segment

	// groups is a slice stores the singleflight keys.
	groups []*singleflight.Group
}

// NewCache returns a new Cache holder for use.
func NewCache(opts ...Option) *Cache {
	c := &Cache{
		conf: *applyOptions(config.NewDefaultConfig(), opts...),
	}

	c.segments = newSegments(c.conf.MapSize, c.conf.SegmentSize)
	if c.conf.EnableSingleflight {
		c.groups = newGroups(c.conf.MapSize, c.conf.SegmentSize)
	}

	if c.conf.GCDuration > 0 {
		c.AutoGC(c.conf.GCDuration)
	}
	return c
}

// newSegments returns a slice of initialized segments.
func newSegments(mapSize int, segmentSize int) []*segment {
	segments := make([]*segment, segmentSize)
	for i := 0; i < segmentSize; i++ {
		segments[i] = newSegment(mapSize)
	}
	return segments
}

// newGroups returns a slice of initialized singleflight groups.
func newGroups(mapSize int, groupSize int) []*singleflight.Group {
	groups := make([]*singleflight.Group, groupSize)
	for i := 0; i < groupSize; i++ {
		groups[i] = singleflight.NewGroup(mapSize)
	}
	return groups
}

// indexOf returns a position of this key.
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
	return c.segments[c.indexOf(key)&(len(c.segments)-1)]
}

// groupOf returns the singleflight group of this key.
func (c *Cache) groupOf(key string) *singleflight.Group {
	return c.groups[c.indexOf(key)&(len(c.groups)-1)]
}

// Get fetches value of key from cache first, and returns it if ok.
// Returns an NotFoundErr if this key is not found, and you can use IsNotFound to judge if this error is not found.
// Also, you can specify a function which will be called if missed, so you can load this entry to cache again.
// See GetOption.
func (c *Cache) Get(key string, opts ...GetOption) (interface{}, error) {
	v, ok := c.segmentOf(key).get(key)
	if ok {
		return v, nil
	}

	if len(opts) <= 0 {
		return nil, errNotFound
	}

	conf := applyGetOptions(config.NewDefaultGetConfig(), opts...)
	if conf.OnMissed == nil {
		return nil, errNotFound
	}

	var data interface{}
	var err error
	if c.conf.EnableSingleflight && conf.Singleflight {
		data, err = c.groupOf(key).Call(conf.Ctx, key, conf.OnMissed)
	} else {
		data, err = conf.OnMissed(conf.Ctx)
	}

	if err != nil {
		return nil, err
	}

	c.Set(key, data, WithSetTTL(conf.TTL))
	return data, nil
}

// Set sets key and value to cache.
// In default, this entry will not expire, so if you want it to expire, see SetOption.
func (c *Cache) Set(key string, value interface{}, opts ...SetOption) {
	conf := applySetOptions(config.NewDefaultSetConfig(), opts...)
	c.segmentOf(key).set(key, value, conf.TTL)
}

// AutoSet starts a goroutine to execute Set() at fixed duration.
// It returns a channel which can be used to stop this goroutine.
// See AutoSetOption.
func (c *Cache) AutoSet(key string, fn func(ctx context.Context) (interface{}, error), opts ...AutoSetOption) chan<- struct{} {
	conf := applyAutoSetOptions(config.NewDefaultAutoSetConfig(), opts...)

	quitChan := make(chan struct{})
	go func() {
		ticker := time.NewTicker(conf.Gap)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if data, err := fn(conf.Ctx); err == nil {
					c.Set(key, data, WithSetTTL(conf.TTL))
				}
			case <-quitChan:
				return
			}
		}
	}()

	return quitChan
}

// Delete removes the value of key.
// If this key is not existed, nothing will happen.
func (c *Cache) Delete(key string) {
	c.segmentOf(key).delete(key)
	c.groupOf(key).Delete(key)
}

// DeleteAll removes all keys in cache.
// Notice that this method is weak-consistency.
func (c *Cache) DeleteAll() {
	for _, segment := range c.segments {
		segment.deleteAll()
	}

	for _, group := range c.groups {
		group.DeleteAll()
	}
}

// Size returns the size of cache.
// Notice that this method is weak-consistency.
func (c *Cache) Size() int {
	size := 0
	for _, segment := range c.segments {
		size += segment.size()
	}
	return size
}

// GC removes dead entries in cache.
// Notice that this method is weak-consistency, and it doesn't guarantee 100% removed.
func (c *Cache) GC() {
	for _, segment := range c.segments {
		segment.gc()
	}
}

// AutoGC starts a goroutine to execute GC() at fixed duration.
// It returns a channel which can be used to stop this goroutine.
func (c *Cache) AutoGC(duration time.Duration) chan<- struct{} {
	quitChan := make(chan struct{})

	go func() {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				c.GC()
			case <-quitChan:
				return
			}
		}
	}()

	return quitChan
}
