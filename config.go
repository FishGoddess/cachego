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
	cacheName    string
	cacheType    CacheType
	shardings    int
	singleflight bool
	gcDuration   time.Duration

	maxScans   int
	maxEntries int

	now  func() int64
	hash func(key string) int

	recordMissed bool
	recordHit    bool
	recordGC     bool
	recordLoad   bool

	reportMissed func(reporter *Reporter, key string)
	reportHit    func(reporter *Reporter, key string, value interface{})
	reportGC     func(reporter *Reporter, cost time.Duration, cleans int)
	reportLoad   func(reporter *Reporter, key string, value interface{}, ttl time.Duration, err error)
}

func newDefaultConfig() *config {
	return &config{
		cacheName:    "",
		cacheType:    standard,
		shardings:    0,
		singleflight: true,
		gcDuration:   10 * time.Minute,
		maxScans:     10000,
		maxEntries:   100000,
		now:          now,
		hash:         hash,
		recordMissed: true,
		recordHit:    true,
		recordGC:     true,
		recordLoad:   true,
	}
}
