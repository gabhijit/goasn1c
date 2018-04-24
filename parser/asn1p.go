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
	marker       *asn1types.Asn1Marker
}

const Tok_ALL = 57346
const Tok_APPLICATION = 57347
const Tok_ASSIGNMENT = 57348
const Tok_AUTOMATIC = 57349
const Tok_BEGIN = 57350
const Tok_BIT = 57351
const Tok_BOOLEAN = 57352
const Tok_CHARACTER = 57353
const Tok_CHOICE = 57354
const Tok_COMPONENTS = 57355
const Tok_DEFAULT = 57356
const Tok_DEFINITIONS = 57357
const Tok_EMBEDDED = 57358
const Tok_END = 57359
const Tok_ENUMERATED = 57360
const Tok_EXCEPT = 57361
const Tok_EXPLICIT = 57362
const Tok_EXPORTS = 57363
const Tok_EXTENSIBILITY = 57364
const Tok_EXTERNAL = 57365
const Tok_FALSE = 57366
const Tok_FROM = 57367
const Tok_GeneralizedTime = 57368
const Tok_IDENTIFIER = 57369
const Tok_IMPLICIT = 57370
const Tok_IMPLIED = 57371
const Tok_IMPORTS = 57372
const Tok_INTEGER = 57373
const Tok_INTERSECTION = 57374
const Tok_NULL = 57375
const Tok_OF = 57376
const Tok_OBJECT = 57377
const Tok_OCTET = 57378
const Tok_OPTIONAL = 57379
const Tok_PDV = 57380
const Tok_PRIVATE = 57381
const Tok_REAL = 57382
const Tok_RELATIVE_OID = 57383
const Tok_SEQUENCE = 57384
const Tok_SET = 57385
const Tok_STRING = 57386
const Tok_TAGS = 57387
const Tok_TRUE = 57388
const Tok_UNION = 57389
const Tok_UNIVERSAL = 57390
const Tok_UTCTime = 57391
const Tok_BMPString = 57392
const Tok_GeneralString = 57393
const Tok_GraphicString = 57394
const Tok_IA5String = 57395
const Tok_ISO646String = 57396
const Tok_NumericString = 57397
const Tok_PrintableString = 57398
const Tok_T61String = 57399
const Tok_TeletexString = 57400
const Tok_UniversalString = 57401
const Tok_UTF8String = 57402
const Tok_VideotexString = 57403
const Tok_VisibleString = 57404
const Tok_ObjectDescriptor = 57405
const Tok_Ellipsis = 57406
const Tok_TwoLeftBrackets = 57407
const Tok_TwoRightBrackets = 57408
const Tok_TypeReference = 57409
const Tok_CAPITALREFERENCE = 57410
const Tok_Number = 57411
const Tok_Identifier = 57412
const Tok_CString = 57413
const Tok_BString = 57414
const Tok_HString = 57415

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
	"Tok_CHOICE",
	"Tok_COMPONENTS",
	"Tok_DEFAULT",
	"Tok_DEFINITIONS",
	"Tok_EMBEDDED",
	"Tok_END",
	"Tok_ENUMERATED",
	"Tok_EXCEPT",
	"Tok_EXPLICIT",
	"Tok_EXPORTS",
	"Tok_EXTENSIBILITY",
	"Tok_EXTERNAL",
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
	"Tok_OF",
	"Tok_OBJECT",
	"Tok_OCTET",
	"Tok_OPTIONAL",
	"Tok_PDV",
	"Tok_PRIVATE",
	"Tok_REAL",
	"Tok_RELATIVE_OID",
	"Tok_SEQUENCE",
	"Tok_SET",
	"Tok_STRING",
	"Tok_TAGS",
	"Tok_TRUE",
	"Tok_UNION",
	"Tok_UNIVERSAL",
	"Tok_UTCTime",
	"Tok_BMPString",
	"Tok_GeneralString",
	"Tok_GraphicString",
	"Tok_IA5String",
	"Tok_ISO646String",
	"Tok_NumericString",
	"Tok_PrintableString",
	"Tok_T61String",
	"Tok_TeletexString",
	"Tok_UniversalString",
	"Tok_UTF8String",
	"Tok_VideotexString",
	"Tok_VisibleString",
	"Tok_ObjectDescriptor",
	"Tok_Ellipsis",
	"Tok_TwoLeftBrackets",
	"Tok_TwoRightBrackets",
	"Tok_TypeReference",
	"Tok_CAPITALREFERENCE",
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
	"'!'",
}
var asn1Statenames = [...]string{}

const asn1EofCode = 1
const asn1ErrCode = 2
const asn1InitialStackSize = 16

//line parser/asn1p.y:1400

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	17, 26,
	-2, 45,
	-1, 146,
	75, 164,
	-2, 63,
	-1, 147,
	75, 164,
	-2, 63,
}

const asn1Private = 57344

const asn1Last = 371

var asn1Act = [...]int{

	174, 179, 196, 163, 159, 188, 208, 234, 201, 190,
	204, 182, 16, 183, 16, 199, 195, 10, 230, 65,
	197, 224, 220, 194, 167, 4, 214, 4, 150, 18,
	213, 278, 275, 200, 184, 79, 276, 280, 240, 79,
	200, 73, 241, 52, 232, 59, 227, 67, 68, 69,
	228, 269, 247, 244, 59, 232, 217, 81, 219, 70,
	274, 273, 83, 67, 272, 271, 245, 51, 253, 58,
	52, 66, 37, 223, 75, 75, 242, 28, 58, 131,
	104, 109, 95, 282, 194, 67, 108, 66, 101, 75,
	18, 194, 258, 114, 51, 74, 111, 18, 243, 239,
	79, 100, 48, 105, 141, 112, 107, 79, 140, 66,
	106, 113, 96, 97, 181, 79, 84, 231, 216, 110,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 161, 17, 18, 102, 103, 142,
	87, 26, 11, 160, 18, 157, 191, 198, 198, 205,
	209, 161, 149, 210, 192, 80, 162, 148, 176, 147,
	160, 193, 146, 145, 202, 5, 6, 144, 18, 88,
	85, 71, 281, 162, 178, 212, 49, 5, 6, 168,
	18, 169, 171, 172, 170, 279, 186, 268, 226, 5,
	6, 168, 18, 215, 5, 6, 168, 18, 169, 171,
	172, 170, 50, 186, 17, 18, 211, 229, 5, 6,
	14, 168, 18, 156, 233, 246, 5, 6, 248, 18,
	61, 35, 63, 54, 137, 33, 32, 31, 155, 191,
	249, 153, 151, 198, 256, 255, 251, 192, 254, 152,
	238, 259, 205, 260, 193, 209, 263, 264, 210, 266,
	267, 265, 262, 257, 261, 181, 250, 161, 138, 198,
	252, 270, 5, 6, 237, 18, 160, 136, 24, 46,
	34, 277, 86, 154, 134, 161, 42, 72, 21, 162,
	43, 22, 133, 25, 160, 82, 225, 236, 221, 23,
	12, 36, 139, 89, 15, 29, 222, 162, 218, 30,
	5, 6, 168, 18, 169, 171, 172, 170, 27, 3,
	8, 1, 7, 20, 19, 13, 9, 207, 5, 6,
	168, 18, 169, 171, 172, 170, 206, 186, 115, 235,
	189, 187, 185, 180, 177, 143, 57, 203, 98, 173,
	166, 165, 164, 158, 56, 135, 132, 78, 77, 76,
	116, 99, 93, 92, 91, 94, 90, 55, 53, 47,
	41, 40, 175, 64, 62, 60, 45, 44, 39, 38,
	2,
}
var asn1Pact = [...]int{

	141, -1000, 141, -1000, -1000, -1000, -1000, -1000, 68, 275,
	-1000, 135, 261, 66, -1000, -1000, 1, -1000, -1000, 289,
	261, -1000, 182, 181, 180, 241, -1000, -1000, 152, 283,
	-1000, -1000, -1000, -1000, -1000, -5, 255, -1000, 263, -1000,
	239, -1000, 98, -1000, 149, -1000, 195, -30, -19, -1000,
	-1000, 97, -1000, 149, -1000, -1000, -1000, -1000, 35, -45,
	-21, -1000, 149, -1000, 37, -1000, 96, -1000, -1000, 149,
	-1000, 65, -1000, -45, 287, -1000, 70, -1000, 254, 219,
	286, -1000, -1000, 141, 149, 64, -1000, -1000, -1000, 93,
	-1000, -1000, -1000, -1000, -1000, 89, 88, 85, -1000, -1000,
	83, 78, -54, -1000, -1000, -1000, -1000, 188, 201, 187,
	-1000, -1000, 246, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 184, -1000, -1000, -1000, 144, -1000, -1000, -1000, 233,
	68, -1000, -1000, -1000, 110, -41, 20, 20, 74, 142,
	141, -1000, -1000, -1000, -1000, -1000, -51, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -56, -1000, -1000,
	124, -1000, -1000, -1000, -1000, -1000, -1000, 43, -1000, -23,
	-25, 269, -11, -1000, 267, -1000, 251, -1000, -1000, -29,
	-1000, -45, -1000, -1000, -67, 42, -24, -1000, -45, 250,
	206, -1000, 24, -37, -1000, 0, 23, -26, -1000, -10,
	-1000, -1000, -1000, -1000, 74, -27, -1000, 154, 127, -1000,
	-1000, 127, 127, -1000, -1000, 127, -9, -1000, -41, -1000,
	122, -1000, 27, 250, -1000, -1000, -1000, 233, -45, -1000,
	-1000, 74, 122, -1000, 142, 122, -1000, 118, -28, -11,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 20, -1000,
	-1000, -1000, -1000, -12, -13, -1000, -16, -17, -43, 251,
	-35, -1000, -1000, -1000, -1000, -1000, 116, -1000, -1000, -42,
	103, 8, -1000,
}
var asn1Pgo = [...]int{

	0, 370, 24, 0, 309, 369, 368, 367, 366, 365,
	364, 222, 363, 19, 362, 361, 360, 359, 202, 358,
	223, 357, 95, 15, 356, 355, 354, 353, 352, 351,
	350, 349, 348, 347, 346, 345, 344, 5, 343, 4,
	3, 342, 341, 340, 339, 338, 337, 10, 336, 335,
	334, 1, 333, 11, 13, 34, 332, 331, 330, 9,
	8, 16, 2, 20, 7, 329, 328, 326, 317, 6,
	316, 17, 315, 294, 314, 313, 278, 311, 310, 298,
	296,
}
var asn1R1 = [...]int{

	0, 77, 1, 1, 78, 4, 70, 70, 71, 71,
	72, 72, 73, 73, 73, 3, 2, 2, 74, 74,
	75, 75, 76, 76, 76, 76, 5, 5, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 14, 14, 11,
	12, 12, 13, 13, 13, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 19, 20, 20, 20,
	21, 22, 23, 31, 31, 32, 33, 35, 35, 35,
	35, 34, 34, 34, 24, 26, 26, 27, 27, 27,
	27, 25, 25, 25, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 30, 30, 30,
	36, 37, 37, 38, 38, 38, 38, 38, 38, 39,
	39, 41, 41, 41, 40, 44, 43, 42, 42, 67,
	68, 68, 69, 69, 69, 69, 69, 28, 45, 45,
	45, 46, 46, 47, 47, 79, 79, 80, 80, 49,
	48, 50, 50, 50, 50, 51, 51, 52, 52, 53,
	53, 54, 54, 55, 55, 56, 58, 58, 59, 59,
	59, 60, 60, 60, 61, 61, 62, 62, 62, 63,
	63, 63, 63, 64, 64, 65, 65, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 57,
}
var asn1R2 = [...]int{

	0, 1, 1, 2, 0, 9, 0, 1, 3, 2,
	1, 2, 1, 4, 1, 1, 1, 1, 0, 1,
	1, 2, 2, 2, 2, 2, 0, 1, 3, 0,
	1, 3, 2, 0, 1, 1, 2, 0, 1, 4,
	1, 3, 1, 3, 1, 0, 1, 3, 3, 2,
	1, 3, 1, 3, 1, 1, 2, 1, 1, 1,
	3, 1, 2, 0, 1, 2, 4, 0, 1, 1,
	1, 0, 1, 1, 1, 1, 1, 1, 4, 4,
	4, 1, 4, 4, 1, 1, 1, 2, 2, 2,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 2,
	4, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 5, 9, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 4, 4, 1, 1, 1, 1, 1,
	3, 1, 3, 4, 4, 1, 1, 1, 1, 3,
	4, 1, 1, 3, 5, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 3, 1, 1, 3, 2, 1,
	1, 1, 3, 3, 0, 1, 1, 3, 5, 3,
	2, 3, 1, 0, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1,
}
var asn1Chk = [...]int{

	-1000, -77, -1, -4, -2, 67, 68, -4, -78, -70,
	-71, 74, 15, -72, 75, -73, -3, 69, 70, -74,
	-75, -76, 20, 28, 7, 22, 75, -73, 76, 6,
	-76, 45, 45, 45, 29, 69, 8, 77, -5, -6,
	-15, -16, 21, 17, -7, -8, 30, -17, 4, 78,
	-18, -2, -3, -19, -20, -21, -36, -48, -2, -3,
	-9, 25, -10, -11, -12, -13, -2, -3, 78, 79,
	78, 74, -20, 6, -22, -23, -31, -32, -33, 80,
	-22, 78, -11, 25, 79, 74, -18, 75, -22, 6,
	-24, -26, -27, -28, -25, 12, 42, 43, -45, -29,
	31, 18, 67, 68, 10, 33, 40, 36, 16, 11,
	49, 26, 35, 41, 23, -66, -30, 50, 51, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 9, -34, 28, 20, -35, 48, 5, 39, 6,
	-2, -13, 75, -49, 74, 74, 74, 74, 74, 74,
	82, 44, 38, 44, 27, 44, 69, -37, -38, -39,
	33, 24, 46, -40, -41, -42, -43, -2, 69, 71,
	74, 72, 73, -44, -3, -14, -71, -50, 64, -51,
	-52, 4, -53, -54, -55, -56, 76, -57, -37, -58,
	-59, -3, -60, -23, 64, -61, -62, -63, -3, -23,
	13, -60, -61, -46, -47, -3, -67, -68, -69, -3,
	-40, 64, -2, 81, 82, 69, 75, 79, -79, 83,
	47, 19, -80, 84, 32, 19, -51, 75, 79, -23,
	85, 75, 79, -23, -64, -65, 37, 14, 34, 75,
	75, 79, 76, 75, 79, 76, -3, 79, 64, -53,
	-55, -54, -55, 77, -59, -39, -40, -63, 65, -64,
	-37, -23, -47, -40, -39, -69, -40, -39, 69, 79,
	-62, 77, 77, 77, 77, 75, 79, -51, 66, 69,
	79, 69, 75,
}
var asn1Def = [...]int{

	0, -2, 1, 2, 4, 16, 17, 3, 6, 0,
	7, 0, 18, 0, 9, 10, 12, 14, 15, 0,
	19, 20, 0, 0, 0, 0, 8, 11, 0, 0,
	21, 22, 23, 24, 25, 0, -2, 13, 0, 27,
	29, 46, 0, 5, 0, 30, 33, 0, 0, 49,
	50, 52, 54, 28, 55, 57, 58, 59, 63, 63,
	0, 32, 34, 35, 0, 40, 42, 44, 47, 0,
	48, 0, 56, 63, 0, 61, 0, 64, 71, 67,
	0, 31, 36, 0, 0, 0, 51, 53, 60, 0,
	62, 74, 75, 76, 77, 0, 0, 0, 127, 81,
	97, 98, 128, 129, 84, 85, 86, 0, 0, 0,
	90, 91, 0, 93, 94, 95, 96, 177, 178, 179,
	180, 181, 182, 183, 184, 185, 186, 187, 188, 189,
	190, 0, 65, 72, 73, 0, 68, 69, 70, 0,
	37, 41, 43, 140, 0, 63, -2, -2, 0, 0,
	0, 87, 88, 89, 92, 99, 0, 100, 101, 102,
	103, 104, 105, 106, 107, 108, 109, 0, 114, 111,
	0, 117, 118, 116, 115, 39, 38, 0, 141, 142,
	145, 0, 147, 149, 151, 153, 0, 155, 191, 0,
	156, 63, 159, 160, 161, 0, 165, 166, 63, 173,
	0, 172, 0, 0, 131, 0, 0, 119, 120, 122,
	125, 126, 130, 66, 0, 0, 139, 0, 0, 135,
	136, 0, 0, 137, 138, 0, 0, 78, 63, 158,
	0, 79, 63, 173, 170, 174, 175, 0, 63, 80,
	82, 0, 0, 83, 0, 0, 110, 0, 143, 148,
	146, 150, 152, 154, 157, 162, 163, 167, 63, 169,
	176, 171, 132, 0, 0, 121, 0, 0, 0, 0,
	0, 133, 134, 123, 124, 112, 0, 144, 168, 0,
	0, 0, 113,
}
var asn1Tok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 85, 3, 3, 3, 3, 3, 3,
	76, 77, 3, 3, 79, 3, 82, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 78,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 80, 3, 81, 84, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 74, 83, 75,
}
var asn1Tok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73,
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
		//line parser/asn1p.y:238
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:245
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:249
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:256
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:260
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
		//line parser/asn1p.y:276
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:277
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:280
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:283
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:289
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:293
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:299
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:302
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:305
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:309
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:314
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:317
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:322
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:323
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:329
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:332
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:338
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:341
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:344
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:347
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:352
		{
			asn1VAL.module = nil
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:353
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:358
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
	case 29:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:377
		{
			asn1VAL.module = nil
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:381
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:387
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:394
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:398
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:402
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:409
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:410
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:413
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:423
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:427
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:434
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:439
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:443
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:451
		{
			asn1VAL.module = nil
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:452
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			if asn1Dollar[1].xports == nil {
				asn1VAL.module.Exports = append(asn1VAL.module.Exports, asn1Dollar[1].xports)
			} else {
				/* "EXPORTS ALL;" */
			}
		}
	case 47:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:462
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:465
		{
			asn1VAL.xports = nil
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:468
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:474
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:479
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:487
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:492
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:497
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:506
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:509
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:519
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:523
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:527
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:534
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:544
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:549
		{
			asn1VAL.tag = nil
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:550
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:554
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:561
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:567
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:568
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:569
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:570
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:574
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:575
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:576
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 74:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:580
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:584
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 78:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:589
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrChoice
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:594
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSequence
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:599
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSet
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:607
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 82:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:612
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:617
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 84:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:625
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:626
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:627
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:628
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:629
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:630
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:631
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:632
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:633
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:634
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:635
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 97:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:641
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 98:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:642
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 99:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:643
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 100:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:652
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 103:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:673
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 104:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:677
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 105:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:681
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 110:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:693
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 111:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:702
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 112:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:705
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
	case 113:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:717
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
	case 114:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:736
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:742
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 116:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:747
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 117:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:752
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 118:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:755
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 119:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:761
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:767
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:770
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:777
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 123:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:784
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 124:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:791
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 125:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:798
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 126:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:805
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:825
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 128:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:845
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 129:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:848
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 130:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:851
		{
			fmt.Println("XXX")
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 131:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:858
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:861
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:867
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:873
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 139:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:886
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 140:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:889
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 141:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:897
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeExtensibilityMark
		}
	case 143:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:902
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeExtensibilityMark
		}
	case 144:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:906
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeExtensibilityMark
		}
	case 146:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:914
		{
		}
	case 148:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:920
		{
		}
	case 150:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:926
		{
		}
	case 152:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:933
		{
		}
	case 154:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:939
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeSet
		}
	case 155:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:946
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.ConstraintTypeValue
		}
	case 156:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:954
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 157:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:957
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 158:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:963
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 159:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:966
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 160:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:969
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 161:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:975
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 162:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:981
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 163:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:987
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 164:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:999
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 165:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1000
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 166:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1003
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 167:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1006
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 168:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1009
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 169:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1015
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 170:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1019
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 171:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1022
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			//$$.Meta = $3.Meta
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeComponentsOf
		}
	case 172:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1027
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 173:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1037
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagNoMark
			asn1VAL.marker.Value = nil
		}
	case 174:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1042
		{
			asn1VAL.marker = asn1Dollar[1].marker
		}
	case 175:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1046
		{
			fmt.Println("YYY")
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect
			asn1VAL.marker.Value = nil
		}
	case 176:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1052
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagDefault
			asn1VAL.marker.Value = asn1Dollar[2].value
		}
	case 177:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1060
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBMPString
		}
	case 178:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1061
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralString
			fmt.Println("WARNING: GeneralString is not fully supported")
		}
	case 179:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1065
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGraphicsString
			fmt.Println("WARNING: GraphicString is not fully supported")
		}
	case 180:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1069
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeIA5String
		}
	case 181:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1070
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeISO646String
		}
	case 182:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1071
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNumericString
		}
	case 183:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1072
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypePrintableString
		}
	case 184:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1073
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeT61String
			fmt.Println("WARNING: T61String is not fully supported")
		}
	case 185:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1077
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeTeletexString
		}
	case 186:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1078
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUniversalString
		}
	case 187:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1079
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtf8String
		}
	case 188:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1080
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVideoTexString
			fmt.Println("WARNING: VideotexString is not fully supported")
		}
	case 189:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1084
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVisibleString
		}
	case 190:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1085
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectDescriptor
		}
	}
	goto asn1stack /* stack new state and value */
}
