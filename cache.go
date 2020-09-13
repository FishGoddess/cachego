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
	defaultMapSize     = 1024
	defaultSegmentSize = 1024
)

// Cache is a struct of cache.
type Cache struct {
	mapSize     int
	segmentSize int
	segments    []*segment
}

// NewCache returns a new Cache holder.
func NewCache() *Cache {
	return &Cache{
		mapSize:     defaultMapSize,
		segmentSize: defaultSegmentSize,
		segments:    newSegments(defaultMapSize, defaultSegmentSize),
	}
}

func newSegments(mapSize int, segmentSize int) []*segment {
	segments := make([]*segment, segmentSize)
	for i := 0; i < segmentSize; i++ {
		segments[i] = newSegment(mapSize)
	}
	return segments
}

func index(key string) int {
	index := 0
	keyBytes := []byte(key)
	for _, b := range keyBytes {
		index = 31*index + int(b&0xff)
	}
	return index
}

func (c *Cache) segmentOf(key string) *segment {
	return c.segments[index(key)&(c.segmentSize-1)]
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.segmentOf(key).get(key)
}

func (c *Cache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, NeverDie)
}

func (c *Cache) SetWithTTL(key string, value interface{}, ttl int64) {
	c.segmentOf(key).set(key, value, ttl)
}

// Remove removes the value of key.
// If this key is not existed, nothing will happen.
func (c *Cache) Remove(key string) {
	c.segmentOf(key).remove(key)
}

func (c *Cache) RemoveAll() {
	for _, segment := range c.segments {
		segment.removeAll()
	}
}

func (c *Cache) Size() int {
	size := 0
	for _, segment := range c.segments {
		size += segment.size()
	}
	return size
}

// Gc is for cleaning up dead data.
// Notice that this method will take lots of time to remove all dead data
// if there are many entries in cache. So it is not recommended to call Gc()
// manually. Let cachego do this automatically will be better.
func (c *Cache) Gc() {
	for _, segment := range c.segments {
		segment.gc()
	}
}

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
