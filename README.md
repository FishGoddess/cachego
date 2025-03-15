# 🍰 cachego

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/cachego)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/cachego/actions/workflows/test.yml/badge.svg)

**cachego** 是一个拥有分片机制的轻量级内存缓存库，API 友好，支持多种数据淘汰机制，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

> 目前已经在多个线上服务中运行稳定，服务日常请求过万 qps，最高抵御过 96w/s qps 的冲击，欢迎使用！👏🏻

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
* 自带快速时钟，支持纳秒级获取时间

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
	// By default, it creates a standard cache which evicts entries randomly.
	// Use WithShardings to shard cache to several parts for higher performance.
	// Use WithGC to clean expired entries every 10 minutes.
	cache := cachego.NewCache(cachego.WithGC(10*time.Minute), cachego.WithShardings(64))

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
		// By default, it will use singleflight.
		value, _ = cache.Load("key", time.Second, func() (value interface{}, err error) {
			return 666, nil
		})
	}

	fmt.Println(value) // 666

	// You can use WithLRU to specify the type of cache to lru.
	// Also, try WithLFU if you want to use lfu to evict data.
	cache = cachego.NewCache(cachego.WithLRU(100))
	cache = cachego.NewCache(cachego.WithLFU(100))

	// Use NewCacheWithReport to create a cache with report.
	cache, reporter := cachego.NewCacheWithReport(cachego.WithCacheName("test"))
	fmt.Println(reporter.CacheName())
	fmt.Println(reporter.CacheType())
}
```

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

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

> 注：Ecache 只有 LRU 模式，v1 和 v2 两个版本；Freecache 默认是 256 分片，无法调节为 1 个分片进行对比测试。

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)。

### 👥 贡献者

* [cristiane](https://gitee.com/cristiane)：提供 hash 算法的优化建议
* [hzy15610046011](https://gitee.com/hzy15610046011)：提供架构设计文档和图片
* [chen661](https://gitee.com/chen661)：提供 segmentSize 设置选项的参数限制想法

如果您觉得 cachego 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

[![Star History Chart](https://api.star-history.com/svg?repos=fishgoddess/cachego&type=Date)](https://star-history.com/#fishgoddess/cachego&Date)
