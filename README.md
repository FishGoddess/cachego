# 📜 cachego

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**cachego** 是一个拥有高性能分段锁机制的轻量级内存缓存，拥有懒清理和哨兵清理两种清理机制，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md).

### 🕹 功能特性

* 以键值对形式缓存数据，极简的 API 设计风格
* 使用粒度更细的分段锁机制进行设计，具有非常高的并发性能
* 支持懒清理机制，每一次访问的时候判断是否过期
* 支持哨兵清理机制，每隔一定的时间间隔进行清理

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

具体设计可以参考 [架构设计介绍](./docs/架构介绍.md) 文档。

### 🚀 安装方式

cachego 没有任何其他额外的依赖，唯一需要的依赖就是 [Golang 运行环境](https://golang.org)。

```bash
$ go get -u github.com/FishGoddess/cachego
```

### 💡 参考案例

```go
package main

import (
	"fmt"

	"github.com/FishGoddess/cachego"
)

func main() {

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
}
```

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)

```bash
$ go test -v ./_examples/performance_test.go
```

> 总缓存数据为 100w 条，并发数为 10w，循环测试写入和读取次数为 50 次

> 测试环境：R7-5800X CPU @ 3.8GHZ GHZ，32 GB RAM

| 测试 | 写入消耗时间 (越小越好) | 读取消耗时间 (越小越好) | 混合操作消耗时间 (越小越好) |
|-----------|-------------|-------------|-------------|
| **cachego** | **965ms** | **949ms** | **991ms** |
| go-cache | 3216ms | 980ms | 4508ms |
| freeCache | 954ms | 968ms | 987ms |

可以看出，由于使用了分段锁机制，读写性能在并发下依然非常高，但是分段锁会多一次定位的操作，如果加锁的消耗小于定位的消耗，那分段锁就不占优势。 这也是为什么 cachego 在写入性能上比 go-cache
强一大截，但是读取性能却没强多少的原因。后续会着重优化读取性能！

### 👥 贡献者

* [cristiane](https://gitee.com/cristiane)：提供 hash 算法的优化建议
* [hzy15610046011](https://gitee.com/hzy15610046011)：提供架构设计文档和图片

如果您觉得 cachego 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。
