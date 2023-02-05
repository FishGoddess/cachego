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
	"time"
)

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

// Option applies to config and sets some values to config.
type Option func(conf *config)

func (o Option) applyTo(conf *config) {
	o(conf)
}

func applyOptions(conf *config, opts []Option) {
	for _, opt := range opts {
		opt.applyTo(conf)
	}
}

// WithLRU returns an option setting the type of cache to lru.
// Notice that lru cache must have max entries limit, so you have to specify a maxEntries.
func WithLRU(maxEntries int) Option {
	return func(conf *config) {
		conf.cacheType = lru
		conf.maxEntries = maxEntries
	}
}

// WithLFU returns an option setting the type of cache to lfu.
// Notice that lfu cache must have max entries limit, so you have to specify a maxEntries.
func WithLFU(maxEntries int) Option {
	return func(conf *config) {
		conf.cacheType = lfu
		conf.maxEntries = maxEntries
	}
}

// WithShardings returns an option setting the sharding count of cache.
// Negative value means no sharding.
func WithShardings(shardings int) Option {
	return func(conf *config) {
		if shardings > 0 {
			conf.shardings = shardings
		}
	}
}

// WithDisableSingleflight returns an option turning off singleflight mode of cache.
func WithDisableSingleflight() Option {
	return func(conf *config) {
		conf.singleflight = false
	}
}

// WithGC returns an option setting the duration of cache gc.
// Negative value means no gc.
func WithGC(gcDuration time.Duration) Option {
	return func(conf *config) {
		if gcDuration > 0 {
			conf.gcDuration = gcDuration
		}
	}
}

// WithMaxScans returns an option setting the max scans of cache.
// Negative value means no limit.
func WithMaxScans(maxScans int) Option {
	return func(conf *config) {
		if maxScans > 0 {
			conf.maxScans = maxScans
		}
	}
}

// WithMaxEntries returns an option setting the max entries of cache.
// Negative value means no limit.
func WithMaxEntries(maxEntries int) Option {
	return func(conf *config) {
		if maxEntries > 0 {
			conf.maxEntries = maxEntries
		}
	}
}

// WithNow returns an option setting the now function of cache.
// A now function should return a nanosecond unix time.
func WithNow(now func() int64) Option {
	return func(conf *config) {
		if now != nil {
			conf.now = now
		}
	}
}

// WithHash returns an option setting the hash function of cache.
// A hash function should return the hash code of key.
func WithHash(hash func(key string) int) Option {
	return func(conf *config) {
		if hash != nil {
			conf.hash = hash
		}
	}
}

// WithReportMissed returns an option setting the reportMissed of cache.
func WithReportMissed(reportMissed func(key string)) Option {
	return func(conf *config) {
		if reportMissed != nil {
			conf.reportMissed = reportMissed
		}
	}
}

// WithReportHit returns an option setting the reportHit of cache.
func WithReportHit(reportHit func(key string, value interface{})) Option {
	return func(conf *config) {
		if reportHit != nil {
			conf.reportHit = reportHit
		}
	}
}

// WithReportGC returns an option setting the reportGC of cache.
func WithReportGC(reportGC func(cost time.Duration, cleans int)) Option {
	return func(conf *config) {
		if reportGC != nil {
			conf.reportGC = reportGC
		}
	}
}

// WithReportLoad returns an option setting the reportLoad of cache.
func WithReportLoad(reportLoad func(key string, value interface{}, ttl time.Duration, err error)) Option {
	return func(conf *config) {
		if reportLoad != nil {
			conf.reportLoad = reportLoad
		}
	}
}
