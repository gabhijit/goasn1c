
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

	str string
	a_module *Asn1Module
}

%token Tok_BEGIN
%token Tok_END
%token Tok_DEFINITIONS
%token Tok_ASSIGNMENT
%token <str> Tok_TypeReference


%type <str> TypeRefName
%type <a_module> ModuleDefinition

%%

ModuleDefinition:
		TypeRefName { currentModule = NewAsn1Module(); fmt.Println(currentModule);}
		Tok_DEFINITIONS Tok_ASSIGNMENT
		Tok_BEGIN Tok_END {

			$$ = currentModule
			$$.name = $1
			fmt.Println($$)


		};

TypeRefName:
	   Tok_TypeReference {
		$$ = $1
	};

%%
