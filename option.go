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

// Option is a function which initializes cache.
type Option func(cache *Cache)

func applyOptions(cache *Cache, opts ...Option) *Cache {
	for _, applyOption := range opts {
		applyOption(cache)
	}
	return cache
}

// WithMapSize is an option setting initializing map size of cache.
func WithMapSize(mapSize uint) Option {
	return func(cache *Cache) {
		cache.mapSize = int(mapSize)
	}
}

// WithSegmentSize is an option setting initializing segment size of cache.
// segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.
func WithSegmentSize(segmentSize uint) Option {
	if bits.OnesCount(segmentSize) > 1 {
		panic("segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.")
	}

	return func(cache *Cache) {
		cache.segmentSize = int(segmentSize)
	}
}

// WithAutoGC is an option turning on automatically gc.
func WithAutoGC(gcDuration time.Duration) Option {
	return func(cache *Cache) {
		cache.AutoGc(gcDuration)
	}
}

// SetOption is a function which initializes SetConfig.
type SetOption func(conf *config.SetConfig)

func applySetOptions(conf *config.SetConfig, opts ...SetOption) *config.SetConfig {
	for _, applyOption := range opts {
		applyOption(conf)
	}
	return conf
}

func WithSetTTL(ttl time.Duration) SetOption {
	return func(conf *config.SetConfig) {
		if ttl > 0 {
			conf.TTL = ttl
		}
	}
}

func WithSetNoTTL() SetOption {
	return func(conf *config.SetConfig) {
		conf.TTL = config.NoTTL
	}
}

// AutoSetOption is a function which initializes AutoSetConfig.
type AutoSetOption func(conf *config.AutoSetConfig)

func applyAutoSetOptions(conf *config.AutoSetConfig, opts ...AutoSetOption) *config.AutoSetConfig {
	for _, applyOption := range opts {
		applyOption(conf)
	}
	return conf
}

func WithAutoSetContext(ctx context.Context) AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		conf.Ctx = ctx
	}
}

func WithAutoSetGap(gap time.Duration) AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		if gap > 0 {
			conf.Gap = gap
		}
	}
}

func WithAutoSetTTL(ttl time.Duration) AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		if ttl > 0 {
			conf.TTL = ttl
		}
	}
}

func WithAutoSetNoTTL() AutoSetOption {
	return func(conf *config.AutoSetConfig) {
		conf.TTL = config.NoTTL
	}
}

// GetOption is a function which initializes GetConfig.
type GetOption func(conf *config.GetConfig)

func applyGetOptions(conf *config.GetConfig, opts ...GetOption) *config.GetConfig {
	for _, applyOption := range opts {
		applyOption(conf)
	}
	return conf
}

func WithGetContext(ctx context.Context) GetOption {
	return func(conf *config.GetConfig) {
		conf.Ctx = ctx
	}
}

func WithGetOnMissed(onMissed func(ctx context.Context) (data interface{}, err error)) GetOption {
	return func(conf *config.GetConfig) {
		conf.OnMissed = onMissed
	}
}

func WithGetOnMissedSet(need bool) GetOption {
	return func(conf *config.GetConfig) {
		conf.OnMissedSet = need
	}
}

func WithGetOnMissedSetTTL(ttl time.Duration) GetOption {
	return func(conf *config.GetConfig) {
		conf.OnMissedSetTTL = ttl
	}
}

func WithGetOnMissedSetNoTTL() GetOption {
	return func(conf *config.GetConfig) {
		conf.OnMissedSetTTL = config.NoTTL
	}
}
