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
// Created at 2020/09/13 19:13:33

package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {

	// Create a cache and set an entry to cache.
	// The ttl is 3 seconds.
	cache := cachego.NewCache()
	cache.SetWithTTL("key", "value", 3)

	// Check if the key is alive.
	value, ok := cache.Get("key")
	fmt.Println(value, ok) // Output: value true

	// Wait for 5 seconds and check again.
	// Now the key is gone.
	time.Sleep(5 * time.Second)
	value, ok = cache.Get("key")
	fmt.Println(value, ok) // Output: <nil> false
}
