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
	// Both of them are set a key-value with no TTL.
	//cache.Set("key", 666, cachego.WithSetNoTTL())
	cache.Set("key", 666)

	// Get returns the value of this key.
	v, err := cache.Get("key")
	fmt.Println(v, err) // Output: 666 <nil>

	// If you pass a not existed key to of method, nil and false will be returned.
	v, err = cache.Get("not existed key")
	fmt.Println(v, err) // Output: <nil> cachego: key not found

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

2. the ttl usage:

	// Create a cache and set an entry to cache.
	cache := cachego.NewCache()
	cache.Set("key", "value", cachego.WithSetTTL(3*time.Second))

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

3. the gc usage:

	// Create a cache and set an entry to cache.
	cache := cachego.NewCache()
	cache.Set("key", "value", cachego.WithSetTTL(1*time.Second))

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

4. the option usage:

	// We use option function to customize the creation of cache.
	// You can just new one without options.
	cache := cachego.NewCache()
	cache.Set("key", "value")

	// You can set it to a cache with automatic gc if you want
	//  Try WithAutoGC.
	cache = cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Also, you can add more than one option to cache.
	cache = cachego.NewCache(cachego.WithAutoGC(10 * time.Minute), cachego.WithMapSize(64), cachego.WithSegmentSize(4096))

*/
package cachego // import "github.com/FishGoddess/cachego"

// Version is the version string representation of cachego.
const Version = "v0.3.0-alpha"
