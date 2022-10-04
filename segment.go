// Copyright 2020 FishGoddess. All Rights Reserved.
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
	"sync"
	"time"
)

// segment is the struct storing the real data.
type segment struct {
	// data stores all entries.
	data map[string]*value

	// mapSize is the initialized size of map inside.
	mapSize int

	// lock is for concurrency.
	lock sync.RWMutex
}

// newSegment returns a segment holder with mapSize.
func newSegment(mapSize int) *segment {
	return &segment{
		data:    make(map[string]*value, mapSize),
		mapSize: mapSize,
		lock:    sync.RWMutex{},
	}
}

// get returns the value of key and a false if not found.
func (s *segment) get(key string) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if value, ok := s.data[key]; ok && value.alive() {
		return value.data, true
	}

	return nil, false
}

// set puts a value of key with a ttl.
// If you want this key to be alive forever, just give it a noTTL.
func (s *segment) set(key string, value interface{}, ttl time.Duration) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if v, ok := s.data[key]; ok {
		v.renew(value, ttl) // Reuse value memory
		return
	}

	s.data[key] = newValue(value, ttl)
}

// delete will delete the key in segment.
func (s *segment) delete(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.data, key)
}

// deleteAll removes all keys in segment.
func (s *segment) deleteAll() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = make(map[string]*value, s.mapSize)
}

// size returns the size of entries of segment.
func (s *segment) size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.data)
}

// gc removes all dead entries in segment.
func (s *segment) gc() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for key, value := range s.data {
		if !value.alive() {
			delete(s.data, key)
		}
	}
}
