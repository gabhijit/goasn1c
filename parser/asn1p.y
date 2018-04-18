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

	grammar      *asn1types.Asn1Grammar
	module       *asn1types.Asn1Module
	oid	     *asn1types.Asn1Oid
	oid_arc      asn1types.Asn1OidArc
	module_flags asn1types.ModuleFlagType
	xports       *asn1types.Asn1Xports
	expr         *asn1types.Asn1Expression
	aid          *asn1types.Asn1AssignedIdentifier
	expr_type    asn1types.Asn1ExprType
	tag          *asn1types.Asn1Tag
}

%token       Tok_BEGIN
%token       Tok_END
%token       Tok_DEFINITIONS
%token       Tok_ASSIGNMENT
%token       Tok_IMPLICIT
%token       Tok_IMPLIED
%token       Tok_EXPLICIT
%token       Tok_EXTENSIBILITY
%token       Tok_TAGS
%token       Tok_AUTOMATIC
%token       Tok_IMPORTS
%token       Tok_FROM
%token       Tok_ALL
%token       Tok_EXPORTS
%token       Tok_INTEGER
%token       Tok_UNIVERSAL
%token       Tok_APPLICATION
%token       Tok_PRIVATE
%token       Tok_BOOLEAN
%token <str> Tok_TypeReference
%token <num> Tok_Number
%token <str> Tok_Identifier


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
%type <expr>         ConcreteTypeDeclaration


%type <tag>          optTag
%type <tag>          Tag
%type <tag>          TagTypeValue
%type <tag>          TagPlicit
%type <tag>          TagClass

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
	{ $$.Oid = nil; $$.Value = nil;}
	| ObjectIdentifier { $$.Oid = $1;  $$.Value = nil};

ImportsBundle:
	ImportsList Tok_FROM TypeRefName AssignedIdentifier {
		$$ = $1;

		$$.Type = asn1types.Asn1XportsTypeImport;
		$$.FromModule = $3;
		$$.Identifier = $4;
	}
	;
// From here
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
	};

DataTypeReference:
	TypeRefName Tok_ASSIGNMENT Type {
		$$ = asn1types.NewAsn1Expression()
		$$.Identifier = $1
		// FIXME : Need to add code for type of expression
	};

Type:TaggedType;


TaggedType:
	optTag UntaggedType {
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
	    TypeDeclaration {
	};

TypeDeclaration:
	ConcreteTypeDeclaration;
	/*| DefinedType ; */

ConcreteTypeDeclaration: BuiltinType;

BuiltinType:
	Tok_BOOLEAN {
		$$ = asn1types.NewAsn1Expression();
		$$.Type = asn1types.Asn1ExprTypeBoolean;

	}
	| Tok_INTEGER {
		$$ = asn1types.NewAsn1Expression();
		$$.Type = asn1types.Asn1ExprTypeInteger;

	};
%%
