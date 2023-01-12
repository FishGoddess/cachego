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

import "github.com/FishGoddess/cachego/options"

// Cache is the core interface of cachego.
// We provide some implements including simple cache and segment cache.
type Cache interface {
	// Get gets key from cache with options and returns value if found.
	// A nil value will be returned if key doesn't exist in cache.
	Get(key string, opts ...options.GetOption) (value interface{}, found bool)

	// Set sets key and value to cache with options.
	Set(key string, value interface{}, opts ...options.SetOption)

	// Remove removes key with options and returns the removed value of key.
	// A nil value will be returned if key doesn't exist in cache.
	Remove(key string, opts ...options.RemoveOption) (value interface{})

	// Clean cleans some keys in cache with options.
	// It returns the exact count cleaned by cache.
	Clean(opts ...options.CleanOption) (cleans int)

	// Size returns the size of cache with options.
	Size(opts ...options.SizeOption) (size int)
}
