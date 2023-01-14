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

// go test -v -cover=^TestNewStandard$
func TestNewStandard(t *testing.T) {
	cache := NewStandardCache()

	sc1, ok := cache.(*standardCache)
	if !ok {
		t.Errorf("cache.(*standardCache) %T not ok", cache)
	}

	if sc1 == nil {
		t.Error("sc1 == nil")
	}

	cache = NewStandardCache(WithShardings(64))

	sc2, ok := cache.(*shardingCache)
	if !ok {
		t.Errorf("cache.(*shardingCache) %T not ok", cache)
	}

	if sc2 == nil {
		t.Error("sc2 == nil")
	}
}
