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
	"testing"
	"time"
)

// go test -v -cover -run=^TestHash$
func TestHash(t *testing.T) {
	hash := Hash("test")
	if hash <= 0 {
		t.Errorf("hash %d <= 0", hash)
	}
}

// go test -v -cover -run=^TestNow$
func TestNow(t *testing.T) {
	got := Now()
	expect := time.Now().UnixNano()

	if got > expect || got < expect-time.Microsecond.Nanoseconds() {
		t.Errorf("got %d != expect %d", got, expect)
	}
}