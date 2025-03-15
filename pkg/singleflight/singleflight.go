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

package singleflight

import (
	"sync"
)

type call struct {
	fn     func() (result interface{}, err error)
	result interface{}
	err    error

	// deleted is a flag checking if this call has been deleted from Group.
	deleted bool

	wg sync.WaitGroup
}

func newCall(fn func() (result interface{}, err error)) *call {
	return &call{
		fn:      fn,
		deleted: false,
	}
}

func (c *call) do() {
	defer c.wg.Done()

	// Notice: Any panics or runtime.Goexit() happening in fn() will be ignored.
	c.result, c.err = c.fn()
}

// Group stores all function calls in it.
type Group struct {
	calls map[string]*call
	lock  sync.Mutex
}

// NewGroup returns a new Group with initialCap.
func NewGroup(initialCap int) *Group {
	return &Group{
		calls: make(map[string]*call, initialCap),
	}
}

// Call calls fn in singleflight mode and returns its result and error.
func (g *Group) Call(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.lock.Lock()

	if c, ok := g.calls[key]; ok {
		g.lock.Unlock()

		c.wg.Wait()
		return c.result, c.err
	}

	c := newCall(fn)
	c.wg.Add(1)

	g.calls[key] = c
	g.lock.Unlock()

	c.do()
	g.lock.Lock()

	if !c.deleted {
		delete(g.calls, key)
	}

	g.lock.Unlock()
	return c.result, c.err
}

// Delete deletes the call of key so a new call can be called.
func (g *Group) Delete(key string) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if c, ok := g.calls[key]; ok {
		delete(g.calls, key)
		c.deleted = true
	}
}

// Reset resets group to initial status.
func (g *Group) Reset() {
	g.lock.Lock()
	defer g.lock.Unlock()

	for key, c := range g.calls {
		delete(g.calls, key)
		c.deleted = true
	}
}
