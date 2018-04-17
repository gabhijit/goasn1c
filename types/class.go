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

// Information Obect Classes related types constants

package asn1types

type Asn1InfoObjectClassCell struct {
	field *Asn1Expression
	value *Asn1Expression
}

//
type Asn1InfoObjectClassRow struct {
	cols     []Asn1InfoObjectClassCell
	ncols    int
	maxidlen int
}

type Asn1WithSyntaxChunk struct {
	typ     WithClauseType
	content interface{}
	next    *Asn1WithSyntaxChunk
}

type ASn1WithSyntax struct {
	chunks []Asn1WithSyntaxChunk
}

type WithClauseType int

const (
	WithClauseLiteral WithClauseType = iota
	WithClauseWhitespace
	WithClauseField
	WithClauseOptionalGroup
)
