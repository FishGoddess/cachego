// Copyright 2025 FishGoddess. All Rights Reserved.
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

// Task runs a function at fixed duration.
type Task struct {
	ctx      context.Context
	duration time.Duration

	before func(ctx context.Context)
	fn     func(ctx context.Context)
	after  func(ctx context.Context)
}

// New returns a new task for use.
// fn is main function which will called in task loop.
// By default, its duration is 1min, and you can change it by Duration().
func New(fn func(ctx context.Context)) *Task {
	return &Task{
		ctx:      context.Background(),
		duration: time.Minute,
		fn:       fn,
	}
}

// Context sets ctx to task which will be passed to its functions in order to control context.
func (t *Task) Context(ctx context.Context) *Task {
	t.ctx = ctx
	return t
}

// Duration sets duration to task which controls the duration between two task loops.
func (t *Task) Duration(duration time.Duration) *Task {
	t.duration = duration
	return t
}

// Before sets fn to task which will be called before task starting.
func (t *Task) Before(fn func(ctx context.Context)) *Task {
	t.before = fn
	return t
}

// After sets fn to task which will be called after task stopping.
func (t *Task) After(fn func(ctx context.Context)) *Task {
	t.after = fn
	return t
}

// Run runs task.
// You can use context to stop this task, see context.Context.
func (t *Task) Run() {
	if t.fn == nil {
		return
	}

	if t.before != nil {
		t.before(t.ctx)
	}

	if t.after != nil {
		defer t.after(t.ctx)
	}

	ticker := time.NewTicker(t.duration)
	defer ticker.Stop()

	for {
		select {
		case <-t.ctx.Done():
			return
		case <-ticker.C:
			t.fn(t.ctx)
		}
	}
}
