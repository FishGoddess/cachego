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
// Created at 2021/04/05 22:55:22

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/FishGoddess/cachego"
)

// runSetServer runs a set server for demo.
func runSetServer(cache *cachego.Cache, port string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`
							<html>
							<head>
								<meta charset="UTF-8" />
								<title>set to cache</title>
							</head>
							<body>
								<form action="http://127.0.0.1:` + port + `/set" method="POST">
									key:<input type="text" name="key"/><br>
									value:<input type="text" name="value"/><br>
									ttl:<input type="number" name="ttl"/><br>
									<input type="submit" value="set to cache">
								</form>
							</body>
							</html>`))
	})

	http.HandleFunc("/set", func(writer http.ResponseWriter, request *http.Request) {
		key := request.FormValue("key")
		value := request.FormValue("value")

		ttlStr := request.FormValue("ttl")
		ttl, err := strconv.ParseInt(ttlStr, 10, 64)
		if err != nil {
			log.Printf("inputed ttl %d is not an integer! Error: %s\n", ttlStr, err)
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("inputed ttl %d is not an integer!", ttlStr)))
			return
		}

		cache.SetWithTTL(key, value, ttl)
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`<a href="http://127.0.0.1:` + port + `">http://127.0.0.1:` + port + `</a>`))
	})

	err := http.ListenAndServe("127.0.0.1:"+port, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("You can visit http://127.0.0.1:8080 first and set some data to cache.")
	fmt.Println("Then, you should visit http://127.0.0.1:6789 to use debug point.")
	fmt.Println("(Press ctrl+c / control+c to stop this demo)")

	// Turn on debug point creating cache.
	// :6789 is the address of debug server.
	cache := cachego.NewCache(cachego.WithAutoGC(10 * time.Second), cachego.WithSegmentSize(4), cachego.WithDebugPoint("127.0.0.1:6789")) // try to visit 127.0.0.1:6789

	// Run a set server for demo.
	// 8080 is the address of demo server.
	runSetServer(cache, "8080") // try to visit 127.0.0.1:8080
}
