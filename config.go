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

import "time"

type config struct {
	cacheType    cacheType
	shardings    int
	singleflight bool
	gcDuration   time.Duration

	maxScans   int
	maxEntries int

	now  func() int64
	hash func(key string) int

	reportMissed func(key string)
	reportHit    func(key string, value interface{})
	reportGC     func(cost time.Duration, cleans int)
	reportLoad   func(key string, value interface{}, ttl time.Duration, err error)
}

func newDefaultConfig() *config {
	return &config{
		cacheType:    standard,
		shardings:    0,
		singleflight: true,
		gcDuration:   0,
		maxScans:     10000,
		maxEntries:   0,
		now:          now,
		hash:         hash,
	}
}

type reportConfig struct {
	now func() int64

	reportMissed func(key string)
	reportHit    func(key string, value interface{})
	reportGC     func(cost time.Duration, cleans int)
	reportLoad   func(key string, value interface{}, ttl time.Duration, err error)
}

func newDefaultReportConfig() *reportConfig {
	return &reportConfig{
		now:          now,
		reportMissed: nil,
		reportHit:    nil,
		reportGC:     nil,
		reportLoad:   nil,
	}
}
