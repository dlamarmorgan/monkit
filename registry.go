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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
)

// Registry encapsulates all of the top-level state for a monitoring system.
// In general, only the Default registry is ever used.
type Registry struct {
	meterProvider *metric.MeterProvider
}

// NewRegistry creates a NewRegistry, though you almost certainly just want
// to use Default.
func NewRegistry() *Registry {
	return &Registry{
		meterProvider: metric.NewMeterProvider(),
	}
}

// WithTransformers returns a copy of Registry but with the additional
// CallbackTransformers applied to the Stats method.
func (r *Registry) WithTransformers(t ...CallbackTransformer) *Registry {
	panic("not implemented")
}

// Package creates a new monitoring Scope, named after the top level package.
// It's expected that you'll have something like
//
//	var mon = monkit.Package()
//
// at the top of each package.
func (r *Registry) Package() *Scope {
	return r.ScopeNamed(callerPackage(1))
}

// ScopeNamed is like Package, but lets you choose the name.
func (r *Registry) ScopeNamed(name string) *Scope {
	return &Scope{
		name:   name,
		meter:  r.meterProvider.Meter(name),
		tracer: otel.GetTracerProvider().Tracer(""),
	}
}

// ObserveTraces lets you observe all traces flowing through the system.
// The passed in callback 'cb' will be called for every new trace as soon as
// it starts, until the returned cancel method is called.
// Note: this only applies to all new traces. If you want to find existing
// or running traces, please pull them off of live RootSpans.
func (r *Registry) ObserveTraces(cb func(*Trace)) (cancel func()) {
	// even though observeTrace doesn't get a mutex, it's only ever loading
	// the traceWatcher pointer, so we can use this mutex here to safely
	// coordinate the setting of the traceWatcher pointer.
	//do nothing, no hooks are provided for tracing
	return func() {}
}

// RootSpans will call 'cb' on all currently executing Spans with no live or
// reachable parent. See also AllSpans.
func (r *Registry) RootSpans(cb func(s *Span)) {
	panic("not implemented")
}

// AllSpans calls 'cb' on all currently known Spans. See also RootSpans.
func (r *Registry) AllSpans(cb func(s *Span)) {
	panic("not implemented")
}

// Scopes calls 'cb' on all currently known Scopes.
func (r *Registry) Scopes(cb func(s *Scope)) {
	panic("not implemented")
}

// Funcs calls 'cb' on all currently known Funcs.
func (r *Registry) Funcs(cb func(f *Func)) {
	panic("not implemented")
}

// Stats implements the StatSource interface.
func (r *Registry) Stats(cb func(key SeriesKey, field string, val float64)) {
	panic("not implemented")
}

var _ StatSource = (*Registry)(nil)

// Default is the default Registry
var Default = NewRegistry()

// ScopeNamed is just a wrapper around Default.ScopeNamed
func ScopeNamed(name string) *Scope {
	return Default.ScopeNamed(name)
}

// RootSpans is just a wrapper around Default.RootSpans
func RootSpans(cb func(s Span)) {
	panic("not implemented")
}

// Scopes is just a wrapper around Default.Scopes
func Scopes(cb func(s *Scope)) {
	panic("not implemented")
}

// Funcs is just a wrapper around Default.Funcs
func Funcs(cb func(f *Func)) {
	panic("not implemented")
}

// Package is just a wrapper around Default.Package
func Package() *Scope {
	return Default.ScopeNamed(callerPackage(1))
}

// Stats is just a wrapper around Default.Stats
func Stats(cb func(key SeriesKey, field string, val float64)) {
	panic("not implemented")
}
