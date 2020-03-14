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

// cacheValue is the struct representation of cached value
type cacheValue struct {

    // item 是实际缓存的数据。
    item interface{}

    // valid 指这个数据是否有效，false 表示这个数据不存在，item 是无效的。
    valid bool

    // createdTime 是数据的创建时间，以 ns 为单位。
    createdTime int64

    // deadline 是数据寿命的期限，以 ns 为单位。
    // 当 time.Now().UnixNano() >= deadline 的时候，这个数据就死亡了，等待被回收。
    // 当 deadline == 0 的时候，就意味着这个数据永生，不会被回收。
    deadline int64
}

// NewCacheValue 方法创建一个 cacheValue 对象。
func NewCacheValue(item interface{}, valid bool, life time.Duration) *cacheValue {
    now := time.Now()
    return &cacheValue{
        item:        item,
        valid:       valid,
        createdTime: now.UnixNano(),
        deadline:    now.Add(life).UnixNano(),
    }
}

func (cv *cacheValue) Item() interface{} {
    return cv.item
}

func (cv *cacheValue) Valid() bool {
    return cv.valid
}

func (cv *cacheValue) Value() (interface{}, bool) {
    return cv.Item(), cv.Valid()
}

func (cv *cacheValue) Life() time.Duration {
    return time.Duration(cv.deadline - cv.createdTime)
}
