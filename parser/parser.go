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

// Driver for the `go yacc` generated parser and lexer

package parser

import (
	"errors"
	"fmt"
)

type Parser struct {
	name string
	ap   asn1Parser
}

var ParserError = errors.New("parser error")

func NewParser(name string) *Parser {
	ap := asn1NewParser()
	return &Parser{name: name, ap: ap}
}

func (p *Parser) Parse(name, input string, lexdebug bool) error {

	l := lex(name, input, lexdebug)

	err := p.ap.Parse(l)
	if err == 0 {
		return nil
	} else {
		return ParserError
	}
}

func (p *Parser) RunLexer(name, input string, lexdebug bool) error {

	l := lex(name, input, lexdebug)

	for i := l.nextItem(); i.typ != itemEOF; i = l.nextItem() {
		if lexdebug {
			fmt.Println(i)
		}
		if i.typ == itemError {
			return errors.New("Error parsing..")
		}
	}

	return nil
}
