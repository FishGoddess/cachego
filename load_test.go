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

import "testing"

// go test -v -cover -run=^TestNewLoader$
func TestNewLoader(t *testing.T) {
	l := NewLoader(nil, false)

	loader1, ok := l.(*loader)
	if !ok {
		t.Errorf("l.(*loader) %T not ok", l)
	}

	if loader1.group != nil {
		t.Errorf("loader1.group %+v != nil", loader1.group)
	}

	l = NewLoader(nil, true)

	loader2, ok := l.(*loader)
	if !ok {
		t.Errorf("l.(*loader) %T not ok", l)
	}

	if loader2.group == nil {
		t.Error("loader2.group == nil")
	}
}
