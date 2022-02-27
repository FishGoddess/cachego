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
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
	// Create a cache and set an entry to cache.
	cache := cachego.NewCache()
	cache.Set("key", "value", cachego.WithOpTTL(3*time.Second))

	// Check if the key is alive.
	value, err := cache.Get("key")
	fmt.Println(value, err) // Output: value <nil>

	// Wait for 5 seconds and check again.
	// Now the key is gone.
	time.Sleep(5 * time.Second)
	value, err = cache.Get("key")
	fmt.Println(value, err) // Output: <nil> cachego: key not found

	// However, the key is still in cache, and you should remove it by Delete() or DeleteAll().
	// So, we provide an automatic way to remove those who are dead. See more information in example of gc.
	cache.AutoGC(10 * time.Minute)
}
