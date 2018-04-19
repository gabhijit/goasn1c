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

// Parsed Values

package asn1types

type Asn1Value struct {
	Type Asn1ValueType
}

type Asn1ValueType int

const (
	Asn1ValueTypeNoValue Asn1ValueType = iota
	Asn1ValueTypeContainingType
	Asn1ValueTypeNull
	Asn1ValueTypeReal
	Asn1ValueTypeInteger
	Asn1ValueTypeMax
	Asn1ValueTypeMin
	Asn1ValueTypeTrue
	Asn1ValueTypeFalse
	Asn1ValueTypeTuple
	Asn1ValueTypeQuadruple
	Asn1ValueTypeUnparsed
	Asn1ValueTypeBitVector
	Asn1ValueTypeValueSet
	Asn1ValueTypeReferenced
	Asn1ValueTypeChoiceIdentifier
)

func NewAsn1Value() *Asn1Value {

	n := &Asn1Value{}

	return n
}
