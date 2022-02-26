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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/12/25 23:30:29

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
	// In default, cachego enables single-flight mode in get operations.
	// Just use WithGetOnMissed option to enjoy the flight of data.
	cache := cachego.NewCache()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cache.Get("key1", cachego.WithGetOnMissed(func(ctx context.Context) (data interface{}, err error) {
				time.Sleep(30 * time.Millisecond) // Assume I/O costs 30ms
				fmt.Println("key1: single-flight")
				return 123, nil
			}))
		}()
	}
	wg.Wait()

	// If you want to disable single-flight mode in some Get operations, try this:
	wg = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cache.Get("key2", cachego.WithGetOnMissed(func(ctx context.Context) (data interface{}, err error) {
				time.Sleep(30 * time.Millisecond) // Assume I/O costs 30ms
				fmt.Println("key2: multi-flight")
				return 456, nil
			}), cachego.WithGetDisableSingleflight())
		}()
	}
	wg.Wait()

	// Of course, we all know single-flight mode will decrease the success rate of loading data.
	// So you can disable it globally if you need.
	cache = cachego.NewCache(cachego.WithDisableSingleflight())

	wg = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cache.Get("key3", cachego.WithGetOnMissed(func(ctx context.Context) (data interface{}, err error) {
				time.Sleep(30 * time.Millisecond) // Assume I/O costs 30ms
				fmt.Println("key3: multi-flight")
				return 666, nil
			}))
		}()
	}
	wg.Wait()
}
