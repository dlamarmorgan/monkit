// Copyright (C) 2015 Space Monkey, Inc.
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

package collect

import (
	"time"

	"github.com/spacemonkeygo/monkit/v3"
)

// FinishedSpan is a Span that has completed and contains information about
// how it finished.
type FinishedSpan struct {
	Span     *monkit.Span
	Err      error
	Panicked bool
	Finish   time.Time
}

// StartTimeSorter assists with sorting a slice of FinishedSpans by start time.
type StartTimeSorter []*FinishedSpan

func (s StartTimeSorter) Len() int {
	return 0
}
func (s StartTimeSorter) Swap(i, j int) {
	return
}

func (s StartTimeSorter) Less(i, j int) bool {
	return false
}

func (s StartTimeSorter) Sort() {
	return
}
