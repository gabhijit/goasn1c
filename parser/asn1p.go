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

//line parser/asn1p.y:1537

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	25, 26,
	-2, 45,
	-1, 150,
	97, 173,
	-2, 63,
	-1, 155,
	97, 173,
	-2, 63,
	-1, 196,
	86, 208,
	108, 208,
	-2, 160,
	-1, 209,
	104, 17,
	-2, 134,
}

const asn1Private = 57344

const asn1Last = 471

var asn1Act = [...]int{

	183, 188, 290, 168, 206, 221, 91, 247, 232, 172,
	274, 236, 16, 201, 16, 196, 193, 226, 215, 279,
	192, 75, 76, 147, 74, 145, 191, 220, 222, 176,
	4, 262, 4, 186, 10, 329, 305, 270, 253, 257,
	160, 243, 242, 52, 73, 59, 170, 67, 79, 225,
	219, 190, 343, 263, 59, 340, 323, 18, 225, 341,
	324, 304, 79, 67, 169, 272, 81, 79, 252, 294,
	52, 283, 51, 295, 58, 284, 66, 170, 205, 83,
	48, 93, 347, 58, 80, 67, 171, 202, 68, 69,
	267, 328, 66, 207, 268, 169, 293, 287, 88, 51,
	256, 272, 250, 5, 6, 177, 18, 178, 180, 181,
	179, 70, 154, 140, 66, 345, 219, 171, 190, 344,
	152, 152, 336, 18, 335, 219, 154, 311, 334, 333,
	187, 301, 18, 79, 208, 209, 177, 18, 178, 180,
	181, 179, 79, 195, 170, 205, 266, 84, 37, 199,
	216, 223, 325, 166, 202, 155, 223, 148, 233, 237,
	207, 288, 169, 5, 6, 154, 18, 217, 238, 285,
	210, 218, 224, 148, 49, 185, 28, 224, 228, 154,
	248, 151, 213, 229, 171, 244, 18, 17, 18, 349,
	241, 246, 286, 26, 150, 240, 148, 259, 281, 170,
	205, 208, 209, 177, 18, 178, 180, 181, 179, 202,
	195, 271, 17, 18, 249, 207, 142, 169, 14, 87,
	11, 158, 157, 149, 144, 85, 71, 260, 280, 264,
	239, 280, 265, 18, 154, 348, 177, 18, 269, 171,
	346, 291, 326, 165, 292, 273, 153, 153, 5, 6,
	282, 5, 6, 338, 18, 35, 208, 209, 177, 18,
	178, 180, 181, 179, 296, 195, 306, 5, 6, 216,
	18, 298, 65, 223, 308, 300, 299, 303, 297, 156,
	309, 139, 89, 29, 312, 233, 217, 307, 237, 319,
	218, 313, 322, 317, 224, 318, 248, 238, 321, 320,
	314, 310, 315, 327, 33, 316, 5, 6, 177, 18,
	32, 137, 223, 31, 50, 163, 161, 330, 159, 63,
	93, 93, 331, 332, 162, 291, 278, 337, 230, 339,
	342, 227, 277, 224, 102, 118, 105, 54, 110, 95,
	61, 46, 24, 134, 164, 34, 42, 109, 258, 254,
	101, 43, 12, 36, 133, 115, 255, 141, 119, 112,
	120, 121, 22, 251, 25, 138, 276, 170, 100, 15,
	122, 21, 8, 23, 106, 123, 113, 108, 1, 20,
	131, 136, 82, 27, 86, 169, 124, 107, 114, 96,
	97, 72, 30, 19, 125, 13, 5, 6, 126, 18,
	9, 235, 111, 128, 127, 129, 130, 171, 3, 234,
	245, 7, 289, 103, 104, 200, 203, 197, 302, 204,
	261, 198, 212, 211, 5, 6, 177, 18, 178, 180,
	181, 179, 146, 116, 275, 214, 194, 189, 143, 57,
	231, 98, 182, 175, 174, 173, 167, 56, 135, 132,
	78, 77, 117, 99, 92, 94, 90, 55, 53, 47,
	41, 40, 184, 64, 62, 60, 45, 44, 39, 38,
	2,
}
var asn1Pact = [...]int{

	159, -1000, 159, -1000, -1000, -1000, -1000, -1000, 124, 330,
	-1000, 121, 334, 96, -1000, -1000, 78, -1000, -1000, 199,
	334, -1000, 242, 239, 233, 305, -1000, -1000, 164, 344,
	-1000, -1000, -1000, -1000, -1000, 49, 317, -1000, 326, -1000,
	300, -1000, 74, -1000, 178, -1000, 307, -12, 11, -1000,
	-1000, 130, -1000, 178, -1000, -1000, -1000, -1000, -40, -54,
	-34, -1000, 178, -1000, 46, -1000, 129, -1000, -1000, 178,
	-1000, 122, -1000, -54, 198, -1000, 324, -1000, 315, 304,
	197, -1000, -1000, 159, 178, 119, -1000, -1000, -1000, 128,
	-1000, 75, -1000, -1000, -1000, 127, 98, 59, -1000, -1000,
	126, 125, 250, -64, -1000, -1000, -1000, -1000, 248, 266,
	247, -1000, -1000, 306, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 152, -1000, -1000, -1000, 335,
	124, -1000, -1000, -1000, 45, -1000, 75, -1000, 45, -35,
	31, 277, -1000, -1000, 75, 31, 274, 141, 145, 99,
	159, -1000, -1000, -1000, -1000, -61, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -63, -1000, -1000, 94,
	-1000, -1000, -1000, -1000, -1000, -1000, 117, -1000, 1, -37,
	322, -6, -1000, 321, -1000, 112, -1000, -1000, -1000, -1000,
	-1000, -1000, -54, -1000, -55, 75, 75, -1000, -64, -1000,
	-1000, 47, -1000, -1000, -7, -1000, -54, -1000, -1000, -70,
	114, 0, -1000, -54, 311, 272, -1000, 141, -1000, 101,
	141, -26, -1000, 71, 95, -4, -1000, 63, -1000, -1000,
	141, -1000, -1000, 141, -5, -28, -1000, -1000, -1000, -1000,
	179, 167, -1000, -1000, 167, 167, -1000, -1000, 167, 32,
	-1000, 14, -72, 180, -1000, -1000, -1000, -1000, -35, -1000,
	217, -1000, 40, 311, -1000, -1000, -1000, 335, -54, -54,
	-1000, -1000, -54, -1000, 141, 217, -1000, 145, 217, -41,
	-1000, 54, -1000, 151, -1000, 141, -10, -6, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -73, -1000, -1000, -1000,
	-1000, 31, -1000, -1000, -1000, 324, 324, -1000, 30, 29,
	-1000, 25, 23, -1000, 141, 162, -42, -1000, 112, -1000,
	-36, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 20, 16,
	-1000, 149, -1000, -1000, -1000, -1000, -19, 144, 92, -1000,
}
var asn1Pgo = [...]int{

	0, 470, 29, 0, 408, 469, 468, 467, 466, 465,
	464, 319, 463, 272, 462, 461, 460, 459, 314, 458,
	337, 457, 24, 21, 456, 455, 6, 454, 4, 453,
	452, 22, 451, 450, 449, 448, 447, 13, 446, 3,
	9, 445, 444, 443, 442, 441, 440, 8, 439, 438,
	33, 1, 437, 26, 20, 16, 436, 15, 435, 18,
	17, 27, 5, 28, 10, 434, 433, 25, 432, 23,
	423, 422, 421, 420, 419, 418, 417, 416, 149, 415,
	19, 181, 412, 2, 410, 7, 409, 401, 11, 400,
	34, 395, 369, 393, 379, 371, 378, 372, 363, 356,
}
var asn1R1 = [...]int{

	0, 96, 1, 1, 97, 4, 89, 89, 90, 90,
	91, 91, 92, 92, 92, 3, 2, 2, 93, 93,
	94, 94, 95, 95, 95, 95, 5, 5, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 14, 14, 11,
	12, 12, 13, 13, 13, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 19, 20, 20, 20,
	21, 22, 23, 31, 31, 32, 33, 35, 35, 35,
	35, 34, 34, 34, 24, 26, 26, 27, 27, 27,
	27, 27, 27, 25, 25, 25, 25, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	30, 30, 30, 36, 37, 37, 38, 38, 38, 38,
	38, 38, 39, 39, 41, 41, 41, 40, 44, 43,
	42, 42, 42, 42, 86, 87, 87, 88, 88, 88,
	88, 88, 28, 45, 45, 45, 46, 46, 47, 47,
	98, 98, 99, 99, 49, 48, 50, 50, 50, 50,
	51, 51, 52, 52, 53, 53, 54, 54, 55, 55,
	56, 56, 56, 56, 56, 58, 58, 59, 59, 59,
	60, 60, 60, 61, 61, 62, 62, 62, 63, 63,
	63, 63, 64, 64, 65, 65, 66, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	67, 67, 68, 68, 69, 70, 71, 72, 74, 74,
	75, 75, 73, 73, 73, 73, 76, 76, 77, 78,
	79, 80, 80, 81, 81, 81, 82, 82, 83, 83,
	84, 84, 85, 57,
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
	1, 1, 3, 2, 1, 1, 3, 1, 4, 4,
	1, 1, 1, 1, 1, 3, 1, 3, 4, 4,
	1, 1, 1, 1, 3, 4, 1, 1, 3, 5,
	1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
	1, 1, 1, 1, 1, 1, 3, 2, 1, 1,
	1, 3, 3, 0, 1, 1, 3, 5, 3, 2,
	3, 1, 0, 1, 1, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	0, 1, 1, 2, 3, 1, 1, 3, 1, 1,
	1, 1, 1, 2, 2, 3, 2, 1, 2, 2,
	2, 0, 1, 0, 1, 1, 1, 3, 4, 4,
	1, 3, 1, 1,
}
var asn1Chk = [...]int{

	-1000, -96, -1, -4, -2, 89, 90, -4, -97, -89,
	-90, 96, 22, -91, 97, -92, -3, 91, 92, -93,
	-94, -95, 28, 39, 8, 30, 97, -92, 98, 84,
	-95, 71, 71, 71, 40, 91, 9, 99, -5, -6,
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
	94, 95, -44, -3, -14, -90, -50, 85, -51, -52,
	6, -53, -54, -55, -56, 98, -57, -76, -72, -78,
	-79, -37, 42, -77, -74, 33, -28, 48, 89, 90,
	-69, -70, -71, -50, -58, -59, -3, -60, -23, 85,
	-61, -62, -63, -3, -23, 18, -60, 54, -69, -61,
	54, -46, -47, -3, -86, -87, -88, -3, -40, 85,
	96, -2, 103, 104, 91, -84, 97, -85, -3, 97,
	101, -98, 105, 75, 27, -99, 106, 45, 27, -51,
	-22, -73, 86, 108, -69, -67, 99, 97, 101, -23,
	107, 97, 101, -23, -64, -65, 55, 21, 54, -80,
	-3, 97, -80, 97, 101, 98, 97, 101, 98, -82,
	-83, -3, -3, 101, 97, 101, 85, -53, -55, -54,
	-55, 99, -75, -57, 47, 108, 86, -59, -39, -40,
	-63, 87, -64, -37, -23, -31, -31, -47, -40, -39,
	-88, -40, -39, 97, 101, 98, 91, -85, 101, 108,
	-62, -26, -26, 99, 99, 99, 99, -83, 91, -39,
	97, 101, -51, 88, 99, 99, 91, 101, 91, 97,
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
	62, 200, 75, 76, 77, 0, 223, 223, 132, 83,
	100, 101, 0, 133, 134, 87, 88, 89, 0, 0,
	0, 93, 94, 0, 96, 97, 98, 99, 186, 187,
	188, 189, 190, 191, 192, 193, 194, 195, 196, 197,
	198, 199, 65, 72, 73, 0, 68, 69, 70, 0,
	37, 41, 43, 145, 0, 74, 201, 202, 0, 63,
	-2, 0, 224, 225, 0, -2, 0, 0, 0, 102,
	0, 90, 91, 92, 95, 0, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 0, 117, 114, 0,
	120, 121, 119, 118, 39, 38, 0, 146, 147, 150,
	0, 152, 154, 156, 158, 0, -2, 161, 162, 163,
	164, 233, 63, 217, 0, 0, 200, 209, 133, -2,
	203, 0, 205, 206, 0, 165, 63, 168, 169, 170,
	0, 174, 175, 63, 182, 0, 181, 221, 219, 0,
	221, 0, 136, 0, 0, 124, 125, 127, 130, 131,
	0, 135, 66, 0, 0, 0, 123, 230, 232, 144,
	0, 0, 140, 141, 0, 0, 142, 143, 0, 0,
	216, 0, 212, 0, 220, 218, 204, 78, 63, 167,
	0, 79, 63, 182, 179, 183, 184, 0, 63, 63,
	222, 80, 63, 84, 0, 0, 85, 0, 0, 0,
	226, 0, 113, 0, 122, 0, 148, 153, 151, 155,
	157, 159, 207, 210, 211, 213, 214, 166, 171, 172,
	176, 63, 178, 185, 180, 0, 0, 137, 0, 0,
	126, 0, 0, 86, 0, 0, 0, 231, 0, 215,
	0, 81, 82, 138, 139, 128, 129, 227, 0, 0,
	115, 0, 149, 177, 228, 229, 0, 0, 0, 116,
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
		//line parser/asn1p.y:286
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:293
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:297
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:304
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:308
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
		//line parser/asn1p.y:324
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:325
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:328
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:331
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:337
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:341
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:347
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:350
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:353
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:357
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:362
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:365
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:370
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:371
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:377
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:380
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:386
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:389
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:392
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:395
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:400
		{
			asn1VAL.module = nil
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:401
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:406
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
		//line parser/asn1p.y:425
		{
			asn1VAL.module = nil
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:429
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:435
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:442
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:446
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:450
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:457
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:458
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:461
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:471
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:475
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:482
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:487
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:491
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:499
		{
			asn1VAL.module = nil
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:500
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
		//line parser/asn1p.y:510
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:513
		{
			asn1VAL.xports = nil
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:516
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:522
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:527
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:535
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:540
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:545
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:554
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:557
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:567
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:571
		{
			fmt.Println("DDD")
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:576
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:583
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:593
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:598
		{
			asn1VAL.tag = nil
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:599
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:603
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:610
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:616
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:617
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:618
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:619
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:623
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:624
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:625
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 74:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:629
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:633
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 78:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:638
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrChoice
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:643
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSequence
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:648
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSet
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:654
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 82:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:660
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:668
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 84:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:673
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:678
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:683
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[4].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeBitString
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:692
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:693
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:694
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:695
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:696
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:697
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:698
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:699
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 95:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:700
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 96:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:701
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 97:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:702
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 100:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:708
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 101:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:709
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 102:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:710
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 103:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:719
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 106:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:740
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 107:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:744
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 108:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:748
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 113:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:760
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 114:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:769
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:772
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
		//line parser/asn1p.y:784
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
		//line parser/asn1p.y:803
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 118:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:809
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 119:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:814
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:819
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:822
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:825
		{
			asn1VAL.value = asn1types.NewAsn1Value()

		}
	case 123:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:830
		{
			asn1VAL.value = asn1types.NewAsn1Value()

		}
	case 124:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:837
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 125:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:843
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 126:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:846
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:853
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 128:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:860
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 129:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:867
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 130:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:874
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 131:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:881
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:901
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:921
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:924
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 135:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:927
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 136:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:933
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 137:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:936
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 138:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:942
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 139:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:948
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 144:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:961
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 145:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:964
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 146:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:972
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 148:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:977
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 149:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:981
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 151:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:989
		{
		}
	case 153:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:995
		{
		}
	case 155:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1001
		{
		}
	case 157:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1008
		{
		}
	case 159:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1014
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSet
		}
	case 160:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1021
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeValue
		}
	case 161:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1025
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSubType
			/* FIXME: */
		}
	case 165:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1037
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 166:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1040
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 167:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1046
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 168:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1049
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 169:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1052
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 170:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1058
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 171:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1064
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 172:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1070
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 173:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1082
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 174:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1083
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 175:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1086
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 176:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1089
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 177:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1092
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 178:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1098
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 179:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1102
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 180:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1105
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			//$$.Meta = $3.Meta
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeComponentsOf
		}
	case 181:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1110
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 182:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1120
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagNoMark
			asn1VAL.marker.Value = nil
		}
	case 183:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1125
		{
			asn1VAL.marker = asn1Dollar[1].marker
		}
	case 184:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1129
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect
			asn1VAL.marker.Value = nil
		}
	case 185:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1134
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagDefault
			asn1VAL.marker.Value = asn1Dollar[2].value
		}
	case 186:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1142
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBMPString
		}
	case 187:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1143
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralString
			fmt.Println("WARNING: GeneralString is not fully supported")
		}
	case 188:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1147
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGraphicsString
			fmt.Println("WARNING: GraphicString is not fully supported")
		}
	case 189:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1151
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeIA5String
		}
	case 190:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1152
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeISO646String
		}
	case 191:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1153
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNumericString
		}
	case 192:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1154
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypePrintableString
		}
	case 193:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1155
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeT61String
			fmt.Println("WARNING: T61String is not fully supported")
		}
	case 194:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1159
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeTeletexString
		}
	case 195:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1160
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUniversalString
		}
	case 196:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1161
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtf8String
		}
	case 197:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1162
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVideoTexString
			fmt.Println("WARNING: VideotexString is not fully supported")
		}
	case 198:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1166
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVisibleString
		}
	case 199:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1167
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectDescriptor
		}
	case 200:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1173
		{
			asn1VAL.constraint = nil
		}
	case 203:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1178
		{
			// FIXME : Implementation incomplete
		}
	case 204:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1184
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
			// FIXME: Needs implementation
		}
	case 207:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1195
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1Dollar[2].constraint_type
		}
	case 209:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1202
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMin
		}
	case 211:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1209
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMax
		}
	case 212:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1215
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeSimpleRange
		}
	case 213:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1216
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeRightExcludedRange
		}
	case 214:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1217
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftExcludedRange
		}
	case 215:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1218
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftRightExcludedRange
		}
	case 216:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1222
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 217:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1226
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 218:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1232
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// FIXME : Implementation
		}
	case 219:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1239
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			/* FIXME : implementation not yet complete */
		}
	case 220:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1244
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
		}
	case 221:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1249
		{
			asn1VAL.str = ""
		}
	case 222:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1250
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 223:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1256
		{
			asn1VAL.constraint = nil
		}
	case 226:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1262
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 227:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1265
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 228:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1271
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 229:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1277
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 230:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1285
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 231:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1288
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 232:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1293
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	}
	goto asn1stack /* stack new state and value */
}
