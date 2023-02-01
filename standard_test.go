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
	"time"
)

func newTestStandardCache() *standardCache {
	conf := newDefaultConfig()
	conf.maxEntries = maxTestEntries
	return newStandardCache(conf).(*standardCache)
}

// go test -v -cover -run=^TestStandardCache$
func TestStandardCache(t *testing.T) {
	cache := newTestStandardCache()
	testCacheImplement(t, cache)
}

// go test -v -cover -run=^TestStandardCacheEvict$
func TestStandardCacheEvict(t *testing.T) {
	cache := newTestStandardCache()

	for i := 0; i < cache.maxEntries*10; i++ {
		data := strconv.Itoa(i)
		evictedValue := cache.Set(data, data, time.Duration(i)*time.Second)

		if i >= cache.maxEntries && evictedValue == nil {
			t.Errorf("i %d >= cache.maxEntries %d && evictedValue == nil", i, cache.maxEntries)
		}
	}

	if cache.Size() != cache.maxEntries {
		t.Errorf("cache.Size() %d != cache.maxEntries %d", cache.Size(), cache.maxEntries)
	}
}
