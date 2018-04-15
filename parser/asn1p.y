
%{

package parser

import (
	"fmt"
/*	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"unicode/utf8"
*/
)

var currentModule *Asn1Module

%}

%union {
	ch rune
	str string
	a_module *Asn1Module
	a_oid	*Asn1Oid
	a_oid_arc Asn1OidArc
	num int64
}

%token Tok_BEGIN
%token Tok_END
%token Tok_DEFINITIONS
%token Tok_ASSIGNMENT
%token <str> Tok_TypeReference
%token <num> Tok_Number
%token <str> Tok_Identifier


%type <str> TypeRefName
%type <str> Identifier
%type <a_module> ModuleDefinition
%type <a_oid> optObjectIdentifier
%type <a_oid> ObjectIdentifier
%type <a_oid> ObjectIdentifierBody
%type <a_oid_arc> ObjectIdentifierElement

%%

ModuleDefinition:
		TypeRefName { currentModule = NewAsn1Module();}
		optObjectIdentifier Tok_DEFINITIONS Tok_ASSIGNMENT
		Tok_BEGIN Tok_END {

			$$ = currentModule
			$$.name = $1
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
		$$ = NewAsn1Oid();
		$$.arcs = append($$.arcs, $1)
	}
	| ObjectIdentifierBody ObjectIdentifierElement {
		$$ = $1;
		$$.arcs = append($$.arcs, $2)
	};

ObjectIdentifierElement:
	Identifier {					/* iso */
		$$.name = $1;
		$$.num = -1;
	}
	| Identifier '(' Tok_Number ')' {		/* iso(1) */
		$$.name = $1;
		$$.num = $3;
	}
	| Tok_Number {					/* 1 */
		$$.name = "";
		$$.num = $1;
	};

Identifier: Tok_Identifier {
	  $$ = $1
};

TypeRefName:
	   Tok_TypeReference {
		$$ = $1
	};

%%
