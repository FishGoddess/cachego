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
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func newTestLRUCache() *lruCache {
	conf := newDefaultConfig()
	conf.maxEntries = maxTestEntries
	return newLRUCache(conf).(*lruCache)
}

// go test -v -cover -run=^TestLRUCache$
func TestLRUCache(t *testing.T) {
	cache := newTestLRUCache()
	testCacheImplement(t, cache)
}

// go test -v -cover -run=^TestLRUCacheEvict$
func TestLRUCacheEvict(t *testing.T) {
	cache := newTestLRUCache()

	for i := 0; i < cache.maxEntries*10; i++ {
		data := strconv.Itoa(i)
		evictedValue := cache.Set(data, data, time.Duration(i)*time.Second)

		if i >= cache.maxEntries && evictedValue == nil {
			t.Fatalf("i %d >= cache.maxEntries %d && evictedValue == nil", i, cache.maxEntries)
		}
	}

	if cache.Size() != cache.maxEntries {
		t.Fatalf("cache.Size() %d != cache.maxEntries %d", cache.Size(), cache.maxEntries)
	}

	for i := cache.maxEntries*10 - cache.maxEntries; i < cache.maxEntries*10; i++ {
		data := strconv.Itoa(i)
		value, ok := cache.Get(data)
		if !ok || value.(string) != data {
			t.Fatalf("!ok %+v || value.(string) %s != data %s", !ok, value.(string), data)
		}
	}

	i := cache.maxEntries*10 - cache.maxEntries
	element := cache.elementList.Back()
	for element != nil {
		entry := element.Value.(*entry)
		data := strconv.Itoa(i)

		if entry.key != data || entry.value.(string) != data {
			t.Fatalf("entry.key %s != data %s || entry.value.(string) %s != data %s", entry.key, data, entry.value.(string), data)
		}

		element = element.Prev()
		i++
	}
}

// go test -v -cover -run=^TestLRUCacheEvictSimulate$
func TestLRUCacheEvictSimulate(t *testing.T) {
	cache := newTestLRUCache()

	for i := 0; i < maxTestEntries; i++ {
		data := strconv.Itoa(i)
		cache.Set(data, data, NoTTL)
	}

	maxKeys := 10000
	keys := make([]string, 0, maxKeys)
	random := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < maxKeys; i++ {
		key := strconv.Itoa(random.Intn(maxTestEntries))
		keys = append(keys, key)
	}

	for _, key := range keys {
		cache.Get(key)
	}

	expectKeys := make([]string, maxTestEntries)
	index := len(expectKeys) - 1
	for i := len(keys) - 1; i > 0; i-- {
		exist := false

		for _, expectKey := range expectKeys {
			if keys[i] == expectKey {
				exist = true
			}
		}

		if !exist {
			expectKeys[index] = keys[i]
			index--
		}
	}

	t.Log(expectKeys)

	var got strings.Builder
	element := cache.elementList.Back()
	for element != nil {
		got.WriteString(element.Value.(*entry).key)
		element = element.Prev()
	}

	expect := strings.Join(expectKeys, "")
	if strings.Compare(got.String(), expect) != 0 {
		t.Fatalf("got %s != expect %s", got.String(), expect)
	}

	for i := 0; i < maxTestEntries; i++ {
		data := strconv.Itoa(maxTestEntries*10 + i)
		evictedValue := cache.Set(data, data, NoTTL)

		if evictedValue.(string) != expectKeys[i] {
			t.Fatalf("evictedValue.(string) %s != expectKeys[i] %s", evictedValue.(string), expectKeys[i])
		}
	}
}
