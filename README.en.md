# 📝 cachego

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**cachego** is a high-performance and memory-based cache for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🕹 Features

* Cache as entries with minimalist API design
* Use fine-grained and segmented lock mechanism to provide a high performance in concurrency
* Lazy cleanup supports, expired before accessing
* Sentinel cleanup supports, cleaning up at fixed duration

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to get more information._

### 🚀 Installation

cachego has no more external dependencies, the only requirement is the [Golang Programming Language](https://golang.org)
.

```bash
$ go get -u github.com/FishGoddess/cachego
```

### 💡 Examples

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

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

> Benchmark file：[_examples/performance_test.go](./_examples/performance_test.go)

```bash
$ go test -v ./_examples/performance_test.go
```

> Data size is 1 million, concurrency is 100 thousands, loop is 50

> Environment：R7-5800X CPU @ 3.8GHZ GHZ, 32 GB RAM

| tests | write time (less is better) | read time (less is better) | mixed-operation time (less is better) |
|-----------|-------------|-------------|-------------|
| **cachego** | **965ms** | **949ms** | **991ms** |
| go-cache | 3216ms | 980ms | 4508ms |
| freeCache | 954ms | 968ms | 987ms |

As you can see, cachego has a high performance in concurrent, but segmented lock mechanism has one-more-time positioning
operation, so if the price of locking is less than the cost of positioning, this mechanism is dragging. The reading
performance will be optimized in the future version!

### 👥 Contributors

* [cristiane](https://gitee.com/cristiane): Provide some optimizations about hash

If you find that something is not working as expected please open an _**issue**_.
