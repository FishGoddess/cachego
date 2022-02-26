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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/03/16 22:29:02

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
	// Create a cache for use.
	// We use option function to customize the creation of cache.
	// WithAutoGC means it will do gc automatically.
	cache := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Set a new entry to cache.
	// Both of them are set a key-value with no TTL.
	//cache.Set("key", 666, cachego.WithSetNoTTL())
	cache.Set("key", 666)

	// Get returns the value of this key.
	v, err := cache.Get("key")
	fmt.Println(v, err) // Output: 666 <nil>

	// If you pass a not existed key to of method, nil and errNotFound will be returned.
	v, err = cache.Get("not existed key")
	if cachego.IsNotFound(err) {
		fmt.Println(v, err) // Output: <nil> cachego: key not found
	}

	// SetWithTTL sets an entry with expired time.
	// See more information in example of ttl.
	cache.Set("ttlKey", 123, cachego.WithSetTTL(10*time.Second))

	// Also, you can get value from cache first, then load it to cache if missed.
	// OnMissed is usually used to get data from db or somewhere, so you can refresh the value in cache.
	// Notice ctx in onMissed is passed by Get option.
	onMissed := func(ctx context.Context) (data interface{}, err error) {
		return "newValue", nil
	}

	v, err = cache.Get("newKey", cachego.WithGetOnMissed(onMissed), cachego.WithGetTTL(3*time.Second))
	fmt.Println(v, err) // Output: newValue <nil>

	// We provide a way to set data to cache automatically, so you can access some hottest data extremely fast.
	loadFunc := func(ctx context.Context) (interface{}, error) {
		fmt.Println("AutoSet invoking...")
		return nil, nil
	}

	stopCh := cache.AutoSet("autoKey", loadFunc, cachego.WithAutoSetGap(1*time.Second))

	// Keep main running in order to see what AutoSet did.
	time.Sleep(5 * time.Second)
	stopCh <- struct{}{} // Stop AutoSet task
}
