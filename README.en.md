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
