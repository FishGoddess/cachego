# 📝 cachego

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**cachego** is a high-performance and memory-based cache for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🕹 Features

* Cache as entries with minimalist API design
* Use option function mode to customize the creation of cache
* Provide debug point for developing and checking cache
* Use fine-grained and segmented lock mechanism to provide a high performance in concurrency
* Lazy cleanup supports, expired before accessing
* Sentinel cleanup supports, cleaning up at fixed duration

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to get more information._

See more designing detail in [architecture design introduction](./docs/架构介绍.md).

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
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {

	// Create a cache for use.
	// We use option function to customize the creation of cache.
	// WithAutoGC means it will do gc automatically.
	cache := cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Set a new entry to cache.
	cache.Set("key", 666)

	// Get returns the value of this key.
	v, ok := cache.Get("key")
	fmt.Println(v, ok) // Output: 666 true

	// If you pass a not existed key to of method, nil and false will be returned.
	v, ok = cache.Get("not existed key")
	fmt.Println(v, ok) // Output: <nil> false

	// SetWithTTL sets an entry with expired time.
	// The unit of expired time is second.
	// See more information in example of ttl.
	cache.SetWithTTL("ttlKey", 123, 10)
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
* [hzy15610046011](https://gitee.com/hzy15610046011): Provide architecture design documents and pictures

Please open an _**issue**_ if you find something is not working as expected.
