// Copyright (c) 2018 Abhijit Gadgil <gabhijit@iitbombay.org>. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"errors"
	"fmt"
)

type Parser struct {
	name string
}

func NewParser(name string) *Parser {
	return &Parser{name: name}
}

func (p *Parser) Parse(name, input string, output bool) error {

	l := NewLexer(name, input)

	for i := l.nextItem(); i.typ != itemEOF; i = l.nextItem() {
		if output {
			fmt.Println(i)
		}
		if i.typ == itemError {
			return errors.New("Error parsing..")
		}
	}
	return nil
}
