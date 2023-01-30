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

const (
	maxTestEntries = 100
)

func testCacheGet(t *testing.T, cache Cache) {
	value, found := cache.Get("key")
	if found {
		t.Errorf("get %+v should be not found", value)
	}

	cache.Set("key", "value", time.Millisecond)

	value, found = cache.Get("key")
	if !found {
		t.Error("get should be found")
	}

	if value.(string) != "value" {
		t.Errorf("value %+v is wrong", value)
	}

	time.Sleep(2 * time.Millisecond)

	value, found = cache.Get("key")
	if found {
		t.Errorf("get %+v should be not found", value)
	}
}

func testCacheSet(t *testing.T, cache Cache) {
	value, found := cache.Get("key")
	if found {
		t.Errorf("get %+v should be not found", value)
	}

	cache.Set("key", "value", time.Millisecond)

	value, found = cache.Get("key")
	if !found {
		t.Error("get should be found")
	}

	if value.(string) != "value" {
		t.Errorf("value %+v is wrong", value)
	}

	time.Sleep(2 * time.Millisecond)

	value, found = cache.Get("key")
	if found {
		t.Errorf("get %+v should be not found", value)
	}

	cache.Set("key", "value", NoTTL)

	value, found = cache.Get("key")
	if !found {
		t.Error("get should be found")
	}

	if value.(string) != "value" {
		t.Errorf("value %+v is wrong", value)
	}

	time.Sleep(2 * time.Millisecond)

	value, found = cache.Get("key")
	if !found {
		t.Error("get should be found")
	}
}

func testCacheRemove(t *testing.T, cache Cache) {
	removedValue := cache.Remove("key")
	if removedValue != nil {
		t.Errorf("removedValue %+v is wrong", removedValue)
	}

	cache.Set("key", "value", NoTTL)

	removedValue = cache.Remove("key")
	if removedValue.(string) != "value" {
		t.Errorf("removedValue %+v is wrong", removedValue)
	}

	cache.Set("key", "value", time.Millisecond)
	time.Sleep(2 * time.Millisecond)

	removedValue = cache.Remove("key")
	if removedValue == nil {
		t.Error("removedValue == nil")
	}
}

func testCacheSize(t *testing.T, cache Cache) {
	size := cache.Size()
	if size != 0 {
		t.Errorf("size %d is wrong", size)
	}

	for i := int64(0); i < maxTestEntries; i++ {
		cache.Set(strconv.FormatInt(i, 10), i, NoTTL)
	}

	size = cache.Size()
	if size != maxTestEntries {
		t.Errorf("size %d is wrong", size)
	}
}

func testCacheGC(t *testing.T, cache Cache) {
	size := cache.Size()
	if size != 0 {
		t.Errorf("size %d is wrong", size)
	}

	for i := int64(0); i < maxTestEntries; i++ {
		if i&1 == 0 {
			cache.Set(strconv.FormatInt(i, 10), i, NoTTL)
		} else {
			cache.Set(strconv.FormatInt(i, 10), i, time.Millisecond)
		}
	}

	size = cache.Size()
	if size != maxTestEntries {
		t.Errorf("size %d is wrong", size)
	}

	cache.GC()

	size = cache.Size()
	if size != maxTestEntries {
		t.Errorf("size %d is wrong", size)
	}

	time.Sleep(2 * time.Millisecond)

	cache.GC()

	size = cache.Size()
	if size != maxTestEntries/2 {
		t.Errorf("size %d is wrong", size)
	}
}

func testCacheReset(t *testing.T, cache Cache) {
	for i := int64(0); i < maxTestEntries; i++ {
		cache.Set(strconv.FormatInt(i, 10), i, NoTTL)
	}

	for i := int64(0); i < maxTestEntries; i++ {
		value, found := cache.Get(strconv.FormatInt(i, 10))
		if !found {
			t.Errorf("get %d should be found", i)
		}

		if value.(int64) != i {
			t.Errorf("value %+v is wrong", value)
		}
	}

	size := cache.Size()
	if size != maxTestEntries {
		t.Errorf("size %d is wrong", size)
	}

	cache.Reset()

	for i := int64(0); i < maxTestEntries; i++ {
		value, found := cache.Get(strconv.FormatInt(i, 10))
		if found {
			t.Errorf("get %d, %+v should be not found", i, value)
		}
	}

	size = cache.Size()
	if size != 0 {
		t.Errorf("size %d is wrong", size)
	}
}

// go test -v -cover=^TestNewStandardCache$
func TestNew(t *testing.T) {
	cache := NewCache()

	sc1, ok := cache.(*standardCache)
	if !ok {
		t.Errorf("cache.(*standardCache) %T not ok", cache)
	}

	if sc1 == nil {
		t.Error("sc1 == nil")
	}

	cache = NewCache(WithLRU(16))

	sc2, ok := cache.(*lruCache)
	if !ok {
		t.Errorf("cache.(*lruCache) %T not ok", cache)
	}

	if sc2 == nil {
		t.Error("sc2 == nil")
	}

	cache = NewCache(WithShardings(64))

	sc, ok := cache.(*shardingCache)
	if !ok {
		t.Errorf("cache.(*shardingCache) %T not ok", cache)
	}

	if sc == nil {
		t.Error("sc == nil")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("new should panic")
		}
	}()

	cache = NewCache(WithLRU(0))
}
