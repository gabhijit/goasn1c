//line asn1p.y:3
package parser

import __yyfmt__ "fmt"

//line asn1p.y:4
import (
	"fmt" /*	"bufio"
		"bytes"
		"io"
		"log"
		"os"
		"unicode/utf8"
	*/)

var currentModule *Asn1Module

//line asn1p.y:21
type Asn1SymType struct {
	yys int

	str      string
	a_module *Asn1Module
}

const Tok_BEGIN = 57346
const Tok_END = 57347
const Tok_DEFINITIONS = 57348
const Tok_ASSIGNMENT = 57349
const Tok_TypeReference = 57350

var Asn1Toknames = [...]string{
	"$end",
	"error",
	"$unk",
	"Tok_BEGIN",
	"Tok_END",
	"Tok_DEFINITIONS",
	"Tok_ASSIGNMENT",
	"Tok_TypeReference",
}
var Asn1Statenames = [...]string{}

const Asn1EofCode = 1
const Asn1ErrCode = 2
const Asn1InitialStackSize = 16

//line asn1p.y:56

//line yacctab:1
var Asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const Asn1Private = 57344

const Asn1Last = 8

var Asn1Act = [...]int{

	3, 6, 5, 8, 7, 4, 1, 2,
}
var Asn1Pact = [...]int{

	-8, -1000, -1000, -1000, -4, -6, 0, -2, -1000,
}
var Asn1Pgo = [...]int{

	0, 7, 6, 5,
}
var Asn1R1 = [...]int{

	0, 3, 2, 1,
}
var Asn1R2 = [...]int{

	0, 0, 6, 1,
}
var Asn1Chk = [...]int{

	-1000, -2, -1, 8, -3, 6, 7, 4, 5,
}
var Asn1Def = [...]int{

	0, -2, 1, 3, 0, 0, 0, 0, 2,
}
var Asn1Tok1 = [...]int{

	1,
}
var Asn1Tok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8,
}
var Asn1Tok3 = [...]int{
	0,
}

var Asn1ErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	Asn1Debug        = 0
	Asn1ErrorVerbose = false
)

type Asn1Lexer interface {
	Lex(lval *Asn1SymType) int
	Error(s string)
}

type Asn1Parser interface {
	Parse(Asn1Lexer) int
	Lookahead() int
}

type Asn1ParserImpl struct {
	lval  Asn1SymType
	stack [Asn1InitialStackSize]Asn1SymType
	char  int
}

func (p *Asn1ParserImpl) Lookahead() int {
	return p.char
}

func Asn1NewParser() Asn1Parser {
	return &Asn1ParserImpl{}
}

const Asn1Flag = -1000

func Asn1Tokname(c int) string {
	if c >= 1 && c-1 < len(Asn1Toknames) {
		if Asn1Toknames[c-1] != "" {
			return Asn1Toknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func Asn1Statname(s int) string {
	if s >= 0 && s < len(Asn1Statenames) {
		if Asn1Statenames[s] != "" {
			return Asn1Statenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Asn1ErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !Asn1ErrorVerbose {
		return "syntax error"
	}

	for _, e := range Asn1ErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + Asn1Tokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := Asn1Pact[state]
	for tok := TOKSTART; tok-1 < len(Asn1Toknames); tok++ {
		if n := base + tok; n >= 0 && n < Asn1Last && Asn1Chk[Asn1Act[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if Asn1Def[state] == -2 {
		i := 0
		for Asn1Exca[i] != -1 || Asn1Exca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; Asn1Exca[i] >= 0; i += 2 {
			tok := Asn1Exca[i]
			if tok < TOKSTART || Asn1Exca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if Asn1Exca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += Asn1Tokname(tok)
	}
	return res
}

func Asn1lex1(lex Asn1Lexer, lval *Asn1SymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = Asn1Tok1[0]
		goto out
	}
	if char < len(Asn1Tok1) {
		token = Asn1Tok1[char]
		goto out
	}
	if char >= Asn1Private {
		if char < Asn1Private+len(Asn1Tok2) {
			token = Asn1Tok2[char-Asn1Private]
			goto out
		}
	}
	for i := 0; i < len(Asn1Tok3); i += 2 {
		token = Asn1Tok3[i+0]
		if token == char {
			token = Asn1Tok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = Asn1Tok2[1] /* unknown char */
	}
	if Asn1Debug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", Asn1Tokname(token), uint(char))
	}
	return char, token
}

func Asn1Parse(Asn1lex Asn1Lexer) int {
	return Asn1NewParser().Parse(Asn1lex)
}

func (Asn1rcvr *Asn1ParserImpl) Parse(Asn1lex Asn1Lexer) int {
	var Asn1n int
	var Asn1VAL Asn1SymType
	var Asn1Dollar []Asn1SymType
	_ = Asn1Dollar // silence set and not used
	Asn1S := Asn1rcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Asn1state := 0
	Asn1rcvr.char = -1
	Asn1token := -1 // Asn1rcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Asn1state = -1
		Asn1rcvr.char = -1
		Asn1token = -1
	}()
	Asn1p := -1
	goto Asn1stack

ret0:
	return 0

ret1:
	return 1

Asn1stack:
	/* put a state and value onto the stack */
	if Asn1Debug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", Asn1Tokname(Asn1token), Asn1Statname(Asn1state))
	}

	Asn1p++
	if Asn1p >= len(Asn1S) {
		nyys := make([]Asn1SymType, len(Asn1S)*2)
		copy(nyys, Asn1S)
		Asn1S = nyys
	}
	Asn1S[Asn1p] = Asn1VAL
	Asn1S[Asn1p].yys = Asn1state

Asn1newstate:
	Asn1n = Asn1Pact[Asn1state]
	if Asn1n <= Asn1Flag {
		goto Asn1default /* simple state */
	}
	if Asn1rcvr.char < 0 {
		Asn1rcvr.char, Asn1token = Asn1lex1(Asn1lex, &Asn1rcvr.lval)
	}
	Asn1n += Asn1token
	if Asn1n < 0 || Asn1n >= Asn1Last {
		goto Asn1default
	}
	Asn1n = Asn1Act[Asn1n]
	if Asn1Chk[Asn1n] == Asn1token { /* valid shift */
		Asn1rcvr.char = -1
		Asn1token = -1
		Asn1VAL = Asn1rcvr.lval
		Asn1state = Asn1n
		if Errflag > 0 {
			Errflag--
		}
		goto Asn1stack
	}

Asn1default:
	/* default state action */
	Asn1n = Asn1Def[Asn1state]
	if Asn1n == -2 {
		if Asn1rcvr.char < 0 {
			Asn1rcvr.char, Asn1token = Asn1lex1(Asn1lex, &Asn1rcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if Asn1Exca[xi+0] == -1 && Asn1Exca[xi+1] == Asn1state {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Asn1n = Asn1Exca[xi+0]
			if Asn1n < 0 || Asn1n == Asn1token {
				break
			}
		}
		Asn1n = Asn1Exca[xi+1]
		if Asn1n < 0 {
			goto ret0
		}
	}
	if Asn1n == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Asn1lex.Error(Asn1ErrorMessage(Asn1state, Asn1token))
			Nerrs++
			if Asn1Debug >= 1 {
				__yyfmt__.Printf("%s", Asn1Statname(Asn1state))
				__yyfmt__.Printf(" saw %s\n", Asn1Tokname(Asn1token))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Asn1p >= 0 {
				Asn1n = Asn1Pact[Asn1S[Asn1p].yys] + Asn1ErrCode
				if Asn1n >= 0 && Asn1n < Asn1Last {
					Asn1state = Asn1Act[Asn1n] /* simulate a shift of "error" */
					if Asn1Chk[Asn1state] == Asn1ErrCode {
						goto Asn1stack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if Asn1Debug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", Asn1S[Asn1p].yys)
				}
				Asn1p--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if Asn1Debug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", Asn1Tokname(Asn1token))
			}
			if Asn1token == Asn1EofCode {
				goto ret1
			}
			Asn1rcvr.char = -1
			Asn1token = -1
			goto Asn1newstate /* try again in the same state */
		}
	}

	/* reduction by production Asn1n */
	if Asn1Debug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Asn1n, Asn1Statname(Asn1state))
	}

	Asn1nt := Asn1n
	Asn1pt := Asn1p
	_ = Asn1pt // guard against "declared and not used"

	Asn1p -= Asn1R2[Asn1n]
	// Asn1p is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Asn1p+1 >= len(Asn1S) {
		nyys := make([]Asn1SymType, len(Asn1S)*2)
		copy(nyys, Asn1S)
		Asn1S = nyys
	}
	Asn1VAL = Asn1S[Asn1p+1]

	/* consult goto table to find next state */
	Asn1n = Asn1R1[Asn1n]
	Asn1g := Asn1Pgo[Asn1n]
	Asn1j := Asn1g + Asn1S[Asn1p].yys + 1

	if Asn1j >= Asn1Last {
		Asn1state = Asn1Act[Asn1g]
	} else {
		Asn1state = Asn1Act[Asn1j]
		if Asn1Chk[Asn1state] != -Asn1n {
			Asn1state = Asn1Act[Asn1g]
		}
	}
	// dummy call; replaced with literal code
	switch Asn1nt {

	case 1:
		Asn1Dollar = Asn1S[Asn1pt-1 : Asn1pt+1]
		//line asn1p.y:40
		{
			currentModule = NewAsn1Module()
			fmt.Println(currentModule)
		}
	case 2:
		Asn1Dollar = Asn1S[Asn1pt-6 : Asn1pt+1]
		//line asn1p.y:42
		{

			Asn1VAL.a_module = currentModule
			Asn1VAL.a_module.name = Asn1Dollar[1].str
			fmt.Println(Asn1VAL.a_module)

		}
	case 3:
		Asn1Dollar = Asn1S[Asn1pt-1 : Asn1pt+1]
		//line asn1p.y:52
		{
			Asn1VAL.str = Asn1Dollar[1].str
		}
	}
	goto Asn1stack /* stack new state and value */
}
