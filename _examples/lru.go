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

	// Sometimes we want it evicts entries by lru, try WithLRU.
	// You need to specify the max entries storing in lru cache.
	// More details see https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU).
	cache = cachego.NewCache(cachego.WithLRU(10))

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, cachego.NoTTL)
	}

	// Only the least recently used entries can be got in a lru cache.
	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(key, value, ok)
	}

	// By default, lru will share one lock to do all operations.
	// You can sharding cache to several parts for higher performance.
	// Notice that max entries only effect to one part in sharding mode.
	// For example, the total max entries will be 2*10 if shardings is 2 and max entries is 10 in WithLRU or WithMaxEntries.
	// In some cache libraries, they will calculate max entries in each parts of shardings, like 10/2.
	// However, the result divided by max entries and shardings may be not an integer which will make the total max entries incorrect.
	// So we let users decide the exact max entries in each parts of shardings.
	cache = cachego.NewCache(cachego.WithShardings(2), cachego.WithLRU(10))
}
