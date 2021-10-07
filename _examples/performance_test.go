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
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	//"runtime/debug"
	"time"

	"github.com/FishGoddess/cachego"
	//"github.com/coocood/freecache"
	//gocache "github.com/patrickmn/go-cache"
)

//TestCacheGoWrite   spent  965ms
//TestCacheGoRead    spent  949ms
//TestCacheGo        spent  991ms
//TestGoCacheWrite   spent 3216ms
//TestGoCacheRead    spent  980ms
//TestGoCache        spent 4508ms
//TestFreeCacheWrite spent  954ms
//TestFreeCacheRead  spent  968ms
//TestFreeCache      spent  987ms

const (
	dataSize = 100_0000

	loop        = 50
	concurrency = 100000
)

// testTask is the task of benchmark.
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

// timeTask is for recording time of task.
// unit is ms.
func timeTask(task func()) int64 {
	begin := time.Now().UnixNano()
	task()
	return (time.Now().UnixNano() - begin) / 1000_000
}

// go test -v -run=^TestCacheGoWrite$
func TestCacheGoWrite(t *testing.T) {
	c := cachego.NewCache()
	c.AutoGc(30 * time.Minute)

	for i := 0; i < dataSize; i++ {
		key := strconv.Itoa(i)
		c.Set(key, key)
	}

	spent := timeTask(func() {
		for i := 0; i < loop; i++ {
			testTask(func() {
				key := strconv.Itoa(rand.Intn(dataSize))
				c.Set(key, key)
			})
		}
	})

	fmt.Printf("%s spent %dms\n", t.Name(), spent)
}

// go test -v -run=^TestCacheGoRead$
func TestCacheGoRead(t *testing.T) {
	c := cachego.NewCache()
	c.AutoGc(30 * time.Minute)

	for i := 0; i < dataSize; i++ {
		key := strconv.Itoa(i)
		c.Set(key, key)
	}

	spent := timeTask(func() {
		for i := 0; i < loop; i++ {
			testTask(func() {
				key := strconv.Itoa(rand.Intn(dataSize))
				c.Get(key)
			})
		}
	})

	fmt.Printf("%s spent %dms\n", t.Name(), spent)
}

// go test -v -run=^TestCacheGo$
func TestCacheGo(t *testing.T) {
	c := cachego.NewCache()

	for i := 0; i < dataSize; i++ {
		key := strconv.Itoa(i)
		c.Set(key, key)
	}

	spent := timeTask(func() {
		for i := 0; i < loop; i++ {
			testTask(func() {
				key := strconv.Itoa(rand.Intn(dataSize))
				c.Set(key, key)
				c.Get(key)
			})
		}
	})

	fmt.Printf("%s spent %dms\n", t.Name(), spent)
}

//// go test -v -run=^TestGoCacheWrite$
//func TestGoCacheWrite(t *testing.T) {
//	c := gocache.New(gocache.NoExpiration, 10*time.Minute)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set(key, key, gocache.NoExpiration)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Set(key, key, gocache.NoExpiration)
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestGoCacheRead$
//func TestGoCacheRead(t *testing.T) {
//	c := gocache.New(gocache.NoExpiration, 10*time.Minute)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set(key, key, gocache.NoExpiration)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Get(key)
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestGoCache$
//func TestGoCache(t *testing.T) {
//	c := gocache.New(gocache.NoExpiration, 10*time.Minute)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set(key, key, gocache.NoExpiration)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Set(key, key, gocache.NoExpiration)
//				c.Get(key)
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestFreeCacheWrite$
//func TestFreeCacheWrite(t *testing.T) {
//	fcache := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		fcache.Set([]byte(key), []byte(key), 0)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				fcache.Set([]byte(key), []byte(key), 0)
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestFreeCacheRead$
//func TestFreeCacheRead(t *testing.T) {
//	fcache := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		fcache.Set([]byte(key), []byte(key), 0)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				fcache.Get([]byte(key))
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestFreeCache$
//func TestFreeCache(t *testing.T) {
//	fcache := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		fcache.Set([]byte(key), []byte(key), 0)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				fcache.Set([]byte(key), []byte(key), 0)
//				fcache.Get([]byte(key))
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
