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

package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
	// By default, singleflight is enabled in cache.
	// Use WithDisableSingleflight to disable if you want.
	// cache := cachego.NewCache(cachego.WithDisableSingleflight())
	cache := cachego.NewCache()

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // <nil> false

	if !ok {
		// Load loads a value of key to cache with ttl.
		// Use cachego.NoTTL if you want this value is no ttl.
		// After loading value to cache, it returns the loaded value and error if failed.
		value, _ = cache.Load("key", time.Second, func() (value interface{}, err error) {
			return 666, nil
		})
	}

	fmt.Println(value) // 666

	value, ok = cache.Get("key")
	fmt.Println(value, ok) // 666, true

	time.Sleep(2 * time.Second)

	value, ok = cache.Get("key")
	fmt.Println(value, ok) // <nil>, false
}
