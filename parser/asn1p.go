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
	with_syntax     *asn1types.Asn1WithSyntax
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
const Tok_TypeFieldReference = 57438
const Tok_ValueFieldReference = 57439

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
	"Tok_TypeFieldReference",
	"Tok_ValueFieldReference",
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

//line parser/asn1p.y:1656

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	25, 26,
	-2, 45,
	-1, 153,
	99, 174,
	-2, 64,
	-1, 158,
	99, 174,
	-2, 64,
	-1, 191,
	99, 183,
	103, 183,
	-2, 64,
	-1, 203,
	86, 209,
	110, 209,
	-2, 161,
	-1, 216,
	106, 17,
	-2, 135,
}

const asn1Private = 57344

const asn1Last = 511

var asn1Act = [...]int{

	186, 195, 308, 171, 213, 228, 93, 254, 239, 258,
	243, 233, 16, 203, 16, 75, 175, 76, 198, 229,
	199, 208, 200, 190, 297, 259, 148, 222, 261, 150,
	227, 283, 193, 74, 10, 179, 4, 360, 4, 334,
	278, 291, 226, 52, 232, 59, 163, 67, 197, 18,
	264, 274, 250, 266, 59, 284, 263, 249, 374, 73,
	79, 79, 90, 67, 79, 371, 378, 351, 312, 372,
	52, 352, 313, 293, 173, 212, 68, 69, 51, 79,
	58, 95, 66, 273, 209, 67, 173, 265, 83, 58,
	214, 301, 172, 80, 359, 302, 288, 311, 66, 232,
	289, 333, 256, 277, 172, 51, 257, 88, 305, 157,
	48, 226, 293, 340, 174, 271, 81, 266, 18, 142,
	66, 70, 376, 375, 157, 367, 174, 194, 155, 155,
	79, 215, 216, 180, 18, 181, 183, 184, 317, 157,
	182, 366, 202, 5, 6, 180, 18, 181, 183, 184,
	79, 265, 182, 223, 230, 158, 353, 151, 84, 230,
	365, 240, 244, 169, 224, 364, 226, 330, 225, 231,
	153, 197, 151, 18, 231, 287, 37, 188, 245, 217,
	251, 18, 306, 255, 220, 79, 264, 235, 253, 236,
	303, 151, 263, 5, 6, 28, 18, 173, 212, 248,
	79, 17, 18, 247, 280, 380, 49, 209, 304, 26,
	299, 17, 18, 214, 292, 172, 270, 144, 268, 14,
	87, 269, 356, 357, 154, 260, 267, 206, 191, 192,
	61, 11, 157, 161, 160, 298, 152, 174, 298, 290,
	286, 295, 285, 281, 147, 145, 294, 85, 309, 71,
	18, 310, 379, 377, 215, 216, 180, 18, 181, 183,
	184, 246, 300, 182, 354, 202, 168, 180, 18, 316,
	318, 319, 5, 6, 369, 18, 5, 6, 323, 324,
	35, 315, 5, 6, 180, 18, 5, 6, 320, 18,
	223, 326, 335, 325, 230, 337, 332, 328, 327, 141,
	91, 224, 329, 240, 341, 225, 244, 347, 338, 231,
	350, 345, 342, 339, 255, 343, 348, 336, 344, 29,
	346, 355, 245, 349, 159, 322, 156, 156, 33, 5,
	6, 358, 18, 139, 65, 32, 31, 50, 166, 164,
	162, 230, 165, 266, 296, 63, 361, 237, 95, 95,
	362, 363, 234, 309, 46, 368, 231, 370, 104, 120,
	107, 373, 112, 97, 136, 34, 54, 167, 42, 279,
	275, 111, 43, 12, 103, 135, 21, 265, 36, 117,
	276, 15, 121, 114, 122, 123, 3, 140, 272, 7,
	8, 173, 102, 1, 124, 27, 20, 30, 108, 125,
	115, 110, 19, 138, 133, 173, 212, 86, 82, 172,
	126, 109, 116, 98, 99, 209, 13, 9, 127, 143,
	72, 214, 128, 172, 24, 242, 113, 130, 129, 131,
	132, 174, 241, 321, 314, 189, 89, 105, 106, 252,
	157, 307, 207, 210, 22, 174, 25, 204, 5, 6,
	180, 18, 181, 183, 184, 23, 331, 182, 211, 282,
	205, 219, 215, 216, 180, 18, 181, 183, 184, 218,
	149, 182, 118, 202, 262, 221, 201, 196, 146, 57,
	238, 100, 185, 178, 177, 176, 170, 56, 137, 134,
	78, 77, 119, 101, 94, 96, 92, 55, 53, 47,
	41, 40, 187, 64, 62, 60, 45, 44, 39, 38,
	2,
}
var asn1Pact = [...]int{

	187, -1000, 187, -1000, -1000, -1000, -1000, -1000, 133, 351,
	-1000, 120, 416, 110, -1000, -1000, 95, -1000, -1000, 235,
	416, -1000, 265, 264, 257, 325, -1000, -1000, 189, 369,
	-1000, -1000, -1000, -1000, -1000, 75, 339, -1000, 347, -1000,
	313, -1000, 104, -1000, 240, -1000, 197, -26, 19, -1000,
	-1000, 151, -1000, 240, -1000, -1000, -1000, -1000, -25, -44,
	14, -1000, 240, -1000, 55, -1000, 149, -1000, -1000, 240,
	-1000, 121, -1000, 46, 216, -1000, 348, -1000, 336, 326,
	215, -1000, -1000, 187, 240, 118, -1000, -1000, -1000, -1000,
	147, 146, -1000, 91, -1000, -1000, -1000, 138, 72, 57,
	-1000, -1000, 136, 135, 272, -60, -1000, -1000, -1000, -1000,
	271, 284, 270, -1000, -1000, 329, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 175, -1000, -1000,
	-1000, 359, 133, -1000, -1000, 132, -1000, 42, -1000, 91,
	-1000, 42, -43, 81, 298, -1000, -1000, 91, 81, 293,
	158, 176, 105, 187, -1000, -1000, -1000, -1000, -48, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -54,
	-1000, -1000, 89, -1000, -1000, -1000, -1000, -1000, -1000, 3,
	-1000, 96, -40, 117, -1000, 12, -24, 343, -5, -1000,
	342, -1000, 165, -1000, -1000, -1000, -1000, -1000, -1000, -44,
	-1000, -55, 91, 91, -1000, -60, -1000, -1000, 74, -1000,
	-1000, -3, -1000, -44, -1000, -1000, -68, 115, 9, -1000,
	-44, 322, 290, -1000, 158, -1000, 111, 158, -8, -1000,
	90, 109, 5, -1000, 82, -1000, -1000, 158, -1000, -1000,
	158, -6, -31, -1000, -1000, -1000, -1000, 132, -1000, 32,
	322, 322, -1000, -1000, -1000, -1000, 359, 249, 32, 322,
	-1000, 208, 373, -1000, -1000, 373, 373, -1000, -1000, 373,
	66, -1000, 54, -71, 206, -1000, -1000, -1000, -1000, -43,
	-1000, 193, -1000, 26, 322, -1000, -44, -44, -1000, -1000,
	-44, -1000, 158, 193, -1000, 176, 193, -32, -1000, 56,
	-1000, 173, -1000, 158, -1000, -1000, -1000, 126, -1000, -1000,
	-1000, 322, -1000, -1000, -1000, -9, -5, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -73, -1000, -1000, -1000, -1000,
	81, -1000, -1000, 348, 348, -1000, 64, 59, -1000, 40,
	24, -1000, 158, 183, -34, -1000, -1000, -1000, -1000, 165,
	-1000, -30, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 22,
	21, -1000, 162, -1000, -1000, -1000, -1000, -37, 161, 106,
	-1000,
}
var asn1Pgo = [...]int{

	0, 510, 35, 0, 386, 509, 508, 507, 506, 505,
	504, 345, 503, 334, 502, 501, 500, 499, 337, 498,
	366, 497, 33, 15, 496, 495, 6, 494, 4, 493,
	492, 17, 491, 490, 489, 488, 487, 21, 486, 3,
	16, 485, 484, 483, 482, 481, 480, 8, 479, 478,
	32, 1, 477, 18, 20, 22, 476, 13, 475, 27,
	11, 30, 5, 19, 9, 474, 472, 26, 470, 29,
	469, 461, 460, 459, 458, 456, 447, 443, 227, 442,
	24, 224, 441, 2, 439, 7, 436, 435, 23, 28,
	434, 433, 25, 432, 425, 10, 417, 34, 416, 381,
	402, 396, 376, 393, 390, 388, 380,
}
var asn1R1 = [...]int{

	0, 103, 1, 1, 104, 4, 96, 96, 97, 97,
	98, 98, 99, 99, 99, 3, 2, 2, 100, 100,
	101, 101, 102, 102, 102, 102, 5, 5, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 14, 14, 11,
	12, 12, 13, 13, 13, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 19, 20, 20, 20,
	21, 21, 22, 23, 31, 31, 32, 33, 35, 35,
	35, 35, 34, 34, 34, 24, 26, 26, 27, 27,
	27, 27, 27, 27, 25, 25, 25, 25, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 30, 30, 30, 36, 37, 37, 38, 38, 38,
	38, 38, 38, 39, 39, 41, 41, 41, 40, 44,
	43, 42, 42, 42, 42, 93, 94, 94, 95, 95,
	95, 95, 95, 28, 45, 45, 45, 46, 46, 47,
	47, 105, 105, 106, 106, 49, 48, 50, 50, 50,
	50, 51, 51, 52, 52, 53, 53, 54, 54, 55,
	55, 56, 56, 56, 56, 56, 58, 58, 59, 59,
	59, 60, 60, 60, 61, 61, 62, 62, 62, 63,
	63, 63, 63, 64, 64, 65, 65, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 67, 67, 68, 68, 69, 70, 71, 72, 74,
	74, 75, 75, 73, 73, 73, 73, 76, 76, 77,
	78, 79, 80, 80, 81, 81, 81, 82, 82, 83,
	83, 84, 84, 85, 86, 87, 87, 88, 88, 88,
	88, 88, 88, 88, 91, 91, 90, 92, 92, 92,
	89, 57,
}
var asn1R2 = [...]int{

	0, 1, 1, 2, 0, 9, 0, 1, 3, 2,
	1, 2, 1, 4, 1, 1, 1, 1, 0, 1,
	1, 2, 2, 2, 2, 2, 0, 1, 3, 0,
	1, 3, 2, 0, 1, 1, 2, 0, 1, 4,
	1, 3, 1, 3, 1, 0, 1, 3, 3, 2,
	1, 3, 1, 3, 1, 1, 2, 1, 1, 1,
	3, 3, 1, 2, 0, 1, 2, 4, 0, 1,
	1, 1, 0, 1, 1, 2, 1, 1, 1, 4,
	4, 4, 6, 6, 1, 4, 4, 5, 1, 1,
	1, 2, 2, 2, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 2, 4, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 5, 9, 1, 1,
	1, 1, 1, 3, 2, 1, 1, 3, 1, 4,
	4, 1, 1, 1, 1, 1, 3, 1, 3, 4,
	4, 1, 1, 1, 1, 3, 4, 1, 1, 3,
	5, 1, 3, 1, 3, 1, 3, 1, 3, 1,
	3, 1, 1, 1, 1, 1, 1, 3, 2, 1,
	1, 1, 3, 3, 0, 1, 1, 3, 5, 3,
	2, 3, 1, 0, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 0, 1, 1, 2, 3, 1, 1, 3, 1,
	1, 1, 1, 1, 2, 2, 3, 2, 1, 2,
	2, 2, 0, 1, 0, 1, 1, 1, 3, 4,
	4, 1, 3, 1, 5, 1, 3, 2, 4, 3,
	3, 3, 3, 3, 0, 1, 0, 1, 3, 3,
	1, 1,
}
var asn1Chk = [...]int{

	-1000, -103, -1, -4, -2, 89, 90, -4, -104, -96,
	-97, 98, 22, -98, 99, -99, -3, 91, 92, -100,
	-101, -102, 28, 39, 8, 30, 99, -99, 100, 84,
	-102, 71, 71, 71, 40, 91, 9, 101, -5, -6,
	-15, -16, 29, 25, -7, -8, 41, -17, 6, 102,
	-18, -2, -3, -19, -20, -21, -36, -48, -2, -3,
	-9, 33, -10, -11, -12, -13, -2, -3, 102, 103,
	102, 98, -20, 84, -22, -23, -31, -32, -33, 104,
	-22, 102, -11, 33, 103, 98, -18, 99, -22, -86,
	16, 84, -24, -26, -27, -28, -25, 15, 65, 66,
	-45, -29, 44, 26, 10, 89, 90, 12, 50, 63,
	53, 23, 14, 78, 35, 52, 64, 31, -66, -30,
	11, 34, 36, 37, 46, 51, 62, 70, 74, 80,
	79, 81, 82, 56, -34, 39, 28, -35, 77, 7,
	61, 84, -2, -13, 99, 98, -49, 98, -67, -68,
	-69, 100, 98, 98, -81, -69, -78, 67, 98, -81,
	98, 98, 68, 106, 68, 58, 68, 38, 91, -37,
	-38, -39, 50, 32, 72, -40, -41, -42, -43, -2,
	91, 93, 98, 94, 95, -44, -3, -14, -97, -87,
	-88, 96, 97, -50, 85, -51, -52, 6, -53, -54,
	-55, -56, 100, -57, -76, -72, -78, -79, -37, 42,
	-77, -74, 33, -28, 48, 89, 90, -69, -70, -71,
	-50, -58, -59, -3, -60, -23, 85, -61, -62, -63,
	-3, -23, 18, -60, 54, -69, -61, 54, -46, -47,
	-3, -93, -94, -95, -3, -40, 85, 98, -2, 105,
	106, 91, -84, 99, -85, -3, 99, 103, -64, -92,
	-22, -89, -65, 96, 90, 55, 21, -22, -92, -89,
	99, 103, -105, 107, 75, 27, -106, 108, 45, 27,
	-51, -22, -73, 86, 110, -69, -67, 101, 99, 103,
	-23, 109, 99, 103, -23, -64, 54, -80, -3, 99,
	-80, 99, 103, 100, 99, 103, 100, -82, -83, -3,
	-3, 103, 99, 103, -90, -88, -64, 106, -64, -64,
	-37, -91, 76, -64, -64, 85, -53, -55, -54, -55,
	101, -75, -57, 47, 110, 86, -59, -39, -40, -63,
	87, -64, -23, -31, -31, -47, -40, -39, -95, -40,
	-39, 99, 103, 100, 91, -85, 96, 97, -64, 103,
	110, -62, -26, -26, 101, 101, 101, 101, -83, 91,
	-39, 99, 103, -51, 88, 101, 101, 91, 103, 91,
	99,
}
var asn1Def = [...]int{

	0, -2, 1, 2, 4, 16, 17, 3, 6, 0,
	7, 0, 18, 0, 9, 10, 12, 14, 15, 0,
	19, 20, 0, 0, 0, 0, 8, 11, 0, 0,
	21, 22, 23, 24, 25, 0, -2, 13, 0, 27,
	29, 46, 0, 5, 0, 30, 33, 0, 0, 49,
	50, 52, 54, 28, 55, 57, 58, 59, 64, 64,
	0, 32, 34, 35, 0, 40, 42, 44, 47, 0,
	48, 0, 56, 64, 0, 62, 0, 65, 72, 68,
	0, 31, 36, 0, 0, 0, 51, 53, 60, 61,
	0, 0, 63, 201, 76, 77, 78, 0, 224, 224,
	133, 84, 101, 102, 0, 134, 135, 88, 89, 90,
	0, 0, 0, 94, 95, 0, 97, 98, 99, 100,
	187, 188, 189, 190, 191, 192, 193, 194, 195, 196,
	197, 198, 199, 200, 66, 73, 74, 0, 69, 70,
	71, 0, 37, 41, 43, 0, 146, 0, 75, 202,
	203, 0, 64, -2, 0, 225, 226, 0, -2, 0,
	0, 0, 103, 0, 91, 92, 93, 96, 0, 104,
	105, 106, 107, 108, 109, 110, 111, 112, 113, 0,
	118, 115, 0, 121, 122, 120, 119, 39, 38, 0,
	235, -2, 64, 0, 147, 148, 151, 0, 153, 155,
	157, 159, 0, -2, 162, 163, 164, 165, 251, 64,
	218, 0, 0, 201, 210, 134, -2, 204, 0, 206,
	207, 0, 166, 64, 169, 170, 171, 0, 175, 176,
	64, 183, 0, 182, 222, 220, 0, 222, 0, 137,
	0, 0, 125, 126, 128, 131, 132, 0, 136, 67,
	0, 0, 0, 124, 231, 233, 246, 0, 237, 183,
	183, 183, 184, 247, 250, 185, 0, 244, 183, 183,
	145, 0, 0, 141, 142, 0, 0, 143, 144, 0,
	0, 217, 0, 213, 0, 221, 219, 205, 79, 64,
	168, 0, 80, 64, 183, 180, 64, 64, 223, 81,
	64, 85, 0, 0, 86, 0, 0, 0, 227, 0,
	114, 0, 123, 0, 234, 236, 241, 0, 242, 243,
	186, 183, 245, 239, 240, 149, 154, 152, 156, 158,
	160, 208, 211, 212, 214, 215, 167, 172, 173, 177,
	64, 179, 181, 0, 0, 138, 0, 0, 127, 0,
	0, 87, 0, 0, 0, 232, 248, 249, 238, 0,
	216, 0, 82, 83, 139, 140, 129, 130, 228, 0,
	0, 116, 0, 150, 178, 229, 230, 0, 0, 0,
	117,
}
var asn1Tok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 109, 3, 3, 3, 3, 3, 3,
	100, 101, 3, 3, 103, 3, 106, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 102,
	110, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 104, 3, 105, 108, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 98, 107, 99,
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
	92, 93, 94, 95, 96, 97,
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
		//line parser/asn1p.y:297
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:304
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:308
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:315
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:319
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
		//line parser/asn1p.y:335
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:336
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:339
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:342
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:348
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:352
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:358
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:361
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:364
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:368
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:373
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:376
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:381
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:382
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:388
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:391
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:397
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:400
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:403
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:406
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:411
		{
			asn1VAL.module = nil
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:412
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:417
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
		//line parser/asn1p.y:436
		{
			asn1VAL.module = nil
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:440
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:446
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:453
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:457
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:461
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:468
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:469
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:472
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:482
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:486
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:493
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:498
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:502
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:510
		{
			asn1VAL.module = nil
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:511
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
		//line parser/asn1p.y:521
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:524
		{
			asn1VAL.xports = nil
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:527
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:533
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:538
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:546
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:551
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:556
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:565
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:568
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:578
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:582
		{
			fmt.Println("DDD")
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:587
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:594
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 61:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:599
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:609
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:614
		{
			asn1VAL.tag = nil
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:615
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:619
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:626
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:632
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:633
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:634
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:635
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:639
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:640
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 74:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:641
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:645
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 76:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:649
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:654
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrChoice
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:659
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSequence
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:664
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSet
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 82:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:670
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:676
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 84:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:684
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 85:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:689
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:694
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:699
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[4].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeBitString
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:708
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:709
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:710
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:711
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:712
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:713
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:714
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 95:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:715
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 96:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:716
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 97:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:717
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 98:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:718
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 101:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:724
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 102:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:725
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 103:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:726
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 104:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:735
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 107:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:756
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 108:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:760
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 109:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:764
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 114:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:776
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:785
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 116:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:788
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
	case 117:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:800
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
	case 118:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:819
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 119:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:825
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:830
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:835
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:838
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 123:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:841
		{
			asn1VAL.value = asn1types.NewAsn1Value()

		}
	case 124:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:846
		{
			asn1VAL.value = asn1types.NewAsn1Value()

		}
	case 125:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:853
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 126:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:859
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:862
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 128:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:869
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 129:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:876
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 130:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:883
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 131:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:890
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:897
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:917
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:937
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 135:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:940
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 136:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:943
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 137:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:949
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 138:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:952
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 139:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:958
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 140:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:964
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 145:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:977
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 146:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:980
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 147:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:988
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 149:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:993
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 150:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:997
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 152:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1005
		{
		}
	case 154:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1011
		{
		}
	case 156:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1017
		{
		}
	case 158:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1024
		{
		}
	case 160:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1030
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSet
		}
	case 161:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1037
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeValue
		}
	case 162:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1041
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSubType
			/* FIXME: */
		}
	case 166:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1053
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 167:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1056
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 168:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1062
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 169:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1065
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 170:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1068
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 171:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1074
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 172:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1080
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 173:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1086
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 174:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1098
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 175:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1099
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 176:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1102
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 177:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1105
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 178:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1108
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 179:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1114
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 180:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1118
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 181:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1121
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			//$$.Meta = $3.Meta
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeComponentsOf
		}
	case 182:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1126
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 183:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1136
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagNoMark
			asn1VAL.marker.Value = nil
		}
	case 184:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1141
		{
			asn1VAL.marker = asn1Dollar[1].marker
		}
	case 185:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1145
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect
			asn1VAL.marker.Value = nil
		}
	case 186:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1150
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagDefault
			asn1VAL.marker.Value = asn1Dollar[2].value
		}
	case 187:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1158
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBMPString
		}
	case 188:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1159
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralString
			fmt.Println("WARNING: GeneralString is not fully supported")
		}
	case 189:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1163
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGraphicsString
			fmt.Println("WARNING: GraphicString is not fully supported")
		}
	case 190:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1167
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeIA5String
		}
	case 191:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1168
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeISO646String
		}
	case 192:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1169
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNumericString
		}
	case 193:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1170
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypePrintableString
		}
	case 194:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1171
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeT61String
			fmt.Println("WARNING: T61String is not fully supported")
		}
	case 195:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1175
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeTeletexString
		}
	case 196:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1176
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUniversalString
		}
	case 197:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1177
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtf8String
		}
	case 198:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1178
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVideoTexString
			fmt.Println("WARNING: VideotexString is not fully supported")
		}
	case 199:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1182
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVisibleString
		}
	case 200:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1183
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectDescriptor
		}
	case 201:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1189
		{
			asn1VAL.constraint = nil
		}
	case 204:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1194
		{
			// FIXME : Implementation incomplete
		}
	case 205:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1200
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
			// FIXME: Needs implementation
		}
	case 208:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1211
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1Dollar[2].constraint_type
		}
	case 210:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1218
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMin
		}
	case 212:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1225
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMax
		}
	case 213:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1231
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeSimpleRange
		}
	case 214:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1232
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeRightExcludedRange
		}
	case 215:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1233
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftExcludedRange
		}
	case 216:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1234
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftRightExcludedRange
		}
	case 217:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1238
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 218:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1242
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 219:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1248
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// FIXME : Implementation
		}
	case 220:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1255
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			/* FIXME : implementation not yet complete */
		}
	case 221:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1260
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
		}
	case 222:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1265
		{
			asn1VAL.str = ""
		}
	case 223:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1266
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 224:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1272
		{
			asn1VAL.constraint = nil
		}
	case 227:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1278
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 228:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1281
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 229:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1287
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 230:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1293
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 231:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1301
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 232:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1304
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 233:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1309
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 234:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1317
		{
			asn1VAL.expr = asn1Dollar[3].expr
		}
	case 235:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1323
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectClass
		}
	case 236:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1328
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 237:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1338
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 238:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1346
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 239:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1354
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 240:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1362
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 241:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1370
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 242:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1378
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 243:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1386
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 244:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1395
		{
			asn1VAL.num = 0
		}
	case 245:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1396
		{
			asn1VAL.num = 1
		}
	case 246:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1399
		{
			asn1VAL.with_syntax = nil
		}
	case 247:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1403
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 248:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1406
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 249:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1409
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 250:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1415
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	}
	goto asn1stack /* stack new state and value */
}
