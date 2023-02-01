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
	"fmt"
	"testing"
	"time"
)

func isConfigEquals(conf1 *config, conf2 *config) bool {
	if conf1.cacheType != conf2.cacheType {
		return false
	}

	if conf1.shardings != conf2.shardings {
		return false
	}

	if conf1.singleflight != conf2.singleflight {
		return false
	}

	if conf1.gcDuration != conf2.gcDuration {
		return false
	}

	if conf1.maxScans != conf2.maxScans {
		return false
	}

	if conf1.maxEntries != conf2.maxEntries {
		return false
	}

	if fmt.Sprintf("%p", conf1.reportMissed) != fmt.Sprintf("%p", conf2.reportMissed) {
		return false
	}

	if fmt.Sprintf("%p", conf1.reportHit) != fmt.Sprintf("%p", conf2.reportHit) {
		return false
	}

	if fmt.Sprintf("%p", conf1.reportGC) != fmt.Sprintf("%p", conf2.reportGC) {
		return false
	}

	if fmt.Sprintf("%p", conf1.reportLoad) != fmt.Sprintf("%p", conf2.reportLoad) {
		return false
	}

	return true
}

// go test -v -cover -run=^TestApplyOptions$
func TestApplyOptions(t *testing.T) {
	got := &config{
		shardings:    0,
		singleflight: true,
		gcDuration:   0,
		maxScans:     0,
		maxEntries:   0,
	}

	expect := &config{
		shardings:    1,
		singleflight: false,
		gcDuration:   2,
		maxScans:     3,
		maxEntries:   4,
	}

	applyOptions(got, []Option{
		WithShardings(1),
		WithDisableSingleflight(),
		WithGC(2),
		WithMaxScans(3),
		WithMaxEntries(4),
	})

	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

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

// go test -v -cover -run=^TestWithReportMissed$
func TestWithReportMissed(t *testing.T) {
	reportMissed := func(key string) {}

	got := &config{reportMissed: nil}
	expect := &config{reportMissed: reportMissed}

	WithReportMissed(reportMissed).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReportHit$
func TestWithReportHit(t *testing.T) {
	reportHit := func(key string, value interface{}) {}

	got := &config{reportHit: nil}
	expect := &config{reportHit: reportHit}

	WithReportHit(reportHit).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReportGC$
func TestWithReportGC(t *testing.T) {
	reportGC := func(cost time.Duration, cleans int) {}

	got := &config{reportGC: nil}
	expect := &config{reportGC: reportGC}

	WithReportGC(reportGC).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithReportLoad$
func TestWithReportLoad(t *testing.T) {
	reportLoad := func(key string, value interface{}, ttl time.Duration, err error) {}

	got := &config{reportLoad: nil}
	expect := &config{reportLoad: reportLoad}

	WithReportLoad(reportLoad).applyTo(got)
	if !isConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}
