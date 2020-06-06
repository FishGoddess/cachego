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
// Created at 2020/03/14 23:43:52

package main

import (
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
	//fmt.Println(c.Extend().Size())
}

//// 测试 go-cache 性能
//func BenchmarkGoCache(b *testing.B) {
//    c := gocache.New(gocache.DefaultExpiration, 5*time.Second)
//    go func() {
//        for i := 0; i < 1000000; i++ {
//            key := strconv.Itoa(i)
//            c.Set(key, key, time.Duration(i*100)*time.Microsecond)
//        }
//    }()
//
//    testTask := func() string {
//        v, ok := c.Get("100000")
//        if !ok {
//            return ""
//        }
//        value := v.(string)
//        return value
//    }
//
//    b.ReportAllocs()
//    b.StartTimer()
//    for i := 0; i < b.N; i++ {
//        testTask()
//    }
//    fmt.Println(c.ItemCount())
//}
//
//// 测试 freecache 性能
//func BenchmarkFreeCache(b *testing.B) {
//    cacheSize := 100 * 1024 * 1024
//    fcache := freecache.NewCache(cacheSize)
//    debug.SetGCPercent(20)
//
//    go func() {
//        for i := 0; i < 1000000; i++ {
//            key := strconv.Itoa(i)
//            // Expired time is less than cachego and go-cache. So this is not fair to the other guys.
//            fcache.Set([]byte(key), []byte(key), i)
//        }
//    }()
//
//    testTask := func() string {
//        v, err := fcache.Get([]byte("100000"))
//        if err != nil {
//            return ""
//        }
//        return string(v)
//    }
//
//    b.ReportAllocs()
//    b.StartTimer()
//    for i := 0; i < b.N; i++ {
//        testTask()
//    }
//    fmt.Println(fcache.EntryCount())
//}
