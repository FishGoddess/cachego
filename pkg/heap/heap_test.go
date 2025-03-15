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

package heap

import (
	"math"
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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestItem$
func TestItem(t *testing.T) {
	heap := New(64)

	item1 := heap.Push(1, 11)
	if item1.index != 0 {
		t.Fatalf("item1.index %d is wrong", item1.index)
	}

	item2 := heap.Push(2, 22)
	if item2.index != 1 {
		t.Fatalf("item2.index %d is wrong", item2.index)
	}

	item3 := heap.Push(3, 33)
	if item3.index != 2 {
		t.Fatalf("item3.index %d is wrong", item3.index)
	}

	if item1.Weight() != 1 || item1.Value.(int) != 11 {
		t.Fatalf("item1.Weight() %d is wrong || item1.Value.(int) %d is wrong", item1.Weight(), item1.Value.(int))
	}

	if item2.Weight() != 2 || item2.Value.(int) != 22 {
		t.Fatalf("item2.Weight() %d is wrong || item2.Value.(int) %d is wrong", item2.Weight(), item2.Value.(int))
	}

	if item3.Weight() != 3 || item3.Value.(int) != 33 {
		t.Fatalf("item3.Weight() %d is wrong || item3.Value.(int) %d is wrong", item3.Weight(), item3.Value.(int))
	}

	item1.Adjust(111)
	if item1.Weight() != 111 || item1.index != 1 {
		t.Fatalf("item1.Weight() %d is wrong || item1.index %d is wrong", item1.Weight(), item1.index)
	}

	if item2.index != 0 {
		t.Fatalf("item2.index %d is wrong", item2.index)
	}

	if item3.index != 2 {
		t.Fatalf("item3.index %d is wrong", item3.index)
	}

	item2.Adjust(222)

	weight := uint64(math.MaxUint64)
	item3.Adjust(weight + 1)

	if item3.Weight() != 0 {
		t.Fatalf("item3.Weight() %d is wrong", item3.Weight())
	}

	expect := []int{33, 11, 22}
	index := 0

	for heap.Size() > 0 {
		num := heap.Pop().Value.(int)
		if num != expect[index] {
			t.Fatalf("num %d != expect[%d] %d", num, index, expect[index])
		}

		index++
	}

	item1.weight = math.MaxUint64
	for i := uint64(1); i <= 3; i++ {
		item1.weight = item1.weight + 1
		if item1.Weight() != i-1 {
			t.Fatalf("item1.Weight() %d != (i %d - 1)", item1.Weight(), i)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestHeap$
func TestHeap(t *testing.T) {
	data := newTestData(10)
	t.Log(data)

	heap := New(64)
	for _, num := range data {
		heap.Push(uint64(num), num)
	}

	if heap.Size() != len(data) {
		t.Fatalf("heap.Size() %d != len(data) %d", heap.Size(), len(data))
	}

	sort.Ints(data)
	t.Log(data)

	index := 0
	for heap.Size() > 0 {
		num := heap.Pop().Value.(int)
		if num != data[index] {
			t.Fatalf("num %d != data[%d] %d", num, index, data[index])
		}

		index++
	}

	if heap.Size() != 0 {
		t.Fatalf("heap.Size() %d is wrong", heap.Size())
	}

	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	items := make([]*Item, 0, len(data))
	for _, num := range data {
		item := heap.Push(uint64(num), num)
		items = append(items, item)
	}

	if heap.Size() != len(data) {
		t.Fatalf("heap.Size() %d != len(data) %d", heap.Size(), len(data))
	}

	if len(items) != len(data) {
		t.Fatalf("len(items) %d != len(data) %d", len(items), len(data))
	}

	for i, num := range data {
		value := heap.Remove(items[i])
		if value.(int) != num {
			t.Fatalf("value.(int) %d != num %d", value.(int), num)
		}
	}

	if heap.Size() != 0 {
		t.Fatalf("heap.Size() %d is wrong", heap.Size())
	}

	item := &Item{heap: heap, index: poppedIndex, Value: 123}
	if value := heap.Remove(item); value.(int) != 123 {
		t.Fatalf("value.(int) %d is wrong", value.(int))
	}
}
