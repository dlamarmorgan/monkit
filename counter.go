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
	"go.opentelemetry.io/otel/metric"
)

// Counter keeps track of running totals, along with the highest and lowest
// values seen. The overall value can increment or decrement. Counter
// implements StatSource. Should be constructed with NewCounter(), though it
// may be more convenient to use the Counter accessor on a given Scope.
// Expected creation is like:
//
//	var mon = monkit.Package()
//
//	func MyFunc() {
//	  mon.Counter("beans").Inc(1)
//	}

type Counter struct {
	otelCounter metric.Int64UpDownCounter
}

// NewCounter constructs a counter
func NewCounter(key SeriesKey) *Counter {
	return &Counter{}
}

// Set will immediately change the value of the counter to whatever val is. It
// will appropriately update the high and low values, and return the former
// value.
func (c *Counter) Set(val int64) (former int64) {
	panic("not implemented")
}

// Inc will atomically increment the counter by delta and return the new value.
func (c *Counter) Inc(delta int64) (current int64) {
	c.otelCounter.Add(context.Background(), delta)
	return 0
}

// Dec will atomically decrement the counter by delta and return the new value.
func (c *Counter) Dec(delta int64) (current int64) {
	c.otelCounter.Add(context.Background(), delta)
	return 0
}

// High returns the highest value seen since construction or the last reset
func (c *Counter) High() (h int64) {
	panic("not implemented")
}

// Low returns the lowest value seen since construction or the last reset
func (c *Counter) Low() (l int64) {
	panic("not implemented")
}

// Current returns the current value
func (c *Counter) Current() (cur int64) {
	panic("not implemented")
}

// Reset resets all values including high/low counters and returns what they
// were.
func (c *Counter) Reset() (val, low, high int64) {
	panic("not implemented")
}

// Stats implements the StatSource interface
func (c *Counter) Stats(cb func(key SeriesKey, field string, val float64)) {
	panic("not implemented")
}
