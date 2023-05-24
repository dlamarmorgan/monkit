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

package monkit

import (
	"go.opentelemetry.io/otel/attribute"
	"time"
)

type ctxKey int

const (
	spanKey ctxKey = iota
)

// Annotation represents an arbitrary name and value string pair
type Annotation struct {
	Name  string
	Value string
}

// Duration returns the current amount of time the Span has been running
func (s *Span) Duration() time.Duration {
	panic("not implemented")
}

// Start returns the time the Span started.
func (s *Span) Start() time.Time {
	panic("not implemented")
}

// Value implements context.Context
func (s *Span) Value(key interface{}) interface{} {
	panic("not implemented")
}

// String implements context.Context
func (s *Span) String() string {
	panic("not implemented")
}

// Children returns all known running child Spans.
func (s *Span) Children(cb func(s *Span)) {
	panic("not implemented")
}

// Args returns the list of strings associated with the args given to the
// Task that created this Span.
func (s *Span) Args() (rv []string) {
	panic("not implemented")
}

// Id returns the Span id.
func (s *Span) Id() int64 {
	panic("not implemented")
}

// ParentId returns the id of the parent Span, if it has a parent.
func (s *Span) ParentId() (int64, bool) {
	panic("not implemented")
}

// Func returns the Func that kicked off this Span.
func (s *Span) Func() *Func {
	panic("not implemented")
}

// Trace returns the Trace this Span is associated with.
func (s *Span) Trace() *Trace {
	return &Trace{}
}

// Annotations returns any added annotations created through the Span Annotate
// method
func (s *Span) Annotations() []Annotation {
	panic("not implemented")
}

// Annotate adds an annotation to the existing Span.
func (s *Span) Annotate(name, val string) {
	s.Span.SetAttributes(attribute.String(name, val))
}

// Orphaned returns true if the Parent span ended before this Span did.
func (s *Span) Orphaned() (rv bool) {
	panic("not implemented")
}
