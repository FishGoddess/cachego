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
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 测试标准缓存结构是否可用
func TestNewCache(t *testing.T) {
	cache := NewCache()
	cache.Put("fish", 123, 3*time.Second)
	for i := 0; i < 3; i++ {
		if !cache.Of("fish").Ok() {
			t.Fatal("Ok() is wrong...")
		}
		value, ok := cache.Of("fish").Value()
		if value != interface{}(123) || !ok {
			t.Fatal("Value() is wrong...")
		}
		fmt.Println(cache.Of("fish").Life())
		if cache.Of("fish").Int() != 123 {
			t.Fatal("Value() is wrong...")
		}
		tryInt, ok := cache.Of("fish").TryInt()
		if tryInt != 123 || !ok {
			t.Fatal("TryInt() is wrong...")
		}
		tryString, ok := cache.Of("fish").TryString()
		if tryString != "" || ok {
			t.Fatal("TryString() is wrong...")
		}

		//fmt.Println(*cache.Of("fish"))
		time.Sleep(time.Second)
	}
	time.Sleep(1 * time.Second)
	if cache.Of("fish").Ok() {
		t.Fatal("Ok() is wrong...")
	}
}

// 测试标准缓存的 Of 方法
func TestStandardCacheOf(t *testing.T) {
	cache := NewCache()
	cache.Put("key", 123, 5*time.Second)
	value := cache.Of("key")
	if value.Int() != 123 || !value.Ok() {
		t.Fatal("Of() is wrong...")
	}

	value = cache.Of("wrong key")
	if value != InvalidCacheValue() || value.Ok() {
		t.Fatal("Of() is wrong...")
	}

	if value.Or("456").String() != "456" {
		t.Fatal("Of() is wrong...")
	}
}

// 测试标准缓存的 Put 方法
func TestStandardCachePut(t *testing.T) {
	cache := NewCache()
	for i := 0; i < 10; i++ {
		if cache.Extend().Size() != i {
			t.Fatal("Put() is wrong...")
		}
		cache.Put("key"+strconv.Itoa(i), 123, 5*time.Second)
	}
}

// 测试标准缓存的 Change 方法
func TestStandardCacheChange(t *testing.T) {
	cache := NewCache()
	cache.Put("key", 123, 5*time.Second)
	value := cache.Of("key")
	if value.Int() != 123 || !value.Ok() {
		t.Fatal("Of() is wrong...")
	}

	cache.Change("key", 456)
	value = cache.Of("key")
	if value.Int() != 456 || !value.Ok() {
		t.Fatal("Change() is wrong...")
	}
}

// 测试标准缓存的 Remove 方法
func TestStandardCacheRemove(t *testing.T) {
	cache := NewCache()
	cache.Put("key", 123, 5*time.Second)
	value := cache.Of("key")
	if value.Int() != 123 || !value.Ok() {
		t.Fatal("Of() is wrong...")
	}

	cache.Remove("key")
	value = cache.Of("key")
	if value != InvalidCacheValue() || value.Ok() {
		t.Fatal("Remove() is wrong...")
	}
}

// 测试标准缓存的 RemoveAll 方法
func TestStandardCacheRemoveAll(t *testing.T) {
	cache := NewCache()
	cache.Put("key", 123, 50*time.Second)
	if cache.Extend().Size() != 1 {
		t.Fatal("Extend().Size() is wrong...")
	}
	for i := 0; i < 10; i++ {
		cache.Put("key"+strconv.Itoa(i), 123, 50*time.Second)
	}
	if cache.Extend().Size() != 11 {
		t.Fatal("Extend().Size() is wrong...")
	}

	value := cache.Of("key")
	if value.Int() != 123 || !value.Ok() {
		t.Fatal("Of() is wrong...")
	}
	cache.RemoveAll()
	if cache.Extend().Size() != 0 {
		t.Fatal("Extend().Size() is wrong...")
	}
	value = cache.Of("key")
	if value != InvalidCacheValue() || value.Ok() {
		t.Fatal("RemoveAll() is wrong...")
	}
	for i := 0; i < 10; i++ {
		value = cache.Of("key" + strconv.Itoa(i))
		if value != InvalidCacheValue() || value.Ok() {
			t.Fatal("RemoveAll() is wrong...")
		}
	}
}

// 测试标准缓存的 Gc 方法
func TestStandardCacheGc(t *testing.T) {
	cache := NewCacheWithGcDuration(5 * time.Second)
	for i := 1; i <= 20; i++ {
		cache.Put("key"+strconv.Itoa(i), i, time.Duration(i)*time.Second)
	}
	value := cache.Of("key3")
	if value.Int() != 3 || !value.Ok() {
		t.Fatal("Of() is wrong...")
	}
	value = cache.Of("key7")
	if value.Int() != 7 || !value.Ok() {
		t.Fatal("Of() is wrong...")
	}
	if cache.Extend().Size() != 20 {
		t.Fatal("Extend().Size() is wrong...")
	}
	time.Sleep(5500 * time.Millisecond)
	value = cache.Of("key3")
	if value != InvalidCacheValue() || value.Ok() {
		t.Fatal("Gc() is wrong...")
	}
	value = cache.Of("key7")
	if value.Int() != 7 || !value.Ok() {
		t.Fatal("Gc() is wrong...")
	}
	if cache.Extend().Size() != 15 {
		t.Fatal("Gc() is wrong...")
	}
	time.Sleep(5 * time.Second)
	value = cache.Of("key7")
	if value != InvalidCacheValue() || value.Ok() {
		t.Fatal("Gc() is wrong...")
	}
	fmt.Println(cache.Extend().Size())
	if cache.Extend().Size() != 10 {
		t.Fatal("Gc() is wrong...")
	}
}
