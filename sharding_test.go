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

func newTestShardingCache() Cache {
	conf := newDefaultConfig()
	conf.shardings = 4
	return newShardingCache(conf, newStandardCache)
}

// go test -v -cover -run=^TestShardingCacheGet$
func TestShardingCacheGet(t *testing.T) {
	cache := newTestShardingCache()

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

// go test -v -cover -run=^TestShardingCacheSet$
func TestShardingCacheSet(t *testing.T) {
	cache := newTestShardingCache()

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

// go test -v -cover -run=^TestShardingCacheRemove$
func TestShardingCacheRemove(t *testing.T) {
	cache := newTestShardingCache()

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
	if removedValue != nil {
		t.Errorf("removedValue %+v is wrong", removedValue)
	}
}

// go test -v -cover -run=^TestShardingCacheSize$
func TestShardingCacheSize(t *testing.T) {
	cache := newTestShardingCache()

	size := cache.Size()
	if size != 0 {
		t.Errorf("size %d is wrong", size)
	}

	for i := int64(0); i < 100; i++ {
		cache.Set(strconv.FormatInt(i, 10), i, NoTTL)
	}

	size = cache.Size()
	if size != 100 {
		t.Errorf("size %d is wrong", size)
	}
}

// go test -v -cover -run=^TestShardingCacheGC$
func TestShardingCacheGC(t *testing.T) {
	cache := newTestShardingCache()

	size := cache.Size()
	if size != 0 {
		t.Errorf("size %d is wrong", size)
	}

	for i := int64(0); i < 100; i++ {
		if i&1 == 0 {
			cache.Set(strconv.FormatInt(i, 10), i, NoTTL)
		} else {
			cache.Set(strconv.FormatInt(i, 10), i, time.Millisecond)
		}
	}

	size = cache.Size()
	if size != 100 {
		t.Errorf("size %d is wrong", size)
	}

	cache.GC()

	size = cache.Size()
	if size != 100 {
		t.Errorf("size %d is wrong", size)
	}

	time.Sleep(2 * time.Millisecond)

	cache.GC()

	size = cache.Size()
	if size != 50 {
		t.Errorf("size %d is wrong", size)
	}
}

// go test -v -cover -run=^TestShardingCacheReset$
func TestShardingCacheReset(t *testing.T) {
	cache := newTestShardingCache()

	for i := int64(0); i < 100; i++ {
		cache.Set(strconv.FormatInt(i, 10), i, NoTTL)
	}

	for i := int64(0); i < 100; i++ {
		value, found := cache.Get(strconv.FormatInt(i, 10))
		if !found {
			t.Errorf("get %d should be found", i)
		}

		if value.(int64) != i {
			t.Errorf("value %+v is wrong", value)
		}
	}

	size := cache.Size()
	if size != 100 {
		t.Errorf("size %d is wrong", size)
	}

	cache.Reset()

	for i := int64(0); i < 100; i++ {
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
