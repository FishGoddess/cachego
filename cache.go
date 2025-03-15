// Copyright 2025 FishGoddess. All Rights Reserved.
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
	"context"
	"time"

	"github.com/FishGoddess/cachego/pkg/task"
)

const (
	// NoTTL means a key is never expired.
	NoTTL = 0
)

var (
	newCaches = map[CacheType]func(conf *config) Cache{
		standard: newStandardCache,
		lru:      newLRUCache,
		lfu:      newLFUCache,
	}
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

	// Set sets key and value to cache with ttl and returns evicted value if exists.
	// See NoTTL if you want your key is never expired.
	Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{})

	// Remove removes key and returns the removed value of key.
	// A nil value will be returned if key doesn't exist in cache.
	Remove(key string) (removedValue interface{})

	// Size returns the count of keys in cache.
	// The result may be different in different implements.
	Size() (size int)

	// GC cleans the expired keys in cache and returns the exact count cleaned.
	// The exact cleans depend on implements, however, all implements should have a limit of scanning.
	GC() (cleans int)

	// Reset resets cache to initial status which is like a new cache.
	Reset()

	// Load loads a key with ttl to cache and returns an error if failed.
	// We recommend you use this method to load missed keys to cache,
	// because it may use singleflight to reduce the times calling load function.
	Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error)
}

func newCache(withReport bool, opts ...Option) (cache Cache, reporter *Reporter) {
	conf := newDefaultConfig()
	applyOptions(conf, opts)

	newCache, ok := newCaches[conf.cacheType]
	if !ok {
		panic("cachego: cache type doesn't exist")
	}

	if conf.shardings > 0 {
		cache = newShardingCache(conf, newCache)
	} else {
		cache = newCache(conf)
	}

	if withReport {
		cache, reporter = report(conf, cache)
	}

	if conf.gcDuration > 0 {
		RunGCTask(cache, conf.gcDuration)
	}

	return cache, reporter
}

// NewCache creates a cache with options.
// By default, it will create a standard cache which uses one lock to solve data race.
// It may cause a big performance problem in high concurrency.
// You can use WithShardings to create a sharding cache which is good for concurrency.
// Also, you can use options to specify the type of cache to others, such as lru.
// Use NewCacheWithReporter to get a reporter for use if you want.
func NewCache(opts ...Option) (cache Cache) {
	cache, _ = newCache(false, opts...)
	return cache
}

// NewCacheWithReport creates a cache and a reporter with options.
// By default, it will create a standard cache which uses one lock to solve data race.
// It may cause a big performance problem in high concurrency.
// You can use WithShardings to create a sharding cache which is good for concurrency.
// Also, you can use options to specify the type of cache to others, such as lru.
func NewCacheWithReport(opts ...Option) (cache Cache, reporter *Reporter) {
	return newCache(true, opts...)
}

// RunGCTask runs a gc task in a new goroutine and returns a cancel function to cancel the task.
// However, you don't need to call it manually for most time, instead, use options is a better choice.
// Making it a public function is for more customizations in some situations.
// For example, using options to run gc task is un-cancelable, so you can use it to run gc task by your own
// and get a cancel function to cancel the gc task.
func RunGCTask(cache Cache, duration time.Duration) (cancel func()) {
	fn := func(ctx context.Context) {
		cache.GC()
	}

	ctx := context.Background()
	ctx, cancel = context.WithCancel(ctx)

	go task.New(fn).Context(ctx).Duration(duration).Run()
	return cancel
}
