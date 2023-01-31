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
	"strconv"
	"sync"
	"testing"
	"time"
	//"github.com/coocood/freecache"
	//"github.com/orca-zhang/ecache"
	//gocache "github.com/patrickmn/go-cache"

	"github.com/FishGoddess/cachego"
)

const (
	benchTTL           = time.Minute
	benchMaxKeys       = 100000
	benchMaxLoops      = 1000
	benchMaxGoroutines = 4096
)

type benchKeys []string

func newBenchKeys() benchKeys {
	keys := make([]string, 0, benchMaxKeys)

	for i := 0; i < benchMaxKeys; i++ {
		keys = append(keys, strconv.Itoa(i))
	}

	return keys
}

func (bks benchKeys) pick() string {
	index := rand.Intn(len(bks))
	return bks[index]
}

func benchmarkCache(t *testing.T, fn func(loop int, key string)) {
	keys := newBenchKeys()
	begin := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < benchMaxGoroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for loop := 0; loop < benchMaxLoops; loop++ {
				fn(loop, keys.pick())
			}
		}()
	}

	wg.Wait()
	fmt.Printf("%s: %s\n", t.Name(), time.Since(begin))
}

// go test -v -run=^TestCachegoSet$ ./_examples/performance_test.go
func TestCachegoSet(t *testing.T) {
	cache := cachego.NewCache()

	benchmarkCache(t, func(loop int, key string) {
		cache.Set(key, loop, benchTTL)
	})
}
