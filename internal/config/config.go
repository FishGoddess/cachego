// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/12/06 00:12:34

package config

import (
	"context"
	"time"
)

const (
	// NoTTL means value lives forever.
	NoTTL = 0
)

var (
	defaultSetConfig = &SetConfig{
		TTL: NoTTL,
	}

	defaultAutoSetConfig = &AutoSetConfig{
		Ctx: context.Background(),
		Gap: time.Minute,
		TTL: NoTTL,
	}

	defaultGetConfig = &GetConfig{
		Ctx:            context.Background(),
		OnMissed:       nil,
		OnMissedSet:    true,
		OnMissedSetTTL: 10 * time.Second,
	}
)

type SetConfig struct {
	TTL time.Duration
}

func NewDefaultSetConfig() *SetConfig {
	return defaultSetConfig
}

type AutoSetConfig struct {
	Ctx context.Context
	Gap time.Duration
	TTL time.Duration
}

func NewDefaultAutoSetConfig() *AutoSetConfig {
	return defaultAutoSetConfig
}

type GetConfig struct {
	Ctx            context.Context
	OnMissed       func(ctx context.Context) (data interface{}, err error)
	OnMissedSet    bool
	OnMissedSetTTL time.Duration
}

func NewDefaultGetConfig() *GetConfig {
	return defaultGetConfig
}
