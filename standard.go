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

	entries map[string]*entry
	loader  *Loader

	lock sync.RWMutex
}

func newStandardCache(conf config) Cache {
	cache := &standardCache{
		config:  conf,
		entries: make(map[string]*entry, MapInitialCap),
	}

	cache.loader = NewLoader(cache, conf.singleflight)
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

	if sc.maxEntries > 0 && sc.count(true) >= sc.maxEntries {
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

func (sc *standardCache) clean(allKeys bool) (cleans int) {
	scans := 0
	for _, entry := range sc.entries {
		scans++

		if allKeys || entry.expired() {
			delete(sc.entries, entry.key)
			cleans++
		}

		if sc.maxScans > 0 && scans >= sc.maxScans {
			break
		}
	}

	return cleans
}

func (sc *standardCache) count(allKeys bool) (count int) {
	if allKeys {
		return len(sc.entries)
	}

	for _, entry := range sc.entries {
		if !entry.expired() {
			count++
		}
	}

	return count
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

// Clean cleans some keys in cache and returns the exact count cleaned by cache.
// See Cache interface.
func (sc *standardCache) Clean(allKeys bool) (cleans int) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.clean(allKeys)
}

// Count returns the count of keys in cache.
// See Cache interface.
func (sc *standardCache) Count(allKeys bool) (count int) {
	sc.lock.RLock()
	defer sc.lock.RUnlock()

	return sc.count(allKeys)
}

// Load loads a key with ttl to cache and returns an error if failed.
// See Cache interface.
func (sc *standardCache) Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error) {
	return sc.loader.Load(key, ttl, load)
}
