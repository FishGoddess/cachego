# 🍰 cachego

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/cachego)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/cachego/actions/workflows/test.yml/badge.svg)

**cachego** 是一个拥有分片机制的轻量级内存缓存库，API 友好，支持多种数据淘汰机制，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

> 目前 v0.3.x 版本已经在多个线上服务中运行稳定，服务日常请求过万 qps，最高抵御过 30w/s qps 的冲击，欢迎使用！👏🏻

[Read me in English](./README.en.md).

### 🕹 功能特性

* 以键值对形式缓存数据，极简的 API 设计风格
* 引入 option function 模式，简化创建缓存参数
* 提供 ttl 过期机制，支持限制键值对数量
* 提供 lru 清理机制，提供 lfu 清理机制
* 提供锁粒度更细的分片缓存，具有非常高的并发性能
* 支持懒清理机制，每一次访问的时候判断是否过期
* 支持哨兵清理机制，每隔一定的时间间隔进行清理
* 自带 singleflight 机制，减少缓存穿透的伤害
* 自带定时任务封装，方便热数据定时加载到缓存
* 支持上报缓存状况，可自定义多个缓存上报点

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

```bash
$ go get -u github.com/FishGoddess/cachego
```

### 💡 参考案例

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

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ make bench
```

```bash
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 5800X 8-Core Processor
mem: DDR4 16GB*2 4000MHZ

BenchmarkCachegoGet-16                  36299835                33.75 ns/op            0 B/op          0 allocs/op
BenchmarkCachegoGetLRU-16               22167474                54.61 ns/op            0 B/op          0 allocs/op
BenchmarkCachegoGetLFU-16               19970642                56.66 ns/op            0 B/op          0 allocs/op
BenchmarkCachegoGetSharding-16         328267346                 4.13 ns/op            0 B/op          0 allocs/op
BenchmarkGcacheGet-16                   17657367                71.44 ns/op           16 B/op          1 allocs/op
BenchmarkGcacheGetLRU-16                14978523                76.25 ns/op           16 B/op          1 allocs/op
BenchmarkGcacheGetLFU-16                11984180               101.10 ns/op           16 B/op          1 allocs/op
BenchmarkEcacheGet-16                   23887726                49.47 ns/op            0 B/op          0 allocs/op
BenchmarkEcache2Get-16                  23970516                57.68 ns/op            0 B/op          0 allocs/op
BenchmarkBigcacheGet-16                 41191674                37.16 ns/op            7 B/op          2 allocs/op
BenchmarkFreecacheGet-16               100525257                11.22 ns/op           27 B/op          2 allocs/op
BenchmarkGoCacheGet-16                  74411682                35.19 ns/op            0 B/op          0 allocs/op

BenchmarkCachegoSet-16                  14683694                70.30 ns/op           16 B/op          1 allocs/op
BenchmarkCachegoSetLRU-16               17116057                70.84 ns/op           16 B/op          1 allocs/op
BenchmarkCachegoSetLFU-16               14976692                81.79 ns/op           16 B/op          1 allocs/op
BenchmarkCachegoSetSharding-16         100000000                10.24 ns/op           16 B/op          1 allocs/op
BenchmarkGcacheSet-16                    9618228               127.00 ns/op           56 B/op          3 allocs/op
BenchmarkGcacheSetLRU-16                 9984690               123.30 ns/op           56 B/op          3 allocs/op
BenchmarkGcacheSetLFU-16                 9982962               127.50 ns/op           56 B/op          3 allocs/op
BenchmarkEcacheSet-16                   11765893               101.30 ns/op           32 B/op          2 allocs/op
BenchmarkEcache2Set-16                  10891723               100.10 ns/op           32 B/op          2 allocs/op
BenchmarkBigcacheSet-16                  9985014               119.70 ns/op           90 B/op          0 allocs/op
BenchmarkFreecacheSet-16               191627598                 6.29 ns/op            0 B/op          0 allocs/op
BenchmarkGoCacheSet-16                  15304384                76.57 ns/op           16 B/op          1 allocs/op
```

> 注：Ecache 只有 LRU 模式，v1 和 v2 两个版本；Freecache 默认是 256 分片，无法调节为 1 个分片进行对比测试。

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)

可以看出，使用分片机制后的读写性能非常高，但是分片会多一次哈希定位的操作，如果加锁的消耗小于定位的消耗，那分片就不占优势。
不过在绝大多数的情况下，分片机制带来的性能提升都是巨大的，尤其是对写操作较多的 lru 和 lfu 实现。

### 👥 贡献者

* [cristiane](https://gitee.com/cristiane)：提供 hash 算法的优化建议
* [hzy15610046011](https://gitee.com/hzy15610046011)：提供架构设计文档和图片
* [chen661](https://gitee.com/chen661)：提供 segmentSize 设置选项的参数限制想法

如果您觉得 cachego 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

最后，我想感谢 JetBrains 公司的 **free JetBrains Open Source license(s)**，因为 `cachego` 是用该计划下的 Idea / GoLand 完成开发的。

<a href="https://www.jetbrains.com/?from=cachego" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>
