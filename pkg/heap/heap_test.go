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

package heap

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func newTestData(count int) []int {
	random := rand.New(rand.NewSource(time.Now().Unix()))

	var data []int
	for i := 0; i < count; i++ {
		data = append(data, random.Intn(count*10))
	}

	return data
}

// go test -v -cover -run=^TestHeap$
func TestHeap(t *testing.T) {
	data := newTestData(20)
	t.Log(data)

	heap := New(64)
	for _, num := range data {
		heap.Push(uint64(num), num)
	}

	if heap.Size() != len(data) {
		t.Errorf("heap.Size() %d is wrong", heap.Size())
	}

	sort.Ints(data)
	t.Log(data)

	index := 0
	for heap.Size() > 0 {
		num := heap.Pop().Value.(int)
		if num != data[index] {
			t.Errorf("num %d != data[%d] %d", num, index, data[index])
		}

		index++
	}

	if heap.Size() != 0 {
		t.Errorf("heap.Size() %d is wrong", heap.Size())
	}
}
