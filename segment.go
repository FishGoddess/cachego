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

import "sync"

type segment struct {
	mapSize int
	data    map[string]*value
	lock    *sync.RWMutex
}

func newSegment(mapSize int) *segment {
	return &segment{
		mapSize: mapSize,
		data:    make(map[string]*value, mapSize),
		lock:    &sync.RWMutex{},
	}
}

func (s *segment) get(key string) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if value, ok := s.data[key]; ok && value.alive() {
		return value.data, true
	}
	return nil, false
}

func (s *segment) set(key string, value interface{}, ttl int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[key] = newValue(value, ttl)
}

func (s *segment) remove(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.data, key)
}

func (s *segment) removeAll() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = make(map[string]*value, s.mapSize)
}

func (s *segment) size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}

func (s *segment) gc() {
	s.lock.Lock()
	defer s.lock.Unlock()
	for key, value := range s.data {
		if !value.alive() {
			delete(s.data, key)
		}
	}
}
