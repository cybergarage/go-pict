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

	"github.com/cybergarage/go-safecast/safecast"
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
func (elem Elem) Cast(to any) (any, error) {
	err := safecast.To(string(elem), to)
	if err != nil {
		return nil, err
	}
	return to, nil
}

// CastTo casts the element to the specified Type.
func (elem Elem) CastTo(t Type) (any, error) { // nolint: gocognit, gocyclo
	switch t {
	case TypeString:
		var v string
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeByte:
		var v byte
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeBytes:
		var v []byte
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeBool:
		var v bool
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeNil:
		return nil, nil // nolint: nilnil
	case TypeInt:
		var v int
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeInt8:
		var v int8
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeInt16:
		var v int16
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeInt32:
		var v int32
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeInt64:
		var v int64
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeUint:
		var v uint
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeUint8:
		var v uint8
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeUint16:
		var v uint16
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeUint32:
		var v uint32
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeUint64:
		var v uint64
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeFloat32:
		var v float32
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	case TypeFloat64:
		var v float64
		if _, err := elem.Cast(&v); err != nil {
			return nil, err
		}
		return v, nil
	}
	return nil, fmt.Errorf("%s %w", t, ErrNotSupported)
}

// String returns the string representation of the element.
func (elem Elem) String() string {
	return string(elem)
}
