// Copyright 2024 FishGoddess. All Rights Reserved.
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

// go test -v -cover -run=^TestCacheType$
func TestCacheType(t *testing.T) {
	if standard.String() != string(standard) {
		t.Fatalf("standard.String() %s is wrong", standard.String())
	}

	if lru.String() != string(lru) {
		t.Fatalf("lru.String() %s is wrong", lru.String())
	}

	if lfu.String() != string(lfu) {
		t.Fatalf("lfu.String() %s is wrong", lfu.String())
	}

	if !standard.IsStandard() {
		t.Fatal("!standard.IsStandard()")
	}

	if !lru.IsLRU() {
		t.Fatal("!standard.IsLRU()")
	}

	if !lfu.IsLFU() {
		t.Fatal("!standard.IsLFU()")
	}
}
