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

// go test -v -cover -run=^TestApplyOptions$
func TestApplyOptions(t *testing.T) {
	got := config{
		shardings:  0,
		gcDuration: 0,
		maxScans:   0,
		maxEntries: 0,
	}

	expect := config{
		shardings:  1,
		gcDuration: 2,
		maxScans:   3,
		maxEntries: 4,
	}

	applyOptions(&got, []Option{
		WithShardings(1),
		WithGC(2),
		WithMaxScans(3),
		WithMaxEntries(4),
	})

	if got != expect {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithShardings$
func TestWithShardings(t *testing.T) {
	got := config{shardings: 0}
	expect := config{shardings: 1024}

	WithShardings(1024).applyTo(&got)
	if got != expect {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithDisableSingleflight$
func TestWithDisableSingleflight(t *testing.T) {
	got := config{singleflight: true}
	expect := config{singleflight: false}

	WithDisableSingleflight().applyTo(&got)
	if got != expect {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithGC$
func TestWithGC(t *testing.T) {
	got := config{gcDuration: 0}
	expect := config{gcDuration: 1024}

	WithGC(1024).applyTo(&got)
	if got != expect {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithMaxScans$
func TestWithMaxScans(t *testing.T) {
	got := config{maxScans: 0}
	expect := config{maxScans: 1024}

	WithMaxScans(1024).applyTo(&got)
	if got != expect {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}

// go test -v -cover -run=^TestWithMaxEntries$
func TestWithMaxEntries(t *testing.T) {
	got := config{maxEntries: 0}
	expect := config{maxEntries: 1024}

	WithMaxEntries(1024).applyTo(&got)
	if got != expect {
		t.Errorf("got %+v != expect %+v", got, expect)
	}
}
