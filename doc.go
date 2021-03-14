// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
// Created at 2020/03/13 16:15:56

/*
Package cache provides an easy way to use foundation for your caching operations.

1. the basic usage:

	// Create a cache for use.
	cache := cachego.NewCache()

	// Set a new entry to cache.
	cache.Set("key", 666)

	// Get returns the value of this key.
	v, ok := cache.Get("key")
	fmt.Println(v, ok) // Output: 666 true

	// If you want to change the value of a key, just set a new value of this key.
	cache.Set("key", "value")

	// See what value it has.
	v, ok = cache.Get("key")
	fmt.Println(v, ok) // Output: value true

	// If you pass a not existed key to of method, nil and false will be returned.
	v, ok = cache.Get("not existed key")
	fmt.Println(v, ok) // Output: <nil> false

2. the ttl usage:

	// Create a cache and set an entry to cache.
	// The ttl is 3 seconds.
	cache := cachego.NewCache()
	cache.SetWithTTL("key", "value", 3)

	// Check if the key is alive.
	value, ok := cache.Get("key")
	fmt.Println(value, ok) // Output: value true

	// Wait for 5 seconds and check again.
	// Now the key is gone.
	time.Sleep(5 * time.Second)
	value, ok = cache.Get("key")
	fmt.Println(value, ok) // Output: <nil> false

3. the gc usage:

	// Create a cache and set an entry to cache.
	// The ttl is 1 second.
	cache := cachego.NewCache()
	cache.SetWithTTL("key", "value", 1)

	// Wait for 2 seconds and check the key.
	time.Sleep(2 * time.Second)

	// We can see this key is gone and we can't get it anymore.
	value, ok := cache.Get("key")
	fmt.Println(value, ok) // Output: <nil> false

	// However, the key still stores in cache and occupies the space.
	size := cache.Size()
	fmt.Println(size) // Output: 1

	// We should call Gc() to clean up these dead entries.
	// Notice that this method will takes some CPU time to finish this task.
	cache.Gc()
	size = cache.Size()
	fmt.Println(size) // Output: 0

	// Also, we provide an automatic way to do this job at fixed duration.
	// It returns a <-chan type which can be used to stop this automatic job.
	// If you want to stop it, just send an true or false to the chan!
	stopAutoGc := cache.AutoGc(10 * time.Minute)
	stopAutoGc <- true

*/
package cachego // import "github.com/FishGoddess/cachego"

// Version is the version string representation of cachego.
const Version = "v0.1.3"
