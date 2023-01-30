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
	"container/list"
	"time"
)

type lruCache struct {
	cache

	elementMap  map[string]*list.Element
	elementList *list.List
}

func newLRUCache(conf config) Cache {
	if conf.maxEntries <= 0 {
		panic("cachego: lru cache must specify max entries")
	}

	cache := &lruCache{
		elementMap:  make(map[string]*list.Element, MapInitialCap),
		elementList: list.New(),
	}

	cache.setup(conf, cache)
	return cache
}

func (lc *lruCache) unwrap(element *list.Element) *entry {
	entry, ok := element.Value.(*entry)
	if !ok {
		panic("cachego: failed to unwrap lru element's value to entry")
	}

	return entry
}

func (lc *lruCache) evict() (evictedValue interface{}) {
	if element := lc.elementList.Back(); element != nil {
		return lc.removeElement(element)
	}

	return nil
}

func (lc *lruCache) get(key string) (value interface{}, found bool) {
	element, ok := lc.elementMap[key]
	if !ok {
		return nil, false
	}

	entry := lc.unwrap(element)
	if entry.expired(0) {
		return nil, false
	}

	lc.elementList.MoveToFront(element)
	return entry.value, true
}

func (lc *lruCache) set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	element, ok := lc.elementMap[key]
	if ok {
		entry := lc.unwrap(element)
		entry.setup(key, value, ttl)

		lc.elementList.MoveToFront(element)
		return nil
	}

	if lc.maxEntries > 0 && lc.elementList.Len() >= lc.maxEntries {
		evictedValue = lc.evict()
	}

	element = lc.elementList.PushFront(newEntry(key, value, ttl))
	lc.elementMap[key] = element

	return evictedValue
}

func (lc *lruCache) removeElement(element *list.Element) (removedValue interface{}) {
	entry := lc.unwrap(element)

	delete(lc.elementMap, entry.key)
	lc.elementList.Remove(element)

	return entry.value
}

func (lc *lruCache) remove(key string) (removedValue interface{}) {
	if element, ok := lc.elementMap[key]; ok {
		return lc.removeElement(element)
	}

	return nil
}

func (lc *lruCache) size() (size int) {
	return len(lc.elementMap)
}

func (lc *lruCache) gc() (cleans int) {
	now := Now()
	scans := 0
	element := lc.elementList.Back()

	for element != nil {
		scans++

		if entry := lc.unwrap(element); entry.expired(now) {
			old := element
			element = element.Prev()

			lc.removeElement(old)
			cleans++
		} else {
			element = element.Prev()
		}

		if lc.maxScans > 0 && scans >= lc.maxScans {
			break
		}
	}

	return cleans
}

func (lc *lruCache) reset() {
	lc.elementMap = make(map[string]*list.Element, MapInitialCap)
	lc.elementList = list.New()
	lc.Loader.Reset()
}

// Get gets the value of key from cache and returns value if found.
// See Cache interface.
func (lc *lruCache) Get(key string) (value interface{}, found bool) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.get(key)
}

// Set sets key and value to cache with ttl and returns evicted value if exists and unexpired.
// See Cache interface.
func (lc *lruCache) Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.set(key, value, ttl)
}

// Remove removes key and returns the removed value of key.
// See Cache interface.
func (lc *lruCache) Remove(key string) (removedValue interface{}) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.remove(key)
}

// Size returns the count of keys in cache.
// See Cache interface.
func (lc *lruCache) Size() (size int) {
	lc.lock.RLock()
	defer lc.lock.RUnlock()

	return lc.size()
}

// GC cleans the expired keys in cache and returns the exact count cleaned.
// See Cache interface.
func (lc *lruCache) GC() (cleans int) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.gc()
}

// Reset resets cache to initial status which is like a new cache.
// See Cache interface.
func (lc *lruCache) Reset() {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	lc.reset()
}
