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

// Exports and Imports related

package asn1types

type Asn1Xports struct {
	Type       Asn1XportsType          // This is an Export or Import
	FromModule string                  // From Module
	Identifier *Asn1AssignedIdentifier //Optional Assigned Identifier
	Members    []*Asn1Expression
}

type Asn1AssignedIdentifier struct {
	Oid   *Asn1Oid   // Object Identifier
	Value *Asn1Value // Defined Value
}

type Asn1XportsType int

const (
	Asn1XportsTypeExport Asn1XportsType = iota
	Asn1XportsTypeImport
)

func NewAsn1Xports() *Asn1Xports {

	n := &Asn1Xports{}

	return n
}
