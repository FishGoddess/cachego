// Copyright 2020 FishGoddess. All Rights Reserved.
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

package cachego

import (
	"errors"
	"time"

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
	conf config

	// segments is a slice stores the real data.
	segments []*segment

	// groups is a slice stores the singleflight keys.
	groups []*singleflight.Group
}

// NewCache returns a new Cache holder for use.
func NewCache(opts ...Option) *Cache {
	c := &Cache{
		conf: *applyOptions(newDefaultConfig(), opts),
	}

	c.segments = newSegments(c.conf.mapSize, c.conf.segmentSize)

	if c.conf.singleflight {
		c.groups = newGroups(c.conf.mapSize, c.conf.segmentSize)
	}

	if c.conf.gcDuration > 0 {
		c.AutoGC(c.conf.gcDuration)
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

// segmentOf returns the segment of this key.
func (c *Cache) segmentOf(key string) *segment {
	return c.segments[Index(key)&(len(c.segments)-1)]
}

// groupOf returns the singleflight group of this key.
func (c *Cache) groupOf(key string) *singleflight.Group {
	return c.groups[Index(key)&(len(c.groups)-1)]
}

// Get fetches value of key from cache first, and returns it if ok.
// Returns an NotFoundErr if this key is not found, and you can use IsNotFound to judge if this error is not found.
// Also, you can specify a function which will be called if missed, so you can load this entry to cache again.
// See OpOption.
func (c *Cache) Get(key string, opts ...OpOption) (interface{}, error) {
	if v, ok := c.segmentOf(key).get(key); ok {
		return v, nil
	}

	if len(opts) <= 0 {
		return nil, errNotFound
	}

	conf := applyOpOptions(newDefaultGetConfig(), opts)
	if conf.onMissed == nil {
		return nil, errNotFound
	}

	var data interface{}
	var err error
	if c.conf.singleflight && conf.singleflight {
		data, err = c.groupOf(key).Call(conf.ctx, key, conf.onMissed)
	} else {
		data, err = conf.onMissed(conf.ctx)
	}

	if err != nil {
		return nil, err
	}

	if conf.reload {
		c.Set(key, data, WithOpTTL(conf.ttl))
	}

	return data, nil
}

// Set sets key and value to cache.
// In default, this entry will not expire, so if you want it to expire, see SetOption.
func (c *Cache) Set(key string, value interface{}, opts ...OpOption) {
	conf := applyOpOptions(newDefaultSetConfig(), opts)
	c.segmentOf(key).set(key, value, conf.ttl)
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
