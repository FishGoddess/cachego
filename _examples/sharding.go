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
	// All operations in cache share one lock for concurrency.
	// Use read lock or write lock is depends on cache implements.
	// Get will use read lock in standard cache, but lru and lfu don't.
	// This may be a big performance problem in high qps.
	cache := cachego.NewCache()

	// We provide a sharding cache wrapper to shard one cache to several parts with hash.
	// Every parts store its entries and all operations of one entry work on one part.
	// This means there are more than one lock when you operate entries.
	// The performance will be better in high qps.
	cache = cachego.NewCache(cachego.WithShardings(64))
	cache.Set("key", 666, cachego.NoTTL)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 666 true

	// Notice that max entries will be the sum of shards.
	// For example, we set WithShardings(4) and WithMaxEntries(100), and the max entries in whole cache will be 4 * 100.
	cache = cachego.NewCache(cachego.WithShardings(4), cachego.WithMaxEntries(100))

	for i := 0; i < 1000; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, cachego.NoTTL)
	}

	size := cache.Size()
	fmt.Println(size) // 400
}
