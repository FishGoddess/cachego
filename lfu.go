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

	"github.com/FishGoddess/cachego/pkg/heap"
)

type lfuCache struct {
	cache

	itemMap  map[string]*heap.Item
	itemHeap *heap.Heap
}

func newLFUCache(conf config) Cache {
	if conf.maxEntries <= 0 {
		panic("cachego: lfu cache must specify max entries")
	}

	cache := &lfuCache{
		itemMap:  make(map[string]*heap.Item, MapInitialCap),
		itemHeap: heap.New(SliceInitialCap),
	}

	cache.setup(conf, cache)
	return cache
}

func (lc *lfuCache) unwrap(item *heap.Item) *entry {
	entry, ok := item.Value.(*entry)
	if !ok {
		panic("cachego: failed to unwrap lfu item's value to entry")
	}

	return entry
}

func (lc *lfuCache) get(key string) (value interface{}, found bool) {
	item, ok := lc.itemMap[key]
	if !ok {
		return nil, false
	}

	entry := lc.unwrap(item)
	if entry.expired(0) {
		return nil, false
	}

	item.Adjust(item.Weight() + 1)
	return entry.value, true
}

func (lc *lfuCache) set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	return nil // TODO
}

func (lc *lfuCache) removeItem(item *heap.Item) (removedValue interface{}) {
	entry := lc.unwrap(item)

	delete(lc.itemMap, entry.key)
	lc.itemHeap.Remove(item)

	return entry.value
}

func (lc *lfuCache) remove(key string) (removedValue interface{}) {
	if item, ok := lc.itemMap[key]; ok {
		return lc.removeItem(item)
	}

	return nil
}

func (lc *lfuCache) size() (size int) {
	return len(lc.itemMap)
}

func (lc *lfuCache) gc() (cleans int) {
	now := Now()
	scans := 0

	for _, item := range lc.itemMap {
		scans++

		if entry := lc.unwrap(item); entry.expired(now) {
			lc.removeItem(item)
			cleans++
		}

		if lc.maxScans > 0 && scans >= lc.maxScans {
			break
		}
	}

	return cleans
}

func (lc *lfuCache) reset() {
	lc.itemMap = make(map[string]*heap.Item, MapInitialCap)
	lc.itemHeap = heap.New(SliceInitialCap)
	lc.Loader.Reset()
}

// Get gets the value of key from cache and returns value if found.
// See Cache interface.
func (lc *lfuCache) Get(key string) (value interface{}, found bool) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.get(key)
}

// Set sets key and value to cache with ttl and returns evicted value if exists and unexpired.
// See Cache interface.
func (lc *lfuCache) Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.set(key, value, ttl)
}

// Remove removes key and returns the removed value of key.
// See Cache interface.
func (lc *lfuCache) Remove(key string) (removedValue interface{}) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.remove(key)
}

// Size returns the count of keys in cache.
// See Cache interface.
func (lc *lfuCache) Size() (size int) {
	lc.lock.RLock()
	defer lc.lock.RUnlock()

	return lc.size()
}

// GC cleans the expired keys in cache and returns the exact count cleaned.
// See Cache interface.
func (lc *lfuCache) GC() (cleans int) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.gc()
}

// Reset resets cache to initial status which is like a new cache.
// See Cache interface.
func (lc *lfuCache) Reset() {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	lc.reset()
}
