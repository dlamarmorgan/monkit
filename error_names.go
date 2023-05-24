// Copyright (C) 2017 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monkit

import (
	"sync"
	"sync/atomic"
)

// errorNameHandlers keeps track of the list of error name handlers monkit will
// call to give errors good metric names.
var errorNameHandlers struct {
	write_mu sync.Mutex
	value    atomic.Value
}

// AddErrorNameHandler adds an error name handler function that will be
// consulted every time an error is captured for a task. The handlers will be
// called in the order they were registered with the most recently added
// handler first, until a handler returns true for the second return value.
// If no handler returns true, the error is checked to see if it implements
// an interface that allows it to name itself, and otherwise, monkit attempts
// to find a good name for most built in Go standard library errors.
func AddErrorNameHandler(f func(error) (string, bool)) {
	return
}
