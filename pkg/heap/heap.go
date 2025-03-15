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
	"container/heap"
)

const (
	poppedIndex = -1
)

// Item stores all information needed by heap including value.
type Item struct {
	heap   *Heap
	index  int
	weight uint64

	// Value is the exact data storing in heap.
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

// Weight returns the weight of item.
func (i *Item) Weight() uint64 {
	return i.weight
}

// Adjust adjusts weight of item in order to adjust heap.
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

func (is *items) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*is)

	*is = append(*is, item)
}

func (is *items) Pop() interface{} {
	n := len(*is)
	item := (*is)[n-1]
	*is = (*is)[0 : n-1]

	item.index = poppedIndex // Already popped flag
	return item
}

// Heap uses items to build a heap which always pops the min weight item first.
// It uses weight of item to sort items which may overflow because weight is an uint64 integer.
// When overflow happens, its weight will turn to 0 and become one of the lightest items in heap.
type Heap struct {
	items *items
	size  int
}

// New creates a heap with initialCap of underlying slice.
func New(initialCap int) *Heap {
	return &Heap{
		items: newItems(initialCap),
		size:  0,
	}
}

// Push pushes a value with weight to item and returns the item.
func (h *Heap) Push(weight uint64, value interface{}) *Item {
	index := len(*h.items)
	item := newItem(h, index, weight, value)

	heap.Push(h.items, item)
	h.size++

	return item
}

// Pop pops the min item.
func (h *Heap) Pop() *Item {
	if pop := heap.Pop(h.items); pop != nil {
		h.size--
		return pop.(*Item)
	}

	return nil
}

// Remove removes item from heap and returns its value.
func (h *Heap) Remove(item *Item) interface{} {
	if item.heap == h && item.index != poppedIndex {
		heap.Remove(h.items, item.index)
		h.size--
	}

	return item.Value
}

// Size returns how many items storing in heap.
func (h *Heap) Size() int {
	return h.size
}
