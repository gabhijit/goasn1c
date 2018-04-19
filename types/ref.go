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

//  Reference to the types defined somewhere else

package asn1types

type Asn1Reference struct {
	Type Asn1ReferenceType
}

type Asn1ReferenceType int

const (
	Asn1ReferenceTypeUnknown Asn1ReferenceType = iota
	Asn1ReferenceTypeCapitals
	Asn1ReferenceTypeUpperCase
	Asn1ReferenceTypeLowerCase
	Asn1ReferenceTypeAmpUpperCase
	Asn1ReferenceTypeAmpLowerCase
	Asn1ReferenceTypeAtLowerCase
	Asn1ReferenceTypeAtDotLowerCase
	Asn1ReferenceTypeMax
)

func NewAsn1Reference() *Asn1Reference {

	n := &Asn1Reference{}

	return n
}
