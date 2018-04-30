%{

/*

// Derived from ans1c/libasn1parser/asn1p_y.y @ https://github.com/vlm/asn1c
// Copyright (c) 2003-2017  Lev Walkin <vlm@lionet.info> and contributors.

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


*/


package parser

import (
	"fmt"
	"github.com/gabhijit/goasn1c/types"
/*	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"unicode/utf8"
*/
)

var currentModule *asn1types.Asn1Module
var AllModules    *asn1types.Asn1Grammar

%}

%union {
	ch           rune
	str          string
	num          int64

	grammar         *asn1types.Asn1Grammar
	module          *asn1types.Asn1Module
	oid	        *asn1types.Asn1Oid
	oid_arc         asn1types.Asn1OidArc
	module_flags    asn1types.Asn1ModuleFlagType
	xports          *asn1types.Asn1Xports
	expr            *asn1types.Asn1Expression
	aid             *asn1types.Asn1AssignedIdentifier
	expr_type       asn1types.Asn1ExprType
	tag             *asn1types.Asn1Tag
	value           *asn1types.Asn1Value
	ref             *asn1types.Asn1Reference
	constraint      *asn1types.Asn1Constraint
	marker          *asn1types.Asn1Marker
	constraint_type asn1types.Asn1ConstraintType
	with_syntax     *asn1types.Asn1WithSyntax
}

/* Reserved Keywords Begin - DO NOT INSERT By Hand */
%token	Tok_ABSENT
%token	Tok_ABSTRACT_SYNTAX
%token	Tok_ALL
%token	Tok_APPLICATION
%token	Tok_AUTOMATIC
%token	Tok_BEGIN
%token	Tok_BIT
%token	Tok_BMPString
%token	Tok_BOOLEAN
%token	Tok_BY
%token	Tok_CHARACTER
%token	Tok_CHOICE
%token	Tok_CLASS
%token	Tok_COMPONENT
%token	Tok_COMPONENTS
%token	Tok_CONSTRAINED
%token	Tok_CONTAINING
%token	Tok_DEFAULT
%token	Tok_DEFINITIONS
%token	Tok_EMBEDDED
%token	Tok_ENCODED
%token	Tok_END
%token	Tok_ENUMERATED
%token	Tok_EXCEPT
%token	Tok_EXPLICIT
%token	Tok_EXPORTS
%token	Tok_EXTENSIBILITY
%token	Tok_EXTERNAL
%token	Tok_FALSE
%token	Tok_FROM
%token	Tok_GeneralString
%token	Tok_GeneralizedTime
%token	Tok_GraphicString
%token	Tok_IA5String
%token	Tok_IDENTIFIER
%token	Tok_IMPLICIT
%token	Tok_IMPLIED
%token	Tok_IMPORTS
%token	Tok_INCLUDES
%token	Tok_INSTANCE
%token	Tok_INTEGER
%token	Tok_INTERSECTION
%token	Tok_ISO646String
%token	Tok_MAX
%token	Tok_MIN
%token	Tok_MINUS_INFINITY
%token	Tok_NULL
%token	Tok_NumericString
%token	Tok_OBJECT
%token	Tok_OCTET
%token	Tok_OF
%token	Tok_OPTIONAL
%token	Tok_ObjectDescriptor
%token	Tok_PATTERN
%token	Tok_PDV
%token	Tok_PLUS_INFINITY
%token	Tok_PRESENT
%token	Tok_PRIVATE
%token	Tok_PrintableString
%token	Tok_REAL
%token	Tok_RELATIVE_OID
%token	Tok_SEQUENCE
%token	Tok_SET
%token	Tok_SIZE
%token	Tok_STRING
%token	Tok_SYNTAX
%token	Tok_T61String
%token	Tok_TAGS
%token	Tok_TRUE
%token	Tok_TYPE_IDENTIFIER
%token	Tok_TeletexString
%token	Tok_UNION
%token	Tok_UNIQUE
%token	Tok_UNIVERSAL
%token	Tok_UTCTime
%token	Tok_UTF8String
%token	Tok_UniversalString
%token	Tok_VideotexString
%token	Tok_VisibleString
%token	Tok_WITH
/* Reserved Keywords End - DO NOT INSERT By Hand */

%token       Tok_Assignment
%token       Tok_Ellipsis
%token       Tok_TwoDots
%token       Tok_TwoLeftBrackets Tok_TwoRightBrackets

%token <str> Tok_TypeReference
%token <str> Tok_CAPITALREFERENCE
%token <num> Tok_Number
%token <str> Tok_Identifier
%token <str> Tok_CString
%token <str> Tok_BString
%token <str> Tok_HString
%token <str> Tok_TypeFieldReference
%token <str> Tok_ValueFieldReference

%type <grammar>      ModuleList

%type <module> ModuleDefinition
%type <module> ModuleIdentifier

%type <str>    modulereference

%type <oid>     DefinitiveIdentifier
%type <oid>     DefinitiveObjIdComponentList
%type <oid_arc> DefinitiveObjIdComponent
%type <oid_arc> NameForm
%type <oid_arc> DefinitiveNumberForm
%type <oid_arc> DefinitiveNameAndNumberForm

%type <num>    number
%type <str>    identifier

%type <module_flags>    TagDefault
%type <module_flags>    ExtensionDefault

/*
%type <oid_arc> NumberForm
%type <oid_arc> NameAndNumberForm
*/
%%

ParsedGrammar:
	ModuleList {
		AllModules = $1
		fmt.Println(AllModules)
	}
	;

ModuleList:
	ModuleDefinition {
		$$ = asn1types.NewAsn1Grammar();
		$$.Modules = append($$.Modules, $1)
	}
	| ModuleList ModuleDefinition {
		$$ = $1;
		$$.Modules = append($$.Modules, $2)
	}
	;

ModuleDefinition:
		ModuleIdentifier
		Tok_DEFINITIONS
		TagDefault
		ExtensionDefault
		Tok_Assignment
		Tok_BEGIN
		ModuleBody
		Tok_END {



		}
		;

ModuleIdentifier:
		modulereference DefinitiveIdentifier {
			$$ = asn1types.NewAsn1Module()
		}
		;


modulereference: Tok_TypeReference ;


DefinitiveIdentifier:
		{ }
		|
		'{' DefinitiveObjIdComponentList '}' {

		}
		;

DefinitiveObjIdComponentList:
		DefinitiveObjIdComponent {
			$$ = asn1types.NewAsn1Oid()
			$$.Arcs = append($$.Arcs, $1)
		}
		| DefinitiveObjIdComponentList DefinitiveObjIdComponent {
			$$ = $1
			$$.Arcs = append($$.Arcs, $2)
		}
		;

DefinitiveObjIdComponent:
		NameForm
		| DefinitiveNumberForm
		| DefinitiveNameAndNumberForm
		;

NameForm:
		identifier {
			$$.Name = $1
			$$.Num = -1
		}
		;

DefinitiveNumberForm:
		number {
			$$.Name = ""
			$$.Num = $1
		}
		;

number: Tok_Number;

DefinitiveNameAndNumberForm:
		identifier '(' DefinitiveNumberForm ')' {
			$$.Name = $1
			$$.Num = $3.Num
		}
		;

identifier: Tok_Identifier;

TagDefault:
		{ /* empty tag is explicit */
			$$ = asn1types.Asn1ModuleFlagExplicitTags

		}
		| Tok_EXPLICIT Tok_TAGS {
			$$ = asn1types.Asn1ModuleFlagExplicitTags
		}
		| Tok_IMPLICIT Tok_TAGS {
			$$ = asn1types.Asn1ModuleFlagImplicitTags

		}
		| Tok_APPLICATION Tok_TAGS {
			$$ = asn1types.Asn1ModuleFlagAutomaticTags
		}
		;
ExtensionDefault:
		{
			$$ = asn1types.Asn1ModuleFlagNoFlags
		}
		| Tok_EXTENSIBILITY Tok_IMPLIED {
			$$ = asn1types.Asn1ModuleFlagExtensibilityImplied
		}
		;

ModuleBody: Exports Imports /* AssignmentList */
		;

Exports:
		| Tok_EXPORTS SymbolsExported ';'
		| Tok_EXPORTS Tok_ALL ';'
		;

SymbolsExported:
		| SymbolList
		;

SymbolList:
		Symbol
		| SymbolList ',' Symbol
		;

Symbol:
		Reference
		/* | ParameterizedReference  */
		;

Reference:
		typereference
		| valuereference
		/*
		| objectclassreference
		| objectreference
		| objectsetreference
		*/
		;

typereference:
		Tok_TypeReference
		;

valuereference:
		identifier
		;

Imports:
		| Tok_IMPORTS SymbolsImported ';'
		;

SymbolsImported:
		SymbolsFromModuleList
		;

SymbolsFromModuleList:
		SymbolsFromModule
		| SymbolsFromModuleList ',' SymbolsFromModule
		;

SymbolsFromModule:
		SymbolList Tok_FROM GlobalModuleReference
		;

GlobalModuleReference:
		modulereference AssignedIdentifier
		;

AssignedIdentifier:
		DefinitiveIdentifier

/*
		ObjectIdentifierValue
		| DefinedValue
		|
*/
		;


/* FIXME : We'll need it again when we deal with OBJECT IDENTIFIER ValueAssignment

DefinedValue:
		ExternalValueReference
		| valuereference
		/* | ParameterizedValue * /
		;

ExternalValueReference:
		modulereference '.' valuereference
		;

ObjectIdentifierValue:
		'{' ObjIdComponentsList '}'
		| '{' DefinedValue ObjIdComponentsList '}'
		;

ObjIdComponentsList:
		ObjIdComponents
		| ObjIdComponents ObjIdComponentsList

ObjIdComponents:
		NameForm
		| NumberForm
		| NameAndNumberForm
		| DefinedValue
		;

NumberForm:
		number {
			$$.Name = ""
			$$.Num = $1
		}
		| DefinedValue {
			// FIXME: assign correct values
			$$.Num = -1
			$$.Name = ""
		}
		;

NameAndNumberForm:
		identifier '(' NumberForm ')' {
			$$.Name = $1
			$$.Num = $3.Num
		}
		;
*/
/* AssignmentList:; */
