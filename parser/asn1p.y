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

%}

%union {
	ch           rune
	str          string
	module       *asn1types.Asn1Module
	oid	     *asn1types.Asn1Oid
	oid_arc      asn1types.Asn1OidArc
	module_flags asn1types.ModuleFlagType
	num          int64
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
%token <str> Tok_TypeReference
%token <num> Tok_Number
%token <str> Tok_Identifier


%type <str>          TypeRefName
%type <str>          Identifier
%type <module>       ModuleDefinition
%type <oid>          optObjectIdentifier
%type <oid>          ObjectIdentifier
%type <oid>          ObjectIdentifierBody
%type <oid_arc>      ObjectIdentifierElement
%type <module_flags> optModuleDefinitionFlags
%type <module_flags> ModuleDefinitionFlags
%type <module_flags> ModuleDefinitionFlag


%%

ModuleDefinition:
		TypeRefName { currentModule = asn1types.NewAsn1Module();}
		optObjectIdentifier Tok_DEFINITIONS
		optModuleDefinitionFlags
		Tok_ASSIGNMENT
		Tok_BEGIN Tok_END {

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

%%
