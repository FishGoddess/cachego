# 🍰 cachego

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/cachego)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/cachego/actions/workflows/test.yml/badge.svg)

**cachego** 是一个拥有分段锁机制的轻量级内存缓存库，API 友好，支持多种数据淘汰机制，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

> 目前 v0.3.x 版本已经在多个线上服务中运行稳定，服务日常请求过万 qps，最高抵御过 17w/s qps 的冲击，欢迎使用！👏🏻

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

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)

```bash
$ go test -v ./_examples/performance_test.go
```

> 总缓存数据为 100w 条，并发数为 10w，循环测试写入和读取次数为 50 次

> 测试环境：R7-5800X CPU @ 3.8GHZ GHZ，32 GB RAM，Manjaro21 OS

| 测试          | 读取消耗时间 (越小越好) | 写入消耗时间 (越小越好) | 混合操作消耗时间 (越小越好) |
|-------------|---------------|---------------|-----------------|
| **cachego** | **1092ms**    | **1107ms**    | **1098ms**      |
| go-cache    | 1111ms        | 3152ms        | 4738ms          |
| freeCache   | 1070ms        | 1123ms        | 1068ms          |
| ECache      | 1083ms        | 1229ms        | 1121ms          |

可以看出，由于使用了分段锁机制，读写性能在并发下依然非常高，但是分段锁会多一次定位的操作，如果加锁的消耗小于定位的消耗，那分段锁就不占优势。 这也是为什么 cachego 在写入性能上比 go-cache
强一大截，但是读取性能却没强多少的原因。后续会着重优化读取性能！

### 👥 贡献者

* [cristiane](https://gitee.com/cristiane)：提供 hash 算法的优化建议
* [hzy15610046011](https://gitee.com/hzy15610046011)：提供架构设计文档和图片
* [chen661](https://gitee.com/chen661)：提供 segmentSize 设置选项的参数限制想法

如果您觉得 cachego 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

最后，我想感谢 JetBrains 公司的 **free JetBrains Open Source license(s)**，因为 `cachego` 是用该计划下的 Idea / GoLand 完成开发的。

<a href="https://www.jetbrains.com/?from=cachego" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>
