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
	"math"
	"math/rand"
	"testing"
	"time"
)

// go test -bench=^BenchmarkTimeNow$ -benchtime=1s ./clock.go ./clock_test.go
func BenchmarkTimeNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		time.Now().Nanosecond()
	}
}

// go test -bench=^BenchmarkClockNow$ -benchtime=1s ./clock.go ./clock_test.go
func BenchmarkClockNow(b *testing.B) {
	clock := New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		clock.Now()
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNew$
func TestNew(t *testing.T) {
	var clocks []*Clock

	for i := 0; i < 100; i++ {
		clocks = append(clocks, New())
	}

	for i := 0; i < 100; i++ {
		if clocks[i] != clock {
			t.Fatalf("clocks[i] %p != clock %p", clocks[i], clock)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestClock$
func TestClock(t *testing.T) {
	testClock := New()

	for i := 0; i < 100; i++ {
		now := testClock.Now()
		t.Log(now)

		expect := time.Now().UnixNano()
		if math.Abs(float64(expect-now)) > float64(duration)*1.5 {
			t.Fatalf("now %d is wrong with expect %d", now, expect)
		}

		time.Sleep(time.Duration(rand.Int63n(int64(duration))))
	}
}
