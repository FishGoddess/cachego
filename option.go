// Copyright 2025 FishGoddess. All Rights Reserved.
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

// WithCacheName returns an option setting the cacheName of config.
func WithCacheName(cacheName string) Option {
	return func(conf *config) {
		conf.cacheName = cacheName
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
		conf.shardings = shardings
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
		conf.gcDuration = gcDuration
	}
}

// WithMaxScans returns an option setting the max scans of cache.
// Negative value means no limit.
func WithMaxScans(maxScans int) Option {
	return func(conf *config) {
		conf.maxScans = maxScans
	}
}

// WithMaxEntries returns an option setting the max entries of cache.
// Negative value means no limit.
func WithMaxEntries(maxEntries int) Option {
	return func(conf *config) {
		conf.maxEntries = maxEntries
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

// WithRecordMissed returns an option setting the recordMissed of config.
func WithRecordMissed(recordMissed bool) Option {
	return func(conf *config) {
		conf.recordMissed = recordMissed
	}
}

// WithRecordHit returns an option setting the recordHit of config.
func WithRecordHit(recordHit bool) Option {
	return func(conf *config) {
		conf.recordHit = recordHit
	}
}

// WithRecordGC returns an option setting the recordGC of config.
func WithRecordGC(recordGC bool) Option {
	return func(conf *config) {
		conf.recordGC = recordGC
	}
}

// WithRecordLoad returns an option setting the recordLoad of config.
func WithRecordLoad(recordLoad bool) Option {
	return func(conf *config) {
		conf.recordLoad = recordLoad
	}
}

// WithReportMissed returns an option setting the reportMissed of config.
func WithReportMissed(reportMissed func(reporter *Reporter, key string)) Option {
	return func(conf *config) {
		conf.reportMissed = reportMissed
	}
}

// WithReportHit returns an option setting the reportHit of config.
func WithReportHit(reportHit func(reporter *Reporter, key string, value interface{})) Option {
	return func(conf *config) {
		conf.reportHit = reportHit
	}
}

// WithReportGC returns an option setting the reportGC of config.
func WithReportGC(reportGC func(reporter *Reporter, cost time.Duration, cleans int)) Option {
	return func(conf *config) {
		conf.reportGC = reportGC
	}
}

// WithReportLoad returns an option setting the reportLoad of config.
func WithReportLoad(reportLoad func(reporter *Reporter, key string, value interface{}, ttl time.Duration, err error)) Option {
	return func(conf *config) {
		conf.reportLoad = reportLoad
	}
}
