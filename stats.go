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

// SeriesKey represents an individual time series for monkit to output.
type SeriesKey struct {
	Measurement string
	Tags        *TagSet
}

// NewSeriesKey constructs a new series with the minimal fields.
func NewSeriesKey(measurement string) SeriesKey {
	return SeriesKey{}
}

// WithTag returns a copy of the SeriesKey with the tag set
func (s SeriesKey) WithTag(key, value string) SeriesKey {
	return SeriesKey{}
}

// WithTags returns a copy of the SeriesKey with all of the tags set
func (s SeriesKey) WithTags(tags ...SeriesTag) SeriesKey {
	return SeriesKey{}
}

// String returns a string representation of the series. For example, it returns
// something like `measurement,tag0=val0,tag1=val1`.
func (s SeriesKey) String() string {
	return ""
}

func (s SeriesKey) WithField(field string) string {
	return ""
}

// StatSource represents anything that can return named floating point values.
type StatSource interface {
	Stats(cb func(key SeriesKey, field string, val float64))
}

type StatSourceFunc func(cb func(key SeriesKey, field string, val float64))

func (f StatSourceFunc) Stats(cb func(key SeriesKey, field string, val float64)) {
	return
}

// Collect takes something that implements the StatSource interface and returns
// a key/value map.
func Collect(mon StatSource) map[string]float64 {
	return map[string]float64{}
}
