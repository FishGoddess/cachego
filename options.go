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
// Created at 2021/04/05 15:58:16

package cachego

import (
	"fmt"
	"net/http"
	"time"
)

// Option is a function which initializes cache.
type Option func(cache *Cache)

// WithMapSize is an option setting initializing map size of cache.
func WithMapSize(mapSize int) Option {
	return func(cache *Cache) {
		cache.mapSize = mapSize
	}
}

// WithSegmentSize is an option setting initializing segment size of cache.
func WithSegmentSize(segmentSize int) Option {
	return func(cache *Cache) {
		cache.segmentSize = segmentSize
	}
}

// WithAutoGC is an option turning on automatically gc.
func WithAutoGC(gcDuration time.Duration) Option {
	return func(cache *Cache) {
		cache.AutoGc(gcDuration)
	}
}

// WithDebugPoint runs a http server and registers some handlers for debug.
func WithDebugPoint(address string) Option {
	return func(cache *Cache) {

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Welcome to visit cachego debug point!"))
		})
		go func() {
			err := http.ListenAndServe(address, mux)
			if err != nil {
				panic(fmt.Errorf("WithDebugPoint() failed to listen on address [%s] due to %s", address, err.Error()))
			}
		}()
	}
}
