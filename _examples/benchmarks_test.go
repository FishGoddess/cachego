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
// Created at 2020/03/14 23:43:52

package _examples

import (
    "fmt"
    "strconv"
    "testing"
    "time"

    "github.com/FishGoddess/cachego"
)

// 测试 Cachego 的性能
func BenchmarkCachego(b *testing.B) {

    c := cache.NewCacheWithGcDuration(5 * time.Second)
    go func() {
        for i := 0; i < 1000000; i++ {
            key := strconv.Itoa(i)
            c.Put(key, i, time.Duration(i*100)*time.Microsecond)
        }
    }()

    testTask := func(i int) {
        c.Of("100000").Int()
    }

    b.ReportAllocs()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        testTask(i)
    }
    fmt.Println(c.Extend().Size())
}
