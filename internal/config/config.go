// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/12/06 00:12:34

package config

import (
	"context"
	"time"
)

const (
	// NoTTL means entry lives forever.
	NoTTL = 0
)

// Config is the config of cache.
type Config struct {
	// MapSize is the size of map inside.
	MapSize int

	// SegmentSize is the size of segments.
	// This value will affect the performance of concurrency.
	// It should be the pow of 2 (such as 64) or the segments may be uneven.
	SegmentSize int

	// GCDuration is the duration of gc.
	GCDuration time.Duration

	// EnableSingleflight means cache will enable single-flight mode.
	EnableSingleflight bool
}

// NewDefaultConfig returns the default config of cache.
func NewDefaultConfig() *Config {
	return &Config{
		MapSize:            256,
		SegmentSize:        256,
		GCDuration:         0,
		EnableSingleflight: true,
	}
}

// GetConfig is the config of Get operations.
type GetConfig struct {
	// Ctx is the context of AutoSet.
	Ctx context.Context

	// TTL is the ttl of entry set to the cache.
	TTL time.Duration

	// OnMissed is the function which will be called if not nil.
	OnMissed func(ctx context.Context) (data interface{}, err error)

	// Singleflight means the call of OnMissed is single-flight mode.
	// This is a recommended way to load data from storages to cache, however,
	// it may decrease the success rate of loading data.
	Singleflight bool
}

// NewDefaultGetConfig returns the default config of Get operations.
func NewDefaultGetConfig() *GetConfig {
	return &GetConfig{
		Ctx:          context.Background(),
		TTL:          10 * time.Second,
		OnMissed:     nil,
		Singleflight: true,
	}
}

// SetConfig is the config of Set operations.
type SetConfig struct {
	// TTL is the ttl of entry set to the cache.
	TTL time.Duration
}

// NewDefaultSetConfig returns the default config of Set operations.
func NewDefaultSetConfig() *SetConfig {
	return &SetConfig{
		TTL: NoTTL,
	}
}

// AutoSetConfig is the config of AutoSet operations.
type AutoSetConfig struct {
	// Ctx is the context of AutoSet.
	Ctx context.Context

	// TTL is the ttl of entry set to the cache.
	TTL time.Duration

	// Gap is the duration between two AutoSet operations.
	Gap time.Duration
}

// NewDefaultAutoSetConfig returns the default config of AutoSet operations.
func NewDefaultAutoSetConfig() *AutoSetConfig {
	return &AutoSetConfig{
		Ctx: context.Background(),
		TTL: NoTTL,
		Gap: time.Minute,
	}
}
