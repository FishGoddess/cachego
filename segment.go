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

package cachego

import (
	"time"
)

type segmentCache struct {
	config

	caches []Cache
}

func newSegmentCache(conf config, newCache func(conf config) Cache) Cache {
	caches := make([]Cache, 0, conf.segments)
	for i := 0; i < conf.segments; i++ {
		caches = append(caches, newCache(conf))
	}

	return &segmentCache{
		config: conf,
		caches: caches,
	}
}

func (sc *segmentCache) Get(key string) (value interface{}, found bool) {
	//TODO implement me
	panic("implement me")
}

func (sc *segmentCache) Set(key string, value interface{}, ttl time.Duration) (oldValue interface{}) {
	//TODO implement me
	panic("implement me")
}

func (sc *segmentCache) Remove(key string) (removedValue interface{}) {
	//TODO implement me
	panic("implement me")
}

func (sc *segmentCache) Clean(allKeys bool) (cleans int) {
	//TODO implement me
	panic("implement me")
}

func (sc *segmentCache) Count(allKeys bool) (count int) {
	//TODO implement me
	panic("implement me")
}
