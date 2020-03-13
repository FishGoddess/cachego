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

import "time"

// Cache is an interface representation of one kind cache.
// TODO 返回值使用 CacheValue 结构体，达到链式编程效果
type Cache interface {

    // of 方法用于获取 key 对应的 value。
    // 如果获取成功，返回获取到的 value 和 true，否则返回 nil 和 false。
    Of(key string) (interface{}, bool)

    // OfDefault 方法用于获取 key 对应的 value，和 of 方法不同的是，
    // 如果获取成功，返回获取到的 value，否则返回 defaultValue。
    OfDefault(key interface{}, defaultValue interface{}) interface{}

    // Put 方法用于将一个 key-value 键值对数据放进缓存。
    // 注意：建议实现类在内部为数据设置一个默认的寿命时间，比如 30s 之类的。
    // 但是实现类并不一定同样可以不设置数据寿命时间，让数据永不死亡。
    Put(key interface{}, value interface{})

    // PutWithLife 方法用于将一个 key-value 键值对数据放进缓存，同时设置这个数据的寿命时间。
    PutWithLife(key interface{}, value interface{}, life time.Duration)

    // Remove 方法用于从缓存中移除指定 key 的数据，并返回这个 key 对应的 value。
    // 如果移除成功，返回这个 key 对应的 value 和 true，否则返回 nil 和 false。
    Remove(key interface{}) (interface{}, bool)

    // Gc 方法清理死亡的数据
    Gc()
}
