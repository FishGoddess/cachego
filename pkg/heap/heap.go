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
	"container/heap"
)

type Item struct {
	heap   *Heap
	index  int
	weight uint64

	Value interface{}
}

func newItem(heap *Heap, index int, weight uint64, value interface{}) *Item {
	return &Item{
		heap:   heap,
		index:  index,
		weight: weight,
		Value:  value,
	}
}

func (i *Item) Adjust(weight uint64) {
	i.weight = weight
	heap.Fix(i.heap.items, i.index)
}

type items []*Item

func newItems(initialCap int) *items {
	is := make(items, 0, initialCap)
	heap.Init(&is)
	return &is
}

func (is *items) Len() int {
	return len(*is)
}

func (is *items) Less(i, j int) bool {
	return (*is)[i].weight < (*is)[j].weight
}

func (is *items) Swap(i, j int) {
	(*is)[i], (*is)[j] = (*is)[j], (*is)[i]
	(*is)[i].index = i
	(*is)[j].index = j
}

func (is *items) Pop() interface{} {
	n := len(*is)
	item := (*is)[n-1]
	*is = (*is)[0 : n-1]

	item.index = -1 // 标识该数据已经出堆
	return item
}

func (is *items) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*is)

	*is = append(*is, item)
}

type Heap struct {
	items *items
	size  int
}

func New(initialCap int) *Heap {
	return &Heap{
		items: newItems(initialCap),
		size:  0,
	}
}

func (h *Heap) Push(weight uint64, value interface{}) *Item {
	index := len(*h.items)
	item := newItem(h, index, weight, value)

	heap.Push(h.items, item)
	h.size++

	return item
}

func (h *Heap) Pop() *Item {
	if pop := heap.Pop(h.items); pop != nil {
		h.size--
		return pop.(*Item)
	}

	return nil
}

func (h *Heap) Size() int {
	return h.size
}
