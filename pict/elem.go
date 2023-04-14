// Copyright (C) 2023 The go-pict Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pict

// Elem represents a PICT element.
type Elem string

// newElemsWith returns a new PICT element with the specified string.
func newElemsWith(strs []string) []Elem {
	params := make([]Elem, len(strs))
	for n, v := range strs {
		params[n] = Elem(v)
	}
	return params
}

// String returns the string representation of the element.
func (elem Elem) String() string {
	return string(elem)
}
