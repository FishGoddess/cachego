// Copyright 2021 FishGoddess. All Rights Reserved.
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

package singleflight

import (
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func testGroupCall(t *testing.T, group *Group, concurrency int) {
	var wg sync.WaitGroup

	key := strconv.Itoa(rand.Int())
	rightResult := int64(0)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)

		go func(index int64) {
			defer wg.Done()

			result, err := group.Call(key, func() (interface{}, error) {
				time.Sleep(time.Second)
				atomic.StoreInt64(&rightResult, index)
				return index, nil
			})

			if err != nil {
				t.Error(err)
			}

			r := atomic.LoadInt64(&rightResult)
			if result != r {
				t.Errorf("result %d != rightResult %d", result, r)
			}
		}(int64(i))
	}

	wg.Wait()
}

// go test -v -cover -run=^TestGroupCall$
func TestGroupCall(t *testing.T) {
	group := NewGroup(128)
	testGroupCall(t, group, 100000)
}

// go test -v -cover -run=^TestGroupCallMultiKey$
func TestGroupCallMultiKey(t *testing.T) {
	group := NewGroup(128)

	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			testGroupCall(t, group, 1000)
		}()
	}

	wg.Wait()
}
