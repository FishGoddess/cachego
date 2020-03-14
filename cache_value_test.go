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
// Created at 2020/03/14 15:51:02

package cache

import (
    "fmt"
    "testing"
)

// 测试 cacheValue 的基础功能是否正常
func TestNewCacheValue(t *testing.T) {
    value := NewCacheValue(123, true, DefaultLife)
    fmt.Println(value.Item())
    fmt.Println(value.Valid())
    fmt.Println(value.Value())
    fmt.Println(value.Life())

    // 自动测试
    if value.Item() != interface{}(123) {
        t.Fatal("Item() is wrong!")
    }
    if value.Valid() != true {
        t.Fatal("Valid() is wrong!")
    }
    if v, ok := value.Value(); v != interface{}(123) || !ok {
        t.Fatal("Value() is wrong!")
    }
    if value.Life() != DefaultLife {
        t.Fatal("Life() is wrong!")
    }
}
