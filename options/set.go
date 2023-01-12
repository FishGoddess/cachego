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

type SetConfig struct {
	TTL time.Duration
}

func newDefaultSetConfig() *SetConfig {
	return &SetConfig{
		TTL: 0,
	}
}

type SetOption func(conf *SetConfig)

func (o SetOption) ApplyTo(conf *SetConfig) {
	o(conf)
}

type SetOptions []SetOption

func Set() SetOptions {
	return nil
}

func (opts SetOptions) TTL(ttl time.Duration) SetOptions {
	opt := func(conf *SetConfig) {
		conf.TTL = ttl
	}

	return append(opts, opt)
}

func (opts SetOptions) Config() *SetConfig {
	conf := newDefaultSetConfig()

	for _, opt := range opts {
		opt.ApplyTo(conf)
	}

	return conf
}
