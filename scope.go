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
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// Scope represents a named collection of StatSources. Scopes are constructed
// through Registries.
type Scope struct {
	//r       *Registry
	name string
	//mtx     sync.RWMutex
	//sources map[string]StatSource
	//chains  []StatSource

	meter  metric.Meter
	tracer trace.Tracer
}

// Func retrieves or creates a Func named after the currently executing
// function name (via runtime.Caller. See FuncNamed to choose your own name.
func (s *Scope) Func() *Func {
	return s.FuncNamed(callerFunc(0))
}

// FuncNamed retrieves or creates a Func named using the given name and
// SeriesTags. See Func() for automatic name determination.
//
// Each unique combination of keys/values in each SeriesTag will result in a
// unique Func. SeriesTags are not sorted, so keep the order consistent to avoid
// unintentionally creating new unique Funcs.
func (s *Scope) FuncNamed(name string, tags ...SeriesTag) *Func {
	return &Func{
		scope: &Scope{
			name:   name,
			meter:  s.meter,
			tracer: s.tracer,
		},
	}
}

// Funcs calls 'cb' for all Funcs registered on this Scope.
func (s *Scope) Funcs(cb func(f *Func)) {
	panic("not implemented")
}

// Meter retrieves or creates a Meter named after the given name. See Event.
func (s *Scope) Meter(name string, tags ...SeriesTag) *Meter {
	return &Meter{
		name:  name,
		meter: s.meter,
	}
}

// Event retrieves or creates a Meter named after the given name and then
// calls Mark(1) on that meter.
func (s *Scope) Event(name string, tags ...SeriesTag) {
	_, span := s.tracer.Start(context.Background(), name)
	defer span.End()
	span.AddEvent(name, trace.WithAttributes(convertTagsToAttributes(tags...)...))
}

// DiffMeter retrieves or creates a DiffMeter after the given name and two
// submeters.
func (s *Scope) DiffMeter(name string, m1, m2 *Meter, tags ...SeriesTag) {
	panic("not implemented")
}

// IntVal retrieves or creates an IntVal after the given name.
func (s *Scope) IntVal(name string, tags ...SeriesTag) *IntVal {
	int64Counter, err := s.meter.Int64Counter(name)
	if err != nil {
		panic(err)
	}
	int64Counter.Add(context.Background(), 1)
	return nil
}

// IntValf retrieves or creates an IntVal after the given printf-formatted
// name.
func (s *Scope) IntValf(template string, args ...interface{}) *IntVal {
	panic("not implemented")
}

// FloatVal retrieves or creates a FloatVal after the given name.
func (s *Scope) FloatVal(name string, tags ...SeriesTag) *FloatVal {
	float64Counter, err := s.meter.Float64Counter(name)
	if err != nil {
		panic(err)
	}
	float64Counter.Add(context.Background(), 1)
	return nil
}

// FloatValf retrieves or creates a FloatVal after the given printf-formatted
// name.
func (s *Scope) FloatValf(template string, args ...interface{}) *FloatVal {
	panic("not implemented")
}

// BoolVal retrieves or creates a BoolVal after the given name.
func (s *Scope) BoolVal(name string, tags ...SeriesTag) *BoolVal {
	panic("not implemented")
}

// BoolValf retrieves or creates a BoolVal after the given printf-formatted
// name.
func (s *Scope) BoolValf(template string, args ...interface{}) *BoolVal {
	panic("not implemented")
}

// StructVal retrieves or creates a StructVal after the given name.
func (s *Scope) StructVal(name string, tags ...SeriesTag) *StructVal {
	panic("not implemented")
}

// DurationVal retrieves or creates a DurationVal after the given name.
func (s *Scope) DurationVal(name string, tags ...SeriesTag) *DurationVal {
	return &DurationVal{}
}

// Timer retrieves or creates a Timer after the given name.
func (s *Scope) Timer(name string, tags ...SeriesTag) *Timer {
	panic("not implemented")
}

// Counter retrieves or creates a Counter after the given name.
func (s *Scope) Counter(name string, tags ...SeriesTag) *Counter {
	counter, err := s.meter.Int64UpDownCounter(name)
	if err != nil {
		panic(err)
	}
	return &Counter{counter}
}

// Gauge registers a callback that returns a float as the given name in the
// Scope's StatSource table.
func (s *Scope) Gauge(name string, cb func() float64) {
	panic("not implemented")
}

// Chain registers a full StatSource as the given name in the Scope's
// StatSource table.
func (s *Scope) Chain(source StatSource) {
	//do nothing
}

// Stats implements the StatSource interface.
func (s *Scope) Stats(cb func(key SeriesKey, field string, val float64)) {
	panic("not implemented")
}

// Name returns the name of the Scope, often the Package name.
func (s *Scope) Name() string {
	return s.name
}

var _ StatSource = (*Scope)(nil)

// convert tags to attributes
func convertTagsToAttributes(tags ...SeriesTag) []attribute.KeyValue {
	attributes := make([]attribute.KeyValue, 0, len(tags))
	for _, tag := range tags {
		attributes = append(attributes, attribute.String(tag.Key, tag.Val))
	}
	return attributes
}
