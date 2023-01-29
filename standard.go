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
	"sync"
	"time"
)

type standardCache struct {
	config
	Loader

	entries map[string]*entry

	lock sync.RWMutex
}

func newStandardCache(conf config) Cache {
	cache := &standardCache{
		config:  conf,
		entries: make(map[string]*entry, MapInitialCap),
	}

	cache.Loader = NewLoader(cache, conf.singleflight)
	return cache
}

func (sc *standardCache) get(key string) (value interface{}, found bool) {
	entry, ok := sc.entries[key]
	if ok && !entry.expired() {
		return entry.value, true
	}

	return nil, false
}

func (sc *standardCache) evict() (evictedValue interface{}) {
	for key := range sc.entries {
		return sc.remove(key)
	}

	return nil
}

func (sc *standardCache) set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	entry, ok := sc.entries[key]
	if ok {
		entry.setup(key, value, ttl)
		return nil
	}

	if sc.maxEntries > 0 && sc.size() >= sc.maxEntries {
		evictedValue = sc.evict()
	}

	sc.entries[key] = newEntry(key, value, ttl)
	return evictedValue
}

func (sc *standardCache) remove(key string) (removedValue interface{}) {
	entry, ok := sc.entries[key]
	if !ok {
		return nil
	}

	if !entry.expired() {
		removedValue = entry.value
	}

	delete(sc.entries, key)
	return removedValue
}

func (sc *standardCache) size() (size int) {
	return len(sc.entries)
}

func (sc *standardCache) gc() (cleans int) {
	scans := 0
	for _, entry := range sc.entries {
		scans++

		if entry.expired() {
			delete(sc.entries, entry.key)
			cleans++
		}

		if sc.maxScans > 0 && scans >= sc.maxScans {
			break
		}
	}

	return cleans
}

func (sc *standardCache) reset() {
	sc.entries = make(map[string]*entry, MapInitialCap)
	sc.Loader.Reset()
}

// Get gets the value of key from cache and returns value if found.
// See Cache interface.
func (sc *standardCache) Get(key string) (value interface{}, found bool) {
	sc.lock.RLock()
	defer sc.lock.RUnlock()

	return sc.get(key)
}

// Set sets key and value to cache with ttl and returns evicted value if exists and unexpired.
// See Cache interface.
func (sc *standardCache) Set(key string, value interface{}, ttl time.Duration) (oldValue interface{}) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.set(key, value, ttl)
}

// Remove removes key and returns the removed value of key.
// See Cache interface.
func (sc *standardCache) Remove(key string) (removedValue interface{}) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.remove(key)
}

// Size returns the count of keys in cache.
// See Cache interface.
func (sc *standardCache) Size() (size int) {
	sc.lock.RLock()
	defer sc.lock.RUnlock()

	return sc.size()
}

// GC cleans the expired keys in cache and returns the exact count cleaned.
// See Cache interface.
func (sc *standardCache) GC() (cleans int) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.gc()
}

// Reset resets cache to initial status which is like a new cache.
// See Cache interface.
func (sc *standardCache) Reset() {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	sc.reset()
}
