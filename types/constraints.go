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

type Asn1ConstraintKind uint8

const (
	ConstraintTypeInvalid Asn1ConstraintKind = iota

	// constraint elements
	ConstraintKindSubType // WeekEnd
	ConstraintKindValue
	ConstraintKindSimpleRange            // 1..2
	ConstraintKindLeftExcludedRange      // 0<..MAX (positive integers)
	ConstraintKindRightExcludedRange     // MIN..<0 (negative integers)
	ConstraintKindLeftRightExcludedRange // MIN<..<Max

	// constrained types
	ConstraintKindSize
	ConstraintKindFrom
	ConstraintKindWithComponent
	ConstraintKindWithComponents
	ConstraintKindContstrainedBy
	ConstraintKindContainingType
	ConstraintKindPattern

	ConstraintKindSet
	ConstraintKindCrc // Not sure what is this?
	ConstraintKindUnion
	ConstraintKindIntersection
	ConstraintKindExcept
	ConstraintKindAllExcept
)

type Asn1ConstraintPresentType uint8

const (
	ConstraintDefault Asn1ConstraintPresentType = iota
	ConstraintPresent
	ConstraintAbsent
	ConstraintOptional
)

type Asn1ConstraintType struct {
}
