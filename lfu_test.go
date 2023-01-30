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

func newTestLFUCache() Cache {
	conf := newDefaultConfig()
	conf.maxEntries = maxTestEntries
	return newLFUCache(conf)
}

// go test -v -cover -run=^TestLFUCacheGet$
func TestLFUCacheGet(t *testing.T) {
	cache := newTestLFUCache()
	testCacheGet(t, cache)
}

// go test -v -cover -run=^TestLFUCacheSet$
func TestLFUCacheSet(t *testing.T) {
	cache := newTestLFUCache()
	testCacheSet(t, cache)
}

// go test -v -cover -run=^TestLFUCacheRemove$
func TestLFUCacheRemove(t *testing.T) {
	cache := newTestLFUCache()
	testCacheRemove(t, cache)
}

// go test -v -cover -run=^TestLFUCacheSize$
func TestLFUCacheSize(t *testing.T) {
	cache := newTestLFUCache()
	testCacheSize(t, cache)
}

// go test -v -cover -run=^TestLFUCacheGC$
func TestLFUCacheGC(t *testing.T) {
	cache := newTestLFUCache()
	testCacheGC(t, cache)
}

// go test -v -cover -run=^TestLFUCacheReset$
func TestLFUCacheReset(t *testing.T) {
	cache := newTestLFUCache()
	testCacheReset(t, cache)
}
