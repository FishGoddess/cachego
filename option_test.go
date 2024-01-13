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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithCacheName$
func TestWithCacheName(t *testing.T) {
	got := &config{cacheName: ""}
	expect := &config{cacheName: "-"}

	WithCacheName("-").applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithLRU$
func TestWithLRU(t *testing.T) {
	got := &config{cacheType: standard, maxEntries: 0}
	expect := &config{cacheType: lru, maxEntries: 666}

	WithLRU(666).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithLFU$
func TestWithLFU(t *testing.T) {
	got := &config{cacheType: standard, maxEntries: 0}
	expect := &config{cacheType: lfu, maxEntries: 999}

	WithLFU(999).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithShardings$
func TestWithShardings(t *testing.T) {
	got := &config{shardings: 0}
	expect := &config{shardings: 1024}

	WithShardings(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithDisableSingleflight$
func TestWithDisableSingleflight(t *testing.T) {
	got := &config{singleflight: true}
	expect := &config{singleflight: false}

	WithDisableSingleflight().applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithGC$
func TestWithGC(t *testing.T) {
	got := &config{gcDuration: 0}
	expect := &config{gcDuration: 1024}

	WithGC(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithMaxScans$
func TestWithMaxScans(t *testing.T) {
	got := &config{maxScans: 0}
	expect := &config{maxScans: 1024}

	WithMaxScans(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithMaxEntries$
func TestWithMaxEntries(t *testing.T) {
	got := &config{maxEntries: 0}
	expect := &config{maxEntries: 1024}

	WithMaxEntries(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithNow$
func TestWithNow(t *testing.T) {
	now := func() int64 {
		return 0
	}

	got := &config{now: nil}
	expect := &config{now: now}

	WithNow(now).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithHash$
func TestWithHash(t *testing.T) {
	hash := func(key string) int {
		return 0
	}

	got := &config{hash: nil}
	expect := &config{hash: hash}

	WithHash(hash).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithRecordMissed$
func TestWithRecordMissed(t *testing.T) {
	got := &config{recordMissed: false}
	expect := &config{recordMissed: true}

	WithRecordMissed(true).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithRecordHit$
func TestWithRecordHit(t *testing.T) {
	got := &config{recordHit: false}
	expect := &config{recordHit: true}

	WithRecordHit(true).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithRecordGC$
func TestWithRecordGC(t *testing.T) {
	got := &config{recordGC: false}
	expect := &config{recordGC: true}

	WithRecordGC(true).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithRecordLoad$
func TestWithRecordLoad(t *testing.T) {
	got := &config{recordLoad: false}
	expect := &config{recordLoad: true}

	WithRecordLoad(true).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithReportMissed$
func TestWithReportMissed(t *testing.T) {
	reportMissed := func(reporter *Reporter, key string) {}

	got := &config{reportMissed: nil}
	expect := &config{reportMissed: reportMissed}

	WithReportMissed(reportMissed).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithReportHit$
func TestWithReportHit(t *testing.T) {
	reportHit := func(reporter *Reporter, key string, value interface{}) {}

	got := &config{reportHit: nil}
	expect := &config{reportHit: reportHit}

	WithReportHit(reportHit).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithReportGC$
func TestWithReportGC(t *testing.T) {
	reportGC := func(reporter *Reporter, cost time.Duration, cleans int) {}

	got := &config{reportGC: nil}
	expect := &config{reportGC: reportGC}

	WithReportGC(reportGC).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithReportLoad$
func TestWithReportLoad(t *testing.T) {
	reportLoad := func(reporter *Reporter, key string, value interface{}, ttl time.Duration, err error) {}

	got := &config{reportLoad: nil}
	expect := &config{reportLoad: reportLoad}

	WithReportLoad(reportLoad).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Fatalf("got %+v != expect %+v", got, expect)
	}
}
