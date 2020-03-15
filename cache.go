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
    return NewCacheWithGcDuration(5 * time.Second)
}

// NewCacheWithGcDuration 返回一个缓存对象
func NewCacheWithGcDuration(gcDuration time.Duration) Cache {
    standardCache := &StandardCache{
        data: make(map[string]cacheValue, 16),
        mu:   &sync.RWMutex{},
    }

    // 开启 GC 后台任务
    standardCache.startGcTask(gcDuration)
    return standardCache
}

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

func (sc *StandardCache) checkResult(key string, value cacheValue, ok bool) bool {

    // 如果 ok 是 false，说明数据无效，检查不通过
    if !ok {
        return false
    }

    // 说明这个数据已经死亡过期，删除数据
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
    if !sc.checkResult(key, result, ok) {
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
    sc.mu.Lock()
    defer sc.mu.Unlock()
    for key, value := range sc.data {
        if value.Life() <= 0 {
            delete(sc.data, key)
            sc.size--
        }
    }
}

func (sc *StandardCache) Extend() AdvancedCache {
    return sc
}
