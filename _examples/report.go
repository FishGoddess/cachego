// Copyright 2023 FishGoddess. All Rights Reserved.
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
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/FishGoddess/cachego"
)

func reportMissed(reporter *cachego.Reporter, key string) {
	fmt.Printf("report: missed key %s, missed rate %.3f\n", key, reporter.MissedRate())
}

func reportHit(reporter *cachego.Reporter, key string, value interface{}) {
	fmt.Printf("report: hit key %s value %+v, hit rate %.3f\n", key, value, reporter.HitRate())
}

func reportGC(reporter *cachego.Reporter, cost time.Duration, cleans int) {
	fmt.Printf("report: gc cost %s cleans %d, gc count %d, cache size %d\n", cost, cleans, reporter.CountGC(), reporter.CacheSize())
}

func reportLoad(reporter *cachego.Reporter, key string, value interface{}, ttl time.Duration, err error) {
	fmt.Printf("report: load key %s value %+v ttl %s, err %+v, load count %d\n", key, value, ttl, err, reporter.CountLoad())
}

func main() {
	// We provide some reporting points for monitor cache.
	// ReportMissed reports the missed key getting from cache.
	// ReportHit reports the hit entry getting from cache.
	// ReportGC reports the status of cache gc.
	// ReportLoad reports the result of loading.
	// Use NewCacheWithReport to create a cache with report.
	cache, reporter := cachego.NewCacheWithReport(
		cachego.WithMaxEntries(3),
		cachego.WithGC(100*time.Millisecond),

		cachego.WithReportMissed(reportMissed),
		cachego.WithReportHit(reportHit),
		cachego.WithReportGC(reportGC),
		cachego.WithReportLoad(reportLoad),
	)

	for i := 0; i < 5; i++ {
		key := strconv.Itoa(i)
		evictedValue := cache.Set(key, key, 10*time.Millisecond)
		fmt.Println(evictedValue)
	}

	for i := 0; i < 5; i++ {
		key := strconv.Itoa(i)
		value, ok := cache.Get(key)
		fmt.Println(value, ok)
	}

	time.Sleep(200 * time.Millisecond)

	value, err := cache.Load("key", time.Second, func() (value interface{}, err error) {
		return 666, io.EOF
	})

	fmt.Println(value, err)

	// These are some methods of reporter.
	fmt.Println("CountMissed:", reporter.CountMissed())
	fmt.Println("CountHit:", reporter.CountHit())
	fmt.Println("CountGC:", reporter.CountGC())
	fmt.Println("CountLoad:", reporter.CountLoad())
	fmt.Println("CacheSize:", reporter.CacheSize())
	fmt.Println("MissedRate:", reporter.MissedRate())
	fmt.Println("HitRate:", reporter.HitRate())

	// Sometimes you may have several caches in one service.
	// You can set each name by WithCacheName and get the name from reporter.
	cachego.WithCacheName("test")
	reporter.CacheName()
}
