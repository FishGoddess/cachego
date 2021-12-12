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
Package cachego provides an easy way to use foundation for your caching operations.

1. the basic usage:

	// Create a cache for use.
	// We use option function to customize the creation of cache.
	// WithAutoGC means it will do gc automatically.
	cache := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Set a new entry to cache.
	cache.Set("key", 666)

	// Get returns the value of this key.
	v, ok := cache.Get("key")
	fmt.Println(v, ok) // Output: 666 true

	// If you pass a not existed key to of method, nil and false will be returned.
	v, ok = cache.Get("not existed key")
	fmt.Println(v, ok) // Output: <nil> false

	// SetWithTTL sets an entry with expired time.
	// The unit of expired time is second.
	// See more information in example of ttl.
	cache.SetWithTTL("ttlKey", 123, 10)

	// Also, you can get value from cache first, then load it to cache if missed.
	// onMissed is usually used to get data from db or somewhere, so you can refresh the value in cache.
	cache.GetWithLoad("newKey", func() (data interface{}, ttl int64, err error) {
		return "newValue", 3, nil
	})

	// We provide a way to set data to cache automatically, so you can access some hottest data extremely fast.
	stopAutoSet := cache.AutoSet("autoKey", 1*time.Second, func() (interface{}, error) {
		fmt.Println("AutoSet invoking...")
		return nil, nil
	})

	// Keep main running in order to see what AutoSet did
	time.Sleep(5 * time.Second)
	stopAutoSet <- struct{}{}

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

	// However, the key is still in cache and you should delete it by Delete() or DeleteAll().
	// So, we provide an automatic way to delete those who are dead. See more information in example of gc.
	cache.AutoGC(10 * time.Minute)

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

	// We should call GC() to clean up these dead entries.
	// Notice that this method will takes some CPU time to finish this task.
	cache.GC()
	size = cache.Size()
	fmt.Println(size) // Output: 0

	// Also, we provide an automatic way to do this job at fixed duration.
	// It returns a channel which can be used to stop this automatic job.
	// If you want to stop it, just send an true or false to the chan!
	stopAutoGc := cache.AutoGC(10 * time.Minute)
	stopAutoGc <- struct{}{}

4. the option usage:

	// We use option function to customize the creation of cache.
	// You can just new it without options.
	cache := cachego.NewCache()
	cache.Set("key", "value")

	// You can set it to a cache with automatic gc if you want
	//  Try WithAutoGC.
	cache = cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Also, you can add more than one option to cache.
	cache = cachego.NewCache(cachego.WithAutoGC(10 * time.Minute), cachego.WithMapSize(64), cachego.WithSegmentSize(4096))

	// Every option has its function, and you should use them for some purposes.
	// WithDebugPoint runs a http server and registers some handlers for debug.
	cachego.WithDebugPoint(":8888") // try to visit :8888
	time.Sleep(time.Minute)

*/
package cachego // import "github.com/FishGoddess/cachego"

// Version is the version string representation of cachego.
const Version = "v0.3.0-alpha"
