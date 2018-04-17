//line parser/asn1p.y:24
package parser

import __yyfmt__ "fmt"

//line parser/asn1p.y:25
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
var AllModules *asn1types.Asn1Grammar

//line parser/asn1p.y:44
type asn1SymType struct {
	yys          int
	ch           rune
	str          string
	num          int64
	grammar      *asn1types.Asn1Grammar
	module       *asn1types.Asn1Module
	oid          *asn1types.Asn1Oid
	oid_arc      asn1types.Asn1OidArc
	module_flags asn1types.ModuleFlagType
	xports       *asn1types.Asn1Xports
	expr         *asn1types.Asn1Expression
	aid          *asn1types.Asn1AssignedIdentifier
	expr_type    asn1types.Asn1ExprType
}

const Tok_BEGIN = 57346
const Tok_END = 57347
const Tok_DEFINITIONS = 57348
const Tok_ASSIGNMENT = 57349
const Tok_IMPLICIT = 57350
const Tok_IMPLIED = 57351
const Tok_EXPLICIT = 57352
const Tok_EXTENSIBILITY = 57353
const Tok_TAGS = 57354
const Tok_AUTOMATIC = 57355
const Tok_IMPORTS = 57356
const Tok_FROM = 57357
const Tok_ALL = 57358
const Tok_EXPORTS = 57359
const Tok_INTEGER = 57360
const Tok_TypeReference = 57361
const Tok_Number = 57362
const Tok_Identifier = 57363

var asn1Toknames = [...]string{
	"$end",
	"error",
	"$unk",
	"Tok_BEGIN",
	"Tok_END",
	"Tok_DEFINITIONS",
	"Tok_ASSIGNMENT",
	"Tok_IMPLICIT",
	"Tok_IMPLIED",
	"Tok_EXPLICIT",
	"Tok_EXTENSIBILITY",
	"Tok_TAGS",
	"Tok_AUTOMATIC",
	"Tok_IMPORTS",
	"Tok_FROM",
	"Tok_ALL",
	"Tok_EXPORTS",
	"Tok_INTEGER",
	"Tok_TypeReference",
	"Tok_Number",
	"Tok_Identifier",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"';'",
	"','",
}
var asn1Statenames = [...]string{}

const asn1EofCode = 1
const asn1ErrCode = 2
const asn1InitialStackSize = 16

//line parser/asn1p.y:406

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 35,
	5, 25,
	-2, 49,
}

const asn1Private = 57344

const asn1Last = 96

var asn1Act = [...]int{

	9, 62, 4, 61, 4, 63, 49, 59, 53, 47,
	72, 70, 5, 66, 17, 36, 15, 34, 15, 48,
	64, 65, 73, 16, 17, 27, 25, 16, 17, 81,
	13, 76, 57, 10, 74, 5, 5, 17, 17, 20,
	67, 5, 78, 50, 41, 55, 22, 51, 21, 24,
	45, 23, 14, 32, 55, 31, 30, 33, 69, 29,
	28, 68, 11, 42, 35, 26, 71, 50, 7, 1,
	3, 51, 75, 6, 79, 19, 18, 80, 12, 8,
	83, 77, 54, 52, 46, 40, 39, 82, 60, 58,
	56, 44, 43, 38, 37, 2,
}
var asn1Pact = [...]int{

	22, -1000, 22, -1000, -1000, -1000, -1000, 11, 56, -1000,
	7, 38, 3, -1000, -1000, 1, -1000, -1000, 53, 38,
	-1000, 44, 43, 41, 48, -1000, -1000, -3, 60, -1000,
	-1000, -1000, -1000, -1000, -10, 27, -1000, 58, -1000, 36,
	-1000, -7, -1000, 22, -1000, 17, -6, -13, -1000, -1000,
	18, -1000, 22, -1000, -1000, 51, -15, -1000, 16, -1000,
	-5, -1000, 12, -1000, -1000, 16, -1000, 8, -1000, 24,
	-1000, -1000, 22, 16, 6, -1000, -1000, -1000, -1000, 11,
	-1000, -1000, -1000, -1000,
}
var asn1Pgo = [...]int{

	0, 95, 1, 5, 70, 94, 93, 92, 91, 90,
	89, 7, 88, 3, 87, 86, 85, 84, 6, 83,
	8, 82, 81, 79, 0, 78, 52, 76, 75, 39,
	69, 68,
}
var asn1R1 = [...]int{

	0, 30, 1, 1, 31, 4, 23, 23, 24, 24,
	25, 25, 26, 26, 26, 3, 2, 27, 27, 28,
	28, 29, 29, 29, 29, 5, 5, 6, 19, 19,
	20, 21, 22, 7, 7, 8, 8, 9, 9, 10,
	10, 14, 14, 11, 12, 12, 13, 13, 13, 15,
	15, 16, 16, 16, 17, 17, 18, 18, 18,
}
var asn1R2 = [...]int{

	0, 1, 1, 2, 0, 9, 0, 1, 3, 2,
	1, 2, 1, 4, 1, 1, 1, 0, 1, 1,
	2, 2, 2, 2, 2, 0, 1, 3, 1, 2,
	1, 3, 1, 0, 1, 3, 2, 0, 1, 1,
	2, 0, 1, 4, 1, 3, 1, 3, 1, 0,
	1, 3, 3, 2, 1, 3, 1, 3, 1,
}
var asn1Chk = [...]int{

	-1000, -30, -1, -4, -2, 19, -4, -31, -23, -24,
	22, 6, -25, 23, -26, -3, 20, 21, -27, -28,
	-29, 10, 8, 13, 11, 23, -26, 24, 7, -29,
	12, 12, 12, 9, 20, 4, 25, -5, -6, -15,
	-16, 17, 5, -7, -8, 14, -17, 16, 26, -18,
	-2, -3, -19, -20, -21, -2, -9, 15, -10, -11,
	-12, -13, -2, -3, 26, 27, 26, 22, -20, 7,
	26, -11, 15, 27, 22, -18, 23, -22, 18, -2,
	-13, 23, -14, -24,
}
var asn1Def = [...]int{

	0, -2, 1, 2, 4, 16, 3, 6, 0, 7,
	0, 17, 0, 9, 10, 12, 14, 15, 0, 18,
	19, 0, 0, 0, 0, 8, 11, 0, 0, 20,
	21, 22, 23, 24, 0, -2, 13, 0, 26, 33,
	50, 0, 5, 0, 34, 37, 0, 0, 53, 54,
	56, 58, 27, 28, 30, 0, 0, 36, 38, 39,
	0, 44, 46, 48, 51, 0, 52, 0, 29, 0,
	35, 40, 0, 0, 0, 55, 57, 31, 32, 41,
	45, 47, 43, 42,
}
var asn1Tok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 25, 3, 3, 27, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 26,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 3, 23,
}
var asn1Tok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
}
var asn1Tok3 = [...]int{
	0,
}

var asn1ErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	asn1Debug        = 0
	asn1ErrorVerbose = false
)

type asn1Lexer interface {
	Lex(lval *asn1SymType) int
	Error(s string)
}

type asn1Parser interface {
	Parse(asn1Lexer) int
	Lookahead() int
}

type asn1ParserImpl struct {
	lval  asn1SymType
	stack [asn1InitialStackSize]asn1SymType
	char  int
}

func (p *asn1ParserImpl) Lookahead() int {
	return p.char
}

func asn1NewParser() asn1Parser {
	return &asn1ParserImpl{}
}

const asn1Flag = -1000

func asn1Tokname(c int) string {
	if c >= 1 && c-1 < len(asn1Toknames) {
		if asn1Toknames[c-1] != "" {
			return asn1Toknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func asn1Statname(s int) string {
	if s >= 0 && s < len(asn1Statenames) {
		if asn1Statenames[s] != "" {
			return asn1Statenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func asn1ErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !asn1ErrorVerbose {
		return "syntax error"
	}

	for _, e := range asn1ErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + asn1Tokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := asn1Pact[state]
	for tok := TOKSTART; tok-1 < len(asn1Toknames); tok++ {
		if n := base + tok; n >= 0 && n < asn1Last && asn1Chk[asn1Act[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if asn1Def[state] == -2 {
		i := 0
		for asn1Exca[i] != -1 || asn1Exca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; asn1Exca[i] >= 0; i += 2 {
			tok := asn1Exca[i]
			if tok < TOKSTART || asn1Exca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if asn1Exca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += asn1Tokname(tok)
	}
	return res
}

func asn1lex1(lex asn1Lexer, lval *asn1SymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = asn1Tok1[0]
		goto out
	}
	if char < len(asn1Tok1) {
		token = asn1Tok1[char]
		goto out
	}
	if char >= asn1Private {
		if char < asn1Private+len(asn1Tok2) {
			token = asn1Tok2[char-asn1Private]
			goto out
		}
	}
	for i := 0; i < len(asn1Tok3); i += 2 {
		token = asn1Tok3[i+0]
		if token == char {
			token = asn1Tok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = asn1Tok2[1] /* unknown char */
	}
	if asn1Debug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", asn1Tokname(token), uint(char))
	}
	return char, token
}

func asn1Parse(asn1lex asn1Lexer) int {
	return asn1NewParser().Parse(asn1lex)
}

func (asn1rcvr *asn1ParserImpl) Parse(asn1lex asn1Lexer) int {
	var asn1n int
	var asn1VAL asn1SymType
	var asn1Dollar []asn1SymType
	_ = asn1Dollar // silence set and not used
	asn1S := asn1rcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	asn1state := 0
	asn1rcvr.char = -1
	asn1token := -1 // asn1rcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		asn1state = -1
		asn1rcvr.char = -1
		asn1token = -1
	}()
	asn1p := -1
	goto asn1stack

ret0:
	return 0

ret1:
	return 1

asn1stack:
	/* put a state and value onto the stack */
	if asn1Debug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", asn1Tokname(asn1token), asn1Statname(asn1state))
	}

	asn1p++
	if asn1p >= len(asn1S) {
		nyys := make([]asn1SymType, len(asn1S)*2)
		copy(nyys, asn1S)
		asn1S = nyys
	}
	asn1S[asn1p] = asn1VAL
	asn1S[asn1p].yys = asn1state

asn1newstate:
	asn1n = asn1Pact[asn1state]
	if asn1n <= asn1Flag {
		goto asn1default /* simple state */
	}
	if asn1rcvr.char < 0 {
		asn1rcvr.char, asn1token = asn1lex1(asn1lex, &asn1rcvr.lval)
	}
	asn1n += asn1token
	if asn1n < 0 || asn1n >= asn1Last {
		goto asn1default
	}
	asn1n = asn1Act[asn1n]
	if asn1Chk[asn1n] == asn1token { /* valid shift */
		asn1rcvr.char = -1
		asn1token = -1
		asn1VAL = asn1rcvr.lval
		asn1state = asn1n
		if Errflag > 0 {
			Errflag--
		}
		goto asn1stack
	}

asn1default:
	/* default state action */
	asn1n = asn1Def[asn1state]
	if asn1n == -2 {
		if asn1rcvr.char < 0 {
			asn1rcvr.char, asn1token = asn1lex1(asn1lex, &asn1rcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if asn1Exca[xi+0] == -1 && asn1Exca[xi+1] == asn1state {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			asn1n = asn1Exca[xi+0]
			if asn1n < 0 || asn1n == asn1token {
				break
			}
		}
		asn1n = asn1Exca[xi+1]
		if asn1n < 0 {
			goto ret0
		}
	}
	if asn1n == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			asn1lex.Error(asn1ErrorMessage(asn1state, asn1token))
			Nerrs++
			if asn1Debug >= 1 {
				__yyfmt__.Printf("%s", asn1Statname(asn1state))
				__yyfmt__.Printf(" saw %s\n", asn1Tokname(asn1token))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for asn1p >= 0 {
				asn1n = asn1Pact[asn1S[asn1p].yys] + asn1ErrCode
				if asn1n >= 0 && asn1n < asn1Last {
					asn1state = asn1Act[asn1n] /* simulate a shift of "error" */
					if asn1Chk[asn1state] == asn1ErrCode {
						goto asn1stack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if asn1Debug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", asn1S[asn1p].yys)
				}
				asn1p--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if asn1Debug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", asn1Tokname(asn1token))
			}
			if asn1token == asn1EofCode {
				goto ret1
			}
			asn1rcvr.char = -1
			asn1token = -1
			goto asn1newstate /* try again in the same state */
		}
	}

	/* reduction by production asn1n */
	if asn1Debug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", asn1n, asn1Statname(asn1state))
	}

	asn1nt := asn1n
	asn1pt := asn1p
	_ = asn1pt // guard against "declared and not used"

	asn1p -= asn1R2[asn1n]
	// asn1p is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if asn1p+1 >= len(asn1S) {
		nyys := make([]asn1SymType, len(asn1S)*2)
		copy(nyys, asn1S)
		asn1S = nyys
	}
	asn1VAL = asn1S[asn1p+1]

	/* consult goto table to find next state */
	asn1n = asn1R1[asn1n]
	asn1g := asn1Pgo[asn1n]
	asn1j := asn1g + asn1S[asn1p].yys + 1

	if asn1j >= asn1Last {
		asn1state = asn1Act[asn1g]
	} else {
		asn1state = asn1Act[asn1j]
		if asn1Chk[asn1state] != -asn1n {
			asn1state = asn1Act[asn1g]
		}
	}
	// dummy call; replaced with literal code
	switch asn1nt {

	case 1:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:120
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:127
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:131
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:138
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:142
		{

			asn1VAL.module = currentModule
			fmt.Println(asn1VAL.module)
			asn1VAL.module.Name = asn1Dollar[1].str
			fmt.Println(asn1VAL.module)
		}
	case 6:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:151
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:152
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:155
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:158
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:164
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:168
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:174
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:177
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:180
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:184
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:189
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:194
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:195
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:201
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:204
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:210
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:213
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:216
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:219
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:224
		{
			asn1VAL.module = nil
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:225
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:230
		{
			asn1VAL.module = asn1types.NewAsn1Module()

			if asn1Dollar[1].module != nil {
				asn1VAL.module.Exports = asn1Dollar[1].module.Exports
			}
			if asn1Dollar[2].module != nil {
				asn1VAL.module.Imports = asn1Dollar[2].module.Imports
			}

			asn1VAL.module.Members = asn1Dollar[3].module.Members
			// FIXME: Use AssignmentList
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:245
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 29:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:248
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 30:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:258
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:264
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:271
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:279
		{
			asn1VAL.module = nil
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:283
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:289
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:296
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:300
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:304
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:311
		{
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:312
		{
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:315
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:325
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:329
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:336
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 47:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:341
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:345
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:353
		{
			asn1VAL.module = nil
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:354
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			if asn1Dollar[1].xports == nil {
				asn1VAL.module.Exports = append(asn1VAL.module.Exports, asn1Dollar[1].xports)
			} else {
				/* "EXPORTS ALL;" */
			}
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:364
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:367
		{
			asn1VAL.xports = nil
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:370
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:376
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:381
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:389
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:394
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:399
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	}
	goto asn1stack /* stack new state and value */
}
