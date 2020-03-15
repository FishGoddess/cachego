# 📝 cachego

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**cachego** 是一个轻量级内存型的缓存组件，支持懒清理机制和哨兵清理机制，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md).

### 🥇 功能特性

* 以键值对形式缓存数据，并发访问安全，支持自动清理过期数据
* 使用更细粒度的分段锁机制保证更高的缓存性能（开发中）
* 支持懒清理机制，每一次访问的时候判断是否过期（开发中）
* 支持哨兵清理机制，每隔一定的时间间隔进行清理过期数据（开发中）
* 支持内存大小限制，防止无上限的使用内存（开发中）
* 支持缓存个数限制，防止数据量太多导致哈希性能下降（开发中）
* 支持用户自定义达到内存限制时的处理策略（开发中）
* 支持用户自定义达到个数限制时的处理策略（开发中）

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
// TODO 使用案例
```

### 📖 参考案例

* 开发中...

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=10s
```

> 测试文件：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

> 写入缓存和读取缓存并发进行，将缓存 GC 时间设置为 5 秒/次，总缓存数据为 100 万条

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | 功能性 | 扩展性 |
| -----------|--------|-------------|-------------|-------------|
| **cachego** | 98780428 | 111 ns/op | 强大 | 高 |

> 测试环境：I7-6700HQ CPU @ 2.6 GHZ，16 GB RAM

### 👥 贡献者

如果您觉得 cachego 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 cachego 的项目

| 项目 | 作者 | 描述 |
| -----------|--------|-------------|
|  |  |  |

