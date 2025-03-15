// Copyright 2025 FishGoddess. All Rights Reserved.
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

import "time"

var (
	mapInitialCap   = 64
	sliceInitialCap = 64
)

func hash(key string) int {
	hash := 1469598103934665603

	for _, r := range key {
		hash = (hash << 5) - hash + int(r&0xffff)
		hash *= 1099511628211
	}

	return hash
}

func now() int64 {
	return time.Now().UnixNano()
}

// SetMapInitialCap sets the initial capacity of map.
func SetMapInitialCap(initialCap int) {
	if initialCap > 0 {
		mapInitialCap = initialCap
	}
}

// SetSliceInitialCap sets the initial capacity of slice.
func SetSliceInitialCap(initialCap int) {
	if initialCap > 0 {
		sliceInitialCap = initialCap
	}
}
