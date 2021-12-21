// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/12/18 14:28:56

package singleflight

import (
	"sync"
)

type result struct {
	value interface{}
	err   error
	wg    sync.WaitGroup // For reading result concurrently.
}

func (r *result) do(fn func() (interface{}, error)) {
	defer r.wg.Done()
	r.value, r.err = fn() // Ignore any panics or runtime.Goexit()
}

type Call struct {
	calling bool
	ret     *result
	pool    sync.Pool
	lock    sync.Mutex
}

func NewCall() *Call {
	return &Call{
		calling: false,
		pool: sync.Pool{New: func() interface{} {
			return new(result)
		}},
	}
}

// Do will call fn in single flight mode.
func (c *Call) Do(fn func() (interface{}, error)) (interface{}, error) {
	c.lock.Lock()

	// Wait for result if c is calling.
	ret := c.ret
	if ret != nil {
		c.lock.Unlock()
		ret.wg.Wait()
		return ret.value, ret.err
	}

	// Set c.Ret to non-nil which will block other goroutines.
	ret = c.pool.Get().(*result)
	ret.wg.Add(1)
	c.ret = ret
	c.lock.Unlock()

	ret.do(fn)

	// Set c.Ret to nil which will continue next call.
	c.lock.Lock()
	c.ret = nil
	c.pool.Put(ret)
	c.lock.Unlock()
	return ret.value, ret.err
}
