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
	"strconv"

	"github.com/FishGoddess/cachego"
)

func main() {
	// By default, NewCache() returns a standard cache which evicts entries randomly.
	cache := cachego.NewCache(cachego.WithMaxEntries(10))

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, cachego.NoTTL)
	}

	// Since we set 20 entries to cache, the size won't be 20 because we limit the max entries to 10.
	size := cache.Size()
	fmt.Println(size) // 10

	// We don't know which entries will be evicted and stayed.
	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(key, value, ok)
	}

	fmt.Println()

	// Sometimes we want it evicts entries by lfu, try WithLFU.
	// You need to specify the max entries storing in lfu cache.
	// More details see https://en.wikipedia.org/wiki/Cache_replacement_policies#Least-frequently_used_(LFU).
	cache = cachego.NewCache(cachego.WithLFU(10))

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)

		// Let entries have some frequently used operations.
		for j := 0; j < i; j++ {
			cache.Set(key, i, cachego.NoTTL)
		}
	}

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(key, value, ok)
	}
}
