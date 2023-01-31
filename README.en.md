# 🍰 cachego

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/cachego)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/cachego/actions/workflows/test.yml/badge.svg)

**cachego** is an api friendly memory-based cache for [GoLang](https://golang.org) applications.

> It has been used by many services in production, all services are running stable, and the highest qps in services is
> 17w/s, so just use it if you want! 👏🏻

[阅读中文版的 Read me](./README.md).

### 🕹 Features

* Cache as entries with minimalist API design
* Use option function mode to customize the creation of cache
* TTL supports and max size of entries in cache
* LRU supports and LFU supports.
* Use sharding lock mechanism to provide a high performance in concurrency
* Lazy cleanup supports, expired before accessing
* Sentinel cleanup supports, cleaning up at fixed duration
* Singleflight supports, which can decrease the times of cache penetration
* Timer task supports, which is convenient to load data to cache

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to get more information._

### 🚀 Installation

```bash
$ go get -u github.com/FishGoddess/cachego
```

### 💡 Examples

```go
package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
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
}
```

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

> Benchmark file：[_examples/performance_test.go](./_examples/performance_test.go)

```bash
$ go test -v ./_examples/performance_test.go
```

> Data size is 1 million, concurrency is 100 thousands, loop is 50

> Environment：R7-5800X CPU @ 3.8GHZ GHZ, 32 GB RAM, Deepin20 OS

| tests       | read time (less is better) | write time (less is better) | mixed-operation time (less is better) |
|-------------|----------------------------|-----------------------------|---------------------------------------|
| **cachego** | **1092ms**                 | **1107ms**                  | **1098ms**                            |
| go-cache    | 1111ms                     | 3152ms                      | 4738ms                                |
| freeCache   | 1070ms                     | 1123ms                      | 1068ms                                |
| ECache      | 1083ms                     | 1229ms                      | 1121ms                                |

As you can see, cachego has a high performance in concurrent, but segmented lock mechanism has one-more-time positioning
operation, so if the price of locking is less than the cost of positioning, this mechanism is dragging. The reading
performance will be optimized in the future version!

### 👥 Contributors

* [cristiane](https://gitee.com/cristiane): Provide some optimizations about hash
* [hzy15610046011](https://gitee.com/hzy15610046011): Provide architecture design documents and pictures
* [chen661](https://gitee.com/chen661): Provide the limit thought of argument in WithSegmentSize Option

Please open an _**issue**_ if you find something is not working as expected.

At last, I want to thank JetBrains for **free JetBrains Open Source license(s)**, because `cachego` is developed with
Idea / GoLand under it.

<a href="https://www.jetbrains.com/?from=cachego" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>
