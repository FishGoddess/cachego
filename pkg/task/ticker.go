// Copyright 2022 Ye Zi Jie. All Rights Reserved.
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

package task

import (
	"context"
	"time"
)

type TickerTask struct {
	beforeFn func(ctx context.Context)
	fn       func(ctx context.Context)
	afterFn  func(ctx context.Context)
}

func NewTickerTask(beforeFn, fn, afterFn func(ctx context.Context)) *TickerTask {
	return &TickerTask{
		beforeFn: beforeFn,
		fn:       fn,
		afterFn:  afterFn,
	}
}

func (tt *TickerTask) Run(ctx context.Context, d time.Duration) error {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	tt.beforeFn(ctx)
	defer tt.afterFn(ctx)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			tt.fn(ctx)
		}
	}
}
