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

func reportMissed(key string) {
	fmt.Printf("report: missed key %s\n", key)
}

func reportHit(key string, value interface{}) {
	fmt.Printf("report: hit key %s value %+v\n", key, value)
}

func reportGC(cost time.Duration, cleans int) {
	fmt.Printf("report: gc cost %s cleans %d\n", cost, cleans)
}

func reportLoad(key string, value interface{}, ttl time.Duration, err error) {
	fmt.Printf("report: load key %s value %+v ttl %s, err %+v\n", key, value, ttl, err)
}

func main() {
	// Create a cache as usual.
	cache := cachego.NewCache(
		cachego.WithMaxEntries(3),
		cachego.WithGC(100*time.Millisecond),
	)

	// Use Report function to wrap a cache with reporting logics.
	// We provide some reporting points for monitor cache.
	// ReportMissed reports the missed key getting from cache.
	// ReportHit reports the hit entry getting from cache.
	// ReportGC reports the status of cache gc.
	// ReportLoad reports the result of loading.
	cache, reporter := cachego.Report(
		cache,
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
	fmt.Println("CacheSize:", reporter.CacheSize())
	fmt.Println("MissedRate:", reporter.MissedRate())
	fmt.Println("HitRate:", reporter.HitRate())
}
