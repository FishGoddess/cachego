// Copyright 2020 FishGoddess. All Rights Reserved.
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
	"context"
	"testing"
	"time"
)

// go test -cover -run=^TestCache$
func TestCache(t *testing.T) {
	cache := NewCache()

	key := "key"
	value := 123
	cache.Set(key, value)
	if v, err := cache.Get(key); IsNotFound(err) || v != value {
		t.Error("Before reset cache, cache.Of(key) returns wrong err or value!")
	}

	cache.DeleteAll()
	if _, err := cache.Get(key); !IsNotFound(err) || cache.Size() != 0 {
		t.Error("Cache should be reset!")
	}

	cache.Set(key, value)
	if v, err := cache.Get(key); IsNotFound(err) || v != value {
		t.Error("Before delete key, cache.Of(key) returns wrong err or value!")
	}

	cache.Delete(key)
	if _, err := cache.Get(key); !IsNotFound(err) {
		t.Error("After deleting key, key should be dead!")
	}
}

// go test -cover -run=^TestCacheTTL$
func TestCacheTTL(t *testing.T) {
	cache := NewCache()
	cache.AutoGC(3 * time.Second)

	key := "key"
	value := 123
	cache.Set(key, value, WithOpTTL(1*time.Second))
	if v, err := cache.Get(key); IsNotFound(err) || cache.Size() != 1 || v != value {
		t.Error("Before ttl, returns wrong err or size or value!")
	}

	time.Sleep(2 * time.Second)
	if _, err := cache.Get(key); !IsNotFound(err) {
		t.Error("After ttl, key should be dead!")
	}

	if cache.Size() != 1 {
		t.Error("After ttl, size should be 1!")
	}

	time.Sleep(2 * time.Second)
	if cache.Size() != 0 {
		t.Error("After gc, size should be 0!")
	}
}

// go test -cover -run=^TestGetWithLoad$
func TestGetWithLoad(t *testing.T) {
	cache := NewCache()

	loadFunc := func(ctx context.Context) (data interface{}, err error) {
		return "loadFunc", nil
	}

	key := "key"
	value := "get"
	cache.Set(key, value, WithOpTTL(1*time.Second))
	if v, err := cache.Get(key, WithOpOnMissed(loadFunc), WithOpTTL(time.Second)); err != nil || v.(string) != value {
		t.Errorf("Before Sleep, cache.Of(key) returns err %+v or wrong value %s!", err, v.(string))
	}

	time.Sleep(2 * time.Second)
	if v, err := cache.Get(key, WithOpOnMissed(loadFunc), WithOpTTL(time.Second)); err != nil || v.(string) != "loadFunc" {
		t.Errorf("After Sleep, cache.Of(key) returns err %+v or wrong value %s!", err, v.(string))
	}
}

// go test -cover -run=^TestCacheAutoGC$
func TestCacheAutoGC(t *testing.T) {
	cache := NewCache()
	cache.AutoGC(2 * time.Second) <- struct{}{}

	key := "key"
	value := 123
	cache.Set(key, value, WithOpTTL(1*time.Second))
	if v, err := cache.Get(key); IsNotFound(err) || cache.Size() != 1 || v != value {
		t.Error("Before gc, returns wrong err or size or value!")
	}

	time.Sleep(3 * time.Second)
	if _, err := cache.Get(key); !IsNotFound(err) || cache.Size() != 1 {
		t.Error("After gc, key should be dead!")
	}
}
