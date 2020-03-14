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
// Created at 2020/03/13 16:15:12

package cache

import (
    "io"
    "time"
)

// Cache is an interface representation of one kind cache.
type Cache interface {

    // of 方法用于获取 key 对应的 value。
    // 如果获取成功，返回获取到的 value 和 true，否则返回 nil 和 false。
    Of(key string) *cacheValue

    // OfDefault 方法用于获取 key 对应的 value，和 of 方法不同的是，
    // 如果获取成功，返回获取到的 value，否则返回 defaultValue。
    OfDefault(key string, defaultValue interface{}) *cacheValue

    // Put 方法用于将一个 key-value 键值对数据放进缓存。
    // 注意：建议实现类在内部为数据设置一个默认的寿命时间，比如 30s 之类的。
    // 但是实现类并不一定同样可以不设置数据寿命时间，让数据永不死亡。
    Put(key string, value interface{})

    // PutWithLife 方法用于将一个 key-value 键值对数据放进缓存，同时设置这个数据的寿命时间。
    PutWithLife(key string, value interface{}, life time.Duration)

    // Change 方法用于将 key 对应的数据更改为 newValue，并返回更改前的数据。
    Change(key string, newValue interface{}) *cacheValue

    // ChangeWithLife 方法用于将 key 对应的数据更改为 newValue，
    // 并设置新的寿命，最后返回更改前的数据。
    ChangeWithLife(key string, newValue interface{}, life time.Duration) *cacheValue

    // Remove 方法用于从缓存中移除指定 key 的数据，并返回这个 key 对应的 value。
    // 如果移除成功，返回这个 key 对应的 value 和 true，否则返回 nil 和 false。
    Remove(key string) *cacheValue

    // RemoveAll 方法用于清空缓存。
    RemoveAll()

    // Gc 方法清理死亡的数据。
    Gc()
}

// AdvancedCache is an extension of Cache interface.
type AdvancedCache interface {

    // Cache means an AdvancedCache implement also has the features of basic cache.
    Cache

    // ChangeFunction is an advanced changing function to change your value on your way.
    // Notice that the howToChange function is safe in concurrency to AdvancedCache, which
    // means the implements should guarantee it on their own way.
    ChangeFunction(key string, howToChange func(value *cacheValue)) *cacheValue

    // Dump is for storing current cache, however, it is still a question that
    // this feature is beWorth？
    Dump(w io.Writer)

    // extend this interface in future versions...
}
