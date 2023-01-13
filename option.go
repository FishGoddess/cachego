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
	// These fields are for creating.
	maps       int
	segments   int
	gcDuration time.Duration

	// These fields are for operating.
	maxEntries int
	maxScans   int
}

func newDefaultConfig() config {
	return config{
		maps:       128,
		segments:   0,
		gcDuration: 0,
		maxEntries: 0,
		maxScans:   100000,
	}
}

type Option func(conf *config)

func (o Option) applyTo(conf *config) {
	o(conf)
}

func applyOptions(conf *config, opts []Option) {
	for _, opt := range opts {
		opt.applyTo(conf)
	}
}

func WithMaps(maps int) Option {
	return func(conf *config) {
		conf.maps = maps
	}
}

func WithSegments(segments int) Option {
	return func(conf *config) {
		conf.segments = segments
	}
}

func WithGC(gcDuration time.Duration) Option {
	return func(conf *config) {
		conf.gcDuration = gcDuration
	}
}

func WithMaxEntries(maxEntries int) Option {
	return func(conf *config) {
		conf.maxEntries = maxEntries
	}
}

func WithMaxScans(maxScans int) Option {
	return func(conf *config) {
		conf.maxScans = maxScans
	}
}
