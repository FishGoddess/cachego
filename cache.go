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
// It is a k-v entry cache that stores in memory. Actually, this cache
// is a concurrency-safe map essentially. That means it can be visited
// with many goroutines at the same time. More than a map does, It keeps
// a background task that removes all dead key and value, also, you can
// call Gc() manually to invoke this clean up process.
type StandardCache struct {

    // data is the map stores all k-v entries.
    // Cache is a concurrency-safe map essentially, remember?
    data map[string]cacheValue

    // mu is for concurrency-safe. It is a lock.
    mu *sync.RWMutex

    // size is a field representation of how many entries are storing in current cache.
    size int
}

// NewCache Returns a cache implemented AdvancedCache interface.
// Notice that default gc duration is ten minutes. The gc duration will affect the performance
// of cache, so do not set it too small.
func NewCache() Cache {
    return NewCacheWithGcDuration(10 * time.Minute)
}

// NewCacheWithGcDuration Returns a cache implemented AdvancedCache interface.
// The gc duration will affect the performance of cache, so do not set it too small.
func NewCacheWithGcDuration(gcDuration time.Duration) Cache {
    standardCache := &StandardCache{
        data: make(map[string]cacheValue, 64),
        mu:   &sync.RWMutex{},
    }

    // 开启 GC 后台任务
    standardCache.startGcTask(gcDuration)
    return standardCache
}

// startGcTask starts a goroutine to clean up dead entries at fixed gcDuration.
func (sc *StandardCache) startGcTask(gcDuration time.Duration) {
    go func() {
        ticker := time.NewTicker(gcDuration)
        for {
            select {
            case <-ticker.C:
                sc.Gc()
            }
        }
    }()
}

// checkResult
func (sc *StandardCache) verifyResult(key string, value cacheValue, ok bool) bool {

    // 如果 ok 是 false，说明数据无效，检查不通过
    if !ok {
        return false
    }

    // 说明这个数据已经死亡过期，删除数据
    if value.Dead() {
        delete(sc.data, key)
        return false
    }
    return true
}

func (sc *StandardCache) Of(key string) *cacheValue {
    sc.mu.RLock()
    defer sc.mu.RUnlock()
    result, ok := sc.data[key]
    if !sc.verifyResult(key, result, ok) {
        return InvalidCacheValue()
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
    oldValue, ok := sc.data[key]
    if ok {
        sc.data[key] = *NewCacheValue(newValue, (&oldValue).Life())
    }
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
    sc.data = make(map[string]cacheValue, 64)
    sc.size = 0
}

func (sc *StandardCache) Gc() {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    for key, value := range sc.data {
        if value.Dead() {
            delete(sc.data, key)
            sc.size--
        }
    }
}

func (sc *StandardCache) Extend() AdvancedCache {
    return sc
}
