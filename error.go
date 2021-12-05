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
// Created at 2021/12/06 00:47:35

package cachego

import "fmt"

type notFoundErr struct {
	key string
}

func (nfe *notFoundErr) Error() string {
	if nfe == nil {
		return ""
	}
	return fmt.Sprintf("cachego: key %s not found", nfe.key)
}

func newNotFoundErr(key string) error {
	return &notFoundErr{key: key}
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*notFoundErr)
	return ok
}
