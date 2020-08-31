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
// Email: fishgoddess@qq.com
// Created at 2020/09/01 00:00:00

package main

import (
	"math/rand"
	//"runtime/debug"
	"strconv"
	"sync"
	"testing"
	//"time"

	cache "github.com/FishGoddess/cachego"
	//"github.com/coocood/freecache"
	//gocache "github.com/patrickmn/go-cache"
)

//--- PASS: TestCacheGoWrite (1.64s)
//--- PASS: TestCacheGoRead (0.98s)
//--- PASS: TestCacheGo (2.52s)
//--- PASS: TestGoCacheWrite (1.12s)
//--- PASS: TestGoCacheRead (1.00s)
//--- PASS: TestGoCache (1.94s)
//--- PASS: TestFreeCacheWrite (1.03s)
//--- PASS: TestFreeCacheRead (0.76s)
//--- PASS: TestFreeCache (0.73s)

const (
	dataSize = 100_0000

	loop        = 100
	concurrency = 10000
)

// 测试任务
func testTask(task func()) {

	wg := &sync.WaitGroup{}
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			task()
		}()
	}
	wg.Wait()
}

// 测试 cachego 写入的性能
func TestCacheGoWrite(t *testing.T) {

	c := cache.NewCache()
	for i := 0; i < dataSize; i++ {
		key := strconv.Itoa(i)
		c.Put(key, key, cache.NeverDie)
	}

	for i := 0; i < loop; i++ {
		testTask(func() {
			key := strconv.Itoa(rand.Intn(dataSize))
			c.Put(key, key, cache.NeverDie)
		})
	}
}

// 测试 cachego 读取的性能
func TestCacheGoRead(t *testing.T) {

	c := cache.NewCache()
	for i := 0; i < dataSize; i++ {
		key := strconv.Itoa(i)
		c.Put(key, key, cache.NeverDie)
	}

	for i := 0; i < loop; i++ {
		testTask(func() {
			key := strconv.Itoa(rand.Intn(dataSize))
			c.Of(key)
		})
	}
}

// 测试 cachego 的性能
func TestCacheGo(t *testing.T) {

	c := cache.NewCache()
	for i := 0; i < dataSize; i++ {
		key := strconv.Itoa(i)
		c.Put(key, key, cache.NeverDie)
	}

	for i := 0; i < loop; i++ {
		testTask(func() {
			key := strconv.Itoa(rand.Intn(dataSize))
			c.Put(key, key, cache.NeverDie)
			c.Of(key)
		})
	}
}

//// 测试 go-cache 写入的性能
//func TestGoCacheWrite(t *testing.T) {
//
//	c := gocache.New(gocache.NoExpiration, 10*time.Minute)
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set(key, key, gocache.NoExpiration)
//	}
//
//	for i := 0; i < loop; i++ {
//		testTask(func() {
//			key := strconv.Itoa(rand.Intn(dataSize))
//			c.Set(key, key, gocache.NoExpiration)
//		})
//	}
//}
//
//// 测试 go-cache 读取的性能
//func TestGoCacheRead(t *testing.T) {
//
//	c := gocache.New(gocache.NoExpiration, 10*time.Minute)
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set(key, key, gocache.NoExpiration)
//	}
//
//	for i := 0; i < loop; i++ {
//		testTask(func() {
//			key := strconv.Itoa(rand.Intn(dataSize))
//			c.Get(key)
//		})
//	}
//}
//
//// 测试 go-cache 的性能
//func TestGoCache(t *testing.T) {
//
//	c := gocache.New(gocache.NoExpiration, 10*time.Minute)
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set(key, key, gocache.NoExpiration)
//	}
//
//	for i := 0; i < loop; i++ {
//		testTask(func() {
//			key := strconv.Itoa(rand.Intn(dataSize))
//			c.Set(key, key, gocache.NoExpiration)
//			c.Get(key)
//		})
//	}
//}
//
//// 测试 freecache 写入的性能
//func TestFreeCacheWrite(t *testing.T) {
//
//	fcache := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		fcache.Set([]byte(key), []byte(key), 0)
//	}
//
//	for i := 0; i < loop; i++ {
//		testTask(func() {
//			key := strconv.Itoa(rand.Intn(dataSize))
//			fcache.Set([]byte(key), []byte(key), 0)
//		})
//	}
//}
//
//// 测试 freecache 读取的性能
//func TestFreeCacheRead(t *testing.T) {
//
//	fcache := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		fcache.Set([]byte(key), []byte(key), 0)
//	}
//
//	for i := 0; i < loop; i++ {
//		testTask(func() {
//			key := strconv.Itoa(rand.Intn(dataSize))
//			fcache.Get([]byte(key))
//		})
//	}
//}
//
//// 测试 freecache 的性能
//func TestFreeCache(t *testing.T) {
//
//	fcache := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		fcache.Set([]byte(key), []byte(key), 0)
//	}
//
//	for i := 0; i < loop; i++ {
//		testTask(func() {
//			key := strconv.Itoa(rand.Intn(dataSize))
//			fcache.Set([]byte(key), []byte(key), 0)
//			fcache.Get([]byte(key))
//		})
//	}
//}
