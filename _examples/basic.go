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
	cache := cachego.NewStandardCache()
	cache.Set("key", 123, time.Second)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 123 true

	size := cache.Size()
	fmt.Println(size) // 1

	time.Sleep(2 * time.Second)

	value, ok = cache.Get("key")
	fmt.Println(value, ok) // <nil> false

	size = cache.Size()
	fmt.Println(size) // 1

	cleans := cache.GC()
	fmt.Println(cleans) // 1

	cache.Set("key", 123, cachego.NoTTL)

	removedValue := cache.Remove("key")
	fmt.Println(removedValue) // 123

	cache.Reset()

	value, ok = cache.Get("key")
	if !ok {
		value, _ = cache.Load("key", time.Second, func() (value interface{}, err error) {
			return 666, nil
		})
	}

	fmt.Println(value) // 666
}
