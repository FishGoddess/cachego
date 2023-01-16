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
	"errors"
	"time"

	"github.com/FishGoddess/cachego/pkg/singleflight"
)

// Loader loads a value to cache.
// We recommend you set enableSingleflight=true in NewLoader for reducing the times calling load function.
type Loader struct {
	cache Cache
	group *singleflight.Group
}

// NewLoader creates a loader with cache.
// It also creates a singleflight group to call load if enableSingleflight is true.
func NewLoader(cache Cache, enableSingleflight bool) *Loader {
	loader := &Loader{
		cache: cache,
	}

	if enableSingleflight {
		loader.group = singleflight.NewGroup(MapInitialCap)
	}

	return loader
}

// Load loads a key with ttl to cache and returns an error if failed.
// See Cache interface.
func (l *Loader) Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error) {
	if load == nil {
		return nil, errors.New("cachego: load function is nil in loader")
	}

	if l.group != nil {
		value, err = l.group.Call(key, load)
	} else {
		value, err = load()
	}

	if err != nil {
		return nil, err
	}

	if l.cache != nil {
		l.cache.Set(key, value, ttl)
	}

	return value, err
}
