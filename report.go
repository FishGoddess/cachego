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
	"time"
)

type reportableCache struct {
	conf  *config
	cache Cache
}

func report(conf *config, cache Cache) Cache {
	return &reportableCache{
		conf:  conf,
		cache: cache,
	}
}

// Get gets the value of key from cache and returns value if found.
func (rc *reportableCache) Get(key string) (value interface{}, found bool) {
	value, found = rc.cache.Get(key)

	if found {
		if rc.conf.reportHit != nil {
			rc.conf.reportHit(key, value)
		}
	} else {
		if rc.conf.reportMissed != nil {
			rc.conf.reportMissed(key)
		}
	}

	return value, found
}

// Set sets key and value to cache with ttl and returns evicted value if exists and unexpired.
// See Cache interface.
func (rc *reportableCache) Set(key string, value interface{}, ttl time.Duration) (evictedValue interface{}) {
	return rc.cache.Set(key, value, ttl)
}

// Remove removes key and returns the removed value of key.
// See Cache interface.
func (rc *reportableCache) Remove(key string) (removedValue interface{}) {
	return rc.cache.Remove(key)
}

// Size returns the count of keys in cache.
// See Cache interface.
func (rc *reportableCache) Size() (size int) {
	return rc.cache.Size()
}

// GC cleans the expired keys in cache and returns the exact count cleaned.
// See Cache interface.
func (rc *reportableCache) GC() (cleans int) {
	if rc.conf.reportGC == nil {
		return rc.cache.GC()
	}

	begin := Now()
	cleans = rc.cache.GC()
	end := Now()

	cost := time.Duration(end - begin)
	rc.conf.reportGC(cost, cleans)

	return cleans
}

// Reset resets cache to initial status which is like a new cache.
// See Cache interface.
func (rc *reportableCache) Reset() {
	rc.cache.Reset()
}

// Load loads a key with ttl to cache and returns an error if failed.
// See Cache interface.
func (rc *reportableCache) Load(key string, ttl time.Duration, load func() (value interface{}, err error)) (value interface{}, err error) {
	value, err = rc.cache.Load(key, ttl, load)

	if rc.conf.reportLoad != nil {
		rc.conf.reportLoad(key, value, ttl, err)
	}

	return value, err
}
