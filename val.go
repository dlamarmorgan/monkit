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

// IntVal is a convenience wrapper around an IntDist. Constructed using
// NewIntVal, though its expected usage is like:
//
//	var mon = monkit.Package()
//
//	func MyFunc() {
//	  ...
//	  mon.IntVal("size").Observe(val)
//	  ...
//	}
type IntVal struct {
	mtx  sync.Mutex
	dist int
}

// NewIntVal creates an IntVal
func NewIntVal(key SeriesKey) (v *IntVal) {
	return &IntVal{}
}

// Observe observes an integer value
func (v *IntVal) Observe(val int64) {
	return
}

// Stats implements the StatSource interface.
func (v *IntVal) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

// Quantile returns an estimate of the requested quantile of observed values.
// 0 <= quantile <= 1
func (v *IntVal) Quantile(quantile float64) (rv int64) {
	return 0
}

// FloatVal is a convenience wrapper around an FloatDist. Constructed using
// NewFloatVal, though its expected usage is like:
//
//	var mon = monkit.Package()
//
//	func MyFunc() {
//	  ...
//	  mon.FloatVal("size").Observe(val)
//	  ...
//	}
type FloatVal struct {
	mtx  sync.Mutex
	dist FloatDist
}

// NewFloatVal creates a FloatVal
func NewFloatVal(key SeriesKey) (v *FloatVal) {
	return &FloatVal{}
}

// Observe observes an floating point value
func (v *FloatVal) Observe(val float64) {
	return
}

// Stats implements the StatSource interface.
func (v *FloatVal) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

// Quantile returns an estimate of the requested quantile of observed values.
// 0 <= quantile <= 1
func (v *FloatVal) Quantile(quantile float64) (rv float64) {
	return 0
}

// BoolVal keeps statistics about boolean values. It keeps the number of trues,
// number of falses, and the disposition (number of trues minus number of
// falses). Constructed using NewBoolVal, though its expected usage is like:
//
//	var mon = monkit.Package()
//
//	func MyFunc() {
//	  ...
//	  mon.BoolVal("flipped").Observe(bool)
//	  ...
//	}
type BoolVal struct {
	trues  int64
	falses int64
	recent int32
	key    SeriesKey
}

// NewBoolVal creates a BoolVal
func NewBoolVal(key SeriesKey) *BoolVal {
	return &BoolVal{}
}

// Observe observes a boolean value
func (v *BoolVal) Observe(val bool) {
	return
}

// Stats implements the StatSource interface.
func (v *BoolVal) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

// StructVal keeps track of a structure of data. Constructed using
// NewStructVal, though its expected usage is like:
//
//	var mon = monkit.Package()
//
//	func MyFunc() {
//	  ...
//	  mon.StructVal("stats").Observe(stats)
//	  ...
//	}
type StructVal struct {
	mtx    sync.Mutex
	recent interface{}
	key    SeriesKey
}

// NewStructVal creates a StructVal
func NewStructVal(key SeriesKey) *StructVal {
	return &StructVal{}
}

// Observe observes a struct value. Only the fields convertable to float64 will
// be monitored. A reference to the most recently called Observe value is kept
// for reading when Stats is called.
func (v *StructVal) Observe(val interface{}) {
	return
}

// Stats implements the StatSource interface.
func (v *StructVal) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

// DurationVal is a convenience wrapper around an DurationVal. Constructed using
// NewDurationVal, though its expected usage is like:
//
//	var mon = monkit.Package()
//
//	func MyFunc() {
//	  ...
//	  mon.DurationVal("time").Observe(val)
//	  ...
//	}
type DurationVal struct {
	mtx  sync.Mutex
	dist DurationDist
}

// NewDurationVal creates an DurationVal
func NewDurationVal(key SeriesKey) (v *DurationVal) {
	return &DurationVal{}
}

// Observe observes an integer value
func (v *DurationVal) Observe(val time.Duration) {
	return
}

// Stats implements the StatSource interface.
func (v *DurationVal) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

// Quantile returns an estimate of the requested quantile of observed values.
// 0 <= quantile <= 1
func (v *DurationVal) Quantile(quantile float64) (rv time.Duration) {
	return time.Duration(0)
}
