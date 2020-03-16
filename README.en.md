# 📝 cachego

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**cachego** is a chain-programming and memory-based cache for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🥇 Features

* Cache as key-value and for concurrency, automatically clean up expired data supports
* Splitting basic features and advanced features design, more friendly to new users
* Chain-programming api supports, more readability to all users
* Lazy cleanup mechanism supports, clean up expired data before visiting
* Sentinel cleanup mechanism supports, clean up expired data at fixed interval
* Memory limit supports, to protect memory from unlimited uses (coming soon)
* Cache count limit supports, to protect hash performance from too many data (coming soon)
* Memory limit strategy supports, you can customize your strategy to handle memory exceeding (coming soon)
* Cache count limit strategy supports, you can customize your strategy to handle cache count exceeding (coming soon)
* Use more fine-grained segmented lock mechanism to guarantee higher performance (coming soon)

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

### 📖 Examples

* [basic](./_examples/basic.go)
* [cache_value](./_examples/cache_value.go)

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=12s
```

> Benchmark file：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

> Write and read are processing concurrently, gc duration is set to 5 seconds, total is 1 million data

| test case | times ran (large is better) |  ns/op (small is better) | features | extension |
| -----------|--------|-------------|-------------|-------------|
| **cachego** | 127152241 | 104 ns/op | powerful | high |
| freeCache | 132629332 | 107 ns/op | normal | normal |
| go-cache | 276515510 | &nbsp; 44 ns/op | normal | normal |

> Environment：I7-6700HQ CPU @ 2.6 GHZ, 16 GB RAM

Notice:
1. freeCache has a bigger expired time so one gc task does less work. So
it is not fair to cachego and go-cache. Actually, freeCache may slower than 
this benchmark shown (my personal thought).

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using cachego

| Project | Author | Description |
| -----------|--------|-------------|
|  |  |  |

