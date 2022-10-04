// Copyright 2020 FishGoddess. All Rights Reserved.
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
	"context"
	"math/bits"
	"time"
)

// Option is a function which initializes config.
type Option func(conf *config)

// applyOptions applies opts to conf.
func applyOptions(conf *config, opts []Option) *config {
	for _, opt := range opts {
		opt(conf)
	}

	return conf
}

// WithMapSize is an option setting initializing map size of cache.
func WithMapSize(mapSize uint) Option {
	return func(conf *config) {
		conf.mapSize = int(mapSize)
	}
}

// WithSegmentSize is an option setting initializing segment size of cache.
// segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.
func WithSegmentSize(segmentSize uint) Option {
	if bits.OnesCount(segmentSize) > 1 {
		panic("cachego: segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.")
	}

	return func(conf *config) {
		conf.segmentSize = int(segmentSize)
	}
}

// WithAutoGC is an option turning on automatically gc.
func WithAutoGC(d time.Duration) Option {
	return func(conf *config) {
		if d > 0 {
			conf.gcDuration = d
		}
	}
}

// WithDisableSingleflight is an option disabling single-flight mode of cache.
func WithDisableSingleflight() Option {
	return func(conf *config) {
		conf.singleflight = false
	}
}

// OpOption is a function which initializes opConfig.
type OpOption func(conf *opConfig)

// applyOpOptions applies opts to conf.
func applyOpOptions(conf *opConfig, opts []OpOption) *opConfig {
	for _, opt := range opts {
		opt(conf)
	}

	return conf
}

// WithOpContext sets context to ctx.
func WithOpContext(ctx context.Context) OpOption {
	return func(conf *opConfig) {
		conf.ctx = ctx
	}
}

// WithOpTTL sets the ttl of missed key if loaded to ttl.
func WithOpTTL(ttl time.Duration) OpOption {
	return func(conf *opConfig) {
		conf.ttl = ttl
	}
}

// WithOpNoTTL sets the ttl of missed key to no ttl.
func WithOpNoTTL() OpOption {
	return func(conf *opConfig) {
		conf.ttl = noTTL
	}
}

// WithOpOnMissed sets onMissed to Get operation.
func WithOpOnMissed(onMissed func(ctx context.Context) (data interface{}, err error)) OpOption {
	return func(conf *opConfig) {
		conf.onMissed = onMissed
	}
}

// WithOpDisableSingleflight sets the single-flight mode to false.
func WithOpDisableSingleflight() OpOption {
	return func(conf *opConfig) {
		conf.singleflight = false
	}
}

// WithOpDisableReload sets the reloading flag to false.
func WithOpDisableReload() OpOption {
	return func(conf *opConfig) {
		conf.reload = false
	}
}
