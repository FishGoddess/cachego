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
* Report supports, providing several reporting points.

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

```bash
$ make bench
```

```bash
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkCachegoGet-12                  25214618               47.2 ns/op             0 B/op          0 allocs/op
BenchmarkCachegoGetLRU-12                8169417              149.0 ns/op             0 B/op          0 allocs/op
BenchmarkCachegoGetLFU-12                7071300              171.6 ns/op             0 B/op          0 allocs/op
BenchmarkCachegoGetSharding-12          72568048               16.8 ns/op             0 B/op          0 allocs/op
BenchmarkGcacheGet-12                    4765129              252.1 ns/op            16 B/op          1 allocs/op
BenchmarkGcacheGetLRU-12                 5735739              214.0 ns/op            16 B/op          1 allocs/op
BenchmarkGcacheGetLFU-12                 4830048              250.8 ns/op            16 B/op          1 allocs/op
BenchmarkEcacheGet-12                   11515140              101.0 ns/op             0 B/op          0 allocs/op
BenchmarkEcache2Get-12                  12255506               95.6 ns/op             0 B/op          0 allocs/op
BenchmarkBigcacheGet-12                 21711988               60.4 ns/op             7 B/op          2 allocs/op
BenchmarkFreecacheGet-12                24903388               44.3 ns/op            27 B/op          2 allocs/op
BenchmarkGoCacheGet-12                  19818014               61.4 ns/op             0 B/op          0 allocs/op

BenchmarkCachegoSet-12                   5743768               209.6 ns/op           16 B/op          1 allocs/op
BenchmarkCachegoSetLRU-12                6105316               189.9 ns/op           16 B/op          1 allocs/op
BenchmarkCachegoSetLFU-12                5505601               217.2 ns/op           16 B/op          1 allocs/op
BenchmarkCachegoSetSharding-12          39012607                31.2 ns/op           16 B/op          1 allocs/op
BenchmarkGcacheSet-12                    3326841               365.3 ns/op           56 B/op          3 allocs/op
BenchmarkGcacheSetLRU-12                 3471307               318.7 ns/op           56 B/op          3 allocs/op
BenchmarkGcacheSetLFU-12                 3896512               335.1 ns/op           56 B/op          3 allocs/op
BenchmarkEcacheSet-12                    7318136               167.5 ns/op           32 B/op          2 allocs/op
BenchmarkEcache2Set-12                   7020867               175.7 ns/op           32 B/op          2 allocs/op
BenchmarkBigcacheSet-12                  4107825               268.9 ns/op           55 B/op          0 allocs/op
BenchmarkFreecacheSet-12                44181687                28.4 ns/op            0 B/op          0 allocs/op
BenchmarkGoCacheSet-12                   4921483               249.0 ns/op           16 B/op          1 allocs/op
```

> Notice: Ecache only has lru mode, including v1 and v2; Freecache has 256 shardings, and we can't reset to 1.

> Benchmarks: [_examples/performance_test.go](./_examples/performance_test.go)

As you can see, cachego has a higher performance with sharding, but sharding has one-more-time positioning
operation, so if the locking cost is less than the cost of positioning, this sharding is dragging.

### 👥 Contributors

* [cristiane](https://gitee.com/cristiane): Provide some optimizations about hash
* [hzy15610046011](https://gitee.com/hzy15610046011): Provide architecture design documents and pictures
* [chen661](https://gitee.com/chen661): Provide the limit thought of argument in WithSegmentSize Option

Please open an _**issue**_ if you find something is not working as expected.

At last, I want to thank JetBrains for **free JetBrains Open Source license(s)**, because `cachego` is developed with
Idea / GoLand under it.

<a href="https://www.jetbrains.com/?from=cachego" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>
