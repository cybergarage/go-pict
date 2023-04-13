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

package picttest

import (
	"testing"

	"github.com/cybergarage/go-pict/pict"
	"github.com/cybergarage/go-pict/test/picts"
)

func TestParser(t *testing.T) {
	for name, pictBytes := range picts.EmbedPicts {
		t.Run(name, func(t *testing.T) {
			parser := pict.NewParserWithBytes(pictBytes)
			err := parser.Parse()
			if err != nil {
				t.Error(err)
			}
		})
	}
}
