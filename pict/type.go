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
	"strings"
)

// Type represents a PICT type.
type Type string

const (
	TypeString  Type = "string"
	TypeBytes   Type = "bytes"
	TypeBool    Type = "bool"
	TypeFloat32 Type = "float32"
	TypeFloat64 Type = "float64"
	TypeUint    Type = "uint"
	TypeUint8   Type = "uint8"
	TypeUint16  Type = "uint16"
	TypeUint32  Type = "uint32"
	TypeUint64  Type = "uint64"
	TypeInt     Type = "int"
	TypeInt8    Type = "int8"
	TypeInt16   Type = "int16"
	TypeInt32   Type = "int32"
	TypeInt64   Type = "int64"
	TypeNil     Type = "nil"
)

// NewTypeFromString returns a new Type from the specified string.
func NewTypeFromString(typeString string) (Type, error) {
	switch strings.ToLower(typeString) {
	case string(TypeString):
		return TypeString, nil
	case string(TypeBytes):
		return TypeBytes, nil
	case string(TypeBool):
		return TypeBool, nil
	case string(TypeFloat32):
		return TypeFloat32, nil
	case string(TypeFloat64):
		return TypeFloat64, nil
	case string(TypeUint):
		return TypeUint, nil
	case string(TypeUint8):
		return TypeUint8, nil
	case string(TypeUint16):
		return TypeUint16, nil
	case string(TypeUint32):
		return TypeUint32, nil
	case string(TypeUint64):
		return TypeUint64, nil
	case string(TypeInt):
		return TypeInt, nil
	case string(TypeInt8):
		return TypeInt8, nil
	case string(TypeInt16):
		return TypeInt16, nil
	case string(TypeInt32):
		return TypeInt32, nil
	case string(TypeInt64):
		return TypeInt64, nil
	case string(TypeNil):
		return TypeNil, nil
	default:
		return "", fmt.Errorf("unsupported type: %s", typeString)
	}
}

// String returns the string representation of the type.
func (t Type) String() string {
	return string(t)
}
