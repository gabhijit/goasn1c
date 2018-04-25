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
	yys             int
	ch              rune
	str             string
	num             int64
	grammar         *asn1types.Asn1Grammar
	module          *asn1types.Asn1Module
	oid             *asn1types.Asn1Oid
	oid_arc         asn1types.Asn1OidArc
	module_flags    asn1types.ModuleFlagType
	xports          *asn1types.Asn1Xports
	expr            *asn1types.Asn1Expression
	aid             *asn1types.Asn1AssignedIdentifier
	expr_type       asn1types.Asn1ExprType
	tag             *asn1types.Asn1Tag
	value           *asn1types.Asn1Value
	ref             *asn1types.Asn1Reference
	constraint      *asn1types.Asn1Constraint
	marker          *asn1types.Asn1Marker
	constraint_type asn1types.Asn1ConstraintType
}

const Tok_ABSENT = 57346
const Tok_ABSTRACT_SYNTAX = 57347
const Tok_ALL = 57348
const Tok_APPLICATION = 57349
const Tok_AUTOMATIC = 57350
const Tok_BEGIN = 57351
const Tok_BIT = 57352
const Tok_BMPString = 57353
const Tok_BOOLEAN = 57354
const Tok_BY = 57355
const Tok_CHARACTER = 57356
const Tok_CHOICE = 57357
const Tok_CLASS = 57358
const Tok_COMPONENT = 57359
const Tok_COMPONENTS = 57360
const Tok_CONSTRAINED = 57361
const Tok_CONTAINING = 57362
const Tok_DEFAULT = 57363
const Tok_DEFINITIONS = 57364
const Tok_EMBEDDED = 57365
const Tok_ENCODED = 57366
const Tok_END = 57367
const Tok_ENUMERATED = 57368
const Tok_EXCEPT = 57369
const Tok_EXPLICIT = 57370
const Tok_EXPORTS = 57371
const Tok_EXTENSIBILITY = 57372
const Tok_EXTERNAL = 57373
const Tok_FALSE = 57374
const Tok_FROM = 57375
const Tok_GeneralString = 57376
const Tok_GeneralizedTime = 57377
const Tok_GraphicString = 57378
const Tok_IA5String = 57379
const Tok_IDENTIFIER = 57380
const Tok_IMPLICIT = 57381
const Tok_IMPLIED = 57382
const Tok_IMPORTS = 57383
const Tok_INCLUDES = 57384
const Tok_INSTANCE = 57385
const Tok_INTEGER = 57386
const Tok_INTERSECTION = 57387
const Tok_ISO646String = 57388
const Tok_MAX = 57389
const Tok_MIN = 57390
const Tok_MINUS_INFINITY = 57391
const Tok_NULL = 57392
const Tok_NumericString = 57393
const Tok_OBJECT = 57394
const Tok_OCTET = 57395
const Tok_OF = 57396
const Tok_OPTIONAL = 57397
const Tok_ObjectDescriptor = 57398
const Tok_PATTERN = 57399
const Tok_PDV = 57400
const Tok_PLUS_INFINITY = 57401
const Tok_PRESENT = 57402
const Tok_PRIVATE = 57403
const Tok_PrintableString = 57404
const Tok_REAL = 57405
const Tok_RELATIVE_OID = 57406
const Tok_SEQUENCE = 57407
const Tok_SET = 57408
const Tok_SIZE = 57409
const Tok_STRING = 57410
const Tok_SYNTAX = 57411
const Tok_T61String = 57412
const Tok_TAGS = 57413
const Tok_TRUE = 57414
const Tok_TYPE_IDENTIFIER = 57415
const Tok_TeletexString = 57416
const Tok_UNION = 57417
const Tok_UNIQUE = 57418
const Tok_UNIVERSAL = 57419
const Tok_UTCTime = 57420
const Tok_UTF8String = 57421
const Tok_UniversalString = 57422
const Tok_VideotexString = 57423
const Tok_VisibleString = 57424
const Tok_WITH = 57425
const Tok_ASSIGNMENT = 57426
const Tok_Ellipsis = 57427
const Tok_TwoDots = 57428
const Tok_TwoLeftBrackets = 57429
const Tok_TwoRightBrackets = 57430
const Tok_TypeReference = 57431
const Tok_CAPITALREFERENCE = 57432
const Tok_Number = 57433
const Tok_Identifier = 57434
const Tok_CString = 57435
const Tok_BString = 57436
const Tok_HString = 57437

var asn1Toknames = [...]string{
	"$end",
	"error",
	"$unk",
	"Tok_ABSENT",
	"Tok_ABSTRACT_SYNTAX",
	"Tok_ALL",
	"Tok_APPLICATION",
	"Tok_AUTOMATIC",
	"Tok_BEGIN",
	"Tok_BIT",
	"Tok_BMPString",
	"Tok_BOOLEAN",
	"Tok_BY",
	"Tok_CHARACTER",
	"Tok_CHOICE",
	"Tok_CLASS",
	"Tok_COMPONENT",
	"Tok_COMPONENTS",
	"Tok_CONSTRAINED",
	"Tok_CONTAINING",
	"Tok_DEFAULT",
	"Tok_DEFINITIONS",
	"Tok_EMBEDDED",
	"Tok_ENCODED",
	"Tok_END",
	"Tok_ENUMERATED",
	"Tok_EXCEPT",
	"Tok_EXPLICIT",
	"Tok_EXPORTS",
	"Tok_EXTENSIBILITY",
	"Tok_EXTERNAL",
	"Tok_FALSE",
	"Tok_FROM",
	"Tok_GeneralString",
	"Tok_GeneralizedTime",
	"Tok_GraphicString",
	"Tok_IA5String",
	"Tok_IDENTIFIER",
	"Tok_IMPLICIT",
	"Tok_IMPLIED",
	"Tok_IMPORTS",
	"Tok_INCLUDES",
	"Tok_INSTANCE",
	"Tok_INTEGER",
	"Tok_INTERSECTION",
	"Tok_ISO646String",
	"Tok_MAX",
	"Tok_MIN",
	"Tok_MINUS_INFINITY",
	"Tok_NULL",
	"Tok_NumericString",
	"Tok_OBJECT",
	"Tok_OCTET",
	"Tok_OF",
	"Tok_OPTIONAL",
	"Tok_ObjectDescriptor",
	"Tok_PATTERN",
	"Tok_PDV",
	"Tok_PLUS_INFINITY",
	"Tok_PRESENT",
	"Tok_PRIVATE",
	"Tok_PrintableString",
	"Tok_REAL",
	"Tok_RELATIVE_OID",
	"Tok_SEQUENCE",
	"Tok_SET",
	"Tok_SIZE",
	"Tok_STRING",
	"Tok_SYNTAX",
	"Tok_T61String",
	"Tok_TAGS",
	"Tok_TRUE",
	"Tok_TYPE_IDENTIFIER",
	"Tok_TeletexString",
	"Tok_UNION",
	"Tok_UNIQUE",
	"Tok_UNIVERSAL",
	"Tok_UTCTime",
	"Tok_UTF8String",
	"Tok_UniversalString",
	"Tok_VideotexString",
	"Tok_VisibleString",
	"Tok_WITH",
	"Tok_ASSIGNMENT",
	"Tok_Ellipsis",
	"Tok_TwoDots",
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
	"'<'",
}
var asn1Statenames = [...]string{}

const asn1EofCode = 1
const asn1ErrCode = 2
const asn1InitialStackSize = 16

//line parser/asn1p.y:1510

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	25, 26,
	-2, 45,
	-1, 150,
	97, 171,
	-2, 63,
	-1, 155,
	97, 171,
	-2, 63,
	-1, 196,
	86, 206,
	108, 206,
	-2, 158,
	-1, 209,
	104, 17,
	-2, 132,
}

const asn1Private = 57344

const asn1Last = 462

var asn1Act = [...]int{

	183, 188, 286, 206, 168, 221, 172, 91, 236, 201,
	75, 222, 16, 232, 16, 193, 196, 76, 275, 191,
	192, 74, 226, 147, 220, 270, 145, 186, 176, 4,
	10, 4, 215, 258, 322, 299, 253, 266, 249, 219,
	160, 243, 242, 52, 170, 59, 18, 67, 225, 73,
	190, 79, 336, 340, 59, 259, 79, 225, 83, 298,
	48, 321, 169, 67, 342, 268, 81, 79, 248, 289,
	52, 51, 283, 58, 333, 66, 170, 205, 334, 317,
	93, 80, 58, 318, 171, 67, 202, 68, 69, 279,
	268, 66, 207, 280, 169, 88, 263, 252, 51, 246,
	264, 5, 6, 177, 18, 178, 180, 181, 179, 70,
	338, 154, 140, 66, 337, 219, 171, 305, 329, 328,
	152, 152, 18, 327, 219, 326, 84, 295, 262, 187,
	37, 18, 79, 208, 209, 177, 18, 178, 180, 181,
	179, 79, 195, 5, 6, 199, 18, 319, 154, 166,
	216, 223, 154, 284, 49, 281, 223, 148, 233, 237,
	218, 224, 17, 18, 28, 238, 224, 282, 26, 190,
	210, 185, 217, 277, 267, 245, 213, 155, 228, 148,
	229, 150, 240, 148, 17, 18, 142, 87, 11, 241,
	14, 158, 157, 149, 144, 170, 205, 255, 85, 71,
	151, 5, 6, 331, 18, 202, 5, 6, 177, 18,
	239, 207, 18, 169, 341, 339, 177, 18, 5, 6,
	300, 18, 320, 244, 256, 5, 6, 265, 276, 260,
	154, 276, 165, 261, 269, 171, 35, 290, 139, 89,
	29, 287, 153, 153, 288, 61, 65, 33, 32, 278,
	31, 163, 208, 209, 177, 18, 178, 180, 181, 179,
	161, 195, 159, 50, 162, 216, 292, 291, 274, 223,
	294, 302, 293, 303, 297, 218, 63, 230, 137, 224,
	304, 233, 227, 307, 237, 308, 313, 217, 312, 316,
	238, 315, 314, 309, 311, 306, 310, 301, 156, 54,
	273, 5, 6, 46, 18, 34, 223, 164, 12, 134,
	43, 323, 42, 93, 93, 254, 224, 324, 325, 287,
	133, 330, 250, 335, 332, 102, 118, 105, 15, 110,
	95, 141, 138, 86, 272, 21, 36, 3, 109, 82,
	7, 101, 27, 251, 247, 8, 115, 1, 136, 119,
	112, 120, 121, 72, 20, 19, 30, 13, 170, 100,
	9, 122, 235, 234, 285, 106, 123, 113, 108, 200,
	203, 131, 170, 205, 197, 296, 169, 124, 107, 114,
	96, 97, 202, 204, 257, 125, 198, 212, 207, 126,
	169, 211, 146, 111, 128, 127, 129, 130, 171, 116,
	271, 214, 194, 189, 103, 104, 24, 154, 143, 57,
	231, 98, 171, 182, 175, 5, 6, 177, 18, 178,
	180, 181, 179, 174, 173, 167, 22, 56, 25, 208,
	209, 177, 18, 178, 180, 181, 179, 23, 195, 135,
	132, 78, 77, 117, 99, 92, 94, 90, 55, 53,
	47, 41, 40, 184, 64, 62, 60, 45, 44, 39,
	38, 2,
}
var asn1Pact = [...]int{

	136, -1000, 136, -1000, -1000, -1000, -1000, -1000, 92, 286,
	-1000, 93, 398, 71, -1000, -1000, 66, -1000, -1000, 156,
	398, -1000, 179, 177, 176, 265, -1000, -1000, 145, 327,
	-1000, -1000, -1000, -1000, -1000, 31, 283, -1000, 285, -1000,
	262, -1000, 54, -1000, 129, -1000, 212, -13, 9, -1000,
	-1000, 103, -1000, 129, -1000, -1000, -1000, -1000, -35, -51,
	-34, -1000, 129, -1000, 25, -1000, 102, -1000, -1000, 129,
	-1000, 90, -1000, -51, 155, -1000, 315, -1000, 281, 271,
	154, -1000, -1000, 136, 129, 89, -1000, -1000, -1000, 98,
	-1000, 59, -1000, -1000, -1000, 97, 85, 81, -1000, -1000,
	96, 95, 194, -64, -1000, -1000, -1000, -1000, 192, 206,
	183, -1000, -1000, 269, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 141, -1000, -1000, -1000, 326,
	92, -1000, -1000, -1000, 44, -1000, 59, -1000, 44, -46,
	39, 228, -1000, -1000, 59, 39, 223, 120, 125, 86,
	136, -1000, -1000, -1000, -1000, -61, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -63, -1000, -1000, 132,
	-1000, -1000, -1000, -1000, -1000, -1000, 78, -1000, -2, -37,
	295, -9, -1000, 288, -1000, 163, -1000, -1000, -1000, -1000,
	-1000, -1000, -51, -1000, -53, 59, 59, -1000, -64, -1000,
	-1000, 29, -1000, -1000, -1, -1000, -51, -1000, -1000, -70,
	77, -11, -1000, -51, 279, 214, -1000, 120, -1000, 76,
	120, -8, -1000, 57, 70, -29, -1000, 55, -1000, -1000,
	120, -1000, -1000, 120, -32, -1000, 152, 340, -1000, -1000,
	340, 340, -1000, -1000, 340, 28, -1000, 12, -73, 134,
	-1000, -1000, -1000, -1000, -46, -1000, 117, -1000, 30, 279,
	-1000, -1000, -1000, 326, -51, -51, -1000, -1000, -51, -1000,
	120, 117, -1000, 125, 117, -18, -1000, 49, -1000, 131,
	-40, -9, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-74, -1000, -1000, -1000, -1000, 39, -1000, -1000, -1000, 315,
	315, -1000, 26, 24, -1000, 20, 19, -1000, 120, 112,
	-23, 163, -1000, -36, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 15, 11, -1000, 124, -1000, -1000, -1000, -1000, -48,
	123, -33, -1000,
}
var asn1Pgo = [...]int{

	0, 461, 28, 0, 337, 460, 459, 458, 457, 456,
	455, 276, 454, 246, 453, 452, 451, 450, 263, 449,
	299, 448, 21, 10, 447, 446, 7, 445, 3, 444,
	443, 17, 442, 441, 440, 439, 427, 9, 425, 4,
	6, 424, 423, 414, 413, 411, 410, 13, 409, 408,
	27, 1, 403, 19, 20, 15, 402, 16, 401, 32,
	22, 24, 5, 11, 25, 400, 399, 26, 392, 23,
	391, 387, 386, 384, 383, 375, 374, 370, 145, 369,
	18, 200, 364, 2, 363, 362, 8, 360, 30, 357,
	328, 355, 354, 335, 347, 345, 344, 343,
}
var asn1R1 = [...]int{

	0, 94, 1, 1, 95, 4, 87, 87, 88, 88,
	89, 89, 90, 90, 90, 3, 2, 2, 91, 91,
	92, 92, 93, 93, 93, 93, 5, 5, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 14, 14, 11,
	12, 12, 13, 13, 13, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 19, 20, 20, 20,
	21, 22, 23, 31, 31, 32, 33, 35, 35, 35,
	35, 34, 34, 34, 24, 26, 26, 27, 27, 27,
	27, 27, 27, 25, 25, 25, 25, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	30, 30, 30, 36, 37, 37, 38, 38, 38, 38,
	38, 38, 39, 39, 41, 41, 41, 40, 44, 43,
	42, 42, 84, 85, 85, 86, 86, 86, 86, 86,
	28, 45, 45, 45, 46, 46, 47, 47, 96, 96,
	97, 97, 49, 48, 50, 50, 50, 50, 51, 51,
	52, 52, 53, 53, 54, 54, 55, 55, 56, 56,
	56, 56, 56, 58, 58, 59, 59, 59, 60, 60,
	60, 61, 61, 62, 62, 62, 63, 63, 63, 63,
	64, 64, 65, 65, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 67, 67,
	68, 68, 69, 70, 71, 72, 74, 74, 75, 75,
	73, 73, 73, 73, 76, 76, 77, 78, 79, 80,
	80, 81, 81, 81, 82, 82, 83, 83, 57,
}
var asn1R2 = [...]int{

	0, 1, 1, 2, 0, 9, 0, 1, 3, 2,
	1, 2, 1, 4, 1, 1, 1, 1, 0, 1,
	1, 2, 2, 2, 2, 2, 0, 1, 3, 0,
	1, 3, 2, 0, 1, 1, 2, 0, 1, 4,
	1, 3, 1, 3, 1, 0, 1, 3, 3, 2,
	1, 3, 1, 3, 1, 1, 2, 1, 1, 1,
	3, 1, 2, 0, 1, 2, 4, 0, 1, 1,
	1, 0, 1, 1, 2, 1, 1, 1, 4, 4,
	4, 6, 6, 1, 4, 4, 5, 1, 1, 1,
	2, 2, 2, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 2, 4, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 5, 9, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 4, 4, 1, 1,
	1, 1, 1, 3, 1, 3, 4, 4, 1, 1,
	1, 1, 3, 4, 1, 1, 3, 5, 1, 3,
	1, 3, 1, 3, 1, 3, 1, 3, 1, 1,
	1, 1, 1, 1, 3, 2, 1, 1, 1, 3,
	3, 0, 1, 1, 3, 5, 3, 2, 3, 1,
	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
	1, 2, 3, 1, 1, 3, 1, 1, 1, 1,
	1, 2, 2, 3, 2, 1, 2, 2, 2, 0,
	1, 0, 1, 1, 1, 3, 4, 4, 1,
}
var asn1Chk = [...]int{

	-1000, -94, -1, -4, -2, 89, 90, -4, -95, -87,
	-88, 96, 22, -89, 97, -90, -3, 91, 92, -91,
	-92, -93, 28, 39, 8, 30, 97, -90, 98, 84,
	-93, 71, 71, 71, 40, 91, 9, 99, -5, -6,
	-15, -16, 29, 25, -7, -8, 41, -17, 6, 100,
	-18, -2, -3, -19, -20, -21, -36, -48, -2, -3,
	-9, 33, -10, -11, -12, -13, -2, -3, 100, 101,
	100, 96, -20, 84, -22, -23, -31, -32, -33, 102,
	-22, 100, -11, 33, 101, 96, -18, 97, -22, 84,
	-24, -26, -27, -28, -25, 15, 65, 66, -45, -29,
	44, 26, 10, 89, 90, 12, 50, 63, 53, 23,
	14, 78, 35, 52, 64, 31, -66, -30, 11, 34,
	36, 37, 46, 51, 62, 70, 74, 80, 79, 81,
	82, 56, -34, 39, 28, -35, 77, 7, 61, 84,
	-2, -13, 97, -49, 96, -67, -68, -69, 98, 96,
	96, -81, -69, -78, 67, 96, -81, 96, 96, 68,
	104, 68, 58, 68, 38, 91, -37, -38, -39, 50,
	32, 72, -40, -41, -42, -43, -2, 91, 93, 96,
	94, 95, -44, -3, -14, -88, -50, 85, -51, -52,
	6, -53, -54, -55, -56, 98, -57, -76, -72, -78,
	-79, -37, 42, -77, -74, 33, -28, 48, 89, 90,
	-69, -70, -71, -50, -58, -59, -3, -60, -23, 85,
	-61, -62, -63, -3, -23, 18, -60, 54, -69, -61,
	54, -46, -47, -3, -84, -85, -86, -3, -40, 85,
	96, -2, 103, 104, 91, 97, 101, -96, 105, 75,
	27, -97, 106, 45, 27, -51, -22, -73, 86, 108,
	-69, -67, 99, 97, 101, -23, 107, 97, 101, -23,
	-64, -65, 55, 21, 54, -80, -3, 97, -80, 97,
	101, 98, 97, 101, 98, -82, -83, -3, -3, 101,
	85, -53, -55, -54, -55, 99, -75, -57, 47, 108,
	86, -59, -39, -40, -63, 87, -64, -37, -23, -31,
	-31, -47, -40, -39, -86, -40, -39, 97, 101, 98,
	91, 101, 108, -62, -26, -26, 99, 99, 99, 99,
	-83, 91, -39, 97, 101, -51, 88, 99, 99, 91,
	101, 91, 97,
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
	62, 198, 75, 76, 77, 0, 221, 221, 130, 83,
	100, 101, 0, 131, 132, 87, 88, 89, 0, 0,
	0, 93, 94, 0, 96, 97, 98, 99, 184, 185,
	186, 187, 188, 189, 190, 191, 192, 193, 194, 195,
	196, 197, 65, 72, 73, 0, 68, 69, 70, 0,
	37, 41, 43, 143, 0, 74, 199, 200, 0, 63,
	-2, 0, 222, 223, 0, -2, 0, 0, 0, 102,
	0, 90, 91, 92, 95, 0, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 0, 117, 114, 0,
	120, 121, 119, 118, 39, 38, 0, 144, 145, 148,
	0, 150, 152, 154, 156, 0, -2, 159, 160, 161,
	162, 228, 63, 215, 0, 0, 198, 207, 131, -2,
	201, 0, 203, 204, 0, 163, 63, 166, 167, 168,
	0, 172, 173, 63, 180, 0, 179, 219, 217, 0,
	219, 0, 134, 0, 0, 122, 123, 125, 128, 129,
	0, 133, 66, 0, 0, 142, 0, 0, 138, 139,
	0, 0, 140, 141, 0, 0, 214, 0, 210, 0,
	218, 216, 202, 78, 63, 165, 0, 79, 63, 180,
	177, 181, 182, 0, 63, 63, 220, 80, 63, 84,
	0, 0, 85, 0, 0, 0, 224, 0, 113, 0,
	146, 151, 149, 153, 155, 157, 205, 208, 209, 211,
	212, 164, 169, 170, 174, 63, 176, 183, 178, 0,
	0, 135, 0, 0, 124, 0, 0, 86, 0, 0,
	0, 0, 213, 0, 81, 82, 136, 137, 126, 127,
	225, 0, 0, 115, 0, 147, 175, 226, 227, 0,
	0, 0, 116,
}
var asn1Tok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 107, 3, 3, 3, 3, 3, 3,
	98, 99, 3, 3, 101, 3, 104, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 100,
	108, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 102, 3, 103, 106, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 96, 105, 97,
}
var asn1Tok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95,
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
		//line parser/asn1p.y:284
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:291
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:295
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:302
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:306
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
		//line parser/asn1p.y:322
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:323
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:326
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:329
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:335
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:339
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:345
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:348
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:351
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:355
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:360
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:363
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:368
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:369
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:375
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:378
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:384
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:387
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:390
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:393
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:398
		{
			asn1VAL.module = nil
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:399
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:404
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
		//line parser/asn1p.y:423
		{
			asn1VAL.module = nil
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:427
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:433
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:440
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:444
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:448
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:455
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:456
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:459
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:469
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:473
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:480
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:485
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:489
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:497
		{
			asn1VAL.module = nil
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:498
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
		//line parser/asn1p.y:508
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:511
		{
			asn1VAL.xports = nil
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:514
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:520
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:525
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:533
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:538
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:543
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:552
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:555
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:565
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:569
		{
			fmt.Println("DDD")
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:574
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:581
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:591
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:596
		{
			asn1VAL.tag = nil
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:597
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:601
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:608
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:614
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:615
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:616
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:617
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:621
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:622
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:623
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 74:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:627
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:631
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 78:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:636
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrChoice
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:641
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSequence
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:646
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSet
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:652
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 82:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:658
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:666
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 84:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:671
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:676
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:681
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[4].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeBitString
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:690
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:691
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:692
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:693
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:694
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:695
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:696
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:697
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 95:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:698
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 96:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:699
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 97:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:700
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 100:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:706
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 101:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:707
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 102:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:708
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 103:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:717
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 106:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:738
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 107:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:742
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 108:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:746
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 113:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:758
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 114:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:767
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:770
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
	case 116:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:782
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
	case 117:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:801
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 118:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:807
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 119:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:812
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:817
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:820
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:826
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 123:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:832
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 124:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:835
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 125:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:842
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 126:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:849
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:856
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 128:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:863
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 129:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:870
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 130:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:890
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 131:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:910
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:913
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:916
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:922
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 135:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:925
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 136:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:931
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 137:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:937
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 142:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:950
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 143:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:953
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 144:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:961
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 146:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:966
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 147:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:970
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 149:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:978
		{
		}
	case 151:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:984
		{
		}
	case 153:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:990
		{
		}
	case 155:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:997
		{
		}
	case 157:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1003
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSet
		}
	case 158:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1010
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeValue
		}
	case 159:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1014
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSubType
			/* FIXME: */
		}
	case 163:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1026
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 164:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1029
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 165:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1035
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 166:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1038
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 167:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1041
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 168:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1047
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 169:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1053
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 170:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1059
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 171:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1071
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 172:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1072
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 173:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1075
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 174:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1078
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 175:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1081
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 176:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1087
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 177:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1091
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 178:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1094
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			//$$.Meta = $3.Meta
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeComponentsOf
		}
	case 179:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1099
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 180:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1109
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagNoMark
			asn1VAL.marker.Value = nil
		}
	case 181:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1114
		{
			asn1VAL.marker = asn1Dollar[1].marker
		}
	case 182:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1118
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect
			asn1VAL.marker.Value = nil
		}
	case 183:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1123
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagDefault
			asn1VAL.marker.Value = asn1Dollar[2].value
		}
	case 184:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1131
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBMPString
		}
	case 185:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1132
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralString
			fmt.Println("WARNING: GeneralString is not fully supported")
		}
	case 186:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1136
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGraphicsString
			fmt.Println("WARNING: GraphicString is not fully supported")
		}
	case 187:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1140
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeIA5String
		}
	case 188:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1141
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeISO646String
		}
	case 189:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1142
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNumericString
		}
	case 190:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1143
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypePrintableString
		}
	case 191:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1144
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeT61String
			fmt.Println("WARNING: T61String is not fully supported")
		}
	case 192:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1148
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeTeletexString
		}
	case 193:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1149
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUniversalString
		}
	case 194:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1150
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtf8String
		}
	case 195:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1151
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVideoTexString
			fmt.Println("WARNING: VideotexString is not fully supported")
		}
	case 196:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1155
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVisibleString
		}
	case 197:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1156
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectDescriptor
		}
	case 198:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1162
		{
			asn1VAL.constraint = nil
		}
	case 201:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1167
		{
			// FIXME : Implementation incomplete
		}
	case 202:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1173
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
			// FIXME: Needs implementation
		}
	case 205:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1184
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1Dollar[2].constraint_type
		}
	case 207:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1191
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMin
		}
	case 209:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1198
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMax
		}
	case 210:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1204
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeSimpleRange
		}
	case 211:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1205
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeRightExcludedRange
		}
	case 212:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1206
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftExcludedRange
		}
	case 213:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1207
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftRightExcludedRange
		}
	case 214:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1211
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 215:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1215
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 216:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1221
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// FIXME : Implementation
		}
	case 217:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1228
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			/* FIXME : implementation not yet complete */
		}
	case 218:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1233
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
		}
	case 219:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1238
		{
			asn1VAL.str = ""
		}
	case 220:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1239
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 221:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1245
		{
			asn1VAL.constraint = nil
		}
	case 224:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1251
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 225:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1254
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 226:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1260
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 227:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1266
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	}
	goto asn1stack /* stack new state and value */
}
