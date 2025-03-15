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
	"errors"
	"time"

	flight "github.com/FishGoddess/cachego/pkg/singleflight"
)

// loader loads values from somewhere.
type loader struct {
	group *flight.Group
}

// newLoader creates a loader.
// It also creates a singleflight group to call load if singleflight is true.
func newLoader(singleflight bool) *loader {
	loader := new(loader)

	if singleflight {
		loader.group = flight.NewGroup(mapInitialCap)
	}

	return loader
}

// Load loads a value of key with ttl and returns an error if failed.
func (l *loader) Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error) {
	if load == nil {
		return nil, errors.New("cachego: load function is nil in loader")
	}

	if l.group == nil {
		return load()
	}

	return l.group.Call(key, load)
}

// Reset resets loader to initial status which is like a new loader.
func (l *loader) Reset() {
	if l.group != nil {
		l.group.Reset()
	}
}
