# 📝 cachego

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**cachego** 是一个轻量级内存型并支持链式编程的缓存组件，拥有懒清理和哨兵清理两种机制，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md).

### 🥇 功能特性

* 以键值对形式缓存数据，并发访问安全，支持自动清理过期数据
* 基础特性和高级特性分离设计模式，减少新用户学习上手难度
* 链式编程友好的 API 设计，在一定程度上提供了很高的代码可读性
* 支持懒清理机制，每一次访问的时候判断是否过期
* 支持哨兵清理机制，每隔一定的时间间隔进行清理过期数据
* 支持内存大小限制，防止无上限的使用内存（开发中）
* 支持缓存个数限制，防止数据量太多导致哈希性能下降（开发中）
* 支持用户自定义达到内存限制时的处理策略（开发中）
* 支持用户自定义达到个数限制时的处理策略（开发中）
* 使用更细粒度的分段锁机制保证更高的缓存性能（开发中）

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

唯一需要的依赖就是 [Golang 运行环境](https://golang.org).

> Go modules

```bash
$ go get -u github.com/FishGoddess/cachego
```

您也可以直接编辑 go.mod 文件，然后执行 _**go build**_。

```bash
module your_project_name

go 1.14

require (
    github.com/FishGoddess/cachego v0.0.1
)
```

> Go path

```bash
$ go get -u github.com/FishGoddess/cachego
```

cachego 没有任何其他额外的依赖，纯使用 [Golang 标准库](https://golang.org) 完成。

```go
package main

import (
    "fmt"
    "time"

    cache "github.com/FishGoddess/cachego"
)

func main() {

    // Create a cache with default gc duration (10 minutes).
    newCache := cache.NewCache()

    // Put a new entry in cache.
    // This entry will be dead after 5 seconds.
    // However, it will be deleted after 10 minutes if you never access.
    newCache.Put("key", 666, 5 * time.Second)

    // Of returns the value of this key.
    // As you know, this is chain-programming api.
    // If you need int type, just call Int().
    v := newCache.Of("key").Int()
    fmt.Println(v) // Output: 666

    // If you want change the value of key, try this:
    newCache.Change("key", "value")

    // Then you can call String() behind Of().
    s := newCache.Of("key").String()
    fmt.Println(s) // Output: value

    // After 5 seconds, this entry will dead, then an invalidCacheValue will be returned.
    time.Sleep(5 * time.Second)
    ok := newCache.Of("key").Ok()
    fmt.Println(ok) // Output: false

    // Maybe you want a default value for some situations, such as the code above.
    // Use Or() to help you to do that:
    s = newCache.Of("key").Or("default value").String()
    fmt.Println(s) // Output: default value
}
```

### 📖 参考案例

* [basic](./_examples/basic.go)
* [cache_value](./_examples/cache_value.go)

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=12s
```

> 测试文件：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

> 写入缓存和读取缓存并发进行，将缓存 GC 时间设置为 5 秒/次，总缓存数据为 100 万条

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | 功能性 | 扩展性 |
| -----------|--------|-------------|-------------|-------------|
| **cachego** | 127152241 | 104 ns/op | 强大 | 高 |
| freeCache | 132629332 | 107 ns/op | 正常 | 正常 |
| go-cache | 276515510 | &nbsp; 44 ns/op | 正常 | 正常 |

> 测试环境：I7-6700HQ CPU @ 2.6 GHZ，16 GB RAM

注意：
1. freeCache 的过期时间远大于 cachego 和 go-cache，也就意味着单次 GC 的量要少得多，
所以这个结果应该是偏好的，实际生产环境可能要更差（个人想法）。

### 👥 贡献者

如果您觉得 cachego 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 cachego 的项目

| 项目 | 作者 | 描述 |
| -----------|--------|-------------|
|  |  |  |

