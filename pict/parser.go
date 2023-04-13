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
	"bytes"
	"io"
)

// Parser represents a PICT parser.
type Parser struct {
	reader io.Reader
}

// NewParserWithReader returns a new PICT parser with the specified reader.
func NewParserWithReader(msgReader io.Reader) *Parser {
	Parser := &Parser{
		reader: msgReader,
	}
	return Parser
}

// NewParserWithBytes returns a new PICT parser with the specified bytes.
func NewParserWithBytes(msgBytes []byte) *Parser {
	return NewParserWithReader(bytes.NewBuffer(msgBytes))
}

// NewParserWithString returns a new PICT parser with the specified string.
func NewParserWithString(msgString string) *Parser {
	return NewParserWithReader(bytes.NewBuffer([]byte(msgString)))
}

func (parser *Parser) Parse() error {
	return nil
}
