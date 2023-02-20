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

/*
Package cachego provides an easy way to use foundation for your caching operations.

1. basic:

	// Use NewCache function to create a cache.
	// You can use WithLRU to specify the type of cache to lru.
	// Also, try WithLFU if you want to use lfu to evict data.
	cache := cachego.NewCache(cachego.WithLRU(100))
	cache = cachego.NewCache(cachego.WithLFU(100))

	// By default, it creates a standard cache which evicts entries randomly.
	// Use WithShardings to shard cache to several parts for higher performance.
	cache = cachego.NewCache(cachego.WithShardings(64))
	cache = cachego.NewCache()

	// Set an entry to cache with ttl.
	cache.Set("key", 123, time.Second)

	// Get an entry from cache.
	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 123 true

	// Check how many entries stores in cache.
	size := cache.Size()
	fmt.Println(size) // 1

	// Clean expired entries.
	cleans := cache.GC()
	fmt.Println(cleans) // 1

	// Set an entry which doesn't have ttl.
	cache.Set("key", 123, cachego.NoTTL)

	// Remove an entry.
	removedValue := cache.Remove("key")
	fmt.Println(removedValue) // 123

	// Reset resets cache to initial status.
	cache.Reset()

	// Get value from cache and load it to cache if not found.
	value, ok = cache.Get("key")
	if !ok {
		// Loaded entry will be set to cache and returned.
		value, _ = cache.Load("key", time.Second, func() (value interface{}, err error) {
			return 666, nil
		})
	}

	fmt.Println(value) // 666

2. ttl:

	cache := cachego.NewCache()

	// We think most of the entries in cache should have its ttl.
	// So set an entry to cache should specify a ttl.
	cache.Set("key", 666, time.Second)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 666 true

	time.Sleep(2 * time.Second)

	// The entry is expired after ttl.
	value, ok = cache.Get("key")
	fmt.Println(value, ok) // <nil> false

	// Notice that the entry still stores in cache even if it's expired.
	// This is because we think you will reset entry to cache after cache missing in most situations.
	// So we can reuse this entry and just reset its value and ttl.
	size := cache.Size()
	fmt.Println(size) // 1

	// What should I do if I want an expired entry never storing in cache? Try GC:
	cleans := cache.GC()
	fmt.Println(cleans) // 1

	size = cache.Size()
	fmt.Println(size) // 0

	// However, not all entries have ttl, and you can specify a NoTTL constant to do so.
	// In fact, the entry won't expire as long as its ttl is <= 0.
	// So you may have known NoTTL is a "readable" value of "<= 0".
	cache.Set("key", 666, cachego.NoTTL)

3. lru:

	// By default, NewCache() returns a standard cache which evicts entries randomly.
	cache := cachego.NewCache(cachego.WithMaxEntries(10))

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, cachego.NoTTL)
	}

	// Since we set 20 entries to cache, the size won't be 20 because we limit the max entries to 10.
	size := cache.Size()
	fmt.Println(size) // 10

	// We don't know which entries will be evicted and stayed.
	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(key, value, ok)
	}

	fmt.Println()

	// Sometimes we want it evicts entries by lru, try WithLRU.
	// You need to specify the max entries storing in lru cache.
	// More details see https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU).
	cache = cachego.NewCache(cachego.WithLRU(10))

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, cachego.NoTTL)
	}

	// Only the least recently used entries can be got in a lru cache.
	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(key, value, ok)
	}

	// By default, lru will share one lock to do all operations.
	// You can sharding cache to several parts for higher performance.
	// Notice that max entries only effect to one part in sharding mode.
	// For example, the total max entries will be 2*10 if shardings is 2 and max entries is 10 in WithLRU or WithMaxEntries.
	// In some cache libraries, they will calculate max entries in each parts of shardings, like 10/2.
	// However, the result divided by max entries and shardings may be not an integer which will make the total max entries incorrect.
	// So we let users decide the exact max entries in each parts of shardings.
	cache = cachego.NewCache(cachego.WithShardings(2), cachego.WithLRU(10))

4. lfu:

	// By default, NewCache() returns a standard cache which evicts entries randomly.
	cache := cachego.NewCache(cachego.WithMaxEntries(10))

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, cachego.NoTTL)
	}

	// Since we set 20 entries to cache, the size won't be 20 because we limit the max entries to 10.
	size := cache.Size()
	fmt.Println(size) // 10

	// We don't know which entries will be evicted and stayed.
	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(key, value, ok)
	}

	fmt.Println()

	// Sometimes we want it evicts entries by lfu, try WithLFU.
	// You need to specify the max entries storing in lfu cache.
	// More details see https://en.wikipedia.org/wiki/Cache_replacement_policies#Least-frequently_used_(LFU).
	cache = cachego.NewCache(cachego.WithLFU(10))

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)

		// Let entries have some frequently used operations.
		for j := 0; j < i; j++ {
			cache.Set(key, i, cachego.NoTTL)
		}
	}

	for i := 0; i < 20; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(key, value, ok)
	}

	// By default, lfu will share one lock to do all operations.
	// You can sharding cache to several parts for higher performance.
	// Notice that max entries only effect to one part in sharding mode.
	// For example, the total max entries will be 2*10 if shardings is 2 and max entries is 10 in WithLFU or WithMaxEntries.
	// In some cache libraries, they will calculate max entries in each parts of shardings, like 10/2.
	// However, the result divided by max entries and shardings may be not an integer which will make the total max entries incorrect.
	// So we let users decide the exact max entries in each parts of shardings.
	cache = cachego.NewCache(cachego.WithShardings(2), cachego.WithLFU(10))

5. sharding:

	// All operations in cache share one lock for concurrency.
	// Use read lock or write lock is depends on cache implements.
	// Get will use read lock in standard cache, but lru and lfu don't.
	// This may be a serious performance problem in high qps.
	cache := cachego.NewCache()

	// We provide a sharding cache wrapper to shard one cache to several parts with hash.
	// Every parts store its entries and all operations of one entry work on one part.
	// This means there are more than one lock when you operate entries.
	// The performance will be better in high qps.
	cache = cachego.NewCache(cachego.WithShardings(64))
	cache.Set("key", 666, cachego.NoTTL)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 666 true

	// Notice that max entries will be the sum of shards.
	// For example, we set WithShardings(4) and WithMaxEntries(100), and the max entries in whole cache will be 4 * 100.
	cache = cachego.NewCache(cachego.WithShardings(4), cachego.WithMaxEntries(100))

	for i := 0; i < 1000; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, cachego.NoTTL)
	}

	size := cache.Size()
	fmt.Println(size) // 400

6. gc:

	cache := cachego.NewCache()
	cache.Set("key", 666, time.Second)

	time.Sleep(2 * time.Second)

	// The entry is expired after ttl.
	value, ok := cache.Get("key")
	fmt.Println(value, ok) // <nil> false

	// As you know the entry still stores in cache even if it's expired.
	// This is because we think you will reset entry to cache after cache missing in most situations.
	// So we can reuse this entry and just reset its value and ttl.
	size := cache.Size()
	fmt.Println(size) // 1

	// What should I do if I want an expired entry never storing in cache? Try GC:
	cleans := cache.GC()
	fmt.Println(cleans) // 1

	// Is there a smart way to do that? Try WithGC:
	// For testing, we set a small duration of gc.
	// You should set at least 3 minutes in production for performance.
	cache = cachego.NewCache(cachego.WithGC(2 * time.Second))
	cache.Set("key", 666, time.Second)

	size = cache.Size()
	fmt.Println(size) // 1

	time.Sleep(3 * time.Second)

	size = cache.Size()
	fmt.Println(size) // 0

	// Or you want a cancalable gc task? Try RunGCTask:
	cache = cachego.NewCache()
	cancel := cachego.RunGCTask(cache, 2*time.Second)

	cache.Set("key", 666, time.Second)

	size = cache.Size()
	fmt.Println(size) // 1

	time.Sleep(3 * time.Second)

	size = cache.Size()
	fmt.Println(size) // 0

	cancel()

	cache.Set("key", 666, time.Second)

	size = cache.Size()
	fmt.Println(size) // 1

	time.Sleep(3 * time.Second)

	size = cache.Size()
	fmt.Println(size) // 1

	// By default, gc only scans at most maxScans entries one time to remove expired entries.
	// This is because scans all entries may cost much time if there is so many entries in cache, and a "stw" will happen.
	// This can be a serious problem in some situations.
	// Use WithMaxScans to set this value, remember, a value <= 0 means no scan limit.
	cache = cachego.NewCache(cachego.WithGC(10*time.Minute), cachego.WithMaxScans(0))

7. load:

	// By default, singleflight is enabled in cache.
	// Use WithDisableSingleflight to disable if you want.
	cache := cachego.NewCache(cachego.WithDisableSingleflight())

	// We recommend you to use singleflight.
	cache = cachego.NewCache()

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

8. report:

	func reportMissed(key string) {
		fmt.Printf("report: missed key %s\n", key)
	}

	func reportHit(key string, value interface{}) {
		fmt.Printf("report: hit key %s value %+v\n", key, value)
	}

	func reportGC(cost time.Duration, cleans int) {
		fmt.Printf("report: gc cost %s cleans %d\n", cost, cleans)
	}

	func reportLoad(key string, value interface{}, ttl time.Duration, err error) {
		fmt.Printf("report: load key %s value %+v ttl %s, err %+v\n", key, value, ttl, err)
	}

	// Create a cache as usual.
	cache := cachego.NewCache(
		cachego.WithMaxEntries(3),
		cachego.WithGC(100*time.Millisecond),
	)

	// Use Report function to wrap a cache with reporting logics.
	// We provide some reporting points for monitor cache.
	// ReportMissed reports the missed key getting from cache.
	// ReportHit reports the hit entry getting from cache.
	// ReportGC reports the status of cache gc.
	// ReportLoad reports the result of loading.
	cache, reporter := cachego.Report(
		cache,
		cachego.WithReportMissed(reportMissed),
		cachego.WithReportHit(reportHit),
		cachego.WithReportGC(reportGC),
		cachego.WithReportLoad(reportLoad),
	)

	for i := 0; i < 5; i++ {
		key := strconv.Itoa(i)
		evictedValue := cache.Set(key, key, 10*time.Millisecond)
		fmt.Println(evictedValue)
	}

	for i := 0; i < 5; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(value, ok)
	}

	time.Sleep(200 * time.Millisecond)

	value, err := cache.Load("key", time.Second, func() (value interface{}, err error) {
		return 666, io.EOF
	})

	fmt.Println(value, err)

	// These are some methods of reporter.
	fmt.Println("CountMissed:", reporter.CountMissed())
	fmt.Println("CountHit:", reporter.CountHit())
	fmt.Println("CountGC:", reporter.CountGC())
	fmt.Println("CacheSize:", reporter.CacheSize())
	fmt.Println("MissedRate:", reporter.MissedRate())
	fmt.Println("HitRate:", reporter.HitRate())

9. task:

	// Create a context to stop the task.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Wrap context with key and value
	ctx = context.WithValue(ctx, contextKey, "hello")

	// Use New to create a task and run it.
	// You can use it to load some hot data to cache at fixed duration.
	// Before is called before the task loop, optional.
	// After is called after the task loop, optional.
	// Context is passed to fn include fn/before/after which can stop the task by Done(), optional.
	// Duration is the duration between two loop of fn, optional.
	// Run will start a new goroutine and run the task loop.
	// The task will stop if context is done.
	task.New(printContextValue).
		Before(beforePrint).
		After(afterPrint).
		Context(ctx).
		Duration(time.Second).
		Run()

10. clock:

	// Create a fast clock and get current time in nanosecond by Now.
	c := clock.New()
	c.Now()

	// Fast clock may return an "incorrect" time compared with time.Now.
	// The gap will be smaller than about 100 ms.
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int63n(int64(time.Second))))

		timeNow := time.Now().UnixNano()
		clockNow := c.Now()

		fmt.Println(timeNow)
		fmt.Println(clockNow)
		fmt.Println("gap:", time.Duration(timeNow-clockNow))
		fmt.Println()
	}

	// You can specify the fast clock to cache by WithNow.
	// All getting current time operations in this cache will use fast clock.
	cache := cachego.NewCache(cachego.WithNow(clock.New().Now))
	cache.Set("key", 666, 100*time.Millisecond)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 666, true

	time.Sleep(200 * time.Millisecond)

	value, ok = cache.Get("key")
	fmt.Println(value, ok) // <nil>, false
*/
package cachego // import "github.com/FishGoddess/cachego"

// Version is the version string representation of cachego.
const Version = "v0.4.4-alpha"
