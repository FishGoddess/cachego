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
// Created at 2020/03/14 14:43:24
package cachego

import (
	"time"
)

const (
	// NeverDie means value.alive() returns true forever.
	NeverDie = 0
)

// value is a box of data.
type value struct {

	// data stores the real thing inside.
	data interface{}

	// ttl is the life of value.
	// The unit is second.
	ttl int64

	// ctime is the created time of value.
	ctime time.Time
}

// newValue returns a new value with data and ttl.
func newValue(data interface{}, ttl int64) *value {
	return &value{
		data:  data,
		ttl:   ttl,
		ctime: time.Now(),
	}
}

// alive returns if this value is alive or not.
func (v *value) alive() bool {
	return v.ttl == NeverDie || time.Now().Sub(v.ctime).Milliseconds() < v.ttl*1000
}
