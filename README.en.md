# 📝 cachego

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**cachego** is a easy-to-use and memory-based cache for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🥇 Features

* Cache as key-value and for concurrency, automatically clean up expired data supports
* Use more fine-grained segmented lock mechanism to guarantee higher performance (coming soon)
* Lazy cleanup mechanism supports, clean up expired data before visiting (coming soon)
* Sentinel cleanup mechanism supports, clean up expired data at fixed interval (coming soon)
* Memory limit supports, to protect memory from unlimited uses (coming soon)
* Cache count limit supports, to protect hash performance from too many data (coming soon)
* Memory limit strategy supports, you can customize your strategy to handle memory exceeding (coming soon)
* Cache count limit strategy supports, you can customize your strategy to handle cache count exceeding (coming soon)

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to get more information._

### 🚀 Installation

The only requirement is the [Golang Programming Language](https://golang.org).

> Go modules

```bash
$ go get -u github.com/FishGoddess/cachego
```

Or edit your project's go.mod file and execute _**go build**_.

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

cachego has no more external dependencies.

```go
// TODO Using case
```

### 📖 Examples

* coming soon...

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=1s
```

> Benchmark file：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| test case | times ran (large is better) |  ns/op (small is better) | features | extension |
| -----------|--------|-------------|-------------|-------------|
| **cachego** | unknown | unknown | powerful | high |

> Environment：I7-6700HQ CPU @ 2.6 GHZ, 16 GB RAM

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using cachego

| Project | Author | Description |
| -----------|--------|-------------|
|  |  |  |

