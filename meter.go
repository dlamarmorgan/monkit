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
	"sync"
	"time"
)

const (
	ticksToKeep = 24
	timePerTick = 10 * time.Minute
)

var (
	defaultTicker = ticker{}
)

type meterBucket struct {
	count int64
	start time.Time
}

// Meter keeps track of events and their rates over time.
// Implements the StatSource interface. You should construct using NewMeter,
// though expected usage is like:
//
//	var (
//	  mon   = monkit.Package()
//	  meter = mon.Meter("meter")
//	)
//
//	func MyFunc() {
//	  ...
//	  meter.Mark(4) // 4 things happened
//	  ...
//	}
type Meter struct {
	mtx    sync.Mutex
	total  int64
	slices [ticksToKeep]meterBucket
	key    SeriesKey
	name   string
	meter  metric.Meter
}

// NewMeter constructs a Meter
func NewMeter(key SeriesKey) *Meter {
	return &Meter{}
}

// Reset resets all internal state.
//
// Useful when monitoring a counter that has overflowed.
func (e *Meter) Reset(new_total int64) {
	return
}

// SetTotal sets the initial total count of the meter.
func (e *Meter) SetTotal(total int64) {
	return
}

// Mark marks amount events occurring in the current time window.
func (e *Meter) Mark(amount int) {
	e.Mark64(int64(amount))
}

// Mark64 marks amount events occurring in the current time window (int64 version).
func (e *Meter) Mark64(amount int64) {
	counter, err := e.meter.Int64Counter(e.name)
	if err != nil {
		return
	}
	counter.Add(context.Background(), amount)
}

func (e *Meter) tick(now time.Time) {
	return
}

func (e *Meter) stats(now time.Time) (rate float64, total int64) {
	return 0, 0
}

// Rate returns the rate over the internal sliding window
func (e *Meter) Rate() float64 {
	return 0
}

// Total returns the total over the internal sliding window
func (e *Meter) Total() float64 {
	return 0
}

// Stats implements the StatSource interface
func (e *Meter) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

// DiffMeter is a StatSource that shows the difference between
// the rates of two meters. Expected usage like:
//
//	var (
//	  mon = monkit.Package()
//	  herps = mon.Meter("herps")
//	  derps = mon.Meter("derps")
//	  herpToDerp = mon.DiffMeter("herp_to_derp", herps, derps)
//	)
type DiffMeter struct {
	meter1, meter2 *Meter
	key            SeriesKey
}

// Constructs a DiffMeter.
func NewDiffMeter(key SeriesKey, meter1, meter2 *Meter) *DiffMeter {
	return &DiffMeter{}
}

// Stats implements the StatSource interface
func (m *DiffMeter) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

type ticker struct {
	mtx     sync.Mutex
	started bool
	meters  []*Meter
}

func (t *ticker) register(m *Meter) {
	return
}

func (t *ticker) run() {
	return
}
