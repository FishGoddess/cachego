// Copyright 2025 FishGoddess. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/FishGoddess/cachego"
)

const (
	benchTTL        = time.Minute
	benchMaxKeys    = 10000
	benchMaxEntries = 100000
)

type benchKeys []string

func newBenchKeys() benchKeys {
	keys := make([]string, 0, benchMaxKeys)

	for i := 0; i < benchMaxKeys; i++ {
		keys = append(keys, strconv.Itoa(i))
	}

	return keys
}

func (bks benchKeys) pick() string {
	index := rand.Intn(len(bks))
	return bks[index]
}

func (bks benchKeys) loop(fn func(key string)) {
	for _, key := range bks {
		fn(key)
	}
}

func benchmarkCacheGet(b *testing.B, set func(key string, value string), get func(key string)) {
	keys := newBenchKeys()
	keys.loop(func(key string) {
		set(key, key)
	})

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		key := keys.pick()

		for pb.Next() {
			get(key)
		}
	})
}

func benchmarkCacheSet(b *testing.B, set func(key string, value string)) {
	keys := newBenchKeys()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		key := keys.pick()

		for pb.Next() {
			set(key, key)
		}
	})
}

// go test -v -bench=^BenchmarkCachegoGet$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoGet(b *testing.B) {
	cache := cachego.NewCache()

	set := func(key string, value string) {
		cache.Set(key, value, benchTTL)
	}

	get := func(key string) {
		cache.Get(key)
	}

	benchmarkCacheGet(b, set, get)
}

// go test -v -bench=^BenchmarkCachegoGetLRU$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoGetLRU(b *testing.B) {
	cache := cachego.NewCache(cachego.WithLRU(benchMaxEntries))

	set := func(key string, value string) {
		cache.Set(key, value, benchTTL)
	}

	get := func(key string) {
		cache.Get(key)
	}

	benchmarkCacheGet(b, set, get)
}

// go test -v -bench=^BenchmarkCachegoGetLFU$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoGetLFU(b *testing.B) {
	cache := cachego.NewCache(cachego.WithLFU(benchMaxEntries))

	set := func(key string, value string) {
		cache.Set(key, value, benchTTL)
	}

	get := func(key string) {
		cache.Get(key)
	}

	benchmarkCacheGet(b, set, get)
}

// go test -v -bench=^BenchmarkCachegoGetSharding$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoGetSharding(b *testing.B) {
	cache := cachego.NewCache(cachego.WithShardings(16))

	set := func(key string, value string) {
		cache.Set(key, value, benchTTL)
	}

	get := func(key string) {
		cache.Get(key)
	}

	benchmarkCacheGet(b, set, get)
}

// go test -v -bench=^BenchmarkCachegoSet$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoSet(b *testing.B) {
	cache := cachego.NewCache()

	benchmarkCacheSet(b, func(key string, value string) {
		cache.Set(key, value, benchTTL)
	})
}

// go test -v -bench=^BenchmarkCachegoSetLRU$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoSetLRU(b *testing.B) {
	cache := cachego.NewCache(cachego.WithLRU(benchMaxEntries))

	benchmarkCacheSet(b, func(key string, value string) {
		cache.Set(key, value, benchTTL)
	})
}

// go test -v -bench=^BenchmarkCachegoSetLFU$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoSetLFU(b *testing.B) {
	cache := cachego.NewCache(cachego.WithLFU(benchMaxEntries))

	benchmarkCacheSet(b, func(key string, value string) {
		cache.Set(key, value, benchTTL)
	})
}

// go test -v -bench=^BenchmarkCachegoSetSharding$ -benchtime=1s ./_examples/performance_test.go
func BenchmarkCachegoSetSharding(b *testing.B) {
	cache := cachego.NewCache(cachego.WithShardings(16))

	benchmarkCacheSet(b, func(key string, value string) {
		cache.Set(key, value, benchTTL)
	})
}

//// go test -v -bench=^BenchmarkGcacheGet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGcacheGet(b *testing.B) {
//	cache := gcache.New(benchMaxEntries).Expiration(benchTTL).Build()
//
//	set := func(key string, value string) {
//		cache.Set(key, value)
//	}
//
//	get := func(key string) {
//		cache.Get(key)
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkGcacheGetLRU$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGcacheGetLRU(b *testing.B) {
//	cache := gcache.New(benchMaxEntries).Expiration(benchTTL).LRU().Build()
//
//	set := func(key string, value string) {
//		cache.Set(key, value)
//	}
//
//	get := func(key string) {
//		cache.Get(key)
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkGcacheGetLFU$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGcacheGetLFU(b *testing.B) {
//	cache := gcache.New(benchMaxEntries).Expiration(benchTTL).LFU().Build()
//
//	set := func(key string, value string) {
//		cache.Set(key, value)
//	}
//
//	get := func(key string) {
//		cache.Get(key)
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkEcacheGet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkEcacheGet(b *testing.B) {
//	cache := ecache.NewLRUCache(1, math.MaxUint16, benchTTL)
//
//	set := func(key string, value string) {
//		cache.Put(key, value)
//	}
//
//	get := func(key string) {
//		cache.Get(key)
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkEcache2Get$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkEcache2Get(b *testing.B) {
//	cache := ecache.NewLRUCache(1, math.MaxUint16, benchTTL).LRU2(16)
//
//	set := func(key string, value string) {
//		cache.Put(key, value)
//	}
//
//	get := func(key string) {
//		cache.Get(key)
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkBigcacheGet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkBigcacheGet(b *testing.B) {
//	cache, _ := bigcache.New(context.Background(), bigcache.Config{
//		Shards:             1,
//		LifeWindow:         benchTTL,
//		MaxEntriesInWindow: benchMaxEntries,
//		Verbose:            false,
//	})
//
//	set := func(key string, value string) {
//		cache.Set(key, []byte(value))
//	}
//
//	get := func(key string) {
//		cache.Get(key)
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkFreecacheGet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkFreecacheGet(b *testing.B) {
//	cache := freecache.NewCache(benchMaxEntries)
//
//	set := func(key string, value string) {
//		cache.Set([]byte(key), []byte(value), int(benchTTL.Seconds()))
//	}
//
//	get := func(key string) {
//		cache.Get([]byte(key))
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkGoCacheGet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGoCacheGet(b *testing.B) {
//	cache := gocache.New(benchTTL, 0)
//
//	set := func(key string, value string) {
//		cache.Set(key, value, benchTTL)
//	}
//
//	get := func(key string) {
//		cache.Get(key)
//	}
//
//	benchmarkCacheGet(b, set, get)
//}
//
//// go test -v -bench=^BenchmarkGcacheSet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGcacheSet(b *testing.B) {
//	cache := gcache.New(benchMaxEntries).Expiration(benchTTL).Build()
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Set(key, value)
//	})
//}
//
//// go test -v -bench=^BenchmarkGcacheSetLRU$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGcacheSetLRU(b *testing.B) {
//	cache := gcache.New(benchMaxEntries).Expiration(benchTTL).LRU().Build()
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Set(key, value)
//	})
//}
//
//// go test -v -bench=^BenchmarkGcacheSetLFU$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGcacheSetLFU(b *testing.B) {
//	cache := gcache.New(benchMaxEntries).Expiration(benchTTL).LFU().Build()
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Set(key, value)
//	})
//}
//
//// go test -v -bench=^BenchmarkEcacheSet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkEcacheSet(b *testing.B) {
//	cache := ecache.NewLRUCache(1, math.MaxUint16, benchTTL)
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Put(key, value)
//	})
//}
//
//// go test -v -bench=^BenchmarkEcache2Set$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkEcache2Set(b *testing.B) {
//	cache := ecache.NewLRUCache(1, math.MaxUint16, benchTTL).LRU2(16)
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Put(key, value)
//	})
//}
//
//// go test -v -bench=^BenchmarkBigcacheSet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkBigcacheSet(b *testing.B) {
//	cache, _ := bigcache.New(context.Background(), bigcache.Config{
//		Shards:             1,
//		LifeWindow:         benchTTL,
//		MaxEntriesInWindow: benchMaxEntries,
//		Verbose:            false,
//	})
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Set(key, []byte(value))
//	})
//}
//
//// go test -v -bench=^BenchmarkFreecacheSet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkFreecacheSet(b *testing.B) {
//	cache := freecache.NewCache(benchMaxEntries)
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Set([]byte(key), []byte(value), int(benchTTL.Seconds()))
//	})
//}
//
//// go test -v -bench=^BenchmarkGoCacheSet$ -benchtime=1s ./_examples/performance_test.go
//func BenchmarkGoCacheSet(b *testing.B) {
//	cache := gocache.New(benchTTL, 0)
//
//	benchmarkCacheSet(b, func(key string, value string) {
//		cache.Set(key, value, benchTTL)
//	})
//}
