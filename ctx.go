// Copyright (C) 2016 Space Monkey, Inc.
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
	"context"
	"go.opentelemetry.io/otel/trace"
	"sync"
	"time"
)

// Span represents a 'span' of execution. A span is analogous to a stack frame.
// Spans are constructed as a side-effect of Tasks.
type Span struct {
	trace.Span
	//// sync/atomic things
	//mtx spinLock
	//
	//// immutable things from construction
	//id       int64
	//start    time.Time
	//f        *Func
	//trace    *Trace
	//parent   *Span
	//parentId *int64
	//args     []interface{}
	//context.Context
	//
	//// protected by mtx
	//done        bool
	//orphaned    bool
	//children    spanBag
	//annotations []Annotation
}

// SpanFromCtx loads the current Span from the given context. This assumes
// the context already had a Span created through a Task.
func SpanFromCtx(ctx context.Context) *Span {
	otelSpan := trace.SpanFromContext(ctx)
	return &Span{otelSpan}
}

func newSpan(ctx context.Context, f *Func, tags ...SeriesTag) (sctx context.Context, exit func(*error)) {
	s, span := f.scope.tracer.Start(ctx, f.scope.Name())
	span.SetAttributes(convertTagsToAttributes(tags...)...)
	sctx = s
	return sctx, func(errptr *error) {
		var err error
		if errptr != nil {
			err = *errptr
		}
		span.RecordError(err)
		span.End()
	}
}

var emptyContext = context.Background()

// Tasks are created (sometimes implicitly) from Funcs. A Task should be called
// at the start of a monitored task, and its return value should be called
// at the stop of said task.
type Task func(ctx *context.Context, args ...interface{}) func(*error)

// Task returns a new Task for use, creating an associated Func if necessary.
// It also adds a new Span to the given ctx during execution. Expected usage
// like:
//
//	var mon = monkit.Package()
//
//	func MyFunc(ctx context.Context, arg1, arg2 string) (err error) {
//	  defer mon.Task()(&ctx, arg1, arg2)(&err)
//	  ...
//	}
//
// or
//
//	var (
//	  mon = monkit.Package()
//	  funcTask = mon.Task()
//	)
//
//	func MyFunc(ctx context.Context, arg1, arg2 string) (err error) {
//	  defer funcTask(&ctx, arg1, arg2)(&err)
//	  ...
//	}
//
// Task allows you to include SeriesTags. WARNING: Each unique tag key/value
// combination creates a unique Func and a unique series. SeriesTags should
// only be used for low-cardinality values that you intentionally wish to
// result in a unique series. Example:
//
//	func MyFunc(ctx context.Context, arg1, arg2 string) (err error) {
//	  defer mon.Task(monkit.NewSeriesTag("key1", "val1"))(&ctx)(&err)
//	  ...
//	}
//
// Task uses runtime.Caller to determine the associated Func name. See
// TaskNamed if you want to supply your own name. See Func.Task if you already
// have a Func.
//
// If you want to control Trace creation, see Func.ResetTrace and
// Func.RemoteTrace
func (s *Scope) Task(tags ...SeriesTag) Task {
	var initOnce sync.Once
	var f *Func
	return Task(func(ctx *context.Context, args ...interface{}) func(*error) {
		initOnce.Do(func() {
			f = s.FuncNamed(callerFunc(3), tags...)
		})
		ctx = cleanCtx(ctx)
		s, exit := newSpan(*ctx, f, tags...)
		*ctx = s
		return exit
	})
}

// Task returns a new Task for use on this Func. It also adds a new Span to
// the given ctx during execution.
//
//	var mon = monkit.Package()
//
//	func MyFunc(ctx context.Context, arg1, arg2 string) (err error) {
//	  f := mon.Func()
//	  defer f.Task(&ctx, arg1, arg2)(&err)
//	  ...
//	}
//
// It's more expected for you to use mon.Task directly. See RemoteTrace or
// ResetTrace if you want greater control over creating new traces.
func (f *Func) Task(ctx *context.Context, args ...interface{}) func(*error) {
	ctx = cleanCtx(ctx)
	s, exit := newSpan(*ctx, f)
	*ctx = s
	return exit
}

// RemoteTrace is like Func.Task, except you can specify the trace and parent
// span id.
// Needed for things like the Zipkin plugin.
func (f *Func) RemoteTrace(ctx *context.Context, _ int64, _ *Trace, _ ...interface{}) func(*error) {
	ctx = cleanCtx(ctx)
	s, exit := newSpan(*ctx, f)
	*ctx = s
	return exit
}

// ResetTrace is like Func.Task, except it always creates a new Trace.
func (f *Func) ResetTrace(ctx *context.Context, args ...interface{}) func(*error) {
	ctx = cleanCtx(ctx)
	s, exit := newSpan(*ctx, f)
	*ctx = s
	return exit
}

var newContext = context.Background()

func cleanCtx(ctx *context.Context) *context.Context {
	if ctx == nil || *ctx == nil {
		return &newContext
	}
	return ctx
}

// SpanCtxObserver is the interface plugins must implement if they want to observe
// all spans on a given trace as they happen, or add to contexts as they
// pass through mon.Task()(&ctx)(&err) calls.
type SpanCtxObserver interface {
	// Start is called when a Span starts. Start should return the context
	// this span should use going forward. ctx is the context it is currently
	// using.
	Start(ctx context.Context, s *Span) context.Context

	// Finish is called when a Span finishes, along with an error if any, whether
	// or not it panicked, and what time it finished.
	Finish(ctx context.Context, s *Span, err error, panicked bool, finish time.Time)
}

// ResetContextSpan returns a new context with Span information removed.
func ResetContextSpan(ctx context.Context) context.Context {
	return ctx
}
