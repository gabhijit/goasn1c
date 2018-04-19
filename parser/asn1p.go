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
	tag          *asn1types.Asn1Tag
	value        *asn1types.Asn1Value
	ref          *asn1types.Asn1Reference
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
const Tok_UNIVERSAL = 57360
const Tok_APPLICATION = 57361
const Tok_PRIVATE = 57362
const Tok_BOOLEAN = 57363
const Tok_NULL = 57364
const Tok_FALSE = 57365
const Tok_TRUE = 57366
const Tok_REAL = 57367
const Tok_OCTET = 57368
const Tok_STRING = 57369
const Tok_OBJECT = 57370
const Tok_IDENTIFIER = 57371
const Tok_RELATIVE_OID = 57372
const Tok_EXTERNAL = 57373
const Tok_EMBEDDED = 57374
const Tok_PDV = 57375
const Tok_CHARACTER = 57376
const Tok_UTCTime = 57377
const Tok_GeneralizedTime = 57378
const Tok_INTEGER = 57379
const Tok_ENUMERATED = 57380
const Tok_BIT = 57381
const Tok_Ellipsis = 57382
const Tok_TypeReference = 57383
const Tok_Number = 57384
const Tok_Identifier = 57385
const Tok_CString = 57386
const Tok_BString = 57387
const Tok_HString = 57388

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
	"Tok_UNIVERSAL",
	"Tok_APPLICATION",
	"Tok_PRIVATE",
	"Tok_BOOLEAN",
	"Tok_NULL",
	"Tok_FALSE",
	"Tok_TRUE",
	"Tok_REAL",
	"Tok_OCTET",
	"Tok_STRING",
	"Tok_OBJECT",
	"Tok_IDENTIFIER",
	"Tok_RELATIVE_OID",
	"Tok_EXTERNAL",
	"Tok_EMBEDDED",
	"Tok_PDV",
	"Tok_CHARACTER",
	"Tok_UTCTime",
	"Tok_GeneralizedTime",
	"Tok_INTEGER",
	"Tok_ENUMERATED",
	"Tok_BIT",
	"Tok_Ellipsis",
	"Tok_TypeReference",
	"Tok_Number",
	"Tok_Identifier",
	"Tok_CString",
	"Tok_BString",
	"Tok_HString",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"';'",
	"','",
	"'['",
	"']'",
	"'.'",
}
var asn1Statenames = [...]string{}

const asn1EofCode = 1
const asn1ErrCode = 2
const asn1InitialStackSize = 16

//line parser/asn1p.y:731

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 35,
	5, 25,
	-2, 44,
}

const asn1Private = 57344

const asn1Last = 194

var asn1Act = [...]int{

	65, 64, 4, 150, 4, 119, 9, 148, 72, 63,
	49, 15, 61, 15, 120, 121, 122, 144, 152, 162,
	77, 47, 80, 163, 66, 67, 53, 167, 156, 154,
	78, 68, 165, 5, 128, 17, 129, 131, 132, 130,
	164, 36, 51, 50, 57, 56, 5, 157, 17, 27,
	169, 16, 17, 57, 56, 10, 48, 25, 155, 81,
	16, 17, 116, 84, 135, 17, 13, 82, 51, 50,
	69, 168, 59, 79, 5, 128, 17, 166, 83, 70,
	85, 151, 114, 128, 17, 158, 145, 134, 127, 141,
	123, 115, 93, 94, 34, 5, 95, 96, 5, 101,
	17, 102, 103, 97, 139, 98, 99, 100, 105, 92,
	106, 5, 137, 17, 140, 138, 136, 111, 112, 113,
	20, 143, 41, 22, 45, 21, 24, 11, 23, 14,
	32, 31, 30, 108, 86, 109, 149, 33, 71, 28,
	29, 42, 26, 35, 3, 153, 88, 6, 7, 1,
	19, 18, 12, 8, 147, 146, 133, 149, 134, 127,
	126, 160, 125, 161, 159, 124, 118, 117, 55, 110,
	107, 76, 75, 74, 104, 91, 89, 90, 87, 73,
	54, 52, 46, 40, 39, 142, 62, 60, 58, 44,
	43, 38, 37, 2,
}
var asn1Pact = [...]int{

	54, -1000, 54, -1000, -1000, -1000, -1000, 8, 121, -1000,
	18, 115, 9, -1000, -1000, 0, -1000, -1000, 132, 115,
	-1000, 120, 119, 118, 128, -1000, -1000, 52, 139, -1000,
	-1000, -1000, -1000, -1000, -9, 105, -1000, 136, -1000, 110,
	-1000, 5, -1000, 70, -1000, 57, -27, -20, -1000, -1000,
	23, -1000, 70, -1000, -1000, -1000, 131, -33, -21, -1000,
	70, -1000, 7, -1000, 20, -1000, -1000, 70, -1000, 15,
	-1000, -33, 127, -1000, 71, -1000, 125, 99, -1000, -1000,
	54, 70, 14, -1000, -1000, -1000, -8, -1000, -1000, -1000,
	-1000, -1000, 17, -1000, -1000, -1000, 89, 79, 88, -1000,
	-1000, 75, -1000, -1000, -1000, -1000, 87, -1000, -1000, -1000,
	47, -1000, -1000, -1000, 8, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -38, -1000, -1000,
	44, -1000, -1000, -1000, -1000, 41, -1000, -1000, -1000, -1000,
	-1000, -36, -1000, -1000, 22, -23, 10, -24, -1000, -2,
	-1000, -1000, -1000, -1000, 43, -1000, 41, 33, -29, -1000,
	-10, -18, -1000, 35, -1000, -1000, -25, 29, 2, -1000,
}
var asn1Pgo = [...]int{

	0, 193, 1, 0, 144, 192, 191, 190, 189, 188,
	187, 12, 186, 9, 185, 184, 183, 182, 10, 181,
	26, 180, 8, 179, 178, 177, 176, 175, 174, 173,
	172, 171, 170, 169, 168, 167, 166, 5, 3, 165,
	162, 160, 156, 155, 154, 7, 153, 6, 152, 129,
	151, 150, 120, 149, 148, 146,
}
var asn1R1 = [...]int{

	0, 53, 1, 1, 54, 4, 46, 46, 47, 47,
	48, 48, 49, 49, 49, 3, 2, 50, 50, 51,
	51, 52, 52, 52, 52, 5, 5, 6, 7, 7,
	8, 8, 9, 9, 10, 10, 14, 14, 11, 12,
	12, 13, 13, 13, 15, 15, 16, 16, 16, 17,
	17, 18, 18, 18, 19, 19, 20, 20, 21, 22,
	23, 29, 29, 30, 31, 33, 33, 33, 33, 32,
	32, 32, 24, 55, 26, 25, 25, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 28,
	28, 28, 34, 35, 35, 36, 36, 36, 36, 36,
	36, 37, 37, 39, 39, 39, 38, 42, 41, 40,
	40, 43, 44, 44, 45, 45, 45, 45, 45,
}
var asn1R2 = [...]int{

	0, 1, 1, 2, 0, 9, 0, 1, 3, 2,
	1, 2, 1, 4, 1, 1, 1, 0, 1, 1,
	2, 2, 2, 2, 2, 0, 1, 3, 0, 1,
	3, 2, 0, 1, 1, 2, 0, 1, 4, 1,
	3, 1, 3, 1, 0, 1, 3, 3, 2, 1,
	3, 1, 3, 1, 1, 2, 1, 1, 3, 1,
	2, 0, 1, 2, 4, 0, 1, 1, 1, 0,
	1, 1, 1, 1, 1, 1, 4, 1, 1, 1,
	2, 2, 2, 1, 1, 2, 1, 1, 1, 1,
	1, 2, 4, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 5, 9, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 4, 4, 1, 1,
}
var asn1Chk = [...]int{

	-1000, -53, -1, -4, -2, 41, -4, -54, -46, -47,
	47, 6, -48, 48, -49, -3, 42, 43, -50, -51,
	-52, 10, 8, 13, 11, 48, -49, 49, 7, -52,
	12, 12, 12, 9, 42, 4, 50, -5, -6, -15,
	-16, 17, 5, -7, -8, 14, -17, 16, 51, -18,
	-2, -3, -19, -20, -21, -34, -2, -3, -9, 15,
	-10, -11, -12, -13, -2, -3, 51, 52, 51, 47,
	-20, 7, -22, -23, -29, -30, -31, 53, 51, -11,
	15, 52, 47, -18, 48, -22, 7, -24, -55, -26,
	-25, -27, 38, 21, 22, 25, 26, 32, 34, 35,
	36, 28, 30, 31, -28, 37, 39, -32, 8, 10,
	-33, 18, 19, 20, -2, -13, 48, -35, -36, -37,
	22, 23, 24, -38, -39, -40, -41, -2, 42, 44,
	47, 45, 46, -42, -3, 47, 27, 33, 27, 29,
	27, 42, -14, -47, 55, 42, -43, -44, -45, -3,
	-38, 40, 54, -3, 52, 48, 52, 49, 42, -45,
	-38, -37, 48, 52, 50, 50, 42, 52, 42, 48,
}
var asn1Def = [...]int{

	0, -2, 1, 2, 4, 16, 3, 6, 0, 7,
	0, 17, 0, 9, 10, 12, 14, 15, 0, 18,
	19, 0, 0, 0, 0, 8, 11, 0, 0, 20,
	21, 22, 23, 24, 0, -2, 13, 0, 26, 28,
	45, 0, 5, 0, 29, 32, 0, 0, 48, 49,
	51, 53, 27, 54, 56, 57, 0, 61, 0, 31,
	33, 34, 0, 39, 41, 43, 46, 0, 47, 0,
	55, 61, 0, 59, 0, 62, 69, 65, 30, 35,
	0, 0, 0, 50, 52, 58, 0, 60, 72, 73,
	74, 75, 90, 77, 78, 79, 0, 0, 0, 83,
	84, 0, 86, 87, 88, 89, 0, 63, 70, 71,
	0, 66, 67, 68, 36, 40, 42, 92, 93, 94,
	95, 96, 97, 98, 99, 100, 101, 0, 106, 103,
	0, 109, 110, 108, 107, 0, 80, 81, 82, 85,
	91, 0, 38, 37, 0, 0, 0, 111, 112, 114,
	117, 118, 64, 102, 0, 76, 0, 0, 0, 113,
	0, 0, 104, 0, 115, 116, 0, 0, 0, 105,
}
var asn1Tok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 50, 3, 3, 52, 3, 55, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 51,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 53, 3, 54, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 47, 3, 48,
}
var asn1Tok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46,
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
		//line parser/asn1p.y:176
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:183
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:187
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:194
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:198
		{

			asn1VAL.module = currentModule
			fmt.Println(asn1VAL.module)
			asn1VAL.module.Name = asn1Dollar[1].str

			// We have been given a module body
			if asn1Dollar[8].module != nil {
				asn1VAL.module.Imports = asn1Dollar[8].module.Imports
				asn1VAL.module.Exports = asn1Dollar[8].module.Exports
				asn1VAL.module.Members = asn1Dollar[8].module.Members
			}
			fmt.Println(asn1VAL.module)
		}
	case 6:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:214
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:215
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:218
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:221
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:227
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:231
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:237
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:240
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:243
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:247
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:252
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:257
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:258
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:264
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:267
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:273
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:276
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:279
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:282
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:287
		{
			asn1VAL.module = nil
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:288
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:293
		{
			asn1VAL.module = asn1types.NewAsn1Module()

			if asn1Dollar[1].module != nil {
				asn1VAL.module.Exports = asn1Dollar[1].module.Exports
			}
			if asn1Dollar[2].module != nil {
				asn1VAL.module.Imports = asn1Dollar[2].module.Imports
			}

			asn1VAL.module.Members = asn1Dollar[3].module.Members
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:312
		{
			asn1VAL.module = nil
		}
	case 30:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:316
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:322
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:329
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 34:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:333
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:337
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:344
		{
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:345
		{
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:348
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:358
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:362
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:369
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:374
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:378
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:386
		{
			asn1VAL.module = nil
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:387
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			if asn1Dollar[1].xports == nil {
				asn1VAL.module.Exports = append(asn1VAL.module.Exports, asn1Dollar[1].xports)
			} else {
				/* "EXPORTS ALL;" */
			}
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:397
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 47:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:400
		{
			asn1VAL.xports = nil
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:403
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:409
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:414
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:422
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:427
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:432
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:441
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:444
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:454
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:458
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:465
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:475
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 61:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:480
		{
			asn1VAL.tag = nil
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:481
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:485
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:492
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:498
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:499
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:500
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:501
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:505
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:506
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:507
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:511
		{
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:521
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 76:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:533
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 77:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:541
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 78:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:542
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:543
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:544
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:545
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 82:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:546
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:547
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 84:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:548
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:549
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:550
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:551
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:558
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:559
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:560
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:569
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 95:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:590
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 96:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:594
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 97:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:598
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 102:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:610
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 103:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:619
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 104:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:622
		{
			if asn1Dollar[2].num < 0 || asn1Dollar[2].num > 7 {
				return -1
			}
			if asn1Dollar[4].num < 0 || asn1Dollar[4].num > 15 {
				return -1
			}
			//value := $2 << 4 + $4;

			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTuple
		}
	case 105:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:634
		{
			if asn1Dollar[2].num < 0 || asn1Dollar[2].num > 127 {
				return -1
			}
			if asn1Dollar[4].num < 0 || asn1Dollar[4].num > 255 {
				return -1
			}
			if asn1Dollar[6].num < 0 || asn1Dollar[6].num > 255 {
				return -1
			}
			if asn1Dollar[8].num < 0 || asn1Dollar[8].num > 255 {
				return -1
			}
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeQuadruple
		}
	case 106:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:653
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 107:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:659
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 108:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:664
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 109:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:669
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 110:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:672
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 111:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:678
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 112:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:684
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 113:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:687
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 114:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:694
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:701
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 116:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:708
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 117:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:715
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 118:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:722
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	}
	goto asn1stack /* stack new state and value */
}
