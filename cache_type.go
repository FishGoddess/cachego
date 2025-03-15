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

const (
	// standard cache is a simple cache with locked map.
	// It evicts entries randomly if cache size reaches to max entries.
	standard CacheType = "standard"

	// lru cache is a cache using lru to evict entries.
	// More details see https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU).
	lru CacheType = "lru"

	// lfu cache is a cache using lfu to evict entries.
	// More details see https://en.wikipedia.org/wiki/Cache_replacement_policies#Least-frequently_used_(LFU).
	lfu CacheType = "lfu"
)

// CacheType is the type of cache.
type CacheType string

// String returns the cache type in string form.
func (ct CacheType) String() string {
	return string(ct)
}

// IsStandard returns if cache type is standard.
func (ct CacheType) IsStandard() bool {
	return ct == standard
}

// IsLRU returns if cache type is lru.
func (ct CacheType) IsLRU() bool {
	return ct == lru
}

// IsLFU returns if cache type is lfu.
func (ct CacheType) IsLFU() bool {
	return ct == lfu
}
