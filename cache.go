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

const (
    // 默认寿命时间，60 秒
    DefaultLife = 60 * time.Second
)

var (
    // 无效的缓存数据
    InvalidCacheValue = NewCacheValue(nil, false, 0)
)

// StandardCache is a standard cache implements AdvancedCache interface.
type StandardCache struct {
    data map[string]*cacheValue

    size int

    mu sync.RWMutex

    cacheValuePool *sync.Pool
}

// NewCache 创建一个标准的缓存对象，并返回
func NewCache() Cache {
    return NewCacheWithLife(DefaultLife)
}

func NewCacheWithLife(defaultLife time.Duration) Cache {
    return &StandardCache{
        data: make(map[string]*cacheValue, 16),
        cacheValuePool: &sync.Pool{
            New: func() interface{} {
                return NewCacheValue(nil, true, defaultLife)
            },
        },
        mu: sync.RWMutex{},
    }
}

func (sc *StandardCache) wrap(value interface{}) *cacheValue {
    newCacheValue := sc.cacheValuePool.Get().(*cacheValue)
    newCacheValue.item = value
    return newCacheValue
}

func (sc *StandardCache) Of(key string) *cacheValue {
    sc.mu.RLock()
    result, ok := sc.data[key]
    if !ok {
        sc.mu.RUnlock()
        return InvalidCacheValue
    }
    sc.mu.RUnlock()
    return result
}

func (sc *StandardCache) OfDefault(key string, defaultValue interface{}) *cacheValue {
    panic("implement me")
}

func (sc *StandardCache) Put(key string, value interface{}) {
    panic("implement me")
}

func (sc *StandardCache) PutWithLife(key string, value interface{}, life time.Duration) {
    panic("implement me")
}

func (sc *StandardCache) Change(key string, newValue interface{}) *cacheValue {
    panic("implement me")
}

func (sc *StandardCache) ChangeWithLife(key string, newValue interface{}, life time.Duration) *cacheValue {
    panic("implement me")
}

func (sc *StandardCache) Remove(key string) *cacheValue {
    panic("implement me")
}

func (sc *StandardCache) RemoveAll() {
    panic("implement me")
}

func (sc *StandardCache) Gc() {
    panic("implement me")
}
