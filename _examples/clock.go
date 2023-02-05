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
	"math/rand"
	"time"

	"github.com/FishGoddess/cachego"
	"github.com/FishGoddess/cachego/pkg/clock"
)

func main() {
	// Create a fast clock and get current time in nanosecond by Now.
	c := clock.New()
	c.Now()

	// Fast clock may return an "incorrect" time compared with time.Now.
	// The gap will be smaller than about 100 ms.
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int63n(int64(time.Second))))

		timeNow := time.Now().UnixNano()
		clockNow := c.Now()

		fmt.Println(timeNow)
		fmt.Println(clockNow)
		fmt.Println("gap:", time.Duration(timeNow-clockNow))
		fmt.Println()
	}

	// You can specify the fast clock to cache by WithNow.
	// All getting current time operations in this cache will use fast clock.
	cache := cachego.NewCache(cachego.WithNow(clock.New().Now))
	cache.Set("key", 666, 100*time.Millisecond)

	value, ok := cache.Get("key")
	fmt.Println(value, ok) // 666, true

	time.Sleep(200 * time.Millisecond)

	value, ok = cache.Get("key")
	fmt.Println(value, ok) // <nil>, false
}
