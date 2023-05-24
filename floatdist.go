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
//
// WARNING: THE NON-M4 VERSIONS OF THIS FILE ARE GENERATED BY GO GENERATE!
//          ONLY MAKE CHANGES TO THE M4 FILE
//

package monkit

// FloatDist keeps statistics about values such as
// low/high/recent/average/quantiles. Not threadsafe. Construct with
// NewFloatDist(). Fields are expected to be read from but not written to.
type FloatDist struct {
	// Low and High are the lowest and highest values observed since
	// construction or the last reset.
	Low, High float64

	// Recent is the last observed value.
	Recent float64

	// Count is the number of observed values since construction or the last
	// reset.
	Count int64

	// Sum is the sum of all the observed values since construction or the last
	// reset.
	Sum float64

	key       SeriesKey
	reservoir []float32
	rng       string
	sorted    bool
}

// NewFloatDist creates a distribution of float64s.
func NewFloatDist(key SeriesKey) (d *FloatDist) {
	return &FloatDist{}
}

// Insert adds a value to the distribution, updating appropriate values.
func (d *FloatDist) Insert(val float64) {
	return
}

// FullAverage calculates and returns the average of all inserted values.
func (d *FloatDist) FullAverage() float64 {
	return 0
}

// ReservoirAverage calculates the average of the current reservoir.
func (d *FloatDist) ReservoirAverage() float64 {
	return 0
}

// Query will return the approximate value at the given quantile from the
// reservoir, where 0 <= quantile <= 1.
func (d *FloatDist) Query(quantile float64) float64 {
	return 0
}

// Copy returns a full copy of the entire distribution.
func (d *FloatDist) Copy() *FloatDist {
	return &FloatDist{}
}

func (d *FloatDist) Reset() {
	return
}

func (d *FloatDist) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}
