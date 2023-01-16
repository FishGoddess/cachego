// Copyright 2021 FishGoddess. All Rights Reserved.
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

// call wraps fn and its results to a struct.
type call struct {
	// fn is the target function that will be called.
	fn func() (interface{}, error)

	// result is the successful result of fn.
	result interface{}

	// err is the failed result of fn.
	err error

	// deleted means this call has been deleted from Group.
	deleted bool

	// wg is for reading result concurrently.
	wg sync.WaitGroup
}

// newCall wraps fn to a call holder.
func newCall(fn func() (interface{}, error)) *call {
	return &call{
		fn: fn,
	}
}

// do will call fn and fill results to c.
// Notice: Any panics or runtime.Goexit() will be ignored.
func (c *call) do() {
	defer c.wg.Done()
	c.result, c.err = c.fn()
}

// Group stores all calls in flight.
type Group struct {
	// calls stores all calls in flight.
	calls map[string]*call

	// lock is for safe-concurrency.
	lock sync.Mutex
}

// NewGroup returns a new Group holder with maps.
func NewGroup(maps int) *Group {
	return &Group{
		calls: make(map[string]*call, maps),
	}
}

// Call will call fn in single-flight mode.
func (g *Group) Call(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.lock.Lock()

	if c, ok := g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait() // Waiting for result...
		return c.result, c.err
	}

	c := newCall(fn)
	c.wg.Add(1)
	g.calls[key] = c

	g.lock.Unlock()
	c.do() // Call fn to get result...
	g.lock.Lock()

	if !c.deleted {
		delete(g.calls, key)
	}

	g.lock.Unlock()
	return c.result, c.err
}

// Delete removes the flight of key so a new flight will start.
func (g *Group) Delete(key string) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if c, ok := g.calls[key]; ok {
		c.deleted = true
		delete(g.calls, key)
	}
}

// DeleteAll removes all flights.
func (g *Group) DeleteAll() {
	g.lock.Lock()
	defer g.lock.Unlock()

	for key, c := range g.calls {
		c.deleted = true
		delete(g.calls, key)
	}
}
