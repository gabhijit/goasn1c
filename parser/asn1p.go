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
	constraint   *asn1types.Asn1Constraint
}

const Tok_ALL = 57346
const Tok_APPLICATION = 57347
const Tok_ASSIGNMENT = 57348
const Tok_AUTOMATIC = 57349
const Tok_BEGIN = 57350
const Tok_BIT = 57351
const Tok_BOOLEAN = 57352
const Tok_CHARACTER = 57353
const Tok_DEFINITIONS = 57354
const Tok_EMBEDDED = 57355
const Tok_END = 57356
const Tok_ENUMERATED = 57357
const Tok_EXCEPT = 57358
const Tok_EXPLICIT = 57359
const Tok_EXPORTS = 57360
const Tok_EXTENSIBILITY = 57361
const Tok_EXTERNAL = 57362
const Tok_Ellipsis = 57363
const Tok_FALSE = 57364
const Tok_FROM = 57365
const Tok_GeneralizedTime = 57366
const Tok_IDENTIFIER = 57367
const Tok_IMPLICIT = 57368
const Tok_IMPLIED = 57369
const Tok_IMPORTS = 57370
const Tok_INTEGER = 57371
const Tok_INTERSECTION = 57372
const Tok_NULL = 57373
const Tok_OBJECT = 57374
const Tok_OCTET = 57375
const Tok_PDV = 57376
const Tok_PRIVATE = 57377
const Tok_REAL = 57378
const Tok_RELATIVE_OID = 57379
const Tok_STRING = 57380
const Tok_TAGS = 57381
const Tok_TRUE = 57382
const Tok_UNION = 57383
const Tok_UNIVERSAL = 57384
const Tok_UTCTime = 57385
const Tok_TypeReference = 57386
const Tok_Number = 57387
const Tok_Identifier = 57388
const Tok_CString = 57389
const Tok_BString = 57390
const Tok_HString = 57391

var asn1Toknames = [...]string{
	"$end",
	"error",
	"$unk",
	"Tok_ALL",
	"Tok_APPLICATION",
	"Tok_ASSIGNMENT",
	"Tok_AUTOMATIC",
	"Tok_BEGIN",
	"Tok_BIT",
	"Tok_BOOLEAN",
	"Tok_CHARACTER",
	"Tok_DEFINITIONS",
	"Tok_EMBEDDED",
	"Tok_END",
	"Tok_ENUMERATED",
	"Tok_EXCEPT",
	"Tok_EXPLICIT",
	"Tok_EXPORTS",
	"Tok_EXTENSIBILITY",
	"Tok_EXTERNAL",
	"Tok_Ellipsis",
	"Tok_FALSE",
	"Tok_FROM",
	"Tok_GeneralizedTime",
	"Tok_IDENTIFIER",
	"Tok_IMPLICIT",
	"Tok_IMPLIED",
	"Tok_IMPORTS",
	"Tok_INTEGER",
	"Tok_INTERSECTION",
	"Tok_NULL",
	"Tok_OBJECT",
	"Tok_OCTET",
	"Tok_PDV",
	"Tok_PRIVATE",
	"Tok_REAL",
	"Tok_RELATIVE_OID",
	"Tok_STRING",
	"Tok_TAGS",
	"Tok_TRUE",
	"Tok_UNION",
	"Tok_UNIVERSAL",
	"Tok_UTCTime",
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
	"'|'",
	"'^'",
}
var asn1Statenames = [...]string{}

const asn1EofCode = 1
const asn1ErrCode = 2
const asn1InitialStackSize = 16

//line parser/asn1p.y:1193

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 35,
	14, 25,
	-2, 44,
}

const asn1Private = 57344

const asn1Last = 283

var asn1Act = [...]int{

	155, 150, 139, 160, 159, 170, 135, 158, 64, 166,
	9, 49, 15, 164, 15, 143, 4, 73, 4, 185,
	181, 62, 175, 174, 72, 78, 214, 188, 82, 47,
	215, 189, 67, 68, 53, 218, 209, 195, 180, 192,
	178, 80, 69, 51, 213, 58, 212, 66, 211, 184,
	210, 201, 36, 193, 58, 190, 27, 50, 220, 57,
	83, 65, 191, 66, 157, 177, 16, 17, 57, 5,
	51, 17, 25, 122, 78, 157, 79, 65, 86, 48,
	85, 154, 137, 81, 50, 66, 10, 71, 16, 17,
	87, 136, 121, 137, 13, 126, 125, 124, 120, 65,
	138, 84, 136, 70, 5, 144, 17, 145, 147, 148,
	146, 138, 162, 17, 173, 5, 144, 17, 145, 147,
	148, 146, 5, 162, 17, 137, 219, 167, 171, 172,
	217, 152, 208, 133, 136, 5, 144, 17, 144, 17,
	176, 60, 132, 138, 34, 5, 32, 5, 144, 17,
	145, 147, 148, 146, 31, 162, 137, 45, 30, 131,
	129, 127, 5, 187, 17, 136, 128, 33, 130, 196,
	41, 186, 20, 182, 138, 42, 117, 194, 5, 144,
	17, 145, 147, 148, 146, 11, 198, 197, 199, 35,
	200, 167, 29, 203, 171, 172, 206, 204, 205, 202,
	207, 111, 99, 104, 119, 103, 118, 97, 23, 14,
	216, 88, 109, 116, 28, 183, 106, 114, 21, 179,
	24, 96, 26, 100, 107, 102, 113, 22, 101, 108,
	3, 90, 7, 6, 1, 105, 98, 19, 18, 12,
	8, 169, 168, 163, 161, 156, 153, 123, 56, 165,
	94, 149, 142, 141, 140, 134, 55, 115, 112, 77,
	76, 75, 110, 95, 92, 91, 93, 89, 74, 54,
	52, 46, 40, 39, 151, 63, 61, 59, 44, 43,
	38, 37, 2,
}
var asn1Pact = [...]int{

	101, -1000, 101, -1000, -1000, -1000, -1000, 36, 173, -1000,
	43, 201, 21, -1000, -1000, 4, -1000, -1000, 208, 201,
	-1000, 119, 115, 107, 140, -1000, -1000, 99, 181, -1000,
	-1000, -1000, -1000, -1000, -1, 152, -1000, 161, -1000, 129,
	-1000, 25, -1000, 78, -1000, 118, -22, -12, -1000, -1000,
	53, -1000, 78, -1000, -1000, -1000, -1000, 18, -31, -13,
	-1000, 78, -1000, 5, -1000, 51, -1000, -1000, 78, -1000,
	27, -1000, -31, 205, -1000, 192, -1000, 200, 171, 198,
	-1000, -1000, 101, 78, 22, -1000, -1000, -1000, 47, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 46, 45, -1000, -1000,
	-1000, -1000, 123, 132, 122, -1000, -1000, 143, -1000, -1000,
	-1000, 121, -1000, -1000, -1000, 97, -1000, -1000, -1000, 134,
	36, -1000, -1000, -1000, 60, 67, 93, -1000, -1000, -1000,
	-1000, -1000, -34, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -36, -1000, -1000, 95, -1000, -1000, -1000,
	-1000, -1000, -1000, 14, -1000, -15, -21, 157, -11, -1000,
	155, -1000, 71, -1000, -1000, -24, -1000, 3, 11, -16,
	-1000, 1, -1000, -1000, -1000, 67, -18, -1000, 148, 103,
	-1000, -1000, 103, 103, -1000, -1000, 103, -2, -1000, 67,
	91, -1000, 93, 91, -1000, 87, -19, -11, -1000, -1000,
	-1000, -1000, -1000, -3, -5, -1000, -7, -9, -25, 71,
	-1000, -1000, -1000, -1000, -1000, 85, -1000, -20, 81, 7,
	-1000,
}
var asn1Pgo = [...]int{

	0, 282, 15, 1, 230, 281, 280, 279, 278, 277,
	276, 21, 275, 8, 274, 273, 272, 271, 11, 270,
	34, 269, 17, 268, 267, 266, 265, 264, 263, 262,
	261, 260, 259, 258, 257, 256, 13, 255, 6, 2,
	254, 253, 252, 251, 250, 249, 9, 248, 247, 246,
	0, 245, 7, 4, 3, 244, 243, 242, 241, 5,
	240, 10, 239, 209, 238, 237, 172, 234, 232, 231,
	219, 215,
}
var asn1R1 = [...]int{

	0, 67, 1, 1, 68, 4, 60, 60, 61, 61,
	62, 62, 63, 63, 63, 3, 2, 64, 64, 65,
	65, 66, 66, 66, 66, 5, 5, 6, 7, 7,
	8, 8, 9, 9, 10, 10, 14, 14, 11, 12,
	12, 13, 13, 13, 15, 15, 16, 16, 16, 17,
	17, 18, 18, 18, 19, 19, 20, 20, 20, 21,
	22, 23, 30, 30, 31, 32, 34, 34, 34, 34,
	33, 33, 33, 24, 69, 69, 26, 25, 25, 25,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 29, 29, 29, 35, 36, 36, 37, 37,
	37, 37, 37, 37, 38, 38, 40, 40, 40, 39,
	43, 42, 41, 41, 57, 58, 58, 59, 59, 59,
	59, 59, 27, 44, 45, 45, 46, 46, 70, 70,
	71, 71, 48, 47, 49, 49, 49, 49, 50, 50,
	51, 51, 52, 52, 53, 53, 54, 54, 55, 56,
}
var asn1R2 = [...]int{

	0, 1, 1, 2, 0, 9, 0, 1, 3, 2,
	1, 2, 1, 4, 1, 1, 1, 0, 1, 1,
	2, 2, 2, 2, 2, 0, 1, 3, 0, 1,
	3, 2, 0, 1, 1, 2, 0, 1, 4, 1,
	3, 1, 3, 1, 0, 1, 3, 3, 2, 1,
	3, 1, 3, 1, 1, 2, 1, 1, 1, 3,
	1, 2, 0, 1, 2, 4, 0, 1, 1, 1,
	0, 1, 1, 1, 1, 1, 1, 1, 4, 4,
	1, 1, 1, 2, 2, 2, 1, 1, 2, 1,
	1, 1, 1, 1, 2, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 5, 9, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 4, 4,
	1, 1, 1, 1, 1, 3, 4, 4, 1, 1,
	1, 1, 3, 4, 1, 1, 3, 5, 1, 3,
	1, 3, 1, 3, 1, 3, 1, 3, 1, 1,
}
var asn1Chk = [...]int{

	-1000, -67, -1, -4, -2, 44, -4, -68, -60, -61,
	50, 12, -62, 51, -63, -3, 45, 46, -64, -65,
	-66, 17, 26, 7, 19, 51, -63, 52, 6, -66,
	39, 39, 39, 27, 45, 8, 53, -5, -6, -15,
	-16, 18, 14, -7, -8, 28, -17, 4, 54, -18,
	-2, -3, -19, -20, -21, -35, -47, -2, -3, -9,
	23, -10, -11, -12, -13, -2, -3, 54, 55, 54,
	50, -20, 6, -22, -23, -30, -31, -32, 56, -22,
	54, -11, 23, 55, 50, -18, 51, -22, 6, -24,
	-69, -26, -27, -25, -44, -28, 29, 15, 44, 10,
	31, 36, 33, 13, 11, 43, 24, 32, 37, 20,
	-29, 9, -33, 26, 17, -34, 42, 5, 35, 6,
	-2, -13, 51, -48, 50, 50, 50, 38, 34, 38,
	25, 38, 45, -36, -37, -38, 31, 22, 40, -39,
	-40, -41, -42, -2, 45, 47, 50, 48, 49, -43,
	-3, -14, -61, -49, 21, -50, -51, 4, -52, -53,
	-54, -55, 52, -56, -36, -45, -46, -3, -57, -58,
	-59, -3, -39, 21, 57, 58, 45, 51, 55, -70,
	59, 41, 16, -71, 60, 30, 16, -50, 51, 55,
	52, 51, 55, 52, -3, 55, 21, -52, -54, -53,
	-54, 53, -46, -39, -38, -59, -39, -38, 45, 55,
	53, 53, 53, 53, 51, 55, -50, 45, 55, 45,
	51,
}
var asn1Def = [...]int{

	0, -2, 1, 2, 4, 16, 3, 6, 0, 7,
	0, 17, 0, 9, 10, 12, 14, 15, 0, 18,
	19, 0, 0, 0, 0, 8, 11, 0, 0, 20,
	21, 22, 23, 24, 0, -2, 13, 0, 26, 28,
	45, 0, 5, 0, 29, 32, 0, 0, 48, 49,
	51, 53, 27, 54, 56, 57, 58, 62, 62, 0,
	31, 33, 34, 0, 39, 41, 43, 46, 0, 47,
	0, 55, 62, 0, 60, 0, 63, 70, 66, 0,
	30, 35, 0, 0, 0, 50, 52, 59, 0, 61,
	73, 74, 75, 76, 122, 77, 92, 93, 123, 80,
	81, 82, 0, 0, 0, 86, 87, 0, 89, 90,
	91, 0, 64, 71, 72, 0, 67, 68, 69, 0,
	36, 40, 42, 133, 0, 0, 0, 83, 84, 85,
	88, 94, 0, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 0, 109, 106, 0, 112, 113, 111,
	110, 38, 37, 0, 134, 135, 138, 0, 140, 142,
	144, 146, 0, 148, 149, 0, 124, 0, 0, 114,
	115, 117, 120, 121, 65, 0, 0, 132, 0, 0,
	128, 129, 0, 0, 130, 131, 0, 0, 78, 0,
	0, 79, 0, 0, 105, 0, 136, 141, 139, 143,
	145, 147, 125, 0, 0, 116, 0, 0, 0, 0,
	126, 127, 118, 119, 107, 0, 137, 0, 0, 0,
	108,
}
var asn1Tok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	52, 53, 3, 3, 55, 3, 58, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 54,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 56, 3, 57, 60, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 50, 59, 51,
}
var asn1Tok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49,
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
		//line parser/asn1p.y:198
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:205
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:209
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:216
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:220
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
		//line parser/asn1p.y:236
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:237
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:240
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:243
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:249
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:253
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:259
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:262
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:265
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:269
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:274
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:279
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:280
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:286
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:289
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:295
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:298
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:301
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:304
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:309
		{
			asn1VAL.module = nil
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:310
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:315
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
		//line parser/asn1p.y:334
		{
			asn1VAL.module = nil
		}
	case 30:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:338
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:344
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:351
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 34:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:355
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:359
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:366
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:367
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:370
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:380
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:384
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:391
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:396
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:400
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:408
		{
			asn1VAL.module = nil
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:409
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
		//line parser/asn1p.y:419
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 47:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:422
		{
			asn1VAL.xports = nil
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:425
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:431
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:436
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:444
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:449
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:454
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:463
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:466
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:476
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:480
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:484
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:491
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 61:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:501
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:506
		{
			asn1VAL.tag = nil
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:507
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:511
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:518
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:524
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:525
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:526
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:527
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:531
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:532
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:533
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:537
		{
		}
	case 77:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:547
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 78:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:552
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:557
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:565
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:566
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 82:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:567
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:568
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 84:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:569
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:570
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:571
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:572
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:573
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:574
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:575
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:582
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:583
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:584
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 95:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:593
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 98:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:614
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 99:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:618
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 100:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:622
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 105:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:634
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 106:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:643
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 107:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:646
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
	case 108:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:658
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
	case 109:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:677
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 110:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:683
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 111:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:688
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 112:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:693
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 113:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:696
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 114:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:702
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:708
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 116:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:711
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 117:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:718
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 118:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:725
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 119:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:732
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:739
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:746
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:766
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 123:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:785
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 124:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:790
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 125:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:793
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 126:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:799
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:805
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:818
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:821
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:829
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeExtensibilityMark
		}
	case 136:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:834
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeExtensibilityMark
		}
	case 137:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:838
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeExtensibilityMark
		}
	case 139:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:846
		{
		}
	case 141:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:852
		{
		}
	case 143:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:858
		{
		}
	case 145:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:865
		{
		}
	case 147:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:871
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeSet
		}
	case 148:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:878
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeValue
		}
	}
	goto asn1stack /* stack new state and value */
}
