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
	"strconv"
	"testing"
)

const (
	testShardings = 4
)

func newTestShardingCache() *shardingCache {
	conf := newDefaultConfig()
	conf.shardings = testShardings
	return newShardingCache(conf, newStandardCache).(*shardingCache)
}

// go test -v -cover -run=^TestShardingCache$
func TestShardingCache(t *testing.T) {
	cache := newTestShardingCache()
	testCacheImplement(t, cache)
}

// go test -v -cover -run=^TestShardingCacheIndex$
func TestShardingCacheIndex(t *testing.T) {
	cache := newTestShardingCache()
	if len(cache.caches) != testShardings {
		t.Errorf("len(cache.caches) %d is wrong", len(cache.caches))
	}

	for i := 0; i < 100; i++ {
		data := strconv.Itoa(i)
		cache.Set(data, data, NoTTL)
	}

	for i := range cache.caches {
		if cache.caches[i].Size() <= 0 {
			t.Errorf("cache.caches[i].Size() %d <= 0", cache.caches[i].Size())
		}
	}
}
