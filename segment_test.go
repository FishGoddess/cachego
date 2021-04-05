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
// Created at 2021/04/05 16:45:53

package cachego

import (
	"strconv"
	"testing"
)

// go test -cover -run=^TestSegmentSize$
func TestSegmentSize(t *testing.T) {

	checkSize := func(t *testing.T, seq int, s *segment, should int) {
		size := s.size()
		if size != should {
			t.Fatalf("Seq %d ==> Segment should have %d size but got %d", seq, should, size)
		}
	}

	s := newSegment(1024)
	checkSize(t, 1, s, 0)

	for i := int64(0); i < 100; i++ {
		s.set(strconv.FormatInt(i, 10), i, 0)
	}
	checkSize(t, 2, s, 100)

	for i := int64(0); i < 110; i++ {
		s.set(strconv.FormatInt(i, 10), i, 0)
	}
	checkSize(t, 3, s, 110)

	for i := int64(0); i < 50; i++ {
		s.remove(strconv.FormatInt(i, 10))
	}
	checkSize(t, 4, s, 60)

	for i := int64(0); i < 50; i++ {
		s.remove(strconv.FormatInt(999999999 + i, 10))
	}
	checkSize(t, 5, s, 60)

	s.removeAll()
	checkSize(t, 6, s, 0)
}
