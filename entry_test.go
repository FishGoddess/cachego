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
	"fmt"
	"testing"
	"time"
)

const (
	testDurationGap = 10 * time.Microsecond
)

// go test -v -cover -run=^TestNewEntry$
func TestNewEntry(t *testing.T) {
	e := newEntry("key", "value", 0, now)

	if e.key != "key" {
		t.Fatalf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "value" {
		t.Fatalf("e.value %+v is wrong", e.value)
	}

	if e.expiration != 0 {
		t.Fatalf("e.expiration %+v != 0", e.expiration)
	}

	if fmt.Sprintf("%p", e.now) != fmt.Sprintf("%p", now) {
		t.Fatalf("e.now %p is wrong", e.now)
	}

	e = newEntry("k", "v", time.Second, now)
	expiration := time.Now().Add(time.Second).UnixNano()

	if e.key != "k" {
		t.Fatalf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "v" {
		t.Fatalf("e.value %+v is wrong", e.value)
	}

	if e.expiration == 0 {
		t.Fatal("e.expiration == 0")
	}

	if fmt.Sprintf("%p", e.now) != fmt.Sprintf("%p", now) {
		t.Fatalf("e.now %p is wrong", e.now)
	}

	// Keep one us for code running.
	if expiration < e.expiration || e.expiration < expiration-testDurationGap.Nanoseconds() {
		t.Fatalf("e.expiration %d != expiration %d", e.expiration, expiration)
	}
}

// go test -v -cover -run=^TestEntrySetup$
func TestEntrySetup(t *testing.T) {
	e := newEntry("key", "value", 0, now)

	if e.key != "key" {
		t.Fatalf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "value" {
		t.Fatalf("e.value %+v is wrong", e.value)
	}

	if e.expiration != 0 {
		t.Fatalf("e.expiration %+v != 0", e.expiration)
	}

	ee := e
	e.setup("k", "v", time.Second)
	expiration := time.Now().Add(time.Second).UnixNano()

	if ee != e {
		t.Fatalf("ee %p != e %p", ee, e)
	}

	if e.key != "k" {
		t.Fatalf("e.key %s is wrong", e.key)
	}

	if e.value.(string) != "v" {
		t.Fatalf("e.value %+v is wrong", e.value)
	}

	if e.expiration == 0 {
		t.Fatal("e.expiration == 0")
	}

	// Keep one us for code running.
	if expiration < e.expiration || e.expiration < expiration-testDurationGap.Nanoseconds() {
		t.Fatalf("e.expiration %d != expiration %d", e.expiration, expiration)
	}
}

// go test -cover -run=^TestEntryExpired$
func TestEntryExpired(t *testing.T) {
	e := newEntry("", nil, time.Millisecond, now)

	if e.expired(0) {
		t.Fatal("e should be unexpired!")
	}

	time.Sleep(2 * time.Millisecond)

	if !e.expired(0) {
		t.Fatal("e should be expired!")
	}
}
