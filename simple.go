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

type simpleCache struct {
	config

	entries map[string]*entry
	lock    sync.RWMutex
}

func newSimpleCache(conf config) Cache {
	return &simpleCache{
		config:  conf,
		entries: make(map[string]*entry, conf.maps),
	}
}

func (sc *simpleCache) get(key string) (value interface{}, found bool) {
	entry, ok := sc.entries[key]
	if ok && !entry.expired() {
		return entry.value, true
	}

	return nil, false
}

func (sc *simpleCache) evict() (evictedValue interface{}) {
	for key := range sc.entries {
		return sc.remove(key)
	}

	return nil
}

func (sc *simpleCache) set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
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

func (sc *simpleCache) remove(key string) (removedValue interface{}) {
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

func (sc *simpleCache) clean(allKeys bool) (cleans int) {
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

func (sc *simpleCache) count(allKeys bool) (count int) {
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

func (sc *simpleCache) Get(key string) (value interface{}, found bool) {
	sc.lock.RLock()
	defer sc.lock.RUnlock()

	return sc.get(key)
}

func (sc *simpleCache) Set(key string, value interface{}, ttl time.Duration) (oldValue interface{}) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.set(key, value, ttl)
}

func (sc *simpleCache) Remove(key string) (removedValue interface{}) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.remove(key)
}

func (sc *simpleCache) Clean(allKeys bool) (cleans int) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.clean(allKeys)
}

func (sc *simpleCache) Count(allKeys bool) (count int) {
	sc.lock.RLock()
	defer sc.lock.RUnlock()

	return sc.count(allKeys)
}
