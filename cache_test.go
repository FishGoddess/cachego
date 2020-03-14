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
// Created at 2020/03/14 22:58:02

package cache

import (
    "fmt"
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
