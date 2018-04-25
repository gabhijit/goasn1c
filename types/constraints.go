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

// Types and Consts related to constraints

package asn1types

type Asn1ConstraintType uint8

const (
	Asn1ConstraintTypeInvalid Asn1ConstraintType = iota

	// constraint elements
	Asn1ConstraintTypeSubType // WeekEnd
	Asn1ConstraintTypeValue
	Asn1ConstraintTypeSimpleRange            // 1..2
	Asn1ConstraintTypeLeftExcludedRange      // 0<..MAX (positive integers)
	Asn1ConstraintTypeRightExcludedRange     // MIN..<0 (negative integers)
	Asn1ConstraintTypeLeftRightExcludedRange // MIN<..<Max
	Asn1ConstraintTypeExtensibilityMark

	// constrained types
	Asn1ConstraintTypeSize
	Asn1ConstraintTypeFrom
	Asn1ConstraintTypeWithComponent
	Asn1ConstraintTypeWithComponents
	Asn1ConstraintTypeContstrainedBy
	Asn1ConstraintTypeContainingType
	Asn1ConstraintTypePattern

	Asn1ConstraintTypeSet
	Asn1ConstraintTypeCrc // Not sure what is this?
	Asn1ConstraintTypeUnion
	Asn1ConstraintTypeIntersection
	Asn1ConstraintTypeExcept
	Asn1ConstraintTypeAllExcept
)

type Asn1ConstraintPresentType uint8

const (
	ConstraintDefault Asn1ConstraintPresentType = iota
	ConstraintPresent
	ConstraintAbsent
	ConstraintOptional
)

type Asn1Constraint struct {
	Type Asn1ConstraintType
}

func NewAsn1Constraint() *Asn1Constraint {

	n := &Asn1Constraint{}

	return n

}
