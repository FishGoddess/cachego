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

import "time"

type entry struct {
	key   string
	value interface{}

	// Time in nanosecond, valid util 2262 year (enough, right?)
	expiration int64
	now        func() int64
}

func newEntry(key string, value interface{}, ttl time.Duration, now func() int64) *entry {
	e := &entry{
		now: now,
	}

	e.setup(key, value, ttl)
	return e
}

func (e *entry) setup(key string, value interface{}, ttl time.Duration) {
	e.key = key
	e.value = value
	e.expiration = 0

	if ttl > 0 {
		e.expiration = e.now() + ttl.Nanoseconds()
	}
}

func (e *entry) expired(now int64) bool {
	if now > 0 {
		return e.expiration > 0 && e.expiration < now
	}

	return e.expiration > 0 && e.expiration < e.now()
}
