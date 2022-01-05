// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/04/05 15:58:16

package cachego

import (
	"context"
	"math/bits"
	"time"

	"github.com/FishGoddess/cachego/internal/config"
)

// Option is a function which initializes Config.
type Option func(conf *config.Config)

// applyOptions applies opts to conf.
func applyOptions(conf *config.Config, opts ...Option) *config.Config {
	for _, applyOption := range opts {
		applyOption(conf)
	}
	return conf
}

// WithMapSize is an option setting initializing map size of cache.
func WithMapSize(mapSize uint) Option {
	return func(conf *config.Config) {
		conf.MapSize = int(mapSize)
	}
}

// WithSegmentSize is an option setting initializing segment size of cache.
// segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.
func WithSegmentSize(segmentSize uint) Option {
	if bits.OnesCount(segmentSize) > 1 {
		panic("cachego: segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.")
	}

	return func(conf *config.Config) {
		conf.SegmentSize = int(segmentSize)
	}
}

// WithAutoGC is an option turning on automatically gc.
func WithAutoGC(d time.Duration) Option {
	return func(conf *config.Config) {
		if d > 0 {
			conf.GCDuration = d
		}
	}
}

// WithDisableSingleflight is an option disabling single-flight mode of cache.
func WithDisableSingleflight() Option {
	return func(conf *config.Config) {
		conf.EnableSingleflight = false
	}
}

// GetOption is a function which initializes GetConfig.
type GetOption func(conf *config.GetConfig)

// applyGetOptions applies opts to conf.
func applyGetOptions(conf *config.GetConfig, opts ...GetOption) *config.GetConfig {
	for _, applyOption := range opts {
		applyOption(conf)
	}
	return conf
}

// WithGetContext sets context to ctx.
func WithGetContext(ctx context.Context) GetOption {
	return func(conf *config.GetConfig) {
		conf.Ctx = ctx
	}
}

// WithGetTTL sets the ttl of missed key if loaded to ttl.
func WithGetTTL(ttl time.Duration) GetOption {
	return func(conf *config.GetConfig) {
		conf.TTL = ttl
	}
}

// WithGetNoTTL sets the ttl of missed key to no ttl.
func WithGetNoTTL() GetOption {
	return func(conf *config.GetConfig) {
		conf.TTL = config.NoTTL
	}
}

// WithGetOnMissed sets onMissed to Get operation.
func WithGetOnMissed(onMissed func(ctx context.Context) (data interface{}, err error)) GetOption {
	return func(conf *config.GetConfig) {
		conf.OnMissed = onMissed
	}
}

// WithGetDisableSingleflight sets the single-flight mode to false.
func WithGetDisableSingleflight() GetOption {
	return func(conf *config.GetConfig) {
		conf.Singleflight = false
	}
}

// SetOption is a function which initializes SetConfig.
type SetOption func(conf *config.SetConfig)

// applySetOptions applies opts to conf.
func applySetOptions(conf *config.SetConfig, opts ...SetOption) *config.SetConfig {
	for _, applyOption := range opts {
		applyOption(conf)
	}
	return conf
}

// WithSetTTL sets the ttl of key to ttl.
func WithSetTTL(ttl time.Duration) SetOption {
	return func(conf *config.SetConfig) {
		if ttl > 0 {
			conf.TTL = ttl
		}
	}
}

// WithSetNoTTL sets the ttl of key to no ttl.
func WithSetNoTTL() SetOption {
	return func(conf *config.SetConfig) {
		conf.TTL = config.NoTTL
	}
}

// AutoSetOption is a function which initializes AutoSetConfig.
type AutoSetOption func(conf *config.AutoSetConfig)

// applyAutoSetOptions applies opts to conf.
func applyAutoSetOptions(conf *config.AutoSetConfig, opts ...AutoSetOption) *config.AutoSetConfig {
	for _, applyOption := range opts {
		applyOption(conf)
	}
	return conf
}

// WithAutoSetContext sets context to ctx.
func WithAutoSetContext(ctx context.Context) AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		conf.Ctx = ctx
	}
}

// WithAutoSetTTL sets the ttl of key to ttl.
func WithAutoSetTTL(ttl time.Duration) AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		if ttl > 0 {
			conf.TTL = ttl
		}
	}
}

// WithAutoSetNoTTL sets the ttl of key to no ttl.
func WithAutoSetNoTTL() AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		conf.TTL = config.NoTTL
	}
}

// WithAutoSetGap sets the gap between two set operations to gap.
func WithAutoSetGap(gap time.Duration) AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		if gap > 0 {
			conf.Gap = gap
		}
	}
}
