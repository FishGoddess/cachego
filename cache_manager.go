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

// ****************************************
// 中国加油！我相信世界会给我们一个真诚的感谢！    *
//      致谢全体医护人员！感谢您们辛劳的付出！    *
//                           2020-3-16    *
// ****************************************

// Cache is an interface representation of one kind cache.
// In cachego, StandardCache is the only implement of this interface right now.
// However, this interface is not just for cachego. It is for everyone.
// Maybe you are using cachego and StandardCache, but someday your project may
// have more cached data and need higher performance even distributed cache.
// We don't want you to be bother with cachego, so one recommended way to use cachego
// is to use cache interface in code. You just need an implement of another cache.
type Cache interface {

    // Of returns the value of this key.
    Of(key string) *cacheValue

    // Put stores an entry (key, value) to cache, and sets the life of this entry.
    Put(key string, value interface{}, life time.Duration)

    // Change changes the value of key to newValue.
    // If this key is not existed, nothing will happen.
    Change(key string, newValue interface{})

    // Remove removes the value of key.
    // If this key is not existed, nothing will happen.
    Remove(key string)

    // RemoveAll is for removing all data in cache.
    RemoveAll()

    // Gc is for cleaning up dead data.
    // Notice that this method will take lots of time to remove all dead data
    // if there are many entries in cache. So it is not recommended to call Gc()
    // manually. Let cachego do this automatically will be better.
    Gc()

    // Extend returns a cache instance with advanced features.
    // Notice that this method is for extension, so implement it is not required.
    // You can just return a nil in method body.
    Extend() AdvancedCache
}

// AdvancedCache is a Cache interface with advanced features.
// Maybe you are confused for why cachego has two interfaces representation of
// cache, but you should know that some features is for some people, not for everyone.
// So splitting a complex interface into two lightweight interfaces is better.
// For example, some people just need a basic cache for speeding up their data-access process,
// so they use Cache interface is enough. However, some people need more advanced features
// like endurance, so they need a more advanced cache interface to do that.
// This is an isolation measure for basic users and advanced users.
type AdvancedCache interface {

    // Cache means an AdvancedCache implement also has the features of basic cache.
    Cache

    // Size returns the size of current cache.
    Size() int

    // Dump is for endurance, however, it is still a question that
    // this feature is beWorth？
    Dump(w io.Writer)

    // extend this interface in future versions...
}
