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
	"io"
	"testing"
	"time"
)

func newTestReportableCache() (*reportableCache, *Reporter) {
	conf := newDefaultConfig()
	conf.maxEntries = maxTestEntries
	cache, reporter := Report(newStandardCache(conf), WithRecordMissed(true), WithRecordHit(true), WithRecordGC(true), WithRecordLoad(true))
	return cache.(*reportableCache), reporter
}

// go test -v -cover -run=^TestReportableCache$
func TestReportableCache(t *testing.T) {
	cache, _ := newTestReportableCache()
	testCacheImplement(t, cache)
}

// go test -v -cover -run=^TestReportableCacheReportMissed$
func TestReportableCacheReportMissed(t *testing.T) {
	cache, reporter := newTestReportableCache()
	cache.Set("key", 666, NoTTL)

	checked := false
	cache.reportMissed = func(reporter *Reporter, key string) {
		if key == "key" {
			t.Error("key == \"key\"")
		}

		if key != "missed" {
			t.Errorf("key %s != \"missed\"", key)
		}

		checked = true
	}

	cache.Get("key")
	cache.Get("missed")

	if !checked {
		t.Error("reportMissed not checked")
	}

	if reporter.CountMissed() != 1 {
		t.Errorf("CountMissed %d is wrong", reporter.CountMissed())
	}

	missedRate := reporter.MissedRate()
	if missedRate < 0.499 || missedRate > 0.501 {
		t.Errorf("missedRate %.3f is wrong", missedRate)
	}
}

// go test -v -cover -run=^TestReportableCacheReportHit$
func TestReportableCacheReportHit(t *testing.T) {
	cache, reporter := newTestReportableCache()
	cache.Set("key", 666, NoTTL)

	checked := false
	cache.reportHit = func(reporter *Reporter, key string, value interface{}) {
		if key == "missed" {
			t.Error("key == \"missed\"")
		}

		if key != "key" {
			t.Errorf("key %s != \"key\"", key)
		}

		if value.(int) != 666 {
			t.Errorf("value.(int) %d is wrong", value.(int))
		}

		checked = true
	}

	cache.Get("key")
	cache.Get("missed")

	if !checked {
		t.Error("reportHit not checked")
	}

	if reporter.CountHit() != 1 {
		t.Errorf("CountHit %d is wrong", reporter.CountHit())
	}

	hitRate := reporter.HitRate()
	if hitRate < 0.499 || hitRate > 0.501 {
		t.Errorf("hitRate %.3f is wrong", hitRate)
	}
}

// go test -v -cover -run=^TestReportableCacheReportGC$
func TestReportableCacheReportGC(t *testing.T) {
	cache, reporter := newTestReportableCache()
	cache.Set("key1", 1, time.Millisecond)
	cache.Set("key2", 2, time.Millisecond)
	cache.Set("key3", 3, time.Millisecond)
	cache.Set("key4", 4, time.Second)
	cache.Set("key5", 5, time.Second)

	gcCount := uint64(0)
	checked := false
	cache.reportGC = func(reporter *Reporter, cost time.Duration, cleans int) {
		if cost <= 0 {
			t.Errorf("cost %d <= 0", cost)
		}

		if cleans != 3 {
			t.Errorf("cleans %d is wrong", cleans)
		}

		gcCount++
		checked = true
	}

	time.Sleep(10 * time.Millisecond)

	cleans := cache.GC()
	if cleans != 3 {
		t.Errorf("cleans %d is wrong", cleans)
	}

	if !checked {
		t.Error("reportHit not checked")
	}

	if reporter.CountGC() != gcCount {
		t.Errorf("CountGC %d is wrong", reporter.CountGC())
	}
}

// go test -v -cover -run=^TestReportableCacheReportLoad$
func TestReportableCacheReportLoad(t *testing.T) {
	cache, reporter := newTestReportableCache()

	loadCount := uint64(0)
	checked := false
	cache.reportLoad = func(reporter *Reporter, key string, value interface{}, ttl time.Duration, err error) {
		if key != "load" {
			t.Errorf("key %s is wrong", key)
		}

		if value.(int) != 999 {
			t.Errorf("value.(int) %d is wrong", value.(int))
		}

		if ttl != time.Second {
			t.Errorf("ttl %s is wrong", ttl)
		}

		if err != io.EOF {
			t.Errorf("err %+v is wrong", err)
		}

		loadCount++
		checked = true
	}

	value, err := cache.Load("load", time.Second, func() (value interface{}, err error) {
		return 999, io.EOF
	})

	if value.(int) != 999 {
		t.Errorf("value.(int) %d is wrong", value.(int))
	}

	if err != io.EOF {
		t.Errorf("err %+v is wrong", err)
	}

	if !checked {
		t.Error("reportLoad not checked")
	}

	if reporter.CountLoad() != loadCount {
		t.Errorf("CountLoad %d is wrong", reporter.CountLoad())
	}
}

// go test -v -cover -run=^TestReporterCacheSize$
func TestReporterCacheSize(t *testing.T) {
	cache, reporter := newTestReportableCache()
	cache.Set("key1", 1, time.Millisecond)
	cache.Set("key2", 2, time.Millisecond)
	cache.Set("key3", 3, time.Millisecond)
	cache.Set("key4", 4, time.Second)
	cache.Set("key5", 5, time.Second)

	if reporter.CacheSize() != 5 {
		t.Errorf("CacheSize %d is wrong", reporter.CacheSize())
	}

	time.Sleep(100 * time.Millisecond)
	cache.GC()

	if reporter.CacheSize() != 2 {
		t.Errorf("CacheSize %d is wrong", reporter.CacheSize())
	}
}
