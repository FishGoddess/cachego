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
// Created at 2020/03/14 15:28:41

package cache

// *********************************************************
// These functions extend cacheValue to more convenient!     *
//                                     Enjoy yourself!     *
// *********************************************************

// String covers cv.value to string type.
func (cv *cacheValue) String() string {
	result, _ := cv.TryString()
	return result
}

// String covers cv.value to string type.
// If failed, returns "" and false.
func (cv *cacheValue) TryString() (string, bool) {
	result, ok := cv.value.(string)
	return result, ok
}

// Int covers cv.value to int type.
func (cv *cacheValue) Int() int {
	result, _ := cv.TryInt()
	return result
}

// Int covers cv.value to int type.
// If failed, returns 0 and false.
func (cv *cacheValue) TryInt() (int, bool) {
	result, ok := cv.value.(int)
	return result, ok
}

// Int8 covers cv.value to int8 type.
func (cv *cacheValue) Int8() int8 {
	result, _ := cv.TryInt8()
	return result
}

// Int8 covers cv.value to int8 type.
// If failed, returns 0 and false.
func (cv *cacheValue) TryInt8() (int8, bool) {
	result, ok := cv.value.(int8)
	return result, ok
}

// Int16 covers cv.value to int16 type.
func (cv *cacheValue) Int16() int16 {
	result, _ := cv.TryInt16()
	return result
}

// Int16 covers cv.value to int16 type.
// If failed, returns 0 and false.
func (cv *cacheValue) TryInt16() (int16, bool) {
	result, ok := cv.value.(int16)
	return result, ok
}

// Int32 covers cv.value to int32 type.
func (cv *cacheValue) Int32() int32 {
	result, _ := cv.TryInt32()
	return result
}

// Int32 covers cv.value to int32 type.
// If failed, returns 0 and false.
func (cv *cacheValue) TryInt32() (int32, bool) {
	result, ok := cv.value.(int32)
	return result, ok
}

// Int64 covers cv.value to int64 type.
func (cv *cacheValue) Int64() int64 {
	result, _ := cv.TryInt64()
	return result
}

// Int64 covers cv.value to int64 type.
// If failed, returns 0 and false.
func (cv *cacheValue) TryInt64() (int64, bool) {
	result, ok := cv.value.(int64)
	return result, ok
}

// Float32 covers cv.value to float32 type.
func (cv *cacheValue) Float32() float32 {
	result, _ := cv.TryFloat32()
	return result
}

// Float32 covers cv.value to float32 type.
// If failed, returns 0.0 and false.
func (cv *cacheValue) TryFloat32() (float32, bool) {
	result, ok := cv.value.(float32)
	return result, ok
}

// Float64 covers cv.value to float64 type.
func (cv *cacheValue) Float64() float64 {
	result, _ := cv.TryFloat64()
	return result
}

// Float64 covers cv.value to float64 type.
// If failed, returns 0.0 and false.
func (cv *cacheValue) TryFloat64() (float64, bool) {
	result, ok := cv.value.(float64)
	return result, ok
}
