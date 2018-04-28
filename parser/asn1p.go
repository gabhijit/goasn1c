// FIXME: Ignoring to build this (old grammar)

// +build ignore

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

//line parser/asn1p.y:1713

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	25, 26,
	-2, 45,
	-1, 157,
	99, 175,
	-2, 65,
	-1, 162,
	99, 175,
	-2, 65,
	-1, 197,
	99, 184,
	103, 184,
	-2, 65,
	-1, 209,
	86, 210,
	110, 210,
	-2, 162,
	-1, 222,
	106, 17,
	-2, 136,
}

const asn1Private = 57344

const asn1Last = 577

var asn1Act = [...]int{

	190, 201, 319, 175, 219, 234, 96, 263, 245, 269,
	260, 249, 16, 209, 16, 77, 204, 206, 78, 235,
	76, 205, 196, 179, 308, 270, 152, 191, 239, 150,
	228, 272, 214, 10, 154, 183, 4, 233, 4, 199,
	294, 382, 353, 52, 302, 60, 167, 69, 256, 232,
	277, 255, 177, 289, 60, 285, 18, 277, 393, 203,
	275, 396, 394, 86, 295, 69, 274, 93, 81, 81,
	176, 370, 52, 325, 81, 371, 304, 326, 51, 75,
	59, 82, 68, 98, 276, 177, 218, 284, 69, 59,
	323, 276, 178, 238, 324, 215, 91, 238, 48, 81,
	68, 220, 312, 176, 299, 212, 313, 51, 300, 5,
	6, 184, 18, 185, 187, 188, 288, 267, 332, 400,
	161, 268, 146, 68, 81, 178, 275, 70, 71, 322,
	381, 316, 274, 87, 304, 336, 159, 159, 200, 158,
	81, 282, 221, 222, 376, 18, 185, 187, 188, 265,
	266, 186, 259, 208, 84, 81, 72, 229, 236, 398,
	232, 397, 359, 236, 232, 246, 250, 18, 161, 389,
	388, 18, 231, 237, 387, 386, 372, 173, 237, 81,
	194, 5, 6, 81, 18, 230, 349, 261, 223, 251,
	161, 298, 37, 317, 49, 226, 241, 257, 18, 162,
	242, 155, 314, 254, 155, 259, 402, 160, 160, 28,
	291, 315, 17, 18, 310, 303, 281, 148, 271, 278,
	26, 157, 253, 155, 279, 17, 18, 90, 378, 379,
	280, 265, 266, 14, 197, 198, 292, 11, 192, 165,
	164, 309, 163, 156, 309, 301, 297, 306, 151, 149,
	88, 73, 305, 296, 320, 18, 401, 321, 5, 6,
	391, 18, 5, 6, 184, 18, 252, 399, 311, 373,
	172, 35, 184, 18, 5, 6, 83, 18, 5, 6,
	335, 337, 338, 67, 354, 329, 63, 344, 145, 342,
	343, 334, 331, 144, 330, 5, 6, 328, 61, 94,
	345, 229, 29, 142, 346, 236, 356, 351, 348, 347,
	339, 341, 50, 33, 246, 360, 231, 250, 366, 32,
	237, 369, 364, 361, 358, 261, 357, 362, 367, 230,
	363, 355, 31, 377, 375, 374, 170, 168, 365, 166,
	251, 368, 5, 6, 65, 18, 169, 307, 243, 240,
	380, 277, 54, 46, 34, 171, 42, 143, 290, 286,
	236, 43, 12, 36, 264, 383, 24, 98, 98, 384,
	385, 147, 320, 141, 390, 237, 392, 107, 123, 110,
	139, 115, 100, 395, 89, 276, 22, 21, 25, 3,
	114, 138, 7, 106, 287, 283, 8, 23, 120, 1,
	203, 124, 117, 125, 126, 20, 74, 15, 30, 85,
	177, 105, 19, 127, 13, 9, 248, 111, 128, 118,
	113, 27, 247, 136, 262, 352, 177, 218, 176, 129,
	112, 119, 101, 102, 327, 58, 215, 130, 203, 340,
	333, 131, 220, 195, 176, 116, 133, 132, 134, 135,
	178, 92, 258, 318, 213, 216, 108, 109, 210, 350,
	217, 161, 293, 211, 177, 218, 178, 5, 6, 184,
	18, 185, 187, 188, 215, 225, 186, 224, 153, 200,
	220, 121, 176, 221, 222, 184, 18, 185, 187, 188,
	273, 227, 186, 207, 208, 202, 177, 218, 57, 161,
	244, 103, 189, 182, 178, 181, 215, 180, 174, 56,
	177, 140, 220, 137, 176, 80, 79, 122, 104, 97,
	99, 221, 222, 184, 18, 185, 187, 188, 176, 95,
	186, 161, 208, 55, 53, 47, 178, 41, 40, 193,
	66, 64, 62, 45, 44, 39, 38, 2, 0, 0,
	178, 0, 0, 221, 222, 184, 18, 185, 187, 188,
	0, 0, 186, 0, 208, 0, 0, 5, 6, 184,
	18, 185, 187, 188, 0, 0, 186,
}
var asn1Pact = [...]int{

	189, -1000, 189, -1000, -1000, -1000, -1000, -1000, 139, 340,
	-1000, 134, 358, 121, -1000, -1000, 109, -1000, -1000, 218,
	358, -1000, 261, 248, 242, 314, -1000, -1000, 180, 354,
	-1000, -1000, -1000, -1000, -1000, 91, 327, -1000, 336, -1000,
	312, -1000, 92, -1000, 206, -1000, 253, 25, 54, -1000,
	-1000, 153, -1000, 206, -1000, -1000, -1000, -1000, -1000, -5,
	-35, 186, 52, -1000, 185, -1000, 30, -1000, 152, -1000,
	-1000, 185, -1000, 128, -1000, 51, 215, -1000, 367, -1000,
	352, 296, 209, 204, -1000, -1000, 189, 185, 118, -1000,
	-1000, -1000, -1000, 151, 150, -1000, 104, -1000, -1000, -1000,
	145, 123, 101, -1000, -1000, 142, 141, 271, -60, -1000,
	-1000, -1000, -1000, 269, 288, 268, -1000, -1000, 317, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	179, -1000, -1000, -1000, 478, 140, 139, -1000, -1000, 138,
	-1000, 394, -1000, 104, -1000, 394, -36, 79, 295, -1000,
	-1000, 104, 79, 294, 163, 181, 124, 189, -1000, -1000,
	-1000, -1000, -54, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -58, -1000, -1000, 106, -1000, -1000, -1000,
	-1000, -1000, 135, -1000, -1000, 18, -1000, 36, -30, 117,
	-1000, 38, -20, 332, 8, -1000, 331, -1000, 432, -1000,
	-1000, -1000, -1000, -1000, -1000, -35, -1000, -46, 104, 104,
	-1000, -60, -1000, -1000, 90, -1000, -1000, 5, -1000, -35,
	-1000, -1000, -65, 116, 31, -1000, -35, 330, 293, -1000,
	163, -1000, 115, 163, 3, -1000, 102, 112, 28, -1000,
	93, -1000, -1000, 163, -1000, -1000, 163, 26, -9, -1000,
	-1000, -1000, -26, -1000, 20, -1000, -1000, -1000, 138, -1000,
	29, 330, 330, -1000, -1000, -1000, -1000, 478, 235, 29,
	330, -1000, 202, 464, -1000, -1000, 464, 464, -1000, -1000,
	464, 85, -1000, 378, -68, 198, -1000, -1000, -1000, -1000,
	-36, -1000, 173, -1000, 75, 330, -1000, -35, -35, -1000,
	-1000, -35, -1000, 163, 173, -1000, 181, 173, -28, -1000,
	76, -1000, 178, -1000, 163, -1000, 135, -1000, -1000, -1000,
	-1000, -1000, 53, -1000, -1000, -1000, 132, -1000, -1000, -1000,
	330, -1000, -1000, -1000, 27, 8, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -69, -1000, -1000, -1000, -1000, 79,
	-1000, -1000, 367, 367, -1000, 74, 73, -1000, 69, 68,
	-1000, 163, 169, -41, -1000, -1000, 26, -1000, -1000, -1000,
	-1000, 432, -1000, -27, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 60, 58, -1000, 176, -1000, -1000, -1000, -1000, 16,
	165, 107, -1000,
}
var asn1Pgo = [...]int{

	0, 547, 35, 0, 389, 546, 545, 544, 543, 542,
	541, 344, 540, 283, 539, 538, 537, 535, 312, 534,
	352, 533, 20, 15, 529, 520, 6, 519, 4, 518,
	517, 18, 516, 515, 513, 511, 509, 32, 508, 3,
	23, 507, 505, 503, 502, 501, 500, 8, 498, 29,
	39, 1, 495, 16, 21, 17, 493, 13, 491, 30,
	28, 37, 5, 19, 9, 490, 481, 26, 478, 34,
	477, 475, 463, 462, 460, 459, 458, 455, 105, 454,
	24, 139, 453, 2, 452, 10, 451, 443, 22, 31,
	440, 439, 25, 435, 27, 434, 7, 424, 422, 416,
	11, 415, 33, 414, 407, 412, 405, 387, 399, 396,
	395, 394, 364,
}
var asn1R1 = [...]int{

	0, 108, 1, 1, 109, 4, 101, 101, 102, 102,
	103, 103, 104, 104, 104, 3, 2, 2, 105, 105,
	106, 106, 107, 107, 107, 107, 5, 5, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 14, 14, 11,
	12, 12, 13, 13, 13, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 19, 20, 20, 20,
	20, 21, 21, 22, 23, 31, 31, 32, 33, 35,
	35, 35, 35, 34, 34, 34, 24, 26, 26, 27,
	27, 27, 27, 27, 27, 25, 25, 25, 25, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 30, 30, 30, 36, 37, 37, 38, 38,
	38, 38, 38, 38, 39, 39, 41, 41, 41, 40,
	44, 43, 42, 42, 42, 42, 98, 99, 99, 100,
	100, 100, 100, 100, 28, 45, 45, 45, 46, 46,
	47, 47, 110, 110, 111, 111, 49, 48, 50, 50,
	50, 50, 51, 51, 52, 52, 53, 53, 54, 54,
	55, 55, 56, 56, 56, 56, 56, 58, 58, 59,
	59, 59, 60, 60, 60, 61, 61, 62, 62, 62,
	63, 63, 63, 63, 64, 64, 65, 65, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 66, 67, 67, 68, 68, 69, 70, 71, 72,
	74, 74, 75, 75, 73, 73, 73, 73, 76, 76,
	77, 78, 79, 80, 80, 81, 81, 81, 82, 82,
	83, 83, 84, 84, 85, 86, 87, 87, 88, 88,
	88, 88, 88, 88, 88, 91, 91, 90, 92, 92,
	92, 89, 93, 94, 97, 97, 96, 95, 95, 95,
	95, 112, 112, 57,
}
var asn1R2 = [...]int{

	0, 1, 1, 2, 0, 9, 0, 1, 3, 2,
	1, 2, 1, 4, 1, 1, 1, 1, 0, 1,
	1, 2, 2, 2, 2, 2, 0, 1, 3, 0,
	1, 3, 2, 0, 1, 1, 2, 0, 1, 4,
	1, 3, 1, 3, 1, 0, 1, 3, 3, 2,
	1, 3, 1, 3, 1, 1, 2, 1, 1, 1,
	1, 3, 3, 1, 2, 0, 1, 2, 4, 0,
	1, 1, 1, 0, 1, 1, 2, 1, 1, 1,
	4, 4, 4, 6, 6, 1, 4, 4, 5, 1,
	1, 1, 2, 2, 2, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 2, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 5, 9, 1,
	1, 1, 1, 1, 3, 2, 1, 1, 3, 1,
	4, 4, 1, 1, 1, 1, 1, 3, 1, 3,
	4, 4, 1, 1, 1, 1, 3, 4, 1, 1,
	3, 5, 1, 3, 1, 3, 1, 3, 1, 3,
	1, 3, 1, 1, 1, 1, 1, 1, 3, 2,
	1, 1, 1, 3, 3, 0, 1, 1, 3, 5,
	3, 2, 3, 1, 0, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 1, 1, 2, 3, 1, 1, 3,
	1, 1, 1, 1, 1, 2, 2, 3, 2, 1,
	2, 2, 2, 0, 1, 0, 1, 1, 1, 3,
	4, 4, 1, 3, 1, 5, 1, 3, 2, 4,
	3, 3, 3, 3, 3, 0, 1, 0, 1, 3,
	3, 1, 4, 3, 1, 3, 2, 1, 1, 1,
	1, 1, 1, 1,
}
var asn1Chk = [...]int{

	-1000, -108, -1, -4, -2, 89, 90, -4, -109, -101,
	-102, 98, 22, -103, 99, -104, -3, 91, 92, -105,
	-106, -107, 28, 39, 8, 30, 99, -104, 100, 84,
	-107, 71, 71, 71, 40, 91, 9, 101, -5, -6,
	-15, -16, 29, 25, -7, -8, 41, -17, 6, 102,
	-18, -2, -3, -19, -20, -21, -36, -48, -93, -2,
	-3, 92, -9, 33, -10, -11, -12, -13, -2, -3,
	102, 103, 102, 98, -20, 84, -22, -23, -31, -32,
	-33, 104, -22, 90, 102, -11, 33, 103, 98, -18,
	99, -22, -86, 16, 84, -24, -26, -27, -28, -25,
	15, 65, 66, -45, -29, 44, 26, 10, 89, 90,
	12, 50, 63, 53, 23, 14, 78, 35, 52, 64,
	31, -66, -30, 11, 34, 36, 37, 46, 51, 62,
	70, 74, 80, 79, 81, 82, 56, -34, 39, 28,
	-35, 77, 7, 61, 84, 84, -2, -13, 99, 98,
	-49, 98, -67, -68, -69, 100, 98, 98, -81, -69,
	-78, 67, 98, -81, 98, 98, 68, 106, 68, 58,
	68, 38, 91, -37, -38, -39, 50, 32, 72, -40,
	-41, -42, -43, -2, 91, 93, 98, 94, 95, -44,
	-3, -94, 98, -14, -102, -87, -88, 96, 97, -50,
	85, -51, -52, 6, -53, -54, -55, -56, 100, -57,
	-76, -72, -78, -79, -37, 42, -77, -74, 33, -28,
	48, 89, 90, -69, -70, -71, -50, -58, -59, -3,
	-60, -23, 85, -61, -62, -63, -3, -23, 18, -60,
	54, -69, -61, 54, -46, -47, -3, -98, -99, -100,
	-3, -40, 85, 98, -2, 105, 106, 91, -84, 99,
	-85, -3, -97, -96, -112, 96, 97, 99, 103, -64,
	-92, -22, -89, -65, 96, 90, 55, 21, -22, -92,
	-89, 99, 103, -110, 107, 75, 27, -111, 108, 45,
	27, -51, -22, -73, 86, 110, -69, -67, 101, 99,
	103, -23, 109, 99, 103, -23, -64, 54, -80, -3,
	99, -80, 99, 103, 100, 99, 103, 100, -82, -83,
	-3, -3, 103, 99, 103, 99, 103, -95, -37, -22,
	-49, -94, 98, -90, -88, -64, 106, -64, -64, -37,
	-91, 76, -64, -64, 85, -53, -55, -54, -55, 101,
	-75, -57, 47, 110, 86, -59, -39, -40, -63, 87,
	-64, -23, -31, -31, -47, -40, -39, -100, -40, -39,
	99, 103, 100, 91, -85, -96, 91, -3, 96, 97,
	-64, 103, 110, -62, -26, -26, 101, 101, 101, 101,
	-83, 91, -39, 99, 103, -51, 88, 101, 101, 91,
	103, 91, 99,
}
var asn1Def = [...]int{

	0, -2, 1, 2, 4, 16, 17, 3, 6, 0,
	7, 0, 18, 0, 9, 10, 12, 14, 15, 0,
	19, 20, 0, 0, 0, 0, 8, 11, 0, 0,
	21, 22, 23, 24, 25, 0, -2, 13, 0, 27,
	29, 46, 0, 5, 0, 30, 33, 0, 0, 49,
	50, 52, 54, 28, 55, 57, 58, 59, 60, 65,
	65, 15, 0, 32, 34, 35, 0, 40, 42, 44,
	47, 0, 48, 0, 56, 65, 0, 63, 0, 66,
	73, 69, 0, 0, 31, 36, 0, 0, 0, 51,
	53, 61, 62, 0, 0, 64, 202, 77, 78, 79,
	0, 225, 225, 134, 85, 102, 103, 0, 135, 136,
	89, 90, 91, 0, 0, 0, 95, 96, 0, 98,
	99, 100, 101, 188, 189, 190, 191, 192, 193, 194,
	195, 196, 197, 198, 199, 200, 201, 67, 74, 75,
	0, 70, 71, 72, 0, 0, 37, 41, 43, 0,
	147, 0, 76, 203, 204, 0, 65, -2, 0, 226,
	227, 0, -2, 0, 0, 0, 104, 0, 92, 93,
	94, 97, 0, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 0, 119, 116, 0, 122, 123, 121,
	120, 252, 0, 39, 38, 0, 236, -2, 65, 0,
	148, 149, 152, 0, 154, 156, 158, 160, 0, -2,
	163, 164, 165, 166, 263, 65, 219, 0, 0, 202,
	211, 135, -2, 205, 0, 207, 208, 0, 167, 65,
	170, 171, 172, 0, 176, 177, 65, 184, 0, 183,
	223, 221, 0, 223, 0, 138, 0, 0, 126, 127,
	129, 132, 133, 0, 137, 68, 0, 0, 0, 125,
	232, 234, 0, 254, 65, 261, 262, 247, 0, 238,
	184, 184, 184, 185, 248, 251, 186, 0, 245, 184,
	184, 146, 0, 0, 142, 143, 0, 0, 144, 145,
	0, 0, 218, 0, 214, 0, 222, 220, 206, 80,
	65, 169, 0, 81, 65, 184, 181, 65, 65, 224,
	82, 65, 86, 0, 0, 87, 0, 0, 0, 228,
	0, 115, 0, 124, 0, 253, 0, 256, 257, 258,
	259, 260, 0, 235, 237, 242, 0, 243, 244, 187,
	184, 246, 240, 241, 150, 155, 153, 157, 159, 161,
	209, 212, 213, 215, 216, 168, 173, 174, 178, 65,
	180, 182, 0, 0, 139, 0, 0, 128, 0, 0,
	88, 0, 0, 0, 233, 255, 119, 120, 249, 250,
	239, 0, 217, 0, 83, 84, 140, 141, 130, 131,
	229, 0, 0, 117, 0, 151, 179, 230, 231, 0,
	0, 0, 118,
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
	asn1ErrorVerbose = true
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
		//line parser/asn1p.y:304
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:311
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:315
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:322
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:326
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
		//line parser/asn1p.y:342
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:343
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:346
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:349
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:355
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:359
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:365
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:368
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:371
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:375
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:380
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:383
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:388
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:389
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:395
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:398
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:404
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:407
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:410
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:413
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:418
		{
			asn1VAL.module = nil
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:419
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:424
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
		//line parser/asn1p.y:443
		{
			asn1VAL.module = nil
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:447
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:453
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:460
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:464
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:468
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:475
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:476
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:479
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:489
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:493
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:500
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:505
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:509
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:517
		{
			asn1VAL.module = nil
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:518
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
		//line parser/asn1p.y:528
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:531
		{
			asn1VAL.xports = nil
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:534
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:540
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:545
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:553
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:558
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:563
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:572
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:575
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:585
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:589
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:593
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:597
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 61:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:604
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:609
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:619
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:624
		{
			asn1VAL.tag = nil
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:625
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:629
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:636
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:642
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:643
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:644
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:645
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:649
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 74:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:650
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:651
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 76:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:655
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 77:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:659
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:664
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrChoice
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:669
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSequence
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 82:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:674
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSet
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:680
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 84:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:686
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:694
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 86:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:699
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:704
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:709
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[4].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeBitString
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:718
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:719
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:720
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:721
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:722
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:723
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 95:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:724
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 96:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:725
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 97:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:726
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 98:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:727
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 99:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:728
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 102:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:734
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 103:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:735
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 104:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:736
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 105:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:745
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 108:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:766
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 109:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:770
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 110:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:774
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:786
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 116:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:795
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 117:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:798
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
	case 118:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:810
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
	case 119:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:829
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:835
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:840
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:845
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 123:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:848
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 124:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:851
		{
			asn1VAL.value = asn1types.NewAsn1Value()

		}
	case 125:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:856
		{
			asn1VAL.value = asn1types.NewAsn1Value()

		}
	case 126:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:863
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:869
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 128:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:872
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 129:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:879
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 130:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:886
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 131:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:893
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:900
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:907
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:927
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 135:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:947
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 136:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:950
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 137:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:953
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 138:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:959
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 139:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:962
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 140:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:968
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 141:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:974
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 146:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:987
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 147:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:990
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 148:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:998
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 150:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1003
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 151:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1007
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 153:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1015
		{
		}
	case 155:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1021
		{
		}
	case 157:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1027
		{
		}
	case 159:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1034
		{
		}
	case 161:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1040
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSet
		}
	case 162:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1047
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeValue
		}
	case 163:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1051
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSubType
			/* FIXME: */
		}
	case 167:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1063
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 168:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1066
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 169:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1072
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 170:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1075
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 171:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1078
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 172:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1084
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 173:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1090
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 174:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1096
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 175:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1108
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 176:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1109
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 177:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1112
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 178:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1115
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 179:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1118
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 180:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1124
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 181:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1128
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 182:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1131
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			//$$.Meta = $3.Meta
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeComponentsOf
		}
	case 183:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1136
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 184:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1146
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagNoMark
			asn1VAL.marker.Value = nil
		}
	case 185:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1151
		{
			asn1VAL.marker = asn1Dollar[1].marker
		}
	case 186:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1155
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect
			asn1VAL.marker.Value = nil
		}
	case 187:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1160
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagDefault
			asn1VAL.marker.Value = asn1Dollar[2].value
		}
	case 188:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1168
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBMPString
		}
	case 189:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1169
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralString
			fmt.Println("WARNING: GeneralString is not fully supported")
		}
	case 190:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1173
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGraphicsString
			fmt.Println("WARNING: GraphicString is not fully supported")
		}
	case 191:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1177
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeIA5String
		}
	case 192:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1178
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeISO646String
		}
	case 193:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1179
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNumericString
		}
	case 194:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1180
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypePrintableString
		}
	case 195:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1181
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeT61String
			fmt.Println("WARNING: T61String is not fully supported")
		}
	case 196:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1185
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeTeletexString
		}
	case 197:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1186
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUniversalString
		}
	case 198:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1187
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtf8String
		}
	case 199:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1188
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVideoTexString
			fmt.Println("WARNING: VideotexString is not fully supported")
		}
	case 200:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1192
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVisibleString
		}
	case 201:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1193
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectDescriptor
		}
	case 202:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1199
		{
			asn1VAL.constraint = nil
		}
	case 205:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1204
		{
			// FIXME : Implementation incomplete
		}
	case 206:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1210
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
			// FIXME: Needs implementation
		}
	case 209:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1221
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1Dollar[2].constraint_type
		}
	case 211:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1228
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMin
		}
	case 213:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1235
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMax
		}
	case 214:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1241
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeSimpleRange
		}
	case 215:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1242
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeRightExcludedRange
		}
	case 216:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1243
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftExcludedRange
		}
	case 217:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1244
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftRightExcludedRange
		}
	case 218:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1248
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 219:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1252
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 220:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1258
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// FIXME : Implementation
		}
	case 221:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1265
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			/* FIXME : implementation not yet complete */
		}
	case 222:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1270
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
		}
	case 223:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1275
		{
			asn1VAL.str = ""
		}
	case 224:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1276
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 225:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1282
		{
			asn1VAL.constraint = nil
		}
	case 228:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1288
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 229:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1291
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 230:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1297
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 231:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1303
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 232:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1311
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 233:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1314
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 234:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1319
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 235:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1327
		{
			asn1VAL.expr = asn1Dollar[3].expr
		}
	case 236:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1333
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectClass
		}
	case 237:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1338
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 238:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1348
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 239:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1356
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 240:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1364
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 241:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1372
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 242:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1380
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 243:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1388
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 244:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1396
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeClassDef
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObjectField
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 245:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1405
		{
			asn1VAL.num = 0
		}
	case 246:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1406
		{
			asn1VAL.num = 1
		}
	case 247:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1409
		{
			asn1VAL.with_syntax = nil
		}
	case 248:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1413
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 249:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1416
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 250:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1419
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 251:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1425
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 252:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:1431
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeObjectIdentifier
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeObject
		}
	case 253:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1438
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 254:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1444
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()

		}
	case 255:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1448
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 256:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1455
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 257:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1462
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 259:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1466
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	}
	goto asn1stack /* stack new state and value */
}
