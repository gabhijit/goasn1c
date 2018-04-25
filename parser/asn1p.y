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

%{

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
	module_flags    asn1types.ModuleFlagType
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
}

%token       Tok_ALL
%token       Tok_APPLICATION
%token       Tok_ASSIGNMENT
%token       Tok_AUTOMATIC
%token       Tok_BEGIN
%token       Tok_BIT
%token       Tok_BOOLEAN
%token       Tok_CHARACTER
%token       Tok_CHOICE
%token       Tok_COMPONENTS
%token       Tok_DEFAULT
%token       Tok_DEFINITIONS
%token       Tok_EMBEDDED
%token       Tok_END
%token       Tok_ENUMERATED
%token       Tok_EXCEPT
%token       Tok_EXPLICIT
%token       Tok_EXPORTS
%token       Tok_EXTENSIBILITY
%token       Tok_EXTERNAL
%token       Tok_FALSE
%token       Tok_FROM
%token       Tok_GeneralizedTime
%token       Tok_IDENTIFIER
%token       Tok_IMPLICIT
%token       Tok_IMPLIED
%token       Tok_IMPORTS
%token       Tok_INCLUDES
%token       Tok_INTEGER
%token       Tok_INTERSECTION
%token       Tok_MAX
%token       Tok_MIN
%token       Tok_NULL
%token       Tok_OF
%token       Tok_OBJECT
%token       Tok_OCTET
%token       Tok_OPTIONAL
%token       Tok_PDV
%token       Tok_PRIVATE
%token       Tok_REAL
%token       Tok_RELATIVE_OID
%token       Tok_SEQUENCE
%token       Tok_SET
%token       Tok_SIZE
%token       Tok_STRING
%token       Tok_TAGS
%token       Tok_TRUE
%token       Tok_UNION
%token       Tok_UNIVERSAL
%token       Tok_UTCTime

%token       Tok_BMPString
%token       Tok_GeneralString
%token       Tok_GraphicString
%token       Tok_IA5String
%token       Tok_ISO646String
%token       Tok_NumericString
%token       Tok_PrintableString
%token       Tok_T61String
%token       Tok_TeletexString
%token       Tok_UniversalString
%token       Tok_UTF8String
%token       Tok_VideotexString
%token       Tok_VisibleString
%token       Tok_ObjectDescriptor

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

%type <grammar>      ModuleList
%type <str>          TypeRefName
%type <str>          Identifier

%type <module>       ModuleDefinition

%type <module>       optModuleBody
%type <module>       ModuleBody

%type <module>       optImports
%type <module>       ImportsDefinition
%type <module>       optImportsBundleSet
%type <module>       ImportsBundleSet
%type <xports>       ImportsBundle
%type <xports>       ImportsList
%type <expr>         ImportsElement
%type <aid>          AssignedIdentifier

%type <module>       optExports
%type <xports>       ExportsDefinition
%type <xports>       ExportsBody
%type <expr>         ExportsElement

%type <module>       AssignmentList
%type <module>       Assignment
%type <expr>         DataTypeReference

%type <expr>         Type
%type <expr>         TaggedType
%type <expr>         UntaggedType
%type <expr>         BuiltinType
%type <expr>         TypeDeclaration
%type <expr>         ConcreteTypeDeclaration
%type <expr>         DefinedType

%type <expr_type>    BasicTypeId
%type <expr_type>    BasicTypeId_UniverationCompatible

%type <tag>          optTag
%type <tag>          Tag
%type <tag>          TagTypeValue
%type <tag>          TagPlicit
%type <tag>          TagClass

%type <expr>         ValueAssignment
%type <value>        Value
%type <value>        SimpleValue
%type <value>        DefinedValue
%type <value>        SignedNumber
/* %type <value>        RealValue */
%type <value>        RestrictedCharacterStringValue
%type <value>        BitStringValue
%type <value>        IdentifierAsValue
%type <ref>          IdentifierAsReference

%type <ref>          ComplexTypeReference

%type <expr>         NamedNumberList
%type <expr>         NamedNumber

%type <expr>         ValueSetTypeAssignment
%type <constraint>   ValueSet
%type <constraint>   ElementSetSpecs
%type <constraint>   ElementSetSpec
%type <constraint>   Unions
%type <constraint>   Intersections
%type <constraint>   IntersectionElements
%type <constraint>   Elements
%type <constraint>   SubtypeElements
%type <value>        SingleValue

%type <expr>         AlternativeTypeLists
%type <expr>         AlternativeType
%type <expr>         ExtensionAndException

%type <expr>         optComponentTypeLists
%type <expr>         ComponentTypeLists
%type <expr>         ComponentType

%type <marker>       optMarker
%type <marker>       Marker

%type <expr_type>    BasicString

%type <constraint>   optManyConstraints
%type <constraint>   ManyConstraints
%type <constraint>   Constraint
%type <constraint>   ConstraintSpec
%type <constraint>   SubtypeConstraint

%type <constraint>   ValueRange
%type <constraint_type> ConstraintRangeSpec
%type <value>        LowerEndValue
%type <value>        UpperEndValue

%type <value>        ContainedSubtype
%type <expr>         DefinedUntaggedType
%type <constraint>   SizeConstraint
%type <constraint>   PermittedAlphabet

%type <str>          optIdentifier
%type <constraint>   optSizeOrConstraint

%type <expr>         Enumerations
%type <expr>         UniverationList
%type <expr>         UniverationElement

%type <oid>          optObjectIdentifier
%type <oid>          ObjectIdentifier
%type <oid>          ObjectIdentifierBody
%type <oid_arc>      ObjectIdentifierElement
%type <module_flags> optModuleDefinitionFlags
%type <module_flags> ModuleDefinitionFlags
%type <module_flags> ModuleDefinitionFlag


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
		TypeRefName { currentModule = asn1types.NewAsn1Module();}
		optObjectIdentifier Tok_DEFINITIONS
		optModuleDefinitionFlags
		Tok_ASSIGNMENT
		Tok_BEGIN optModuleBody Tok_END {

			$$ = currentModule
			fmt.Println($$)
			$$.Name = $1

			// We have been given a module body
			if ($8 != nil) {
				$$.Imports = $8.Imports
				$$.Exports = $8.Exports
				$$.Members = $8.Members
			}
			fmt.Println($$)
		};

optObjectIdentifier:
	{ $$ = nil; }
	| ObjectIdentifier { $$ = $1; };

ObjectIdentifier:
	'{' ObjectIdentifierBody '}' {
		$$ = $2;
	}
	| '{' '}' {
		$$ = nil;
	}
	;

ObjectIdentifierBody:
	ObjectIdentifierElement {
		$$ = asn1types.NewAsn1Oid();
		$$.Arcs = append($$.Arcs, $1)
	}
	| ObjectIdentifierBody ObjectIdentifierElement {
		$$ = $1;
		$$.Arcs = append($$.Arcs, $2)
	};

ObjectIdentifierElement:
	Identifier {					/* iso */
		$$.Name = $1;
		$$.Num = -1;
	}
	| Identifier '(' Tok_Number ')' {		/* iso(1) */
		$$.Name = $1;
		$$.Num = $3;
	}
	| Tok_Number {					/* 1 */
		$$.Name = "";
		$$.Num = $1;
	};

Identifier: Tok_Identifier {
	  $$ = $1
};

TypeRefName:
	Tok_TypeReference {
		$$ = $1
	}
	| Tok_CAPITALREFERENCE {
		$$ = $1
	};

optModuleDefinitionFlags:
	{ $$ = asn1types.ModuleFlagNoFlags; }
	| ModuleDefinitionFlags {
		$$ = $1;
	}
	;

ModuleDefinitionFlags:
	ModuleDefinitionFlag {
		$$ = $1;
	}
	| ModuleDefinitionFlags ModuleDefinitionFlag {
		$$ = $1 | $2;
	}
	;

ModuleDefinitionFlag:
	Tok_EXPLICIT Tok_TAGS {
		$$ = asn1types.ModuleFlagExplicitTags;
	}
	| Tok_IMPLICIT Tok_TAGS {
		$$ = asn1types.ModuleFlagImplicitTags;
	}
	| Tok_AUTOMATIC Tok_TAGS {
		$$ = asn1types.ModuleFlagAutomaticTags;
	}
	| Tok_EXTENSIBILITY Tok_IMPLIED {
		$$ = asn1types.ModuleFlagExtensibilityImplied;
	} ;

optModuleBody:
	     { $$ = nil;}
	| ModuleBody {
		$$ = $1;
	};

ModuleBody:
	optExports optImports AssignmentList {
		$$ = asn1types.NewAsn1Module()

		if $1 != nil {
			$$.Exports = $1.Exports;
		}
		if $2 != nil {
			$$.Imports = $2.Imports;
		}

		$$.Members = $3.Members
	};

/*
 * === EXAMPLE ===
 * IMPORTS Type1, value FROM Module { iso standard(0) } ;
 * === EOF ===
*/
optImports:
	{ $$ = nil; }
	| ImportsDefinition;

ImportsDefinition:
	Tok_IMPORTS optImportsBundleSet ';' {
		$$ = $2;
	}
	/*
	 * Some error cases.
	 */
	| Tok_IMPORTS Tok_FROM /* ... */ {
		// FIXME: Need to figure out how to call the Lexer's Error()
		return -1;
	}
	;

optImportsBundleSet:
	{ $$ = asn1types.NewAsn1Module(); }
	| ImportsBundleSet;

ImportsBundleSet:
	ImportsBundle {
		$$ = asn1types.NewAsn1Module();
		$$.Imports = append($$.Imports, $1);
	}
	| ImportsBundleSet ImportsBundle {
		$$ = $1;
		$$.Imports = append($$.Imports, $2);
	}
	;

AssignedIdentifier:
	{ $$ = asn1types.NewAsn1AssignedIdentifier(); $$.Oid = nil; $$.Value = nil;}
	| ObjectIdentifier { $$ = asn1types.NewAsn1AssignedIdentifier(); $$.Oid = $1;  $$.Value = nil};

ImportsBundle:
	ImportsList Tok_FROM TypeRefName AssignedIdentifier {
		$$ = $1;

		$$.Type = asn1types.Asn1XportsTypeImport;
		$$.FromModule = $3;
		$$.Identifier = $4;
	}
	;

ImportsList:
	ImportsElement {
		$$ = asn1types.NewAsn1Xports();
		$$.Members = append($$.Members, $1);
	}
	| ImportsList ',' ImportsElement {
		$$ = $1;
		$$.Members = append($$.Members, $3);
	}
	;

ImportsElement:
	TypeRefName {
		$$ = asn1types.NewAsn1Expression();
		$$.Identifier = $1;
		$$.Type = asn1types.Asn1ExprTypeReference;
	}
	| TypeRefName '{' '}' {		/* Completely equivalent to above */
		$$ = asn1types.NewAsn1Expression();
		$$.Identifier = $1;
		$$.Type = asn1types.Asn1ExprTypeReference;
	}
	| Identifier {
		$$ = asn1types.NewAsn1Expression();
		$$.Identifier = $1;
		$$.Type = asn1types.Asn1ExprTypeReference;
	}
	;

optExports:
	{ $$ = nil; }
	| ExportsDefinition {
		$$ = asn1types.NewAsn1Module();
		if($1 == nil) {
			$$.Exports = append($$.Exports, $1);
		} else {
			/* "EXPORTS ALL;" */
		}
	}
	;

ExportsDefinition:
	Tok_EXPORTS ExportsBody ';' {
		$$ = $2;
	}
	| Tok_EXPORTS Tok_ALL ';' {
		$$ = nil;
	}
	| Tok_EXPORTS ';' {
		/* Empty EXPORTS clause effectively prohibits export. */
		$$ = asn1types.NewAsn1Xports();
	}
	;

ExportsBody:
	ExportsElement {
		$$ = asn1types.NewAsn1Xports();
		$$.Type = asn1types.Asn1XportsTypeExport;
		$$.Members = append($$.Members, $1);
	}
	| ExportsBody ',' ExportsElement {
		$$ = $1;
		$$.Type = asn1types.Asn1XportsTypeExport;
		$$.Members = append($$.Members, $3);
	}
	;

ExportsElement:
	TypeRefName {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = $1;
		$$.Type = asn1types.Asn1ExprTypeExportVariable;
	}
	| TypeRefName '{' '}' {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = $1;
		$$.Type = asn1types.Asn1ExprTypeExportVariable;
	}
	| Identifier {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = $1;
		$$.Type = asn1types.Asn1ExprTypeExportVariable;
	}
	;


AssignmentList:
	Assignment {
		$$ = $1;
	}
	| AssignmentList Assignment {
		$$ = $1
		for _, m := range($2.Members) {
			$$.Members = append($$.Members, m)
		}
	}
	;

/* TODO: Big implementation */
Assignment:
	  DataTypeReference {
		$$ = asn1types.NewAsn1Module()
		$$.Members = append($$.Members, $1)
	}
	| ValueAssignment {
		$$ = asn1types.NewAsn1Module();
		$$.Members = append($$.Members, $1)
	}
	| ValueSetTypeAssignment {
		$$ = asn1types.NewAsn1Module();
		$$.Members = append($$.Members, $1)
	}
	;

DataTypeReference:
	TypeRefName Tok_ASSIGNMENT Type {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = $1
		// FIXME : Need to add code for type of expression
	};

Type:TaggedType;


TaggedType:
	optTag UntaggedType {
		$$ = $2
	};

optTag:
	{ $$ = nil; }
	| Tag { $$ = $1; }
	;

Tag:
	TagTypeValue TagPlicit {
		$$ = $1;
		$$.Mode = $2.Mode;
	}
	;

TagTypeValue:
	'[' TagClass Tok_Number ']' {
		$$ = $2;
		$$.Val = $3;
	};

TagClass:
	{ $$ = asn1types.NewAsn1Tag(); $$.Class = asn1types.Asn1TagClassContextSpec; }
	| Tok_UNIVERSAL { $$ = asn1types.NewAsn1Tag(); $$.Class = asn1types.Asn1TagClassUniversal; }
	| Tok_APPLICATION { $$ = asn1types.NewAsn1Tag(); $$.Class = asn1types.Asn1TagClassApplication; }
	| Tok_PRIVATE { $$ = asn1types.NewAsn1Tag(); $$.Class = asn1types.Asn1TagClassPrivate; }
	;

TagPlicit:
	{ $$ = asn1types.NewAsn1Tag(); $$.Mode = asn1types.Asn1TagModeDefault; }
	| Tok_IMPLICIT { $$ = asn1types.NewAsn1Tag(); $$.Mode = asn1types.Asn1TagModeImplicit; }
	| Tok_EXPLICIT { $$ = asn1types.NewAsn1Tag(); $$.Mode = asn1types.Asn1TagModeImplicit; }
	;

UntaggedType:
	    TypeDeclaration optManyConstraints {$$ = $1; }/* FIXME: Constraints not done yet */
	;

TypeDeclaration:
	ConcreteTypeDeclaration { $$ = $1;}
	| DefinedType ;

ConcreteTypeDeclaration:
	BuiltinType
	| Tok_CHOICE '{' AlternativeTypeLists '}' {
		$$ = $3;
		$$.Type = asn1types.Asn1ExprTypeConstrChoice;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
	}
	| Tok_SEQUENCE '{' optComponentTypeLists '}' {
		$$ = $3;
		$$.Type = asn1types.Asn1ExprTypeConstrSequence;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
	}
	| Tok_SET '{' optComponentTypeLists '}' {
		fmt.Println("AAA")
		$$ = $3;
		$$.Type = asn1types.Asn1ExprTypeConstrSet;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
	}
	| Tok_SEQUENCE optSizeOrConstraint Tok_OF optIdentifier optTag TypeDeclaration {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeConstrSetOf;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
		$6.Identifier = $4;
	}
	| Tok_SET optSizeOrConstraint Tok_OF optIdentifier optTag TypeDeclaration {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeConstrSetOf;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
		$6.Identifier = $4;
	};

BuiltinType:
	BasicTypeId {
		$$ = asn1types.NewAsn1Expression();
		$$.Type = $1;

	}
	| Tok_INTEGER '{' NamedNumberList '}' {
		$$ = $3;
		$$.Type = asn1types.Asn1ExprTypeInteger;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
    	}
	| Tok_ENUMERATED '{' Enumerations '}' {
		$$ = $3;
		$$.Type = asn1types.Asn1ExprTypeEnumerated;
		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
	}
	;

BasicTypeId:
	Tok_BOOLEAN { $$ = asn1types.Asn1ExprTypeBoolean; }
	| Tok_NULL { $$ = asn1types.Asn1ExprTypeNull; }
	| Tok_REAL { $$ = asn1types.Asn1ExprTypeReal; }
	| Tok_OCTET Tok_STRING { $$ = asn1types.Asn1ExprTypeOctetString; }
	| Tok_EMBEDDED Tok_PDV { $$ = asn1types.Asn1ExprTypeEmbeddedPdv; }
	| Tok_CHARACTER Tok_STRING { $$ = asn1types.Asn1ExprTypeCharacterString; }
	| Tok_UTCTime { $$ = asn1types.Asn1ExprTypeUtcTime; }
	| Tok_GeneralizedTime { $$ = asn1types.Asn1ExprTypeGeneralizedTime; }
	| Tok_OBJECT Tok_IDENTIFIER { $$ = asn1types.Asn1ExprTypeObjectIdentifier; }
	| Tok_RELATIVE_OID { $$ = asn1types.Asn1ExprTypeRelativeOid; }
	| Tok_EXTERNAL { $$ = asn1types.Asn1ExprTypeExternal; }
	| BasicString
	| BasicTypeId_UniverationCompatible
	;

BasicTypeId_UniverationCompatible:
	Tok_INTEGER { $$ = asn1types.Asn1ExprTypeInteger; }
	| Tok_ENUMERATED { fmt.Println("DDD"); $$ = asn1types.Asn1ExprTypeEnumerated; }
	| Tok_BIT Tok_STRING { $$ = asn1types.Asn1ExprTypeBitString; }
	;

/*
 * === EXAMPLE ===
 * value INTEGER ::= 1
 * === EOF ===
 */
ValueAssignment:
	Identifier Type Tok_ASSIGNMENT Value {
		$$ = $2;
		$$.Identifier = $1;
		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
	}
	;


Value:
	SimpleValue
	| DefinedValue
/* TODO: Opaque not supported yet. Need to understand this first
	| '{' { asn1p_lexer_hack_push_opaque_state(); } Opaque {
		$$ = asn1p_value_frombuf($3.buf, $3.len, 0);
		checkmem($$);
		$$->type = ATV_UNPARSED;
	}
*/
    ;

SimpleValue:
	Tok_NULL {
		$$ = asn1types.NewAsn1Value();
		$$.Type = asn1types.Asn1ValueTypeNull;
	}
	| Tok_FALSE {
		$$ = asn1types.NewAsn1Value();
		$$.Type = asn1types.Asn1ValueTypeFalse;
	}
	| Tok_TRUE {
		$$ = asn1types.NewAsn1Value();
		$$.Type = asn1types.Asn1ValueTypeTrue;
	}
	| SignedNumber
	/* TODO : | RealValue */
	| RestrictedCharacterStringValue
	| BitStringValue
	;

DefinedValue:
	IdentifierAsValue
	| TypeRefName '.' Identifier {
		ref := asn1types.NewAsn1Reference()
		_ = ref
		$$ = asn1types.NewAsn1Value()
	}
	;


RestrictedCharacterStringValue:
	Tok_CString {
		$$ = asn1types.NewAsn1Value()
	}
	| '{' Tok_Number ',' Tok_Number '}' {
		if $2 < 0 || $2 > 7 {
			return -1;
		}
		if $4 < 0 || $4 > 15 {
			return -1;
		}
		//value := $2 << 4 + $4;

		$$ = asn1types.NewAsn1Value()
		$$.Type = asn1types.Asn1ValueTypeTuple;
	}
	| '{' Tok_Number ',' Tok_Number ',' Tok_Number ',' Tok_Number '}' {
		if $2 < 0 || $2 > 127 {
			return -1
		}
		if $4 < 0 || $4 > 255 {
			return -1
		}
		if $6 < 0 || $6 > 255 {
			return -1
		}
		if $8 < 0 || $8 > 255 {
			return -1
		}
		$$ = asn1types.NewAsn1Value()
		$$.Type = asn1types.Asn1ValueTypeQuadruple;
	}
	;

SignedNumber:
	Tok_Number {
		$$ = asn1types.NewAsn1Value()
	}
	;

IdentifierAsReference:
    Identifier {
		$$ = asn1types.NewAsn1Reference()
    };

IdentifierAsValue:
    IdentifierAsReference {
		$$ = asn1types.NewAsn1Value()
    };

BitStringValue:
	Tok_BString {
		$$ = asn1types.NewAsn1Value()
	}
	| Tok_HString {
		$$ = asn1types.NewAsn1Value()
	}
	;

Enumerations:
    UniverationList {
		$$ = $1;
		// TODO handle Enumeration validation
    }

UniverationList:
	UniverationElement {
		$$ = asn1types.NewAsn1Expression()
	}
	| UniverationList ',' UniverationElement {
		$$ = $1;
		// $$.Members = append($$.Members, $3)
	}
	;

UniverationElement:
	Identifier {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeUniversal;

		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
		$$.Identifier = $1;
	}
	| Identifier '(' SignedNumber ')' {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeUniversal;

		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
		$$.Identifier = $1;
	}
	| Identifier '(' DefinedValue ')' {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeUniversal;

		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
		$$.Identifier = $1;
	}
	| SignedNumber {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeUniversal;

		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
		//$$.Identifier = $1;
	}
	| Tok_Ellipsis {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeExtensible;

		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
		$$.Identifier = "...";
	}
	;

DefinedType:
	/*
	 * A DefinedType reference.
	 * "CLASS1.&id.&id2"
	 * or
	 * "Module.Type"
	 * or
	 * "Module.identifier"
	 * or
	 * "Type"
	 */
	ComplexTypeReference {
		$$ = asn1types.NewAsn1Expression();
		$$.Type = asn1types.Asn1ExprTypeReference;
		$$.Meta = asn1types.Asn1ExprMetaTypeTypeRef;
	}
	/* TODO: parameterized Types not supported
	 * A parameterized assignment.
	| ComplexTypeReference '{' ActualParameterList '}' {
		$$ = NEW_EXPR();
		checkmem($$);
		$$->reference = $1;
		$$->rhs_pspecs = $3;
		$$->expr_type = A1TC_REFERENCE;
		$$->meta_type = AMT_TYPEREF;
	}
	 */
	;

ComplexTypeReference:
		    /* FIXME : More needs to be done about references */
	Tok_TypeReference {
		$$ = asn1types.NewAsn1Reference()
	}
	| Tok_CAPITALREFERENCE {
		$$ = asn1types.NewAsn1Reference()
	}
	| Tok_TypeReference '.' TypeRefName {
		fmt.Println("XXX")
		$$ = asn1types.NewAsn1Reference()
	}
	;

NamedNumberList:
	NamedNumber {
		$$ = asn1types.NewAsn1Expression()
	}
	| NamedNumberList ',' NamedNumber {
		$$ = $1;
	}
	;

NamedNumber:
	Identifier '(' SignedNumber ')' {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeUniversal;
		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
		$$.Identifier = $1;
	}
	| Identifier '(' DefinedValue ')' {
		$$ = asn1types.NewAsn1Expression()
		$$.Type = asn1types.Asn1ExprTypeUniversal;
		$$.Meta = asn1types.Asn1ExprMetaTypeValue;
		$$.Identifier = $1;
	};

/*
 * Data type constraints.
 */
UnionMark:		'|' | Tok_UNION;
IntersectionMark:	'^' | Tok_INTERSECTION;

ValueSet: '{' ElementSetSpecs '}' { $$ = $2; };

ValueSetTypeAssignment:
	TypeRefName Type Tok_ASSIGNMENT ValueSet {
		$$ = $2;
		$$.Identifier = $1;
		$$.Meta = asn1types.Asn1ExprMetaTypeValueSet;
	}
	;

ElementSetSpecs:
	Tok_Ellipsis  {
		$$ = asn1types.NewAsn1Constraint()
		$$.Type = asn1types.Asn1ConstraintTypeExtensibilityMark;
	}
   | ElementSetSpec
   | ElementSetSpec ',' Tok_Ellipsis {
       $$ = asn1types.NewAsn1Constraint()
       $$.Type = asn1types.Asn1ConstraintTypeExtensibilityMark;
   }
   | ElementSetSpec ',' Tok_Ellipsis ',' ElementSetSpec {
       $$ = asn1types.NewAsn1Constraint()
       $$.Type = asn1types.Asn1ConstraintTypeExtensibilityMark;
   }
;

ElementSetSpec:
	Unions
	| Tok_ALL Tok_EXCEPT Elements {
	}
	;

Unions:
	Intersections
	| Unions UnionMark Intersections {
	}
	;

Intersections:
	IntersectionElements
	|  Intersections IntersectionMark IntersectionElements {
	}
	;


IntersectionElements:
	Elements
	| Elements Tok_EXCEPT Elements {
	}
	;

Elements:
    SubtypeElements
    | '(' ElementSetSpec ')' {
	$$ = asn1types.NewAsn1Constraint();
	$$.Type = asn1types.Asn1ConstraintTypeSet;
    }
    ;

SubtypeElements:
	SingleValue {
		$$ = asn1types.NewAsn1Constraint();
		$$.Type = asn1types.Asn1ConstraintTypeValue;
	}
	| ContainedSubtype {
		$$ = asn1types.NewAsn1Constraint();
		$$.Type = asn1types.Asn1ConstraintTypeSubType;
		/* FIXME: */
	}
	| ValueRange
    	| SizeConstraint    /* SIZE ... */
	| PermittedAlphabet /* FROM ... */

	;


AlternativeTypeLists:
	AlternativeType {
		$$ = asn1types.NewAsn1Expression()
	}
	| AlternativeTypeLists ',' AlternativeType {
		$$ = asn1types.NewAsn1Expression()
	}
	;

AlternativeType:
	Identifier TaggedType {
		$$ = $2;
	}
	| ExtensionAndException {
		$$ = $1;
	}
	| TaggedType {
		$$ = $1;
	}
	;

ExtensionAndException:
	Tok_Ellipsis {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = "...";
		$$.Type = asn1types.Asn1ExprTypeExtensible;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
	}
	| Tok_Ellipsis '!' DefinedValue {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = "...";
		$$.Type = asn1types.Asn1ExprTypeExtensible;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
	}
	| Tok_Ellipsis '!' SignedNumber {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = "...";
		$$.Type = asn1types.Asn1ExprTypeExtensible;
		$$.Meta = asn1types.Asn1ExprMetaTypeType;
	}
	;

/*
 * A collection of constructed data type members.
 */
optComponentTypeLists:
	{ $$ = asn1types.NewAsn1Expression(); }
	| ComponentTypeLists { $$ = $1; };

ComponentTypeLists:
	ComponentType {
		$$ = asn1types.NewAsn1Expression()
	}
	| ComponentTypeLists ',' ComponentType {
		$$ = $1;
	}
	| ComponentTypeLists ',' Tok_TwoLeftBrackets ComponentTypeLists Tok_TwoRightBrackets {
		$$ = $1;
	}
	;

ComponentType:
	Identifier TaggedType optMarker {
		$$ = $2;
		$$.Identifier = $1;
	}
	| TaggedType optMarker {
		$$ = $1;
	}
	| Tok_COMPONENTS Tok_OF TaggedType {
		$$ = asn1types.NewAsn1Expression()
		//$$.Meta = $3.Meta
		$$.Type = asn1types.Asn1ExprTypeComponentsOf;
	}
	| ExtensionAndException {
		$$ = $1;
	}
	;

/*
 * MARKERS
 */

optMarker:
	{
		$$ = asn1types.NewAsn1Marker()
		$$.Flags = asn1types.Asn1MarkerFlagNoMark;
		$$.Value = nil
	}
	| Marker { $$ = $1; }
	;

Marker:
	Tok_OPTIONAL {
		fmt.Println("YYY")
		$$ = asn1types.NewAsn1Marker()
		$$.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect;
		$$.Value = nil;
	}
	| Tok_DEFAULT Value {
		$$ = asn1types.NewAsn1Marker()
		$$.Flags = asn1types.Asn1MarkerFlagDefault;
		$$.Value = $2;
	}
	;

BasicString:
	Tok_BMPString { $$ = asn1types.Asn1ExprTypeBMPString; }
	| Tok_GeneralString {
		$$ = asn1types.Asn1ExprTypeGeneralString;
		fmt.Println("WARNING: GeneralString is not fully supported");
	}
	| Tok_GraphicString {
		$$ = asn1types.Asn1ExprTypeGraphicsString;
		fmt.Println("WARNING: GraphicString is not fully supported");
	}
	| Tok_IA5String { $$ = asn1types.Asn1ExprTypeIA5String; }
	| Tok_ISO646String { $$ = asn1types.Asn1ExprTypeISO646String; }
	| Tok_NumericString { $$ = asn1types.Asn1ExprTypeNumericString; }
	| Tok_PrintableString { $$ = asn1types.Asn1ExprTypePrintableString; }
	| Tok_T61String {
		$$ = asn1types.Asn1ExprTypeT61String
		fmt.Println("WARNING: T61String is not fully supported");
	}
	| Tok_TeletexString { $$ = asn1types.Asn1ExprTypeTeletexString; }
	| Tok_UniversalString { $$ = asn1types.Asn1ExprTypeUniversalString; }
	| Tok_UTF8String { $$ = asn1types.Asn1ExprTypeUtf8String; }
	| Tok_VideotexString {
		$$ = asn1types.Asn1ExprTypeVideoTexString;
		fmt.Println("WARNING: VideotexString is not fully supported");
	}
	| Tok_VisibleString { $$ = asn1types.Asn1ExprTypeVisibleString; }
	| Tok_ObjectDescriptor { $$ = asn1types.Asn1ExprTypeObjectDescriptor; }
	;


/* empty | Constraint... */
optManyConstraints:
	{ $$ = nil; }
	| ManyConstraints;

ManyConstraints:
    Constraint
	| ManyConstraints Constraint {
	// FIXME : Implementation incomplete
	}
	;

Constraint:
    '(' ConstraintSpec ')' {
		$$ = $2
		// FIXME: Needs implementation
    }
    ;

ConstraintSpec: SubtypeConstraint ; /* | GeneralConstraint; */

SubtypeConstraint: ElementSetSpecs;

ValueRange:
    LowerEndValue ConstraintRangeSpec UpperEndValue {
		$$ = asn1types.NewAsn1Constraint()
		$$.Type = $2;
    };

LowerEndValue:
    SingleValue
    | Tok_MIN {
		$$ = asn1types.NewAsn1Value()
		$$.Type = asn1types.Asn1ValueTypeMin;
    };

UpperEndValue:
    SingleValue
    | Tok_MAX {
		$$ = asn1types.NewAsn1Value()
		$$.Type = asn1types.Asn1ValueTypeMax;
    };

ConstraintRangeSpec:
	Tok_TwoDots		{ $$ = asn1types.Asn1ConstraintTypeSimpleRange; }
	| Tok_TwoDots '<'	{ $$ = asn1types.Asn1ConstraintTypeRightExcludedRange; }
	| '<' Tok_TwoDots	{ $$ = asn1types.Asn1ConstraintTypeLeftExcludedRange; }
	| '<' Tok_TwoDots '<'	{ $$ = asn1types.Asn1ConstraintTypeLeftRightExcludedRange; }
	;

ContainedSubtype:
    Tok_INCLUDES Type {
		$$ = asn1types.NewAsn1Value()
    }
    /* Can't put Type here because of conflicts. Simplified subset */
    | DefinedUntaggedType {
		$$ = asn1types.NewAsn1Value()
    }
	;

DefinedUntaggedType:
	DefinedType optManyConstraints {
		$$ = $1;
		// FIXME : Implementation
	}
	;

SizeConstraint:
	Tok_SIZE Constraint {
		$$ = asn1types.NewAsn1Constraint()
		/* FIXME : implementation not yet complete */
	};

PermittedAlphabet:
	Tok_FROM Constraint {
		$$ = asn1types.NewAsn1Constraint()
	};

optIdentifier:
	{ $$ = ""; }
	| Identifier {
		$$ = $1;
	}
	;
/* empty | Constraint | SIZE(...) */
optSizeOrConstraint:
	{ $$ = nil; }
	| Constraint
	| SizeConstraint
	;

/* XXXXXXXXXX Marker */

/*
    | SizeConstraint    /* SIZE ... */
    /* | TypeConstraint is via ContainedSubtype * /
	| InnerTypeConstraints  /* WITH COMPONENT[S] ... * /
	| PatternConstraint     /* PATTERN ... * /
	| ValueRange
*/
	;


/* TODO: Constraints not implemented yet
PatternConstraint:
	TOK_PATTERN TOK_cstring {
		$$ = asn1p_constraint_new(yylineno, currentModule);
		$$->type = ACT_CT_PATTERN;
		$$->value = asn1p_value_frombuf($2.buf, $2.len, 0);
	}
	| TOK_PATTERN Identifier {
		asn1p_ref_t *ref;
		$$ = asn1p_constraint_new(yylineno, currentModule);
		$$->type = ACT_CT_PATTERN;
		ref = asn1p_ref_new(yylineno, currentModule);
		asn1p_ref_add_component(ref, $2, RLT_lowercase);
		$$->value = asn1p_value_fromref(ref, 0);
		free($2);
	}
	;

*/

SingleValue: Value;

/* TODO: Add support one by one
BitStringValue:
	TOK_bstring {
		$$ = _convert_bitstring2binary($1, 'B');
		checkmem($$);
		free($1);
	}
	| TOK_hstring {
		$$ = _convert_bitstring2binary($1, 'H');
		checkmem($$);
		free($1);
	}
	;

*/

/*
 * X.680 08/2015
 * #51.8.5
 * /
InnerTypeConstraints:
	TOK_WITH TOK_COMPONENT SingleTypeConstraint {
		CONSTRAINT_INSERT($$, ACT_CT_WCOMP, $3, 0);
	}
	| TOK_WITH TOK_COMPONENTS MultipleTypeConstraints {
        assert($3->type == ACT_CA_CSV);
        $3->type = ACT_CT_WCOMPS;
        $$ = $3;
	}
	;
SingleTypeConstraint: Constraint;
MultipleTypeConstraints: FullSpecification | PartialSpecification;
FullSpecification: '{' TypeConstraints '}' { $$ = $2; };
PartialSpecification:
    '{' Tok_Ellipsis ',' TypeConstraints '}' {
        assert($4->type == ACT_CA_CSV);
		$$ = asn1p_constraint_new(yylineno, currentModule);
        $$->type = ACT_CA_CSV;
		asn1p_constraint_t *ct = asn1p_constraint_new(yylineno, currentModule);
		checkmem($$);
		ct->type = asn1types.Asn1ConstraintTypeExtensibilityMark;
        asn1p_constraint_insert($$, ct);
        for(unsigned i = 0; i < $4->el_count; i++) {
            asn1p_constraint_insert($$, $4->elements[i]);
        }
    };
TypeConstraints:
    NamedConstraint {
        $$ = asn1p_constraint_new(yylineno, currentModule);
        $$->type = ACT_CA_CSV;
        asn1p_constraint_insert($$, $1);
    }
    | TypeConstraints ',' NamedConstraint {
        $$ = $1;
        asn1p_constraint_insert($$, $3);
	}
	;
NamedConstraint:
	IdentifierAsValue optConstraint optPresenceConstraint {
        $$ = asn1p_constraint_new(yylineno, currentModule);
        checkmem($$);
        $$->type = ACT_EL_VALUE;
        $$->value = $1;
        if($2) asn1p_constraint_insert($$, $2);
        $$->presence = $3;
    }
    ;

*/

/*
 * presence constraint for NamedConstraint
 * /
optPresenceConstraint:
	{ $$ = ACPRES_DEFAULT; }
	| PresenceConstraint { $$ = $1; }
	;

PresenceConstraint:
	TOK_PRESENT {
		$$ = ACPRES_PRESENT;
	}
	| TOK_ABSENT {
		$$ = ACPRES_ABSENT;
	}
	| TOK_OPTIONAL {
		$$ = ACPRES_OPTIONAL;
	}
	;

*/
/* X.682 * /
GeneralConstraint:
	UserDefinedConstraint
	| TableConstraint
	| ContentsConstraint
	;

UserDefinedConstraint:
	TOK_CONSTRAINED TOK_BY '{'
		{ asn1p_lexer_hack_push_opaque_state(); } Opaque /* '}' * / {
		$$ = asn1p_constraint_new(yylineno, currentModule);
		checkmem($$);
		$$->type = ACT_CT_CTDBY;
		$$->value = asn1p_value_frombuf($5.buf, $5.len, 0);
		checkmem($$->value);
		$$->value->type = ATV_UNPARSED;
	}
	;

ContentsConstraint:
	TOK_CONTAINING Type {
		$$ = asn1p_constraint_new(yylineno, currentModule);
		$$->type = ACT_CT_CTNG;
		$$->value = asn1p_value_fromtype($2);
		asn1p_expr_free($2);
	}
	;

TableConstraint:
	SimpleTableConstraint {
		$$ = $1;
	}
	| ComponentRelationConstraint {
		$$ = $1;
	}
	;

*/
/* TODO :
 * "{ExtensionSet}"
 * /
SimpleTableConstraint:
	'{' TypeRefName '}' {
		asn1p_ref_t *ref = asn1p_ref_new(yylineno, currentModule);
		asn1p_constraint_t *ct;
		int ret;
		ret = asn1p_ref_add_component(ref, $2, 0);
		checkmem(ret == 0);
		ct = asn1p_constraint_new(yylineno, currentModule);
		checkmem($$);
		ct->type = ACT_EL_VALUE;
		ct->value = asn1p_value_fromref(ref, 0);
		CONSTRAINT_INSERT($$, ACT_CA_CRC, ct, 0);
		free($2);
	}
	;

ComponentRelationConstraint:
	SimpleTableConstraint '{' AtNotationList '}' {
		CONSTRAINT_INSERT($$, ACT_CA_CRC, $1, $3);
	}
	;

AtNotationList:
	AtNotationElement {
		$$ = asn1p_constraint_new(yylineno, currentModule);
		checkmem($$);
		$$->type = ACT_EL_VALUE;
		$$->value = asn1p_value_fromref($1, 0);
	}
	| AtNotationList ',' AtNotationElement {
		asn1p_constraint_t *ct;
		ct = asn1p_constraint_new(yylineno, currentModule);
		checkmem(ct);
		ct->type = ACT_EL_VALUE;
		ct->value = asn1p_value_fromref($3, 0);
		CONSTRAINT_INSERT($$, ACT_CA_CSV, $1, ct);
	}
	;

/*
 * @blah
 * /
AtNotationElement:
	'@' ComponentIdList {
		char *p = malloc(strlen($2) + 2);
		int ret;
		*p = '@';
		strcpy(p + 1, $2);
		$$ = asn1p_ref_new(yylineno, currentModule);
		ret = asn1p_ref_add_component($$, p, 0);
		checkmem(ret == 0);
		free(p);
		free($2);
	}
	| '@' '.' ComponentIdList {
		char *p = malloc(strlen($3) + 3);
		int ret;
		p[0] = '@';
		p[1] = '.';
		strcpy(p + 2, $3);
		$$ = asn1p_ref_new(yylineno, currentModule);
		ret = asn1p_ref_add_component($$, p, 0);
		checkmem(ret == 0);
		free(p);
		free($3);
	}
	;

/* identifier "." ... * /
ComponentIdList:
	Identifier {
		$$ = $1;
	}
	| ComponentIdList '.' Identifier {
		int l1 = strlen($1);
		int l3 = strlen($3);
		$$ = malloc(l1 + 1 + l3 + 1);
		memcpy($$, $1, l1);
		$$[l1] = '.';
		memcpy($$ + l1 + 1, $3, l3);
		$$[l1 + 1 + l3] = '\0';
		free($1);
		free($3);
	}
	;

*/
%%
