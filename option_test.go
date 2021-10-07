// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/10/07 20:49:20

package cachego

import "testing"

// go test -v -cover -run=^TestWithMapSize$
func TestWithMapSize(t *testing.T) {
	cache := &Cache{mapSize: 0}

	WithMapSize(16)(cache)
	if cache.mapSize != 16 {
		t.Errorf("cache.mapSize %d should be 16", cache.mapSize)
	}
}

// go test -v -cover -run=^TestWithSegmentSize$
func TestWithSegmentSize(t *testing.T) {
	cache := &Cache{segmentSize: 0}

	WithSegmentSize(16)(cache)
	if cache.segmentSize != 16 {
		t.Errorf("cache.segmentSize %d should be 16", cache.segmentSize)
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
