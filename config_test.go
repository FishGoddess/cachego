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

package cachego

import (
	"context"
	"testing"
	"time"
)

// go test -v -cover -run=^TestConfig$
func TestConfig(t *testing.T) {
	expect := config{
		mapSize:      256,
		segmentSize:  256,
		gcDuration:   0,
		singleflight: true,
	}

	result := *newDefaultConfig()
	if result != expect {
		t.Errorf("result %+v != expect %+v", result, expect)
	}
}

// go test -v -cover -run=^TestGetConfig$
func TestGetConfig(t *testing.T) {
	expect := opConfig{
		ctx:          context.Background(),
		ttl:          10 * time.Second,
		onMissed:     nil,
		singleflight: true,
	}

	result := *newDefaultGetConfig()
	if result.ctx != expect.ctx {
		t.Errorf("result.ctx %+v != expect.ctx %+v", result.ctx, expect.ctx)
	}

	if result.ttl != expect.ttl {
		t.Errorf("result.ttl %+v != expect.ttl %+v", result.ttl, expect.ttl)
	}

	if result.singleflight != expect.singleflight {
		t.Errorf("result.singleflight %+v != expect.singleflight %+v", result.singleflight, expect.singleflight)
	}
}

// go test -v -cover -run=^TestSetConfig$
func TestSetConfig(t *testing.T) {
	expect := opConfig{
		ttl: noTTL,
	}

	result := *newDefaultSetConfig()
	if result.ttl != expect.ttl {
		t.Errorf("result.ttl %+v != expect.ttl %+v", result.ttl, expect.ttl)
	}
}
