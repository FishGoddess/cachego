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

package options

import "time"

type CleanConfig struct {
	MaxScans int
	Timeout  time.Duration
}

func newDefaultCleanConfig() *CleanConfig {
	return &CleanConfig{
		MaxScans: 0,
		Timeout:  0,
	}
}

type CleanOption func(conf *CleanConfig)

func (o CleanOption) ApplyTo(conf *CleanConfig) {
	o(conf)
}

type CleanOptions []CleanOption

func Clean() CleanOptions {
	return nil
}

func (opts CleanOptions) MaxScans(maxScans int) CleanOptions {
	opt := func(conf *CleanConfig) {
		conf.MaxScans = maxScans
	}

	return append(opts, opt)
}

func (opts CleanOptions) Timeout(timeout time.Duration) CleanOptions {
	opt := func(conf *CleanConfig) {
		conf.Timeout = timeout
	}

	return append(opts, opt)
}

func (opts CleanOptions) Config() *CleanConfig {
	conf := newDefaultCleanConfig()

	for _, opt := range opts {
		opt.ApplyTo(conf)
	}

	return conf
}
