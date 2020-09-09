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
// Created at 2020/03/13 16:15:56

/*
Package cache provides an easy way to use foundation for your caching operations.

1. the basic usage:

	// Create a cache for use.
	cache := cachego.NewCache()

	// Put a new entry in cache.
	cache.Put("key", 666)

	// Of returns the value of this key.
	v := cache.Of("key")
	fmt.Println(v) // Output: 666

	// If you want to change the value of a key, just put a new value of this key.
	cache.Put("key", "value")

	// See what value it has.
	s := cache.Of("key")
	fmt.Println(s) // Output: value

*/
package cachego // import "github.com/FishGoddess/cachego"

// Version is the version string representation of cachego.
const Version = "v0.1.0"
