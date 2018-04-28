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

// ASN.1 Module related types and structures

package asn1types

// Asn1Module is a type representing ASN.1 Module
type Asn1Module struct {
	Name    string
	Imports []*Asn1Xports
	Exports []*Asn1Xports
	Members []*Asn1Expression
}

type Asn1ModuleFlagType uint16

const (
	Asn1ModuleFlagUnkInstr Asn1ModuleFlagType = 1 << iota
	Asn1ModuleFlagTagInstr
	Asn1ModuleFlagXerInstr
	_
	Asn1ModuleFlagImplicitTags
	Asn1ModuleFlagExplicitTags
	Asn1ModuleFlagAutomaticTags
	_
	Asn1ModuleFlagExtensibilityImplied

	Asn1ModuleFlagNoFlags          Asn1ModuleFlagType = 0x00
	Asn1ModuleFlagInstructionsMask Asn1ModuleFlagType = 0x0F
	Asn1ModuleFlagsTagMask         Asn1ModuleFlagType = 0xF0
)

func NewAsn1Module() *Asn1Module {

	n := &Asn1Module{}

	return n
}

type Asn1Grammar struct {
	Modules []*Asn1Module
}

func NewAsn1Grammar() *Asn1Grammar {

	g := &Asn1Grammar{}
	return g
}
