// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/10/07 20:49:20

package cachego

import (
	"testing"
	"time"

	"github.com/FishGoddess/cachego/internal/config"
)

// go test -v -cover -run=^TestWithMapSize$
func TestWithMapSize(t *testing.T) {
	conf := &config.Config{MapSize: 0}

	WithMapSize(16)(conf)
	if conf.MapSize != 16 {
		t.Errorf("conf.MapSize %d should be 16", conf.MapSize)
	}
}

// go test -v -cover -run=^TestWithSegmentSize$
func TestWithSegmentSize(t *testing.T) {
	conf := &config.Config{SegmentSize: 0}

	WithSegmentSize(16)(conf)
	if conf.SegmentSize != 16 {
		t.Errorf("conf.SegmentSize %d should be 16", conf.SegmentSize)
	}

	for i := uint(1); i < 100000; i *= 2 {
		WithSegmentSize(i)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("WithSegmentSize(13) should panic")
		}
	}()

	WithSegmentSize(13)
}

// go test -v -cover -run=^TestWithAutoGC$
func TestWithAutoGC(t *testing.T) {
	conf := &config.Config{GCDuration: 0}

	d := 10 * time.Minute
	WithAutoGC(d)(conf)
	if conf.GCDuration != d {
		t.Errorf("conf.GCDuration %d should be %s", conf.GCDuration, d)
	}
}
