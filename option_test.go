// Copyright 2022 FishGoddess. All Rights Reserved.
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

// go test -v -cover -run=^TestWithMapSize$
func TestWithMapSize(t *testing.T) {
	conf := &config{mapSize: 0}

	WithMapSize(16)(conf)
	if conf.mapSize != 16 {
		t.Errorf("conf.mapSize %d should be 16", conf.mapSize)
	}
}

// go test -v -cover -run=^TestWithSegmentSize$
func TestWithSegmentSize(t *testing.T) {
	conf := &config{segmentSize: 0}

	WithSegmentSize(16)(conf)
	if conf.segmentSize != 16 {
		t.Errorf("conf.segmentSize %d should be 16", conf.segmentSize)
	}

	for i := uint(1); i < 100000; i *= 2 {
		WithSegmentSize(i)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("WithSegmentSize(13) should panic")
		}
	}()

	WithSegmentSize(13)
}

// go test -v -cover -run=^TestWithAutoGC$
func TestWithAutoGC(t *testing.T) {
	conf := &config{gcDuration: 0}

	d := 10 * time.Minute
	WithAutoGC(d)(conf)
	if conf.gcDuration != d {
		t.Errorf("conf.gcDuration %d should be %s", conf.gcDuration, d)
	}
}
