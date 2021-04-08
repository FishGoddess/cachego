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
	"encoding/json"
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

// debugPointHandler returns debug http handler.
func debugPointHandler(cache *Cache) http.Handler {

	// for responding Json
	writeAsJson := func(writer http.ResponseWriter, data map[string]interface{}) {

		dataBytes, err := json.Marshal(data)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("err: %+v\ndata: %+v", err, data)))
			return
		}
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.Write(dataBytes)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writeAsJson(writer, map[string]interface{}{
			"info": "Welcome to cachego debug point!",
			"cache": map[string]interface{}{
				"mapSize":     cache.mapSize,
				"segmentSize": cache.segmentSize,
			},
			"version": Version,
			"time":    time.Now().Format("2006-01-02 15:04:05"),
			"points": map[string]string{
				"/get": "Get a value from cache, ex: /get?key=test",
				"/remove": "Remove a value from cache, ex: /remove?key=test",
				"/size": "Get size of cache",
				"/gc": "Do gc task",
				"/detail": "List all segments' detail",
			},
		})
	})
	mux.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		data, ok := cache.Get(key)
		writeAsJson(writer, map[string]interface{}{
			"data": data,
			"ok":   ok,
		})
	})
	mux.HandleFunc("/remove", func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")
		cache.Remove(key)
		writeAsJson(writer, map[string]interface{}{
			"ok": true,
		})
	})
	mux.HandleFunc("/size", func(writer http.ResponseWriter, request *http.Request) {
		writeAsJson(writer, map[string]interface{}{
			"size": cache.Size(),
		})
	})
	mux.HandleFunc("/gc", func(writer http.ResponseWriter, request *http.Request) {
		cache.Gc()
		writeAsJson(writer, map[string]interface{}{
			"ok": true,
		})
	})
	mux.HandleFunc("/detail", func(writer http.ResponseWriter, request *http.Request) {
		data := make(map[string]interface{}, cache.segmentSize)
		for i, segment := range cache.segments {
			values := make(map[string]interface{}, len(segment.data))
			for k, v := range segment.data {
				values[k] = map[string]interface{}{
					"data":  v.data,
					"ttl":   v.ttl,
					"ctime": v.ctime.Format("2006-01-02 15:04:05"),
					"alive": v.alive(),
				}
			}
			data[fmt.Sprintf("segment.%d", i)] = map[string]interface{}{
				"size": segment.aliveSize,
				"values":    values,
			}
		}
		writeAsJson(writer, data)
	})

	return mux
}

// WithDebugPoint runs a http server and registers some handlers for debug.
// Don't use it in production!
func WithDebugPoint(address string) Option {
	return func(cache *Cache) {
		go func() {
			err := http.ListenAndServe(address, debugPointHandler(cache))
			if err != nil {
				panic(fmt.Errorf("WithDebugPoint() failed to listen on address [%s] due to %s", address, err.Error()))
			}
		}()
	}
}
