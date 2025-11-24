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

package pict_test

import (
	"fmt"

	"github.com/cybergarage/go-pict/pict"
)

func Example() {
	pictString :=
		"bool	nil	float32	float64	string	bytes	uint	uint8	uint16	uint32	uint64	int	int8	int16	int32	int64" + "\n" +
			"false	nil	1.0	0.0	\"a\"	\"abc\"	18446744073709551615	0	0	4294967295	0	1	0	-32768	0	9223372036854775807" + "\n" +
			"true	nil	0.0	1.0	\"ab\"	\"ab\"	0	255	65535	0	18446744073709551615	0	1	1	1	9223372036854775807"

	parser := pict.NewParserWithString(pictString)
	err := parser.Parse()
	if err != nil {
		return
	}

	params := parser.Params()
	for i, param := range params {
		paramType, err := param.Type()
		if err != nil {
			return
		}
		if 0 < i {
			fmt.Print(" ")
		}
		fmt.Printf("%v:%v", param.Name(), paramType)
	}
	fmt.Println()

	for _, elems := range parser.Cases() {
		for i, elem := range elems {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(elem)
		}
		fmt.Println()
	}

	// Output:
	// bool:bool nil:nil float32:float32 float64:float64 string:string bytes:bytes uint:uint uint8:uint8 uint16:uint16 uint32:uint32 uint64:uint64 int:int int8:int8 int16:int16 int32:int32 int64:int64
	// false nil 1.0 0.0 "a" "abc" 18446744073709551615 0 0 4294967295 0 1 0 -32768 0 9223372036854775807
	// true nil 0.0 1.0 "ab" "ab" 0 255 65535 0 18446744073709551615 0 1 1 1 9223372036854775807
}
