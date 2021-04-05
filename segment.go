// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/09/11 00:19:07

package cachego

import (
	"sync"
)

// segment is the struct storing the real data.
type segment struct {

	// mapSize is the initialized size of map inside.
	mapSize int

	// aliveSize is the size of alive values in data.
	aliveSize int

	// data stores all entries.
	data map[string]*value

	// lock is for concurrency.
	lock *sync.RWMutex
}

// newSegment returns a segment holder with mapSize.
func newSegment(mapSize int) *segment {
	return &segment{
		mapSize:   mapSize,
		aliveSize: 0,
		data:      make(map[string]*value, mapSize),
		lock: &sync.RWMutex{},
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

// set sets key and value with a ttl.
// If you want this key to be alive forever, just give it a NeverDie ttl.
// See value.
// The unit of ttl is second.
func (s *segment) set(key string, value interface{}, ttl int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.data[key]; !ok {
		s.aliveSize++
	}
	s.data[key] = newValue(value, ttl)
}

// remove removes the key in segment.
func (s *segment) remove(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.data[key]; ok {
		s.aliveSize--
	}
	delete(s.data, key)
}

// removeAll removes all keys in segment.
func (s *segment) removeAll() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.aliveSize = 0
	s.data = make(map[string]*value, s.mapSize)
}

// size returns the size of entries of segment.
func (s *segment) size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.aliveSize
}

// gc removes all dead entries in segment.
func (s *segment) gc() {
	s.lock.Lock()
	defer s.lock.Unlock()
	for key, value := range s.data {
		if !value.alive() {
			s.aliveSize--
			delete(s.data, key)
		}
	}
}
