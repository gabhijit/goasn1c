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
	ch        rune
	str       string
	a_module  *asn1types.Asn1Module
	a_oid	  *asn1types.Asn1Oid
	a_oid_arc asn1types.Asn1OidArc
	num       int64
}

%token       Tok_BEGIN
%token       Tok_END
%token       Tok_DEFINITIONS
%token       Tok_ASSIGNMENT
%token <str> Tok_TypeReference
%token <num> Tok_Number
%token <str> Tok_Identifier


%type <str>       TypeRefName
%type <str>       Identifier
%type <a_module>  ModuleDefinition
%type <a_oid>     optObjectIdentifier
%type <a_oid>     ObjectIdentifier
%type <a_oid>     ObjectIdentifierBody
%type <a_oid_arc> ObjectIdentifierElement

%%

ModuleDefinition:
		TypeRefName { currentModule = asn1types.NewAsn1Module();}
		optObjectIdentifier Tok_DEFINITIONS Tok_ASSIGNMENT
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

%%
