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

	"github.com/FishGoddess/cachego"
)

func main() {

	// Create a cache for use.
	// We use option function to customize the creation of cache.
	// WithAutoGC means it will do gc automatically.
	cache := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Set a new entry to cache.
	cache.Set("key", 666)

	// Get returns the value of this key.
	v, ok := cache.Get("key")
	fmt.Println(v, ok) // Output: 666 true

	// If you pass a not existed key to of method, nil and false will be returned.
	v, ok = cache.Get("not existed key")
	fmt.Println(v, ok) // Output: <nil> false

	// SetWithTTL sets an entry with expired time.
	// The unit of expired time is second.
	// See more information in example of ttl.
	cache.SetWithTTL("ttlKey", 123, 10)

	// Also, you can get value from cache first, then load it to cache if missed.
	// onMissed is usually used to get data from db or somewhere, so you can refresh the value in cache.
	cache.GetWithLoad("newKey", func() (data interface{}, ttl int64, err error) {
		return "newValue", 3, nil
	})

	// We provide a way to set data to cache automatically, so you can access some hottest data extremely fast.
	stopAutoSet := cache.AutoSet("autoKey", 1*time.Second, func() (interface{}, error) {
		fmt.Println("AutoSet invoking...")
		return nil, nil
	})

	// Keep main running in order to see what AutoSet did
	time.Sleep(5 * time.Second)
	stopAutoSet <- struct{}{}
}
