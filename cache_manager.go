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

    // Put 方法用于将一个 key-value 键值对数据放进缓存，同时设置这个数据的寿命时间。
    Put(key string, value interface{}, life time.Duration)

    // Change 方法用于将 key 对应的数据更改为 newValue，并返回更改前的数据。
    Change(key string, newValue interface{})

    // Remove 方法用于从缓存中移除指定 key 的数据。如果 key 不存在，就不会发生任何的事情。
    Remove(key string)

    // RemoveAll 方法用于清空缓存。
    RemoveAll()

    // Gc 方法清理死亡的数据。
    Gc()

    // Extend 方法返回拥有更多高级特性的缓存对象
    Extend() AdvancedCache
}

// AdvancedCache is an extension of Cache interface.
type AdvancedCache interface {

    // Cache means an AdvancedCache implement also has the features of basic cache.
    Cache

    // Size returns the size of current cache.
    Size() int

    // Dump is for storing current cache, however, it is still a question that
    // this feature is beWorth？
    Dump(w io.Writer)

    // extend this interface in future versions...
}
