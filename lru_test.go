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
	"strconv"
	"testing"
	"time"
)

func newTestLRUCache() *lruCache {
	conf := newDefaultConfig()
	conf.maxEntries = maxTestEntries
	return newLRUCache(conf).(*lruCache)
}

// go test -v -cover -run=^TestLRUCacheGet$
func TestLRUCacheGet(t *testing.T) {
	cache := newTestLRUCache()
	testCacheGet(t, cache)
}

// go test -v -cover -run=^TestLRUCacheSet$
func TestLRUCacheSet(t *testing.T) {
	cache := newTestLRUCache()
	testCacheSet(t, cache)
}

// go test -v -cover -run=^TestLRUCacheRemove$
func TestLRUCacheRemove(t *testing.T) {
	cache := newTestLRUCache()
	testCacheRemove(t, cache)
}

// go test -v -cover -run=^TestLRUCacheSize$
func TestLRUCacheSize(t *testing.T) {
	cache := newTestLRUCache()
	testCacheSize(t, cache)
}

// go test -v -cover -run=^TestLRUCacheGC$
func TestLRUCacheGC(t *testing.T) {
	cache := newTestLRUCache()
	testCacheGC(t, cache)
}

// go test -v -cover -run=^TestLRUCacheReset$
func TestLRUCacheReset(t *testing.T) {
	cache := newTestLRUCache()
	testCacheReset(t, cache)
}

// go test -v -cover -run=^TestLRUCacheEvict$
func TestLRUCacheEvict(t *testing.T) {
	cache := newTestLRUCache()

	for i := 0; i < cache.maxEntries*10; i++ {
		data := strconv.Itoa(i)
		evictedValue := cache.Set(data, data, time.Duration(i)*time.Second)

		if i >= cache.maxEntries && evictedValue == nil {
			t.Errorf("i %d >= cache.maxEntries %d && evictedValue == nil", i, cache.maxEntries)
		}
	}

	if cache.Size() != cache.maxEntries {
		t.Errorf("cache.Size() %d != cache.maxEntries %d", cache.Size(), cache.maxEntries)
	}

	for i := cache.maxEntries*10 - cache.maxEntries; i < cache.maxEntries*10; i++ {
		data := strconv.Itoa(i)
		value, ok := cache.Get(data)
		if !ok || value.(string) != data {
			t.Errorf("!ok %+v || value.(string) %s != data %s", !ok, value.(string), data)
		}
	}

	i := cache.maxEntries*10 - cache.maxEntries
	element := cache.elementList.Back()
	for element != nil {
		entry := element.Value.(*entry)
		data := strconv.Itoa(i)

		if entry.key != data || entry.value.(string) != data {
			t.Errorf("entry.key %s != data %s || entry.value.(string) %s != data %s", entry.key, data, entry.value.(string), data)
		}

		element = element.Prev()
		i++
	}
}
