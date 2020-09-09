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
// Created at 2020/09/01 00:46:33

package cachego

import "sync"

const (
	defaultSegmentSize = 1024
)

type segment struct {
	data map[string]interface{}
	lock *sync.RWMutex
}

func newSegment() *segment {
	return &segment{
		data: make(map[string]interface{}, defaultSegmentSize),
		lock: &sync.RWMutex{},
	}
}

func newSegments(segmentSize int) []*segment {
	segments := make([]*segment, segmentSize)
	for i := 0; i < segmentSize; i++ {
		segments[i] = newSegment()
	}
	return segments
}

type Map struct {
	segments    []*segment
	segmentSize int
	size        int
}

func NewMap() *Map {
	return &Map{
		segmentSize: defaultSegmentSize,
		segments:    newSegments(defaultSegmentSize),
		size:        0,
	}
}

func (m *Map) index(key string) int {
	index := 0
	keyBytes := []byte(key)
	for _, b := range keyBytes {
		index = 31*index + int(b&0xff)
	}
	return index
}

func (m *Map) segmentOf(key string) *segment {
	return m.segments[m.index(key)&(m.segmentSize-1)]
}

func (m *Map) Get(key string) (interface{}, bool) {
	segment := m.segmentOf(key)
	segment.lock.RLock()
	result, ok := segment.data[key]
	segment.lock.RUnlock()
	return result, ok
}

func (m *Map) Set(key string, value interface{}) {
	segment := m.segmentOf(key)
	segment.lock.Lock()
	segment.data[key] = value
	segment.lock.Unlock()
}

func (m *Map) Delete(key string) {
	segment := m.segmentOf(key)
	segment.lock.Lock()
	delete(segment.data, key)
	segment.lock.Unlock()
}

func (m *Map) Size() int {
	return 0
}
