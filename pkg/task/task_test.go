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
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type testEntry struct {
	key   string
	value string
}

// go test -v -cover -run=^TestTickerTaskRun$
func TestTickerTaskRun(t *testing.T) {
	before := testEntry{key: "before_key", value: "before_value"}
	fn := testEntry{key: "task_key", value: "task_value"}
	after := testEntry{key: "after_key", value: "after_value"}

	var loop int64
	var result strings.Builder

	beforeFn := func(ctx context.Context) {
		value, ok := ctx.Value(before.key).(string)
		if !ok {
			t.Errorf("ctx.Value(before.key).(string) %+v failed", ctx.Value(before.key))
		}

		if value != before.value {
			t.Errorf("value %s != before.value %s", value, before.value)
		}

		result.WriteString(value)
	}

	mainFn := func(ctx context.Context) {
		value, ok := ctx.Value(fn.key).(string)
		if !ok {
			t.Errorf("ctx.Value(fn.key).(string) %+v failed", ctx.Value(fn.key))
		}

		if value != fn.value {
			t.Errorf("value %s != fn.value %s", value, fn.value)
		}

		atomic.AddInt64(&loop, 1)
		result.WriteString(value)
	}

	afterFn := func(ctx context.Context) {
		value, ok := ctx.Value(after.key).(string)
		if !ok {
			t.Errorf("ctx.Value(after.key).(string) %+v failed", ctx.Value(after.key))
		}

		if value != after.value {
			t.Errorf("value %s != after.value %s", value, after.value)
		}

		result.WriteString(value)
	}

	ctx := context.WithValue(context.Background(), before.key, before.value)
	ctx = context.WithValue(ctx, fn.key, fn.value)
	ctx = context.WithValue(ctx, after.key, after.value)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	go New(mainFn).Context(ctx).Before(beforeFn).After(afterFn).Duration(3 * time.Millisecond).Run()
	time.Sleep(time.Second)

	var expect strings.Builder
	expect.WriteString(before.value)

	for i := int64(0); i < atomic.LoadInt64(&loop); i++ {
		expect.WriteString(fn.value)
	}

	expect.WriteString(after.value)

	if result.String() != expect.String() {
		t.Errorf("result %s != expect %s", result.String(), expect.String())
	}
}
