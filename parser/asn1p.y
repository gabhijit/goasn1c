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
%type <xports>       ImportsElement
%type <aid>          AssignedIdentifier

%type <module>       optExports
%type <xports>       ExportsDefinition
%type <xports>       ExportsBody
%type <expr>         ExportsElement

%type <module>       AssignmentList
%type <str>          Assignment

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
			$$.Name = $1
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
	     { $$ = nil; }
	| ModuleBody {
		$$ = $1;
	};

ModuleBody:
	optExports optImports AssignmentList {
		$$ = asn1types.NewAsn1Module()

		$$.Exports = $1
		$$.Imports = $2
		// FIXME: Use AssignmentList
	};

AssignmentList:
	Assignment {
		$$ = $1;
	}
	| AssignmentList Assignment {
		if($1) {
			$$ = $1;
		} else {
			$$ = $2;
			break;
		}
	}
	;

/* TODO: Big implementation */
Assignment:
	  Tok_Identifier { $$ = nil;};
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
		return yyerror("Empty IMPORTS list");
	}
	;

optImportsBundleSet:
	{ $$ = asn1types.NewAsn1Module(); }
	| ImportsBundleSet;

ImportsBundleSet:
	ImportsBundle {
		$$ = asn1types.NewAsn1Module();
		$$.Imports = append($$.Imports, $1)
	}
	| ImportsBundleSet ImportsBundle {
		$$ = $1;
		$$.Imports = append($$.Imports, $1)
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
		$$.Members = append($$.Members, $1)
	}
	| ImportsList ',' ImportsElement {
		$$ = $1;
		TQ_ADD(&($$->members), $3, next);
	}
	;

ImportsElement:
	TypeRefName {
		$$ = asn1types.NewAsn1Expression()
		checkmem($$);
		$$->Identifier = $1;
		$$->expr_type = A1TC_REFERENCE;
	}
	| TypeRefName '{' '}' {		/* Completely equivalent to above */
		$$ = NEW_EXPR();
		checkmem($$);
		$$->Identifier = $1;
		$$->expr_type = A1TC_REFERENCE;
	}
	| Identifier {
		$$ = NEW_EXPR();
		checkmem($$);
		$$->Identifier = $1;
		$$->expr_type = A1TC_REFERENCE;
	}
	;

optExports:
	{ $$ = nil; }
	| ExportsDefinition {
		$$ = asn1p_module_new();
		checkmem($$);
		if($1) {
			TQ_ADD(&($$->exports), $1, xp_next);
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
		$$ = 0;
	}
	| Tok_EXPORTS ';' {
		/* Empty EXPORTS clause effectively prohibits export. */
		$$ = asn1p_xports_new();
		checkmem($$);
	}
	;

ExportsBody:
	ExportsElement {
		$$ = asn1p_xports_new();
		assert($$);
		TQ_ADD(&($$->members), $1, next);
	}
	| ExportsBody ',' ExportsElement {
		$$ = $1;
		TQ_ADD(&($$->members), $3, next);
	}
	;

ExportsElement:
	TypeRefName {
		$$ = NEW_EXPR();
		checkmem($$);
		$$->Identifier = $1;
		$$->expr_type = A1TC_EXPORTVAR;
	}
	| TypeRefName '{' '}' {
		$$ = NEW_EXPR();
		checkmem($$);
		$$->Identifier = $1;
		$$->expr_type = A1TC_EXPORTVAR;
	}
	| Identifier {
		$$ = NEW_EXPR();
		checkmem($$);
		$$->Identifier = $1;
		$$->expr_type = A1TC_EXPORTVAR;
	}
	;

%%
