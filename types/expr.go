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

// Expressions parsed by the parser and other stuff

package asn1types

// ASN1Expression is a type representing parsed expression
type Asn1Expression struct {
	Type       Asn1ExprType
	Identifier string
	Meta       Asn1ExprMetaType
}

type Asn1ExprType uint16

// Internal Types
const (
	Asn1ExprTypeInvalid Asn1ExprType = iota

	Asn1ExprTypeReference
	Asn1ExprTypeExportVariable
	Asn1ExprTypeUniversal
	Asn1ExprTypeBitVector
	Asn1ExprTypeOpaque
	Asn1ExprTypeExtensible
	Asn1ExprTypeComponentsOf
	Asn1ExprTypeValueSet
	Asn1ExprTypeClassDef
	Asn1ExprTypeInstance
)

// ClassField Types
const (
	Asn1ExprTypeClsFldTypeFld Asn1ExprType = 0x10 + iota
	Asn1ExprTypeClsFldFixedTypeValueFld
	Asn1ExprTypeClsFldVarTypeValueFld
	Asn1ExprTypeClsFldFixedTypeValueSetFld
	Asn1ExprTypeClsFldVarTypeValueSetFld
	Asn1ExprTypeClsFldObjectFld
	Asn1ExprTypeClsFldObjectSetFld
)

// Constructed Types
const (
	Asn1ExprTypeConstrSeq Asn1ExprType = 0x20 + iota
	Asn1ExprTypeConstrChoice
	Asn1ExprTypeConstrSet
	Asn1ExprTypeConstrSeqOf
	Asn1ExprTypeConstrChoiceOf
)

// Basic Types
const (
	Asn1ExprTypeAny Asn1ExprType = 0x40 + iota // (deprecated)
	Asn1ExprTypeBoolean
	Asn1ExprTypeNull
	Asn1ExprTypeInteger
	Asn1ExprTypeReal
	Asn1ExprTypeEnumerated
	Asn1ExprTypeBitString
	Asn1ExprTypeOctetString
	Asn1ExprTypeObjectIdentifier
	Asn1ExprTypeRelativeOid
	Asn1ExprTypeExternal
	Asn1ExprTypeEmbeddedPdv
	Asn1ExprTypeCharacterString
	Asn1ExprTypeUtcTime
	Asn1ExprTypeGeneralizedTime
)

// String Types - Known Multipliers
const (
	Asn1ExprTypeIA5String Asn1ExprType = 0x100 + iota
	Asn1ExprTypePrintableString
	Asn1ExprTypeVisibleString
	Asn1ExprTypeISO646String
	Asn1ExprTypeNumericString
	Asn1ExprTypeUniversalString
	Asn1ExprTypeBMPString
)

// String Types - Not Known Multipliers
const (
	Asn1ExprTypeUtf8String Asn1ExprType = 0x200 + iota
	Asn1ExprTypeGraphicsString
	Asn1ExprTypeTeletexString
	Asn1ExprTypeT61String
	Asn1ExprTypeVideoTexString
	Asn1ExprTypeObjectDescriptor

	Asn1ExprTypeMax
)

func NewAsn1Expression() *Asn1Expression {
	n := &Asn1Expression{}

	return n
}

type Asn1ExprMetaType int

const (
	Asn1ExprMetaTypeInvalid Asn1ExprMetaType = iota
	Asn1ExprMetaTypeType
	Asn1ExprMetaTypeTypeRef
	Asn1ExprMetaTypeValue
	Asn1ExprMetaTypeValueSet
	Asn1ExprMetaTypeObject
	Asn1ExprMetaTypeObjectClass
	Asn1ExprMetaTypeObjectField
	Asn1ExprMetaTypeMax
)

type Asn1Tag struct {
	Class Asn1TagClass
	Mode  Asn1TagMode
	Val   int64
}

type Asn1TagClass int

const (
	Asn1TagClassNoClass Asn1TagClass = iota
	Asn1TagClassUniversal
	Asn1TagClassApplication
	Asn1TagClassContextSpec
	Asn1TagClassPrivate
)

type Asn1TagMode int

const (
	Asn1TagModeDefault Asn1TagMode = iota
	Asn1TagModeImplicit
	Asn1TagModeExplicit
)

func NewAsn1Tag() *Asn1Tag {
	n := &Asn1Tag{}

	return n
}
