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
	"sync/atomic"
	"time"
)

// Reporter stores some values for reporting.
type Reporter struct {
	cache Cache

	missedCount uint64
	hitCount    uint64
	gcCount     uint64
	loadCount   uint64
}

func (r *Reporter) increaseMissedCount() {
	atomic.AddUint64(&r.missedCount, 1)
}

func (r *Reporter) increaseHitCount() {
	atomic.AddUint64(&r.hitCount, 1)
}

func (r *Reporter) increaseGCCount() {
	atomic.AddUint64(&r.gcCount, 1)
}

func (r *Reporter) increaseLoadCount() {
	atomic.AddUint64(&r.loadCount, 1)
}

// CountMissed returns the missed count.
func (r *Reporter) CountMissed() uint64 {
	return atomic.LoadUint64(&r.missedCount)
}

// CountHit returns the hit count.
func (r *Reporter) CountHit() uint64 {
	return atomic.LoadUint64(&r.hitCount)
}

// CountGC returns the gc count.
func (r *Reporter) CountGC() uint64 {
	return atomic.LoadUint64(&r.gcCount)
}

// CountLoad returns the load count.
func (r *Reporter) CountLoad() uint64 {
	return atomic.LoadUint64(&r.loadCount)
}

// CacheSize returns the size of cache.
func (r *Reporter) CacheSize() int {
	return r.cache.Size()
}

// MissedRate returns the missed rate.
func (r *Reporter) MissedRate() float64 {
	hit := r.CountHit()
	missed := r.CountMissed()

	total := hit + missed
	if total <= 0 {
		return 0.0
	}

	return float64(missed) / float64(total)
}

// HitRate returns the hit rate.
func (r *Reporter) HitRate() float64 {
	hit := r.CountHit()
	missed := r.CountMissed()

	total := hit + missed
	if total <= 0 {
		return 0.0
	}

	return float64(hit) / float64(total)
}

type reportableCache struct {
	*config
	*Reporter
}

func report(conf *config, cache Cache) (Cache, *Reporter) {
	reporter := &Reporter{
		cache:       cache,
		hitCount:    0,
		missedCount: 0,
		gcCount:     0,
		loadCount:   0,
	}

	cache = &reportableCache{
		config:   conf,
		Reporter: reporter,
	}

	return cache, reporter
}

// Get gets the value of key from cache and returns value if found.
func (rc *reportableCache) Get(key string) (value interface{}, found bool) {
	value, found = rc.cache.Get(key)

	if found {
		if rc.recordHit {
			rc.increaseHitCount()
		}

		if rc.reportHit != nil {
			rc.reportHit(rc.Reporter, key, value)
		}
	} else {
		if rc.recordMissed {
			rc.increaseMissedCount()
		}

		if rc.reportMissed != nil {
			rc.reportMissed(rc.Reporter, key)
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
	if rc.recordGC {
		rc.increaseGCCount()
	}

	if rc.reportGC == nil {
		return rc.cache.GC()
	}

	begin := rc.now()
	cleans = rc.cache.GC()
	end := rc.now()

	cost := time.Duration(end - begin)
	rc.reportGC(rc.Reporter, cost, cleans)

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

	if rc.recordLoad {
		rc.increaseLoadCount()
	}

	if rc.reportLoad != nil {
		rc.reportLoad(rc.Reporter, key, value, ttl, err)
	}

	return value, err
}
