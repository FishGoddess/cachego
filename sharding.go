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

package cachego

import (
	"math/bits"
	"time"
)

type shardingCache struct {
	*config
	caches []Cache
}

func newShardingCache(conf *config, newCache func(conf *config) Cache) Cache {
	if conf.shardings <= 0 {
		panic("cachego: shardings must be > 0.")
	}

	if bits.OnesCount(uint(conf.shardings)) > 1 {
		panic("cachego: shardings must be the pow of 2 (such as 64).")
	}

	caches := make([]Cache, 0, conf.shardings)
	for i := 0; i < conf.shardings; i++ {
		caches = append(caches, newCache(conf))
	}

	return &shardingCache{
		config: conf,
		caches: caches,
	}
}

func (sc *shardingCache) cacheOf(key string) Cache {
	hash := sc.hash(key)
	mask := len(sc.caches) - 1
	return sc.caches[hash&mask]
}

// Get gets the value of key from cache and returns value if found.
func (sc *shardingCache) Get(key string) (value interface{}, found bool) {
	return sc.cacheOf(key).Get(key)
}

// Set sets key and value to cache with ttl and returns evicted value if exists and unexpired.
// See Cache interface.
func (sc *shardingCache) Set(key string, value interface{}, ttl time.Duration) (oldValue interface{}) {
	return sc.cacheOf(key).Set(key, value, ttl)
}

// Remove removes key and returns the removed value of key.
// See Cache interface.
func (sc *shardingCache) Remove(key string) (removedValue interface{}) {
	return sc.cacheOf(key).Remove(key)
}

// Size returns the count of keys in cache.
// See Cache interface.
func (sc *shardingCache) Size() (size int) {
	for _, cache := range sc.caches {
		size += cache.Size()
	}

	return size
}

// GC cleans the expired keys in cache and returns the exact count cleaned.
// See Cache interface.
func (sc *shardingCache) GC() (cleans int) {
	for _, cache := range sc.caches {
		cleans += cache.GC()
	}

	return cleans
}

// Reset resets cache to initial status which is like a new cache.
// See Cache interface.
func (sc *shardingCache) Reset() {
	for _, cache := range sc.caches {
		cache.Reset()
	}
}

// Load loads a key with ttl to cache and returns an error if failed.
// See Cache interface.
func (sc *shardingCache) Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error) {
	return sc.cacheOf(key).Load(key, ttl, load)
}
