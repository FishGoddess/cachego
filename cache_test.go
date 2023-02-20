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
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxTestEntries = 10
)

type testCache struct {
	cache
	count int32
}

func (tc *testCache) currentCount() int32 {
	return atomic.LoadInt32(&tc.count)
}

func (tc *testCache) Get(key string) (value interface{}, found bool) {
	return nil, false
}

func (tc *testCache) Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	return nil
}

func (tc *testCache) Remove(key string) (removedValue interface{}) {
	return nil
}

func (tc *testCache) Size() (size int) {
	return 0
}

func (tc *testCache) GC() (cleans int) {
	atomic.AddInt32(&tc.count, 1)
	return 0
}

func (tc *testCache) Reset() {}

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

func testCacheImplement(t *testing.T, cache Cache) {
	testCaches := []func(t *testing.T, cache Cache){
		testCacheGet, testCacheSet, testCacheRemove, testCacheSize, testCacheGC, testCacheReset,
	}

	for _, testCache := range testCaches {
		cache.Reset()
		testCache(t, cache)
	}
}

// go test -v -cover=^TestNewCache$
func TestNewCache(t *testing.T) {
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

// go test -v -cover=^TestNewCacheWithReport$
func TestNewCacheWithReport(t *testing.T) {
	cache, reporter := NewCacheWithReport()

	sc1, ok := cache.(*reportableCache)
	if !ok {
		t.Errorf("cache.(*reportableCache) %T not ok", cache)
	}

	if sc1 == nil {
		t.Error("sc1 == nil")
	}

	if reporter == nil {
		t.Error("reporter == nil")
	}
}

// go test -v -cover=^TestRunGCTask$
func TestRunGCTask(t *testing.T) {
	cache := new(testCache)

	count := cache.currentCount()
	if count != 0 {
		t.Errorf("cache.currentCount() %d is wrong", count)
	}

	cancel := RunGCTask(cache, 10*time.Millisecond)

	time.Sleep(105 * time.Millisecond)

	count = cache.currentCount()
	if count != 10 {
		t.Errorf("cache.currentCount() %d is wrong", count)
	}

	time.Sleep(80 * time.Millisecond)
	cancel()

	count = cache.currentCount()
	if count != 18 {
		t.Errorf("cache.currentCount() %d is wrong", count)
	}

	time.Sleep(time.Second)

	count = cache.currentCount()
	if count != 18 {
		t.Errorf("cache.currentCount() %d is wrong", count)
	}
}
