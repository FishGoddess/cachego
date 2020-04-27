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
// Email: fishinlove@163.com
// Created at 2020/03/14 22:43:18

package cache

import "io"

// Size returns the count of entries in cache.
func (sc *StandardCache) Size() int {
	sc.mu.RLock()
	size := sc.size
	sc.mu.RUnlock()
	return size
}

// Dump is for endurance.
// It will write all alive data by w, which means one gc task will be invoked
// before writing. It will be implemented in future versions...
func (sc *StandardCache) Dump(w io.Writer) {

	// 在持久化数据之前先清理一次死亡过期数据，减少不必要的 IO 操作
	sc.Gc()
}
