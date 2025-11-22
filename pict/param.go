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

// Param represents a PICT parameter.
type Param string

// newParamsWith returns a new PICT parameter with the specified string.
func newParamsWith(strs []string) []Param {
	params := make([]Param, len(strs))
	for n, v := range strs {
		params[n] = Param(v)
	}
	return params
}

// Type returns the Type of the parameter.
func (param Param) Type() Type {
	return Type(param)
}

// String returns the string representation of the parameter.
func (param Param) String() string {
	return string(param)
}
