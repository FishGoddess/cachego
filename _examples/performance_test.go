// Copyright 2020 FishGoddess. All Rights Reserved.
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

package main

import (
	"fmt"
	"math/rand"
	//"runtime/debug"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/FishGoddess/cachego"
	//"github.com/coocood/freecache"
	//"github.com/orca-zhang/ecache"
	//gocache "github.com/patrickmn/go-cache"
)

//TestCacheGoRead      spent     945ms
//TestCacheGoWrite     spent     942ms
//TestCacheGo          spent     941ms
//TestGoCacheRead      spent     965ms
//TestGoCacheWrite     spent    3251ms
//TestGoCache          spent    4390ms
//TestFreeCacheRead    spent     935ms
//TestFreeCacheWrite   spent     994ms
//TestFreeCache        spent    1012ms
//TestECacheRead       spent     931ms
//TestECacheWrite      spent    1068ms
//TestECache           spent    1071ms

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
	begin := time.Now()
	task()
	return time.Since(begin).Milliseconds()
}

// go test -v -run=^TestCacheGoRead$
func TestCacheGoRead(t *testing.T) {
	c := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

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

// go test -v -run=^TestCacheGoWrite$
func TestCacheGoWrite(t *testing.T) {
	c := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

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

// go test -v -run=^TestCacheGo$
func TestCacheGo(t *testing.T) {
	c := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

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
//// go test -v -run=^TestFreeCacheRead$
//func TestFreeCacheRead(t *testing.T) {
//	c := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set([]byte(key), []byte(key), 0)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Get([]byte(key))
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestFreeCacheWrite$
//func TestFreeCacheWrite(t *testing.T) {
//	c := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set([]byte(key), []byte(key), 0)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Set([]byte(key), []byte(key), 0)
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestFreeCache$
//func TestFreeCache(t *testing.T) {
//	c := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Set([]byte(key), []byte(key), 0)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Set([]byte(key), []byte(key), 0)
//				c.Get([]byte(key))
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestECacheRead$
//func TestECacheRead(t *testing.T) {
//	c := ecache.NewLRUCache(256, 256, 0)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Put(key, key)
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
//// go test -v -run=^TestECacheWrite$
//func TestECacheWrite(t *testing.T) {
//	c := ecache.NewLRUCache(256, 256, 0)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Put(key, key)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Put(key, key)
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^TestECache$
//func TestECache(t *testing.T) {
//	c := ecache.NewLRUCache(256, 256, 0)
//
//	for i := 0; i < dataSize; i++ {
//		key := strconv.Itoa(i)
//		c.Put(key, key)
//	}
//
//	spent := timeTask(func() {
//		for i := 0; i < loop; i++ {
//			testTask(func() {
//				key := strconv.Itoa(rand.Intn(dataSize))
//				c.Put(key, key)
//				c.Get(key)
//			})
//		}
//	})
//
//	fmt.Printf("%s spent %dms\n", t.Name(), spent)
//}
//
//// go test -v -run=^$ -bench=^BenchmarkCacheGo$ -benchtime=10s
//func BenchmarkCacheGo(b *testing.B) {
//	c := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))
//
//	var wg sync.WaitGroup
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for j := 0; j < 1024; j++ {
//				key := strconv.Itoa(rand.Intn(1024))
//				c.Set(key, key)
//				c.Get(key)
//			}
//		}()
//	}
//	wg.Wait()
//}
//
//// go test -v -run=^$ -bench=^BenchmarkGoCache$ -benchtime=10s
//func BenchmarkGoCache(b *testing.B) {
//	c := gocache.New(gocache.NoExpiration, 10*time.Minute)
//
//	var wg sync.WaitGroup
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for j := 0; j < 1024; j++ {
//				key := strconv.Itoa(rand.Intn(1024))
//				c.Set(key, key, gocache.NoExpiration)
//				c.Get(key)
//			}
//		}()
//	}
//	wg.Wait()
//}
//
//// go test -v -run=^$ -bench=^BenchmarkFreeCache$ -benchtime=10s
//func BenchmarkFreeCache(b *testing.B) {
//	c := freecache.NewCache(10 * dataSize)
//	debug.SetGCPercent(20)
//
//	var wg sync.WaitGroup
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for j := 0; j < 1024; j++ {
//				key := strconv.Itoa(rand.Intn(1024))
//				c.Set([]byte(key), []byte(key), 0)
//				c.Get([]byte(key))
//			}
//		}()
//	}
//	wg.Wait()
//}
//
//// go test -v -run=^$ -bench=^BenchmarkECache$ -benchtime=10s
//func BenchmarkECache(b *testing.B) {
//	c := ecache.NewLRUCache(256, 256, 0)
//
//	b.ReportAllocs()
//	b.ResetTimer()
//	var wg sync.WaitGroup
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for j := 0; j < 1024; j++ {
//				key := strconv.Itoa(rand.Intn(1024))
//				c.Put(key, key)
//				c.Get(key)
//			}
//		}()
//	}
//	wg.Wait()
//}
