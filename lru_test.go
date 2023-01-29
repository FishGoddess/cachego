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

import "testing"

func newTestLRUCache() Cache {
	conf := newDefaultConfig()
	conf.maxEntries = maxTestEntries
	return newLRUCache(conf)
}

// go test -v -cover -run=^TestLRUCacheGet$
func TestLRUCacheGet(t *testing.T) {
	cache := newTestLRUCache()
	testCacheGet(t, cache)
}

// go test -v -cover -run=^TestLRUCacheSet$
func TestLRUCacheSet(t *testing.T) {
	cache := newTestLRUCache()
	testCacheSet(t, cache)
}

// go test -v -cover -run=^TestLRUCacheRemove$
func TestLRUCacheRemove(t *testing.T) {
	cache := newTestLRUCache()
	testCacheRemove(t, cache)
}

// go test -v -cover -run=^TestLRUCacheSize$
func TestLRUCacheSize(t *testing.T) {
	cache := newTestLRUCache()
	testCacheSize(t, cache)
}

// go test -v -cover -run=^TestLRUCacheGC$
func TestLRUCacheGC(t *testing.T) {
	cache := newTestLRUCache()
	testCacheGC(t, cache)
}

// go test -v -cover -run=^TestLRUCacheReset$
func TestLRUCacheReset(t *testing.T) {
	cache := newTestLRUCache()
	testCacheReset(t, cache)
}
