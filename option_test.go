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

// go test -v -cover -run=^TestWithLRU$
func TestWithLRU(t *testing.T) {
	got := &config{cacheType: standard, maxEntries: 0}
	expect := &config{cacheType: lru, maxEntries: 666}

	WithLRU(666).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithLFU$
func TestWithLFU(t *testing.T) {
	got := &config{cacheType: standard, maxEntries: 0}
	expect := &config{cacheType: lfu, maxEntries: 999}

	WithLFU(999).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithShardings$
func TestWithShardings(t *testing.T) {
	got := &config{shardings: 0}
	expect := &config{shardings: 1024}

	WithShardings(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithDisableSingleflight$
func TestWithDisableSingleflight(t *testing.T) {
	got := &config{singleflight: true}
	expect := &config{singleflight: false}

	WithDisableSingleflight().applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithGC$
func TestWithGC(t *testing.T) {
	got := &config{gcDuration: 0}
	expect := &config{gcDuration: 1024}

	WithGC(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithMaxScans$
func TestWithMaxScans(t *testing.T) {
	got := &config{maxScans: 0}
	expect := &config{maxScans: 1024}

	WithMaxScans(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithMaxEntries$
func TestWithMaxEntries(t *testing.T) {
	got := &config{maxEntries: 0}
	expect := &config{maxEntries: 1024}

	WithMaxEntries(1024).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithNow$
func TestWithNow(t *testing.T) {
	now := func() int64 {
		return 0
	}

	got := &config{now: nil}
	expect := &config{now: now}

	WithNow(now).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithHash$
func TestWithHash(t *testing.T) {
	hash := func(key string) int {
		return 0
	}

	got := &config{hash: nil}
	expect := &config{hash: hash}

	WithHash(hash).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReporterNow$
func TestWithReporterNow(t *testing.T) {
	now := func() int64 {
		return 0
	}

	got := &reportConfig{now: nil}
	expect := &reportConfig{now: now}

	WithReporterNow(now).applyTo(got)
	if !isReportConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReportMissed$
func TestWithReportMissed(t *testing.T) {
	reportMissed := func(key string) {}

	got := &reportConfig{reportMissed: nil}
	expect := &reportConfig{reportMissed: reportMissed}

	WithReportMissed(reportMissed).applyTo(got)
	if !isReportConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReportHit$
func TestWithReportHit(t *testing.T) {
	reportHit := func(key string, value interface{}) {}

	got := &reportConfig{reportHit: nil}
	expect := &reportConfig{reportHit: reportHit}

	WithReportHit(reportHit).applyTo(got)
	if !isReportConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReportGC$
func TestWithReportGC(t *testing.T) {
	reportGC := func(cost time.Duration, cleans int) {}

	got := &reportConfig{reportGC: nil}
	expect := &reportConfig{reportGC: reportGC}

	WithReportGC(reportGC).applyTo(got)
	if !isReportConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReportLoad$
func TestWithReportLoad(t *testing.T) {
	reportLoad := func(key string, value interface{}, ttl time.Duration, err error) {}

	got := &reportConfig{reportLoad: nil}
	expect := &reportConfig{reportLoad: reportLoad}

	WithReportLoad(reportLoad).applyTo(got)
	if !isReportConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}
