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
// Created at 2020/03/14 15:51:02

package cachego

import (
	"testing"
	"time"
)

// go test -cover -run=^TestValue$
func TestValue(t *testing.T) {

	v := newValue(nil, 1)
	if !v.alive() {
		t.Fatal("V should be alive!")
	}

	time.Sleep(2 * time.Second)
	if v.alive() {
		t.Fatal("V should be dead!")
	}
}
