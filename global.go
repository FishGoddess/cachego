// Copyright 2022 FishGoddess. All Rights Reserved.
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

var (
	// Index returns an index of this key.
	Index = index
)

// index returns an index of this key.
func index(key string) int {
	index := 1469598103934665603

	keyBytes := []byte(key)
	for _, b := range keyBytes {
		index = (index << 5) - index + int(b&0xff)
		index *= 1099511628211
	}

	return index
}
