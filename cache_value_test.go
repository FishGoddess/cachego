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
// Email: fishinlove@163.com
// Created at 2020/03/14 15:51:02

package cache

import (
	"fmt"
	"testing"
	"time"
)

// 测试 cacheValue 的基础功能是否正常
func TestNewCacheValue(t *testing.T) {
	value := NewCacheValue(123, 30*time.Second)
	fmt.Println(value.Ok())
	fmt.Println(value.Value())
	fmt.Println(value.Life())

	// 自动测试
	if v, ok := value.Value(); v != interface{}(123) || !ok {
		t.Fatal("Value() is wrong!")
	}
	time.Sleep(1500 * time.Millisecond)
	if value.Life() < 28*time.Second && value.Life() > 29*time.Second {
		t.Fatal("Life() is wrong!")
	}
}

// 测试 cacheValue 的 Ok 方法
func TestCacheValueOk(t *testing.T) {
	value := NewCacheValue(123, 30*time.Second)
	if !value.Ok() {
		t.Fatal("Ok() is wrong...")
	}
	value = InvalidCacheValue()
	if value.Ok() {
		t.Fatal("Ok() is wrong...")
	}
}

// 测试 cacheValue 的 Value 方法
func TestCacheValueValue(t *testing.T) {
	value := NewCacheValue(123, 30*time.Second)
	v, ok := value.Value()
	if v != interface{}(123) || !ok {
		t.Fatal("Value() is wrong...")
	}
	value = InvalidCacheValue()
	v, ok = value.Value()
	if v != interface{}(nil) || ok {
		t.Fatal("Value() is wrong...")
	}
}

// 测试 cacheValue 的 Or 方法
func TestCacheValueOr(t *testing.T) {
	value := NewCacheValue(123, 30*time.Second)
	if value.Or(456).Int() != interface{}(123) {
		t.Fatal("Or() is wrong...")
	}
	value = InvalidCacheValue()
	if value.Or(456).Int() != interface{}(456) {
		t.Fatal("Or() is wrong...")
	}
}

// 测试 cacheValue 的 Life 方法
func TestCacheValueLife(t *testing.T) {
	value := NewCacheValue(123, 30*time.Second)
	time.Sleep(1500 * time.Millisecond)
	if value.Life() < 28*time.Second && value.Life() > 29*time.Second {
		t.Fatal("Life() is wrong!")
	}
}
