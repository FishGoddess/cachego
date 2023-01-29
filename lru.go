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
	"sync"
	"time"
)

type lruCache struct {
	config

	elementMap  map[string]*list.Element
	elementList *list.List
	loader      Loader

	lock sync.RWMutex
}

func newLRUCache(conf config) Cache {
	cache := &lruCache{
		config:      conf,
		elementMap:  make(map[string]*list.Element, MapInitialCap),
		elementList: list.New(),
	}

	cache.loader = NewLoader(cache, conf.singleflight)
	return cache
}

func (lc *lruCache) onEvicted(key string, value interface{}) {}

func (lc *lruCache) unwrap(element *list.Element) *entry {
	entry, ok := element.Value.(*entry)
	if !ok {
		panic("cachego: failed to unwrap element's value to entry")
	}

	return entry
}

func (lc *lruCache) evict() (evictedValue interface{}) {
	element := lc.elementList.Back()
	if element == nil {
		return nil
	}

	entry := lc.unwrap(element)
	evictedValue = lc.removeElement(element)

	lc.onEvicted(entry.key, entry.value)
	return evictedValue
}

func (lc *lruCache) get(key string) (value interface{}, found bool) {
	element, ok := lc.elementMap[key]
	if !ok {
		return nil, false
	}

	entry := lc.unwrap(element)
	value = entry.value

	lc.elementList.MoveToFront(element)
	return value, true
}

func (lc *lruCache) set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	element, ok := lc.elementMap[key]
	if ok {
		entry := lc.unwrap(element)
		entry.setup(key, value, ttl)

		lc.elementList.MoveToFront(element)
		return nil
	}

	if lc.elementList.Len() >= lc.maxEntries {
		evictedValue = lc.evict()
	}

	element = lc.elementList.PushFront(newEntry(key, value, ttl))
	lc.elementMap[key] = element

	return evictedValue
}

func (lc *lruCache) removeElement(element *list.Element) (removedValue interface{}) {
	entry := lc.unwrap(element)
	removedValue = lc.elementList.Remove(element)

	delete(lc.elementMap, entry.key)
	return removedValue
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
	scans := 0

	element := lc.elementList.Back()
	for element != nil {
		scans++

		if entry := lc.unwrap(element); entry.expired() {
			delete(lc.elementMap, entry.key)
			cleans++
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
	lc.loader.Reset()
}

func (lc *lruCache) Get(key string) (value interface{}, found bool) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.get(key)
}

func (lc *lruCache) Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.set(key, value, ttl)
}

func (lc *lruCache) Remove(key string) (removedValue interface{}) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.remove(key)
}

func (lc *lruCache) Size() (size int) {
	lc.lock.RLock()
	defer lc.lock.RUnlock()

	return lc.size()
}

func (lc *lruCache) GC() (cleans int) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	return lc.gc()
}

func (lc *lruCache) Reset() {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	lc.reset()
}

func (lc *lruCache) Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error) {
	return lc.loader.Load(key, ttl, load)
}
