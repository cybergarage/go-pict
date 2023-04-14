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

import (
	"fmt"
	"strconv"
)

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

// Cast casts the element to the specified type.
func (elem Elem) Cast(to any) (any, error) { // nolint: goerr113
	var err error
	switch v := to.(type) {
	case *string:
		*v = string(elem)
		return *v, nil
	case *[]byte:
		*v = []byte(elem)
		return *v, nil
	case *bool:
		*v, err = strconv.ParseBool(string(elem))
		return *v, err
	case nil:
		return nil, nil
	case *int:
		i, err := strconv.ParseInt(string(elem), 10, 64)
		if err != nil {
			return nil, err
		}
		*v = int(i)
		return *v, nil
	case *int8:
		i, err := strconv.ParseInt(string(elem), 10, 8)
		if err != nil {
			return nil, err
		}
		*v = int8(i)
		return *v, nil
	case *int16:
		i, err := strconv.ParseInt(string(elem), 10, 16)
		if err != nil {
			return nil, err
		}
		*v = int16(i)
		return *v, nil
	case *int32:
		i, err := strconv.ParseInt(string(elem), 10, 32)
		if err != nil {
			return nil, err
		}
		*v = int32(i)
		return *v, nil
	case *int64:
		i, err := strconv.ParseInt(string(elem), 10, 64)
		if err != nil {
			return nil, err
		}
		*v = int64(i)
		return *v, nil
	case *uint:
		i, err := strconv.ParseUint(string(elem), 10, 64)
		if err != nil {
			return nil, err
		}
		*v = uint(i)
		return *v, nil
	case *uint8:
		i, err := strconv.ParseUint(string(elem), 10, 8)
		if err != nil {
			return nil, err
		}
		*v = uint8(i)
		return *v, nil
	case *uint16:
		i, err := strconv.ParseUint(string(elem), 10, 16)
		if err != nil {
			return nil, err
		}
		*v = uint16(i)
		return *v, nil
	case *uint32:
		i, err := strconv.ParseUint(string(elem), 10, 32)
		if err != nil {
			return nil, err
		}
		*v = uint32(i)
		return *v, nil
	case *uint64:
		*v, err = strconv.ParseUint(string(elem), 10, 64)
		if err != nil {
			return nil, err
		}
		return *v, nil
	case *float32:
		f, err := strconv.ParseFloat(string(elem), 32)
		if err != nil {
			return nil, err
		}
		*v = float32(f)
		return *v, nil
	case *float64:
		*v, err = strconv.ParseFloat(string(elem), 64)
		if err != nil {
			return nil, err
		}
		return *v, nil
	}
	return nil, fmt.Errorf("%T %w", to, ErrNotSupported)
}

// Cast casts the element to the specified type.
func (elem Elem) CastType(t string) (any, error) { // nolint: goerr113
	switch t {
	case "string":
		var v string
		return elem.Cast(&v)
	case "bytes":
		var v []byte
		return elem.Cast(&v)
	case "bool":
		var v bool
		return elem.Cast(&v)
	case "nil":
		return nil, nil
	case "int":
		var v int
		return elem.Cast(&v)
	case "int8":
		var v int8
		return elem.Cast(&v)
	case "int16":
		var v int16
		return elem.Cast(&v)
	case "int32":
		var v int32
		return elem.Cast(&v)
	case "int64":
		var v int64
		return elem.Cast(&v)
	case "uint":
		var v uint
		return elem.Cast(&v)
	case "uint8":
		var v uint8
		return elem.Cast(&v)
	case "uint16":
		var v uint16
		return elem.Cast(&v)
	case "uint32":
		var v uint32
		return elem.Cast(&v)
	case "uint64":
		var v uint64
		return elem.Cast(&v)
	case "float32":
		var v float32
		return elem.Cast(&v)
	case "float64":
		var v float64
		return elem.Cast(&v)
	}
	return nil, fmt.Errorf("%s %w", t, ErrNotSupported)
}

// String returns the string representation of the element.
func (elem Elem) String() string {
	return string(elem)
}
