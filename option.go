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

// ReportOption applies to report config and sets some values to report config.
type ReportOption func(conf *reportConfig)

func (ro ReportOption) applyTo(conf *reportConfig) {
	ro(conf)
}

func applyReportOptions(conf *reportConfig, opts []ReportOption) {
	for _, opt := range opts {
		opt.applyTo(conf)
	}
}

// WithReporterNow returns an option setting the now function of reporter.
// A now function should return a nanosecond unix time.
func WithReporterNow(now func() int64) ReportOption {
	return func(conf *reportConfig) {
		if now != nil {
			conf.now = now
		}
	}
}

// WithRecordMissed returns an option setting the recordMissed of report config.
func WithRecordMissed(recordMissed bool) ReportOption {
	return func(conf *reportConfig) {
		conf.recordMissed = recordMissed
	}
}

// WithRecordHit returns an option setting the recordHit of report config.
func WithRecordHit(recordHit bool) ReportOption {
	return func(conf *reportConfig) {
		conf.recordHit = recordHit
	}
}

// WithRecordGC returns an option setting the recordGC of report config.
func WithRecordGC(recordGC bool) ReportOption {
	return func(conf *reportConfig) {
		conf.recordGC = recordGC
	}
}

// WithRecordLoad returns an option setting the recordLoad of report config.
func WithRecordLoad(recordLoad bool) ReportOption {
	return func(conf *reportConfig) {
		conf.recordLoad = recordLoad
	}
}

// WithReportMissed returns an option setting the reportMissed of report config.
func WithReportMissed(reportMissed func(reporter *Reporter, key string)) ReportOption {
	return func(conf *reportConfig) {
		conf.reportMissed = reportMissed
	}
}

// WithReportHit returns an option setting the reportHit of report config.
func WithReportHit(reportHit func(reporter *Reporter, key string, value interface{})) ReportOption {
	return func(conf *reportConfig) {
		conf.reportHit = reportHit
	}
}

// WithReportGC returns an option setting the reportGC of report config.
func WithReportGC(reportGC func(reporter *Reporter, cost time.Duration, cleans int)) ReportOption {
	return func(conf *reportConfig) {
		conf.reportGC = reportGC
	}
}

// WithReportLoad returns an option setting the reportLoad of report config.
func WithReportLoad(reportLoad func(reporter *Reporter, key string, value interface{}, ttl time.Duration, err error)) ReportOption {
	return func(conf *reportConfig) {
		conf.reportLoad = reportLoad
	}
}
