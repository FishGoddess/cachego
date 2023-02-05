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

package clock

import (
	"sync/atomic"
	"time"
)

const (
	duration = 100 * time.Millisecond
)

// Clock is a fast clock for getting current time.
// It caches time and updates it in fixed duration which may return an "incorrect" time compared with time.Now().UnixNano().
// In fact, we don't recommend you to use it unless you have to...
// According to our benchmarks, it does run faster than time.Now:
//
// In my win10 pc:
// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 5800X 8-Core Processor
// BenchmarkTimeNow-16             338466458                3.523 ns/op           0 B/op          0 allocs/op
// BenchmarkClockNow-16            1000000000               0.2165 ns/op          0 B/op          0 allocs/op
//
// In my macbook with charging:
// goos: darwin
// goarch: amd64
// pkg: github.com/FishGoddess/cachego/pkg/clock
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkTimeNow-12             17841402                67.05 ns/op            0 B/op          0 allocs/op
// BenchmarkClockNow-12            1000000000               0.2528 ns/op          0 B/op          0 allocs/op
//
// In my cloud server using 2 cores and running benchmarks in docker container:
// goos: linux
// goarch: amd64
// pkg: github.com/FishGoddess/cachego/pkg/clock
// cpu: AMD EPYC 7K62 48-Core Processor
// BenchmarkTimeNow-2      17946441                65.62 ns/op            0 B/op          0 allocs/op
// BenchmarkClockNow-2     1000000000               0.3162 ns/op          0 B/op          0 allocs/op
//
// PS: All benchmarks are ran with "go test -bench=. -benchtime=1s".
//
// However, the performance of time.Now is faster enough in many os and is enough for 99.9% situations.
// The another reason choosing to use it is .
// So, better performance should not be the first reason to use it.
// The first reason to use it is reducing gc objects, but we hope you never use it :)
type Clock struct {
	now int64
}

// New creates a new clock which caches time and updates it in fixed duration.
func New() *Clock {
	clock := &Clock{
		now: time.Now().UnixNano(),
	}

	go clock.start()
	return clock
}

func (c *Clock) start() {
	for {
		for i := 0; i < 9; i++ {
			time.Sleep(duration)
			atomic.AddInt64(&c.now, int64(duration))
		}

		time.Sleep(100 * time.Millisecond)
		atomic.StoreInt64(&c.now, time.Now().UnixNano())
	}
}

// Now returns the current time in nanoseconds.
func (c *Clock) Now() int64 {
	return atomic.LoadInt64(&c.now)
}
