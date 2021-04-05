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
// Created at 2021/04/05 17:31:18

package main

import (
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {

	// We use option function to customize the creation of cache.
	// You can just new it without options.
	cache := cachego.NewCache()
	cache.Set("key", "value")

	// You can set it to a cache with automatic gc if you want
	//  Try WithAutoGC.
	cache = cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Also, you can add more than one option to cache.
	cache = cachego.NewCache(cachego.WithAutoGC(10 * time.Minute), cachego.WithMapSize(64), cachego.WithSegmentSize(4096))

	// Every option has its function, and you should use them for some purposes.
	// WithDebugPoint runs a http server and registers some handlers for debug.
	cachego.WithDebugPoint(":8888") // try to visit :8888
	time.Sleep(time.Minute)
}
