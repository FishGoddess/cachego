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

	// We think most of the entries in cache should have its ttl.
	// So set an entry to cache should specify a ttl.
	cache.Set("key", 666, time.Second)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 666 true

	time.Sleep(2 * time.Second)

	// The entry is expired after ttl.
	value, ok = cache.Get("key")
	fmt.Println(value, ok) // <nil> false

	// Notice that the entry still stores in cache even if it's expired.
	// This is because we think you will reset entry to cache after cache missing in most situations.
	// So we can reuse this entry and just reset its value and ttl.
	size := cache.Size()
	fmt.Println(size) // 1

	// What should I do if I want an expired entry never storing in cache? Try GC:
	cleans := cache.GC()
	fmt.Println(cleans) // 1

	size = cache.Size()
	fmt.Println(size) // 0

	// However, not all entries have ttl, and you can specify a NoTTL constant to do so.
	// In fact, the entry won't expire as long as its ttl is <= 0.
	// So you may have known NoTTL is a "readable" value of "<= 0".
	cache.Set("key", 666, cachego.NoTTL)
}
