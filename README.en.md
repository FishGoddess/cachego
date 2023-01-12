# 📝 cachego

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/cachego)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/cachego/actions/workflows/test.yml/badge.svg)

**cachego** is a high-performance and memory-based cache for [GoLang](https://golang.org) applications.

> It has been used by many services in production, and even 17w/s qps is ok for it, so just use it if you want!

[阅读中文版的 Read me](./README.md).

### 🕹 Features

* Cache as entries with minimalist API design
* Use option function mode to customize the operations of cache
* Use fine-grained and segmented lock mechanism to provide a high performance in concurrency
* Lazy cleanup supports, expired before accessing
* Sentinel cleanup supports, cleaning up at fixed duration
* Singleflight supports, which can decrease the times of cache penetration
* ....

_More features in [_examples](_examples) and designing detail in [introduction.md](_examples/docs/introduction.md)._

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to get more information._

### 🚀 Installation

cachego has no more external dependencies, the only requirement is the [Golang Programming Language](https://golang.org)
.

```bash
$ go get -u github.com/FishGoddess/cachego
```

### 💡 Examples

```go
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
