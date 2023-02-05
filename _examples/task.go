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
	"context"
	"fmt"
	"time"

	"github.com/FishGoddess/cachego/pkg/task"
)

var (
	contextKey = struct{}{}
)

func beforePrint(ctx context.Context) {
	fmt.Println("before:", ctx.Value(contextKey))
}

func afterPrint(ctx context.Context) {
	fmt.Println("after:", ctx.Value(contextKey))
}

func printContextValue(ctx context.Context) {
	fmt.Println("context value:", ctx.Value(contextKey))
}

func main() {
	// Create a context to stop the task.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Wrap context with key and value
	ctx = context.WithValue(ctx, contextKey, "hello")

	// Use New to create a task and run it.
	// You can use it to load some hot data to cache at fixed duration.
	// Before is called before the task loop, optional.
	// After is called after the task loop, optional.
	// Context is passed to fn include fn/before/after which can stop the task by Done(), optional.
	// Duration is the duration between two loop of fn, optional.
	// Run will start a new goroutine and run the task loop.
	// The task will stop if context is done.
	task.New(printContextValue).
		Before(beforePrint).
		After(afterPrint).
		Context(ctx).
		Duration(time.Second).
		Run()
}
