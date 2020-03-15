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
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/14 14:43:24

package cache

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

// NewCacheValue 方法创建一个 cacheValue 对象。
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

// Ok 方法返回这个数据是否有效，如果数据无效返回 false
func (cv *cacheValue) Ok() bool {
    return cv != InvalidCacheValue()
}

// Value 方法获取实际的缓存数据
func (cv *cacheValue) Value() (interface{}, bool) {
    return cv.value, cv.Ok()
}

// Or 方法会判断当前数据是否有效，如果无效则返回包装了 value 数据的结果
func (cv *cacheValue) Or(value interface{}) *cacheValue {
    if cv.Ok() {
        return cv
    }
    return NewCacheValue(value, 0)
}

// Life 方法返回当前数据剩余寿命时间
func (cv *cacheValue) Life() time.Duration {
    return cv.deadline.Sub(time.Now())
}

func (cv *cacheValue) Dead() bool {
    // cv.deadline.Unix() != int64(NeverDie) 表示这个数据是凡人，是会死的
    // cv.Life() <= 0 表示阳寿已尽
    return cv.deadline.Unix() != int64(NeverDie) && cv.Life() <= 0
}
