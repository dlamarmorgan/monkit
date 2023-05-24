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

package present

import (
	"github.com/spacemonkeygo/monkit/v3/collect"
	"io"
)

const (
	SampledCBKey = "sampled-cb"
)

// SpansToSVG takes a list of FinishedSpans and writes them to w in SVG format.
// It draws a trace using the Spans where the Spans are ordered by start time.
func SpansToSVG(w io.Writer, spans []*collect.FinishedSpan) error {
	return nil
}

// SpansToJSON turns a list of FinishedSpans into JSON format.
func SpansToJSON(w io.Writer, spans []*collect.FinishedSpan) error {
	return nil
}
