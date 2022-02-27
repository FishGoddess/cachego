// Copyright 2020 FishGoddess. All Rights Reserved.
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
	"time"

	"github.com/FishGoddess/cachego"
)

func main() {
	// We use option function to customize the creation of cache.
	// You can just new one without options.
	cache := cachego.NewCache()
	cache.Set("key", "value")

	// You can set it to a cache with automatic gc if you want
	//  Try WithAutoGC.
	cache = cachego.NewCache(cachego.WithAutoGC(10 * time.Minute))

	// Also, you can add more than one option to cache.
	cache = cachego.NewCache(cachego.WithAutoGC(10*time.Minute), cachego.WithMapSize(64), cachego.WithSegmentSize(4096))

	// Remember, some operations have their options, here is one example:
	cache.Get("key", cachego.WithOpOnMissed(func(ctx context.Context) (data interface{}, err error) {
		return "value", nil
	}))
}
