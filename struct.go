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

type emptyStatSource struct{}

func (emptyStatSource) Stats(cb func(key SeriesKey, field string, val float64)) {}

// StatSourceFromStruct uses the reflect package to implement the Stats call
// across all float64-castable fields of the struct.
func StatSourceFromStruct(key SeriesKey, structData interface{}) StatSource {
	return StatSourceFunc(func(cb func(key SeriesKey, field string, val float64)) {})
}
