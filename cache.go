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

// Cache is a struct of cache.
type Cache struct {

	// data is the map stores all k-v entries.
	// Cache is a concurrency-safe map essentially, remember?
	data *Map
}

func NewCache() *Cache {
	return &Cache{
		data: NewMap(),
	}
}

// Of returns the value of this key.
// Return invalidCacheValue if this key is absent in cache.
func (c *Cache) Of(key string) (interface{}, bool) {
	if v, ok := c.data.Get(key); ok && v.(*value).alive() {
		return v.(*value).data, true
	}
	return nil, false
}

func (c *Cache) Put(key string, value interface{}) {
	c.PutWithTTL(key, value, NeverDie)
}

func (c *Cache) PutWithTTL(key string, value interface{}, ttl int64) {
	c.data.Set(key, newValue(value, ttl))
}

// Remove removes the value of key.
// If this key is not existed, nothing will happen.
func (c *Cache) Delete(key string) {
	c.data.Delete(key)
}

// RemoveAll is for removing all data in cache.
func (c *Cache) Reset() {
	c.data = NewMap()
}

func (c *Cache) Size() int {
	return c.data.Size()
}

// Gc is for cleaning up dead data.
// Notice that this method will take lots of time to remove all dead data
// if there are many entries in cache. So it is not recommended to call Gc()
// manually. Let cachego do this automatically will be better.
func (c *Cache) Gc() {
	// TODO implement gc
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
