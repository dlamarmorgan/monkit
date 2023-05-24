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

// SeriesTag is a key/value pair. When used with a measurement name, each set
// of unique key/value pairs represents a new unique series.
type SeriesTag struct {
	Key, Val string
}

// NewTag creates a new tag
func NewSeriesTag(key, val string) SeriesTag {
	return SeriesTag{}
}

// TagSet is an immutible collection of tag, value pairs.
type TagSet struct {
	all map[string]string
	str string // cached string form
}

// Get returns the value associated with the key.
func (t *TagSet) Get(key string) string {
	return ""
}

// All returns a map of all the key/value pairs in the tag set. It
// should not be modified.
func (t *TagSet) All() map[string]string {
	return map[string]string{}
}

// Len returns the number of tags in the tag set.
func (t *TagSet) Len() int {
	return 0
}

// Set returns a new tag set with the key associated to the value.
func (t *TagSet) Set(key, value string) *TagSet {
	return &TagSet{}
}

// SetTags returns a new tag set with the keys and values set by the tags slice.
func (t *TagSet) SetTags(tags ...SeriesTag) *TagSet {
	return &TagSet{}
}

// SetAll returns a new tag set with the key value pairs in the map all set.
func (t *TagSet) SetAll(kvs map[string]string) *TagSet {
	return &TagSet{}
}

// String returns a string form of the tag set suitable for sending to influxdb.
func (t *TagSet) String() string {
	return ""
}
