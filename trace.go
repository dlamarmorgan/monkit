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
	"sync"
	"time"
)

// SpanObserver is the interface plugins must implement if they want to observe
// all spans on a given trace as they happen.
type SpanObserver interface {
	// Start is called when a Span starts
	Start(s *Span)

	// Finish is called when a Span finishes, along with an error if any, whether
	// or not it panicked, and what time it finished.
	Finish(s *Span, err error, panicked bool, finish time.Time)
}

// Trace represents a 'trace' of execution. A 'trace' is the collection of all
// of the 'spans' kicked off from the same root execution context. A trace is
// a concurrency-supporting analog of a stack trace, where a span is somewhat
// like a stack frame.
type Trace struct {
	// sync/atomic things
	spanCount int64
	//spanObservers *spanObserverTuple

	// immutable things from construction
	id int64

	// protected by mtx
	mtx  sync.Mutex
	vals map[interface{}]interface{}
}

// NewTrace creates a new Trace.
func NewTrace(id int64) *Trace {
	return &Trace{}
}

// ObserveSpans lets you register a SpanObserver for all future Spans on the
// Trace. The returned cancel method will remove your observer from the trace.
func (t *Trace) ObserveSpans(observer SpanObserver) (cancel func()) {
	return func() {}
}

// ObserveSpansCtx lets you register a SpanCtxObserver for all future Spans on the
// Trace. The returned cancel method will remove your observer from the trace.
func (t *Trace) ObserveSpansCtx(observer SpanCtxObserver) (cancel func()) {
	return func() {}
}

// Id returns the id of the Trace
func (t *Trace) Id() int64 {
	return 0
}

// GetAll returns values associated with a trace. See SetAll.
func (t *Trace) GetAll() (val map[interface{}]interface{}) {
	return map[interface{}]interface{}{}
}

// Get returns a value associated with a key on a trace. See Set.
func (t *Trace) Get(key interface{}) (val interface{}) {
	return 0
}

// Set sets a value associated with a key on a trace. See Get.
func (t *Trace) Set(key, val interface{}) {
	return
}

// Spans returns the number of spans currently associated with the Trace.
func (t *Trace) Spans() int64 {
	return 0
}
