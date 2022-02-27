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
	cache.Set("key", "value", cachego.WithOpTTL(1*time.Second))

	value, err := cache.Get("key")
	fmt.Println(value, err) // Output: value <nil>

	// Wait for 2 seconds and check the key.
	time.Sleep(2 * time.Second)

	// We can see this key is gone, and we can't get it anymore.
	value, err = cache.Get("key")
	fmt.Println(value, err) // Output: <nil> cachego: key not found

	// However, the key still stores in cache and occupies the space.
	size := cache.Size()
	fmt.Println(size) // Output: 1

	// We should call GC() to clean up these dead entries.
	// Notice that this method will take some CPU time to finish this task.
	cache.GC()
	size = cache.Size()
	fmt.Println(size) // Output: 0

	// Also, we provide an automatic way to do this job at fixed duration.
	// It returns a channel which can be used to stop this automatic job.
	// If you want to stop it, just send an true or false to the chan!
	stopAutoGc := cache.AutoGC(10 * time.Minute)
	stopAutoGc <- struct{}{}
}
