// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/03/14 22:58:02

package cachego

import (
	"testing"
	"time"
)

// go test -cover -run=^TestCache$
func TestCache(t *testing.T) {

	cache := NewCache()

	key := "key"
	value := 123
	cache.Put(key, value)
	if v, ok := cache.Of(key); !ok || v != value {
		t.Fatal("Before reset cache, cache.Of(key) returns wrong ok or value!")
	}

	cache.Reset()
	if _, ok := cache.Of(key); ok || cache.Size() != 0 {
		t.Fatal("Cache should be reset!")
	}

	cache.Put(key, value)
	if v, ok := cache.Of(key); !ok || v != value {
		t.Fatal("Before delete key, cache.Of(key) returns wrong ok or value!")
	}

	cache.Delete(key)
	if _, ok := cache.Of(key); ok {
		t.Fatal("After deleting key, key should be dead!")
	}
}

// go test -cover -run=^TestCacheTTL$
func TestCacheTTL(t *testing.T) {

	cache := NewCache()
	cache.AutoGc(3 * time.Second)

	key := "key"
	value := 123
	cache.PutWithTTL(key, value, 1)
	if v, ok := cache.Of(key); !ok || cache.Size() != 1 || v != value {
		t.Fatal("Before ttl, returns wrong ok or size or value!")
	}

	time.Sleep(2 * time.Second)
	if _, ok := cache.Of(key); ok {
		t.Fatal("After ttl, key should be dead!")
	}

	if cache.Size() != 1 {
		t.Fatal("After ttl, size should be 1!")
	}

	time.Sleep(2 * time.Second)
	if cache.Size() != 0 {
		t.Fatal("After gc, size should be 0!")
	}
}

// go test -cover -run=^TestCacheAutoGc$
func TestCacheAutoGc(t *testing.T) {

	cache := NewCache()
	cache.AutoGc(2 * time.Second) <- true

	key := "key"
	value := 123
	cache.PutWithTTL(key, value, 1)
	if v, ok := cache.Of(key); !ok || cache.Size() != 1 || v != value {
		t.Fatal("Before gc, returns wrong ok or size or value!")
	}

	time.Sleep(3 * time.Second)
	if _, ok := cache.Of(key); ok || cache.Size() != 1 {
		t.Fatal("After gc, key should be dead!")
	}
}
