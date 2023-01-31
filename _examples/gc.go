// Copyright 2023 FishGoddess. All Rights Reserved.
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

package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
	cache := cachego.NewCache()
	cache.Set("key", 666, time.Second)

	time.Sleep(2 * time.Second)

	// The entry is expired after ttl.
	value, ok := cache.Get("key")
	fmt.Println(value, ok) // <nil> false

	// As you know the entry still stores in cache even if it's expired.
	// This is because we think you will reset entry to cache after cache missing in most situations.
	// So we can reuse this entry and just reset its value and ttl.
	size := cache.Size()
	fmt.Println(size) // 1

	// What should I do if I want an expired entry never storing in cache? Try GC:
	cleans := cache.GC()
	fmt.Println(cleans) // 1

	// Is there a smart way to do that? Try WithGC:
	// For testing, we set a small duration of gc.
	// You should set at least 3 minutes in production for performance.
	cache = cachego.NewCache(cachego.WithGC(2 * time.Second))
	cache.Set("key", 666, time.Second)

	size = cache.Size()
	fmt.Println(size) // 1

	time.Sleep(3 * time.Second)

	size = cache.Size()
	fmt.Println(size) // 0
}
