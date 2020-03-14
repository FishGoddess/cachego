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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/14 16:28:56

package cache

import (
    "sync"
    "time"
)

// StandardCache is a standard cache implements AdvancedCache interface.
type StandardCache struct {
    data map[string]cacheValue
    size int
    mu   *sync.RWMutex
}

// NewCache 返回一个缓存对象
func NewCache() Cache {
    return &StandardCache{
        data: make(map[string]cacheValue, 16),
        mu:   &sync.RWMutex{},
    }
}

func (sc *StandardCache) afterOf(key string, value cacheValue) bool {
    if value.Life() <= 0 {
        delete(sc.data, key)
        return false
    }
    return true
}

func (sc *StandardCache) Of(key string) *cacheValue {
    sc.mu.RLock()
    defer sc.mu.RUnlock()
    result, ok := sc.data[key]
    if !ok || !sc.afterOf(key, result) {
        return InvalidCacheValue
    }
    return &result
}

func (sc *StandardCache) Put(key string, value interface{}, life time.Duration) {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    sc.data[key] = *NewCacheValue(value, life)
    sc.size++
}

func (sc *StandardCache) Change(key string, newValue interface{}) {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    oldValue := sc.data[key]
    sc.data[key] = *NewCacheValue(newValue, (&oldValue).Life())
}

func (sc *StandardCache) Remove(key string) {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    delete(sc.data, key)
    sc.size--
}

func (sc *StandardCache) RemoveAll() {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    for key := range sc.data {
        delete(sc.data, key)
    }
    sc.size = 0
}

func (sc *StandardCache) Gc() {
    // Do nothing...
    // implement in future versions...
}

func (sc *StandardCache) Extend() AdvancedCache {
    return sc
}
