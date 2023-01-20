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

package task

import (
	"context"
	"time"
)

// Task runs a task at fixed duration which will call fn().
type Task struct {
	// Before will be called before running this task.
	Before func(ctx context.Context)

	// Fn is main function which will called in loop.
	Fn func(ctx context.Context)

	// After will be called after the task loop.
	After func(ctx context.Context)
}

// Run runs this task at fixed duration d.
func (t *Task) Run(ctx context.Context, d time.Duration) {
	if t.Fn == nil {
		return
	}

	if t.Before != nil {
		t.Before(ctx)
	}

	if t.After != nil {
		defer t.After(ctx)
	}

	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			t.Fn(ctx)
		}
	}
}
