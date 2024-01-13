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

package cachego

import (
	"testing"
	"time"
)

// go test -v -bench=^BenchmarkHash$ -benchtime=1s ./global.go ./global_test.go
func BenchmarkHash(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hash("key")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestHash$
func TestHash(t *testing.T) {
	hash := hash("test")
	if hash < 0 {
		t.Fatalf("hash %d <= 0", hash)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNow$
func TestNow(t *testing.T) {
	got := now()
	expect := time.Now().UnixNano()

	if got > expect || got < expect-testDurationGap.Nanoseconds() {
		t.Fatalf("got %d != expect %d", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestSetMapInitialCap$
func TestSetMapInitialCap(t *testing.T) {
	oldInitialCap := mapInitialCap

	SetMapInitialCap(-2)
	if mapInitialCap != oldInitialCap {
		t.Fatalf("mapInitialCap %d is wrong", mapInitialCap)
	}

	SetMapInitialCap(0)
	if mapInitialCap != oldInitialCap {
		t.Fatalf("mapInitialCap %d is wrong", mapInitialCap)
	}

	SetMapInitialCap(2)
	if mapInitialCap != 2 {
		t.Fatalf("mapInitialCap %d is wrong", mapInitialCap)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestSetSliceInitialCap$
func TestSetSliceInitialCap(t *testing.T) {
	oldInitialCap := sliceInitialCap

	SetSliceInitialCap(-2)
	if sliceInitialCap != oldInitialCap {
		t.Fatalf("sliceInitialCap %d is wrong", sliceInitialCap)
	}

	SetSliceInitialCap(0)
	if sliceInitialCap != oldInitialCap {
		t.Fatalf("sliceInitialCap %d is wrong", sliceInitialCap)
	}

	SetSliceInitialCap(2)
	if sliceInitialCap != 2 {
		t.Fatalf("sliceInitialCap %d is wrong", sliceInitialCap)
	}
}
