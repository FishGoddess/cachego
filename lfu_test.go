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

func newTestLFUCache() *lfuCache {
	conf := newDefaultConfig()
	conf.maxEntries = maxTestEntries
	return newLFUCache(conf).(*lfuCache)
}

// go test -v -cover -run=^TestLFUCacheGet$
func TestLFUCacheGet(t *testing.T) {
	cache := newTestLFUCache()
	testCacheGet(t, cache)
}

// go test -v -cover -run=^TestLFUCacheSet$
func TestLFUCacheSet(t *testing.T) {
	cache := newTestLFUCache()
	testCacheSet(t, cache)
}

// go test -v -cover -run=^TestLFUCacheRemove$
func TestLFUCacheRemove(t *testing.T) {
	cache := newTestLFUCache()
	testCacheRemove(t, cache)
}

// go test -v -cover -run=^TestLFUCacheSize$
func TestLFUCacheSize(t *testing.T) {
	cache := newTestLFUCache()
	testCacheSize(t, cache)
}

// go test -v -cover -run=^TestLFUCacheGC$
func TestLFUCacheGC(t *testing.T) {
	cache := newTestLFUCache()
	testCacheGC(t, cache)
}

// go test -v -cover -run=^TestLFUCacheReset$
func TestLFUCacheReset(t *testing.T) {
	cache := newTestLFUCache()
	testCacheReset(t, cache)
}

// go test -v -cover -run=^TestLFUCacheEvict$
func TestLFUCacheEvict(t *testing.T) {
	cache := newTestLFUCache()

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
		for j := 0; j < i; j++ {
			data := strconv.Itoa(i)
			cache.Set(data, data, time.Duration(i)*time.Second)
			cache.Get(data)
		}
	}

	for i := cache.maxEntries*10 - cache.maxEntries; i < cache.maxEntries*10; i++ {
		data := strconv.Itoa(i)
		value, ok := cache.Get(data)
		if !ok || value.(string) != data {
			t.Errorf("!ok %+v || value.(string) %s != data %s", !ok, value.(string), data)
		}
	}

	i := cache.maxEntries*10 - cache.maxEntries
	for cache.itemHeap.Size() > 0 {
		item := cache.itemHeap.Pop()
		entry := item.Value.(*entry)
		data := strconv.Itoa(i)

		if entry.key != data || entry.value.(string) != data {
			t.Errorf("entry.key %s != data %s || entry.value.(string) %s != data %s", entry.key, data, entry.value.(string), data)
		}

		i++
	}
}
