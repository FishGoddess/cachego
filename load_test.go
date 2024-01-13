// Copyright 2023 FishGoddess. All Rights Reserved.
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

package cachego

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

type testLoadCache struct {
	key   string
	value interface{}
	ttl   time.Duration

	loader *loader
}

func newTestLoadCache(singleflight bool) Cache {
	cache := &testLoadCache{
		loader: newLoader(singleflight),
	}

	return cache
}

func (tlc *testLoadCache) Get(key string) (value interface{}, found bool) {
	return tlc.value, key == tlc.key
}

func (tlc *testLoadCache) Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	tlc.key = key
	tlc.value = value
	tlc.ttl = ttl

	return nil
}

func (tlc *testLoadCache) Remove(key string) (removedValue interface{}) {
	return nil
}

func (tlc *testLoadCache) Size() (size int) {
	return 1
}

func (tlc *testLoadCache) GC() (cleans int) {
	return 0
}

func (tlc *testLoadCache) Reset() {}

func (tlc *testLoadCache) Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error) {
	return tlc.loader.Load(key, ttl, load)
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewLoader$
func TestNewLoader(t *testing.T) {
	loader := newLoader(false)
	if loader.group != nil {
		t.Fatalf("loader.group %+v != nil", loader.group)
	}

	loader = newLoader(true)
	if loader.group == nil {
		t.Fatal("loader.group == nil")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoaderLoad$
func TestLoaderLoad(t *testing.T) {
	cache := newTestLoadCache(false)
	loadCount := 0

	for i := int64(0); i < 100; i++ {
		str := strconv.FormatInt(i, 10)

		value, err := cache.Load("key", time.Duration(i), func() (value interface{}, err error) {
			loadCount++
			return str, nil
		})

		if err != nil {
			t.Fatal(err)
		}

		if value.(string) != str {
			t.Fatalf("value.(string) %s != str %s", value.(string), str)
		}
	}

	if loadCount != 100 {
		t.Fatalf("loadCount %d != 100", loadCount)
	}

	cache = newTestLoadCache(true)
	loadCount = 0

	var errs []error
	var lock sync.Mutex
	var wg sync.WaitGroup

	for i := int64(0); i < 100; i++ {
		wg.Add(1)

		go func(i int64) {
			defer wg.Done()

			str := strconv.FormatInt(i, 10)

			_, err := cache.Load("key", time.Duration(i), func() (value interface{}, err error) {
				time.Sleep(time.Second)
				loadCount++

				return str, nil
			})

			if err != nil {
				lock.Lock()
				errs = append(errs, err)
				lock.Unlock()
			}
		}(i)
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}

	if loadCount != 1 {
		t.Fatalf("loadCount %d != 1", loadCount)
	}
}
