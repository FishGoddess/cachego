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

	if fmt.Sprintf("%p", conf1.now) != fmt.Sprintf("%p", conf2.now) {
		return false
	}

	if fmt.Sprintf("%p", conf1.hash) != fmt.Sprintf("%p", conf2.hash) {
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

func isReportConfigEquals(conf1 *reportConfig, conf2 *reportConfig) bool {
	if fmt.Sprintf("%p", conf1.now) != fmt.Sprintf("%p", conf2.now) {
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

// go test -v -cover -run=^TestApplyReportOptions$
func TestApplyReportOptions(t *testing.T) {
	reportHit := func(key string, value interface{}) {}

	got := &reportConfig{}

	expect := &reportConfig{
		now:       now,
		reportHit: reportHit,
	}

	applyReportOptions(got, []ReportOption{
		WithReporterNow(now),
		WithReportHit(reportHit),
	})

	if !isReportConfigEquals(got, expect) {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}
