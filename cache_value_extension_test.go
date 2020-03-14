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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/14 16:06:00

package cache

import (
    "testing"
    "time"
)

// 测试 cacheValue 的 String 和 TryString 方法
func TestCacheValueString(t *testing.T) {
    value := NewCacheValue("str", 30*time.Second)
    if value.String() != "str" {
        t.Fatal("String() is wrong!")
    }
    s, ok := value.TryString()
    if s != "str" || !ok {
        t.Fatal("TryString() is wrong!")
    }

    value = NewCacheValue(123, 30*time.Second)
    s, ok = value.TryString()
    if ok {
        t.Fatal("TryString() is wrong!")
    }
}

// 测试 cacheValue 的 Int 和 TryInt 方法
func TestCacheValueInt(t *testing.T) {
    value := NewCacheValue(123, 30*time.Second)
    if value.Int() != 123 {
        t.Fatal("Int() is wrong!")
    }
    i, ok := value.TryInt()
    if i != 123 || !ok {
        t.Fatal("TryInt() is wrong!")
    }

    value = NewCacheValue("str", 30*time.Second)
    i, ok = value.TryInt()
    if ok {
        t.Fatal("TryInt() is wrong!")
    }
}

// 测试 cacheValue 的 Int8 和 TryInt8 方法
func TestCacheValueInt8(t *testing.T) {
    value := NewCacheValue(int8(123), 30*time.Second)
    if value.Int8() != int8(123) {
        t.Fatal("Int8() is wrong!")
    }
    i, ok := value.TryInt8()
    if i != int8(123) || !ok {
        t.Fatal("TryInt8() is wrong!")
    }

    value = NewCacheValue("str", 30*time.Second)
    i, ok = value.TryInt8()
    if ok {
        t.Fatal("TryInt8() is wrong!")
    }
}

// 测试 cacheValue 的 Int16 和 TryInt16 方法
func TestCacheValueInt16(t *testing.T) {
    value := NewCacheValue(int16(123), 30*time.Second)
    if value.Int16() != int16(123) {
        t.Fatal("Int16() is wrong!")
    }
    i, ok := value.TryInt16()
    if i != int16(123) || !ok {
        t.Fatal("TryInt16() is wrong!")
    }

    value = NewCacheValue("str", 30*time.Second)
    i, ok = value.TryInt16()
    if ok {
        t.Fatal("TryInt16() is wrong!")
    }
}

// 测试 cacheValue 的 Int32 和 TryInt32 方法
func TestCacheValueInt32(t *testing.T) {
    value := NewCacheValue(int32(123), 30*time.Second)
    if value.Int32() != int32(123) {
        t.Fatal("Int32() is wrong!")
    }
    i, ok := value.TryInt32()
    if i != int32(123) || !ok {
        t.Fatal("TryInt32() is wrong!")
    }

    value = NewCacheValue("str", 30*time.Second)
    i, ok = value.TryInt32()
    if ok {
        t.Fatal("TryInt32() is wrong!")
    }
}

// 测试 cacheValue 的 Int64 和 TryInt64 方法
func TestCacheValueInt64(t *testing.T) {
    value := NewCacheValue(int64(123), 30*time.Second)
    if value.Int64() != int64(123) {
        t.Fatal("Int64() is wrong!")
    }
    i, ok := value.TryInt64()
    if i != int64(123) || !ok {
        t.Fatal("TryInt64() is wrong!")
    }

    value = NewCacheValue("str", 30*time.Second)
    i, ok = value.TryInt64()
    if ok {
        t.Fatal("TryInt64() is wrong!")
    }
}

// 测试 cacheValue 的 Float32 和 TryFloat32 方法
func TestCacheValueFloat32(t *testing.T) {
    value := NewCacheValue(float32(123.123), 30*time.Second)
    if value.Float32() != float32(123.123) {
        t.Fatal("Float32() is wrong!")
    }
    i, ok := value.TryFloat32()
    if i != float32(123.123) || !ok {
        t.Fatal("TryFloat32() is wrong!")
    }

    value = NewCacheValue("str", 30*time.Second)
    i, ok = value.TryFloat32()
    if ok {
        t.Fatal("TryFloat32() is wrong!")
    }
}

// 测试 cacheValue 的 Float64 和 TryFloat64 方法
func TestCacheValueFloat64(t *testing.T) {
    value := NewCacheValue(123.123, 30*time.Second)
    if value.Float64() != 123.123 {
        t.Fatal("Float64() is wrong!")
    }
    i, ok := value.TryFloat64()
    if i != 123.123 || !ok {
        t.Fatal("TryFloat64() is wrong!")
    }

    value = NewCacheValue("str", 30*time.Second)
    i, ok = value.TryFloat64()
    if ok {
        t.Fatal("TryFloat64() is wrong!")
    }
}
