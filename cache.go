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
	"time"
)

const (
	// NoTTL means a key is never expired.
	NoTTL = 0
)

// Cache is the core interface of cachego.
// We provide some implements including standard cache and sharding cache.
type Cache interface {
	// Get gets the value of key from cache and returns value if found.
	// A nil value will be returned if key doesn't exist in cache.
	// Notice that we won't remove expired keys in get method, so you should remove them manually or set a limit of keys.
	// The reason why we won't remove expired keys in get method is for higher re-usability, because we often set a new value
	// of expired key after getting it (so we can reuse the memory of entry).
	Get(key string) (value interface{}, found bool)

	// Set sets key and value to cache with ttl and returns evicted value if exists and unexpired.
	// See NoTTL if you want your key is never expired.
	Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{})

	// Remove removes key and returns the removed value of key.
	// A nil value will be returned if key doesn't exist in cache or expired.
	Remove(key string) (removedValue interface{})

	// Clean cleans some keys in cache and returns the exact count cleaned by cache.
	// By default, it will only clean expired keys in cache, set allKeys=true if you want to clean all keys.
	// Also, we limit the scans of cache due to scanning keys in cache may cause a performance problem, see newDefaultConfig.
	// As you know we won't remove expired keys in get method, so this method is the only way to clean expired keys except set method.
	// We recommend you to set a limit of keys, so you are no need to clean expired keys manually.
	// Or you can start a goroutine to call Clean(false) at a fixed duration, of course, we provide an option for you, see WithGC.
	Clean(allKeys bool) (cleans int)

	// Count returns the count of keys in cache.
	// By default, expired keys won't be counted, set allKeys=true if you want to count all keys.
	Count(allKeys bool) (count int)

	// Load loads a key with ttl to cache and returns an error if failed.
	// We recommend you use this method to load missed keys to cache because it uses singleflight to reduce the times calling load function.
	Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error)
}

func newCache(conf config, newCache func(conf config) Cache) (cache Cache) {
	if conf.shardings > 0 {
		cache = newShardingCache(conf, newCache)
	} else {
		cache = newCache(conf)
	}

	if conf.gcDuration > 0 {
		go func() {
			for {
				time.Sleep(conf.gcDuration)
				cache.Clean(false)
			}
		}()
	}

	return cache
}

// NewStandardCache creates a standard cache with options.
// By default, it will create a cache which uses one lock to solve data race.
// It may cause a big performance problem in high concurrency.
// You can use WithShardings to create a sharding standard cache which is good for concurrency.
func NewStandardCache(opts ...Option) (cache Cache) {
	conf := newDefaultConfig()
	applyOptions(&conf, opts)
	return newCache(conf, newStandardCache)
}
