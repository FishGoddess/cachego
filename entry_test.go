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

// go test -v -cover -run=^TestNewEntry$
func TestNewEntry(t *testing.T) {
	e := newEntry("key", "value", 0)

	if e.key != "key" {
		t.Errorf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "value" {
		t.Errorf("e.value %+v is wrong", e.value)
	}

	if e.expiration != 0 {
		t.Errorf("e.expiration %+v != 0", e.expiration)
	}

	e = newEntry("k", "v", time.Second)
	expiration := time.Now().Add(time.Second).UnixNano()

	if e.key != "k" {
		t.Errorf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "v" {
		t.Errorf("e.value %+v is wrong", e.value)
	}

	if e.expiration == 0 {
		t.Error("e.expiration == 0")
	}

	// Keep one us for code running.
	if expiration < e.expiration || e.expiration < expiration-time.Microsecond.Nanoseconds() {
		t.Errorf("e.expiration %d != expiration %d", e.expiration, expiration)
	}
}

// go test -v -cover -run=^TestEntrySetup$
func TestEntrySetup(t *testing.T) {
	e := newEntry("key", "value", 0)

	if e.key != "key" {
		t.Errorf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "value" {
		t.Errorf("e.value %+v is wrong", e.value)
	}

	if e.expiration != 0 {
		t.Errorf("e.expiration %+v != 0", e.expiration)
	}

	ee := e
	e.setup("k", "v", time.Second)
	expiration := time.Now().Add(time.Second).UnixNano()

	if ee != e {
		t.Errorf("ee %p != e %p", ee, e)
	}

	if e.key != "k" {
		t.Errorf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "v" {
		t.Errorf("e.value %+v is wrong", e.value)
	}

	if e.expiration == 0 {
		t.Error("e.expiration == 0")
	}

	// Keep one us for code running.
	if expiration < e.expiration || e.expiration < expiration-time.Microsecond.Nanoseconds() {
		t.Errorf("e.expiration %d != expiration %d", e.expiration, expiration)
	}
}

// go test -cover -run=^TestEntryExpired$
func TestEntryExpired(t *testing.T) {
	e := newEntry("", nil, time.Millisecond)

	if e.expired() {
		t.Error("e should be unexpired!")
	}

	time.Sleep(2 * time.Millisecond)

	if !e.expired() {
		t.Error("e should be expired!")
	}
}