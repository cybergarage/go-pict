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
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Parser represents a PICT parser.
type Parser struct {
	reader io.Reader
	params []Param
	cases  []Case
}

// NewParserWithReader returns a new PICT parser with the specified reader.
func NewParserWithReader(msgReader io.Reader) *Parser {
	Parser := &Parser{
		reader: msgReader,
		params: []Param{},
		cases:  []Case{},
	}
	return Parser
}

// NewParserWithBytes returns a new PICT parser with the specified bytes.
func NewParserWithBytes(msgBytes []byte) *Parser {
	return NewParserWithReader(bytes.NewBuffer(msgBytes))
}

// NewParserWithString returns a new PICT parser with the specified string.
func NewParserWithString(msgString string) *Parser {
	return NewParserWithReader(bytes.NewBufferString(msgString))
}

// Params returns the parameters.
func (parser *Parser) Params() []Param {
	return parser.params
}

// Param returns the parameter at the specified index.
func (parser *Parser) Param(n int) Param {
	return parser.params[n]
}

// Cases returns the all cases.
func (parser *Parser) Cases() []Case {
	return parser.cases
}

// Case returns the case at the specified index.
func (parser *Parser) Case(n int) Case {
	return parser.cases[n]
}

func (parser *Parser) Parse() error {
	reader := bufio.NewReader(parser.reader)

	paramLine, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	parser.params = newParamsWith(strings.Split(string(paramLine), "\t"))

	for {
		caseLine, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		parser.cases = append(parser.cases, newCaseWith(strings.Split(string(caseLine), "\t")))
	}

	return nil
}
