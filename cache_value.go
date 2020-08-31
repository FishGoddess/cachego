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

import "time"

var (
	// InvalidCacheValue is the representation of invalid cache value.
	invalidCacheValue = NewCacheValue(nil, 0)
)

const (
	// NeverDie means this value will not be dead (cleaned up by gc).
	NeverDie = time.Duration(0)
)

// cacheValue is the struct representation of cached value.
type cacheValue struct {

	// value is the real value of this cacheValue.
	value interface{}

	// deadline is the time when this value will be dead.
	// When time.Now().After(deadline) == true, it's dead.
	deadline time.Time
}

// NewCacheValue returns a new cache value including real cached value and its life.
func NewCacheValue(value interface{}, life time.Duration) *cacheValue {
	var deadline time.Time
	if life != NeverDie {
		deadline = time.Now().Add(life)
	} else {
		deadline = time.Unix(0, 0)
	}

	return &cacheValue{
		value:    value,
		deadline: deadline,
	}
}

// InvalidCacheValue returns an invalid cache value.
func InvalidCacheValue() *cacheValue {
	return invalidCacheValue
}

// Ok returns if this value is valid.
// In current cache, it means this value existed or not normally.
func (cv *cacheValue) Ok() bool {
	return cv != InvalidCacheValue()
}

// Value returns the real value in cache.
// If this value is invalid, then nil and false will be returned.
func (cv *cacheValue) Value() (interface{}, bool) {
	return cv.value, cv.Ok()
}

// Or is for more elegance. As you know, this value may not exist in cache,
// then we need a default value to let our code have more elasticity.
func (cv *cacheValue) Or(value interface{}) *cacheValue {
	if cv.Ok() {
		return cv
	}
	return NewCacheValue(value, 0)
}

// Life returns the leftover life of this value.
// Notice that this has nothing to do with those never die value (such as invalidCacheValue).
// So do not use this method to judge if this value is dead, which is Dead() method doing.
func (cv *cacheValue) Life() time.Duration {
	return cv.deadline.Sub(time.Now())
}

// Dead returns if this value is dead.
// A dead value will be cleaned up by gc.
func (cv *cacheValue) Dead() bool {
	// cv.deadline.Unix() != int64(NeverDie) 表示这个数据是凡人，是会死的
	// cv.Life() <= 0 表示阳寿已尽
	return cv.deadline.Unix() != int64(NeverDie) && cv.Life() <= 0
}
