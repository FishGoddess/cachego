# 📜 cachego

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/cachego)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**cachego** 是一个拥有高性能分段锁机制的轻量级内存缓存，拥有懒清理和哨兵清理两种清理机制，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

> 目前已经在多个线上服务中运行良好，也抵御过最高 17w/s qps 的冲击，可以稳定使用！

> 我正在开发 v0.3.x 版本，这将在 API 以及功能上达到全新的使用体验，敬请期待，也期待大家的建议！！！

[Read me in English](./README.en.md).

### 🕹 功能特性

* 以键值对形式缓存数据，极简的 API 设计风格
* 引入 option function 模式，可定制化各种操作的过程
* 使用粒度更细的分段锁机制进行设计，具有非常高的并发性能
* 支持懒清理机制，每一次访问的时候判断是否过期
* 支持哨兵清理机制，每隔一定的时间间隔进行清理
* 自带 singleflight 机制，减少缓存穿透的伤害
* ....

_更多功能请参考 [_examples](_examples)。架构设计请参考 [arch.md](_examples/docs/arch.md) 文档。_

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

```bash
$ go get -u github.com/FishGoddess/cachego
```

### 💡 参考案例

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FishGoddess/cachego"
	"github.com/FishGoddess/cachego/pkg/task"
)

func main() {
	// Create a cache for use.
	// We use option function to customize the creation of cache.
	// WithAutoGC means it will do gc automatically.
	cache := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Set a new entry to cache.
	// Both of them are set a key-value with no ttl.
	//cache.Set("key", 666, cachego.WithSetNoTTL())
	cache.Set("key", 666)

	// Get returns the value of this key.
	v, err := cache.Get("key")
	fmt.Println(v, err) // Output: 666 <nil>

	// If you pass a not existed key to of method, nil and errNotFound will be returned.
	v, err = cache.Get("not existed key")
	if cachego.IsNotFound(err) {
		fmt.Println(v, err) // Output: <nil> cachego: key not found
	}

	// SetWithTTL sets an entry with expired time.
	// See more information in example of ttl.
	cache.Set("ttlKey", 123, cachego.WithOpTTL(10*time.Second))

	// Also, you can get value from cache first, then load it to cache if missed.
	// onMissed is usually used to get data from db or somewhere, so you can refresh the value in cache.
	// Notice ctx in onMissed is passed by Get option.
	onMissed := func(ctx context.Context) (data interface{}, err error) {
		return "newValue", nil
	}

	v, err = cache.Get("newKey", cachego.WithOpOnMissed(onMissed), cachego.WithOpTTL(3*time.Second))
	fmt.Println(v, err) // Output: newValue <nil>

	// We provide a way to set data to cache automatically, so you can access some hottest data extremely fast.
	// See pkg/task/Task.
	t := task.Task{
		Before: func(ctx context.Context) {
			cache.Set("before", "value")
		},
		Fn: func(ctx context.Context) {
			cache.Set("fn", "value")
		},
		After: func(ctx context.Context) {
			cache.Set("after", "value")
		},
	}

	// Run this task automatically every second.
	go t.Run(context.Background(), time.Second)
	time.Sleep(5 * time.Second)
}
```

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)

```bash
$ go test -v ./_examples/performance_test.go
```

> 总缓存数据为 100w 条，并发数为 10w，循环测试写入和读取次数为 50 次

> 测试环境：R7-5800X CPU @ 3.8GHZ GHZ，32 GB RAM，Deepin20 OS

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
