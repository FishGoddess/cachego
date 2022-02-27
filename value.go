// Copyright 2020 FishGoddess. All Rights Reserved.
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

package cachego

import (
	"time"
)

// value is a box of data.
type value struct {
	// data stores the real thing inside.
	data interface{}

	// ttl is the life of value.
	ttl time.Duration

	// createTime is the created time of value.
	createTime time.Time
}

// newValue returns a new value with data and ttl.
func newValue(data interface{}, ttl time.Duration) *value {
	if ttl < 0 {
		ttl = noTTL // Should panic if ttl < 0?
	}

	return &value{
		data:       data,
		ttl:        ttl,
		createTime: time.Now(),
	}
}

// alive returns if this value is alive or not.
func (v *value) alive() bool {
	return v != nil && (v.ttl == noTTL || time.Since(v.createTime) <= v.ttl)
}

// renew sets v to a new one.
func (v *value) renew(data interface{}, ttl time.Duration) *value {
	if v == nil {
		return nil
	}

	v.data = data
	v.ttl = ttl
	v.createTime = time.Now()
	return v
}
