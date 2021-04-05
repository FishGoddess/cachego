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
	"time"
)

const (
	// defaultMapSize determines the initialized map size of one segment.
	defaultMapSize = 1024

	// defaultSegmentSize determines the size of segments.
	// This value will affect the performance of concurrency.
	defaultSegmentSize = 1024
)

// Cache is a struct of cache.
type Cache struct {

	// mapSize is the size of map inside.
	mapSize int

	// segmentSize is the size of segments.
	// This value will affect the performance of concurrency.
	segmentSize int

	// segments is a slice stores the real data.
	segments []*segment
}

// NewCache returns a new Cache holder for use.
func NewCache(options ...Option) *Cache {

	cache := &Cache{
		mapSize:     defaultMapSize,
		segmentSize: defaultSegmentSize,
	}

	// Initializing with options
	for _, applyOption := range options {
		applyOption(cache)
	}

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

// index returns a position in segments of this key.
func index(key string) int {
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
	return c.segments[index(key)&(c.segmentSize-1)]
}

// Get returns the value of key and a false if not found.
func (c *Cache) Get(key string) (interface{}, bool) {
	return c.segmentOf(key).get(key)
}

// Set sets key and value to Cache.
// The key will not expire.
func (c *Cache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, NeverDie)
}

// SetWithTTL sets key and value to Cache with a ttl.
// The unit of ttl is second.
func (c *Cache) SetWithTTL(key string, value interface{}, ttl int64) {
	c.segmentOf(key).set(key, value, ttl)
}

// Remove removes the value of key.
// If this key is not existed, nothing will happen.
func (c *Cache) Remove(key string) {
	c.segmentOf(key).remove(key)
}

// RemoveAll removes all keys in Cache.
// Notice that this method is weak-consistency.
func (c *Cache) RemoveAll() {
	for _, segment := range c.segments {
		segment.removeAll()
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
// It returns a <-chan type which can be used to stop this goroutine.
func (c *Cache) AutoGc(duration time.Duration) chan<- bool {
	quitChan := make(chan bool)
	go func() {
		ticker := time.NewTicker(duration)
		for {
			select {
			case <-ticker.C:
				c.Gc()
			case <-quitChan:
				return
			}
		}
	}()
	return quitChan
}
