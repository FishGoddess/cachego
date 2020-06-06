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
// Created at 2020/03/16 22:29:02

package main

import (
	"fmt"
	"time"

	cache "github.com/FishGoddess/cachego"
)

func main() {

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

	// If you want to change the value of key, try this:
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
}
