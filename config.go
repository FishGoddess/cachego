// Copyright 2021 FishGoddess. All Rights Reserved.
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
	"time"
)

const (
	// noTTL means entry lives forever.
	noTTL = 0
)

// config is the config of cache.
type config struct {
	// mapSize is the size of map inside.
	mapSize int

	// segmentSize is the size of segments.
	// This value will affect the performance of concurrency.
	// It should be the pow of 2 (such as 64) or the segments may be uneven.
	segmentSize int

	// gcDuration is the duration of gc.
	gcDuration time.Duration

	// singleflight means cache will enable single-flight mode.
	singleflight bool
}

// newDefaultConfig returns the default config of cache.
func newDefaultConfig() *config {
	return &config{
		mapSize:      256,
		segmentSize:  256,
		gcDuration:   0,
		singleflight: true,
	}
}

// opConfig is the config of operations.
type opConfig struct {
	// ctx is the context of operation.
	ctx context.Context

	// ttl is the ttl of entry set to the cache in operation.
	ttl time.Duration

	// onMissed is the function which will be called if not nil in operation.
	onMissed func(ctx context.Context) (data interface{}, err error)

	// singleflight means the call of onMissed is single-flight mode.
	// This is a recommended way to load data from storages to cache, however,
	// it may decrease the success rate of loading data.
	singleflight bool

	// reload means this operation will reload data from onMissed to cache.
	reload bool
}

// newDefaultGetConfig returns the default config of Get operations.
func newDefaultGetConfig() *opConfig {
	return &opConfig{
		ctx:          context.Background(),
		ttl:          10 * time.Second,
		onMissed:     nil,
		singleflight: true,
		reload:       true,
	}
}

// newDefaultSetConfig returns the default config of Set operations.
func newDefaultSetConfig() *opConfig {
	return &opConfig{
		ttl: noTTL,
	}
}
