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
// Created at 2020/03/13 16:15:56

/*
Package cache provides an easy way to use foundation for your caching operations.

1. the basic usage:

    // Create a cache with default gc duration (10 minutes).
    newCache := cache.NewCache()

    // Put a new entry in cache.
    // This entry will be dead after 5 seconds.
    // However, it will be deleted after 10 minutes if you never access.
    newCache.Put("key", 666, 5*time.Second)

    // Of returns the value of this key.
    // As you know, this is chain-programming api.
    // If you need int type, just call Int().
    v := newCache.Of("key").Int()
    fmt.Println(v) // Output: 666

    // If you want change the value of key, try this:
    newCache.Change("key", "value")

    // Then you can call String() behind Of().
    s := newCache.Of("key").String()
    fmt.Println(s) // Output: value

    // After 5 seconds, this entry will dead, then an invalidCacheValue will be returned.
    time.Sleep(5 * time.Second)
    ok := newCache.Of("key").Ok()
    fmt.Println(ok) // Output: false

    // Maybe you want a default value for some situations, such as the code above.
    // Use Or() to help you to do that:
    s = newCache.Of("key").Or("default value").String()
    fmt.Println(s) // Output: default value

2. cache value usage:

    // Create a cache with default gc duration (10 minutes).
    newCache := cache.NewCache()

    // Put a new entry in cache.
    // This entry will be dead after 5 seconds.
    // However, it will be deleted after 10 minutes if you never access.
    newCache.Put("key", 666, 5*time.Second)

    // Of returns the value of this key.
    // As you know, this is chain-programming api.
    // If you need int type, just call Int().
    v := newCache.Of("key").Int()
    fmt.Println(v) // Output: 666

    // If you need another type like string, just call String().
    // But you should know, this is not gonna work because the value is int
    // type in fact, so it will return "".
    nilStr := newCache.Of("key").String()
    fmt.Printf("%q\n", nilStr) // Output: ""

    // Sometimes you don't know the real type of value, and you try to
    // convert to some type, try this:
    // TryXxx returns two results (value, ok or not). If ok, this value will be valid.
    nilFloat32, ok := newCache.Of("key").TryFloat32()
    fmt.Println(nilFloat32, ok) // Output: 0 false

    // Of cause, there are more functions for other type:
    newCache.Of("key").Int8()
    newCache.Of("key").Int16()
    newCache.Of("key").Int32()
    newCache.Of("key").Int64()
    newCache.Of("key").Float32()
    newCache.Of("key").Float64()
    newCache.Of("key").String()
    newCache.Of("key").TryInt8()
    newCache.Of("key").TryInt16()
    newCache.Of("key").TryInt32()
    newCache.Of("key").TryInt64()
    newCache.Of("key").TryFloat32()
    newCache.Of("key").TryFloat64()
    newCache.Of("key").TryString()

    // Of cause, we have the most original method Value() for you to get the value.
    value, ok := newCache.Of("key").Value()
    fmt.Println(value, ok) // Output: 666 true

    // There are some functions for you to get info of value:
    ok = newCache.Of("key").Ok()      // Return true if this value is valid.
    dead := newCache.Of("key").Dead() // Return true if this value is dead.
    life := newCache.Of("key").Life() // Return the remained life of this value.
    fmt.Println(ok, dead, life)       // Output: true false 4.9990022s

*/
package cache // import "github.com/FishGoddess/cachego"

// Version is the version string representation of cachego.
const Version = "v0.0.1"
