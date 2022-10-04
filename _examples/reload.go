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

package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/FishGoddess/cachego"
)

var (
	counter int64 = 0
)

func onMissed(ctx context.Context) (data interface{}, err error) {
	time.Sleep(time.Second) // Simulate blocking operations.
	return atomic.AddInt64(&counter, 1), nil
}

func main() {
	cache := cachego.NewCache()

	for i := 0; i < 10; i++ {
		value, err := cache.Get("key-reload", cachego.WithOpOnMissed(onMissed), cachego.WithOpTTL(time.Second))
		fmt.Println(value, err)
	}

	fmt.Println("=============================")

	for i := 0; i < 10; i++ {
		value, err := cache.Get("key-reload-sleep", cachego.WithOpOnMissed(onMissed), cachego.WithOpTTL(time.Second))
		fmt.Println(value, err)
		time.Sleep(428 * time.Millisecond)
	}

	fmt.Println("=============================")

	for i := 0; i < 10; i++ {
		value, err := cache.Get("key-not-reload", cachego.WithOpOnMissed(onMissed), cachego.WithOpTTL(time.Second), cachego.WithOpDisableReload())
		fmt.Println(value, err)
	}
}
