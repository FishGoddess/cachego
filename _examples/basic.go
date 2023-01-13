// Copyright 2023 FishGoddess. All Rights Reserved.
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

package main

import (
	"fmt"
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
	cache := cachego.NewSimpleCache(cachego.WithSegments(0))
	cache.Set("key", 123, 100*time.Millisecond)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 123 true

	count := cache.Count(false)
	fmt.Println(count) // 1

	time.Sleep(200 * time.Millisecond)

	value, ok = cache.Get("key")
	fmt.Println(value, ok) // <nil> false

	count = cache.Count(false)
	fmt.Println(count) // 0
}
