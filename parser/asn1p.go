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
const Tok_INCLUDES = 57373
const Tok_INTEGER = 57374
const Tok_INTERSECTION = 57375
const Tok_MAX = 57376
const Tok_MIN = 57377
const Tok_NULL = 57378
const Tok_OF = 57379
const Tok_OBJECT = 57380
const Tok_OCTET = 57381
const Tok_OPTIONAL = 57382
const Tok_PDV = 57383
const Tok_PRIVATE = 57384
const Tok_REAL = 57385
const Tok_RELATIVE_OID = 57386
const Tok_SEQUENCE = 57387
const Tok_SET = 57388
const Tok_SIZE = 57389
const Tok_STRING = 57390
const Tok_TAGS = 57391
const Tok_TRUE = 57392
const Tok_UNION = 57393
const Tok_UNIVERSAL = 57394
const Tok_UTCTime = 57395
const Tok_BMPString = 57396
const Tok_GeneralString = 57397
const Tok_GraphicString = 57398
const Tok_IA5String = 57399
const Tok_ISO646String = 57400
const Tok_NumericString = 57401
const Tok_PrintableString = 57402
const Tok_T61String = 57403
const Tok_TeletexString = 57404
const Tok_UniversalString = 57405
const Tok_UTF8String = 57406
const Tok_VideotexString = 57407
const Tok_VisibleString = 57408
const Tok_ObjectDescriptor = 57409
const Tok_Ellipsis = 57410
const Tok_TwoDots = 57411
const Tok_TwoLeftBrackets = 57412
const Tok_TwoRightBrackets = 57413
const Tok_TypeReference = 57414
const Tok_CAPITALREFERENCE = 57415
const Tok_Number = 57416
const Tok_Identifier = 57417
const Tok_CString = 57418
const Tok_BString = 57419
const Tok_HString = 57420

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
	"Tok_INCLUDES",
	"Tok_INTEGER",
	"Tok_INTERSECTION",
	"Tok_MAX",
	"Tok_MIN",
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
	"Tok_SIZE",
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

//line parser/asn1p.y:1477

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	17, 26,
	-2, 45,
	-1, 150,
	80, 170,
	-2, 63,
	-1, 155,
	80, 170,
	-2, 63,
	-1, 196,
	69, 205,
	91, 205,
	-2, 157,
	-1, 209,
	87, 17,
	-2, 131,
}

const asn1Private = 57344

const asn1Last = 446

var asn1Act = [...]int{

	183, 188, 206, 221, 91, 172, 168, 76, 236, 222,
	75, 215, 16, 232, 16, 226, 196, 274, 192, 191,
	147, 74, 201, 269, 220, 145, 176, 4, 186, 4,
	10, 199, 315, 151, 257, 295, 252, 265, 248, 159,
	242, 219, 241, 52, 193, 59, 326, 67, 18, 79,
	328, 323, 170, 225, 59, 324, 258, 278, 79, 267,
	83, 279, 294, 67, 169, 314, 190, 65, 262, 51,
	52, 58, 263, 66, 285, 247, 48, 282, 171, 93,
	58, 80, 68, 69, 267, 67, 170, 205, 245, 66,
	73, 225, 251, 202, 81, 88, 51, 207, 169, 70,
	5, 6, 177, 18, 178, 180, 181, 179, 219, 154,
	140, 66, 171, 322, 321, 18, 320, 152, 152, 84,
	319, 291, 11, 261, 37, 79, 17, 18, 153, 153,
	187, 156, 26, 330, 208, 209, 177, 18, 178, 180,
	181, 179, 154, 195, 5, 6, 219, 18, 301, 154,
	216, 223, 141, 18, 283, 49, 223, 280, 233, 237,
	218, 224, 166, 79, 238, 217, 224, 210, 148, 79,
	28, 185, 17, 18, 155, 228, 148, 213, 14, 281,
	229, 150, 158, 148, 276, 266, 240, 170, 205, 244,
	142, 87, 157, 149, 202, 144, 85, 254, 207, 169,
	71, 18, 239, 170, 5, 6, 177, 18, 177, 18,
	154, 329, 61, 171, 327, 169, 5, 6, 296, 18,
	313, 243, 5, 6, 255, 165, 259, 264, 275, 171,
	35, 275, 260, 50, 268, 208, 209, 177, 18, 178,
	180, 181, 179, 284, 195, 286, 54, 33, 277, 32,
	63, 5, 6, 177, 18, 178, 180, 181, 179, 5,
	6, 31, 18, 164, 216, 162, 287, 137, 223, 289,
	160, 299, 298, 293, 218, 297, 161, 300, 224, 217,
	233, 272, 305, 237, 304, 306, 308, 309, 238, 311,
	312, 310, 302, 307, 288, 303, 273, 230, 290, 46,
	72, 227, 223, 86, 138, 316, 34, 271, 93, 93,
	317, 318, 224, 82, 136, 134, 325, 131, 104, 109,
	95, 163, 24, 133, 108, 42, 101, 21, 253, 249,
	43, 114, 12, 36, 111, 22, 15, 25, 139, 89,
	100, 29, 3, 23, 105, 7, 112, 107, 30, 250,
	27, 106, 113, 96, 97, 246, 190, 8, 1, 20,
	19, 110, 117, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 129, 130, 170, 205, 13, 9,
	102, 103, 235, 202, 234, 200, 203, 207, 169, 197,
	292, 204, 256, 198, 212, 211, 146, 115, 270, 154,
	214, 194, 171, 189, 143, 57, 231, 98, 182, 175,
	174, 173, 167, 56, 135, 132, 78, 77, 116, 99,
	92, 94, 90, 55, 208, 209, 177, 18, 178, 180,
	181, 179, 53, 195, 47, 41, 40, 184, 64, 62,
	60, 45, 44, 39, 38, 2,
}
var asn1Pact = [...]int{

	150, -1000, 150, -1000, -1000, -1000, -1000, -1000, 43, 317,
	-1000, 98, 315, 52, -1000, -1000, 89, -1000, -1000, 335,
	315, -1000, 212, 200, 198, 277, -1000, -1000, 156, 325,
	-1000, -1000, -1000, -1000, -1000, 42, 304, -1000, 313, -1000,
	269, -1000, 72, -1000, 144, -1000, 187, -1, 16, -1000,
	-1000, 121, -1000, 144, -1000, -1000, -1000, -1000, 84, -36,
	11, -1000, 144, -1000, 35, -1000, 117, -1000, -1000, 144,
	-1000, 111, -1000, -36, 333, -1000, 308, -1000, 295, 262,
	332, -1000, -1000, 150, 144, 110, -1000, -1000, -1000, 116,
	-1000, 87, -1000, -1000, -1000, 114, 102, 95, -1000, -1000,
	113, 103, -48, -1000, -1000, -1000, -1000, 222, 235, 217,
	-1000, -1000, 294, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 215, -1000, -1000, -1000, 151, -1000, -1000, -1000, 179,
	43, -1000, -1000, -1000, 62, -1000, 87, -1000, 62, -27,
	40, 264, -1000, -1000, 87, 40, 260, 126, 134, 150,
	-1000, -1000, -1000, -1000, -1000, -44, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -47, -1000, -1000, 147,
	-1000, -1000, -1000, -1000, -1000, -1000, 109, -1000, 4, -13,
	310, 3, -1000, 309, -1000, 352, -1000, -1000, -1000, -1000,
	-1000, -1000, -36, -1000, -35, 87, 87, -1000, -48, -1000,
	-1000, 41, -1000, -1000, -12, -1000, -36, -1000, -1000, -53,
	105, 0, -1000, -36, 267, 259, -1000, 126, -1000, 104,
	126, -23, -1000, 76, 99, -7, -1000, 73, -1000, -1000,
	-1000, -1000, 126, -10, -1000, 177, 163, -1000, -1000, 163,
	163, -1000, -1000, 163, 39, -1000, 28, -56, 149, -1000,
	-1000, -1000, -1000, -27, -1000, 132, -1000, 78, 267, -1000,
	-1000, -1000, 179, -36, -36, -1000, -1000, -36, -1000, 126,
	132, -1000, 134, 132, -1000, 146, -19, 3, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -59, -1000, -1000, -1000,
	-1000, 40, -1000, -1000, -1000, 308, 308, -1000, 38, 34,
	-1000, 32, 31, -29, 352, -1000, -25, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 140, -1000, -1000, -34, 137, 53,
	-1000,
}
var asn1Pgo = [...]int{

	0, 445, 26, 0, 342, 444, 443, 442, 441, 440,
	439, 250, 438, 67, 437, 436, 435, 434, 233, 432,
	246, 423, 21, 10, 422, 421, 4, 420, 2, 419,
	418, 7, 417, 416, 415, 414, 413, 22, 412, 6,
	5, 411, 410, 409, 408, 407, 406, 13, 405, 404,
	28, 1, 403, 19, 18, 44, 401, 16, 400, 11,
	15, 24, 3, 9, 23, 398, 397, 25, 396, 20,
	395, 394, 393, 392, 391, 390, 389, 386, 31, 385,
	17, 33, 384, 382, 8, 379, 30, 378, 336, 360,
	359, 327, 358, 357, 355, 349,
}
var asn1R1 = [...]int{

	0, 92, 1, 1, 93, 4, 85, 85, 86, 86,
	87, 87, 88, 88, 88, 3, 2, 2, 89, 89,
	90, 90, 91, 91, 91, 91, 5, 5, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 14, 14, 11,
	12, 12, 13, 13, 13, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 19, 20, 20, 20,
	21, 22, 23, 31, 31, 32, 33, 35, 35, 35,
	35, 34, 34, 34, 24, 26, 26, 27, 27, 27,
	27, 27, 27, 25, 25, 25, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 30,
	30, 30, 36, 37, 37, 38, 38, 38, 38, 38,
	38, 39, 39, 41, 41, 41, 40, 44, 43, 42,
	42, 82, 83, 83, 84, 84, 84, 84, 84, 28,
	45, 45, 45, 46, 46, 47, 47, 94, 94, 95,
	95, 49, 48, 50, 50, 50, 50, 51, 51, 52,
	52, 53, 53, 54, 54, 55, 55, 56, 56, 56,
	56, 56, 58, 58, 59, 59, 59, 60, 60, 60,
	61, 61, 62, 62, 62, 63, 63, 63, 63, 64,
	64, 65, 65, 66, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 66, 66, 66, 67, 67, 68,
	68, 69, 70, 71, 72, 74, 74, 75, 75, 73,
	73, 73, 73, 76, 76, 77, 78, 79, 80, 80,
	81, 81, 81, 57,
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
	4, 6, 6, 1, 4, 4, 1, 1, 1, 2,
	2, 2, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 2, 4, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 5, 9, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 4, 4, 1, 1, 1,
	1, 1, 3, 1, 3, 4, 4, 1, 1, 1,
	1, 3, 4, 1, 1, 3, 5, 1, 3, 1,
	3, 1, 3, 1, 3, 1, 3, 1, 1, 1,
	1, 1, 1, 3, 2, 1, 1, 1, 3, 3,
	0, 1, 1, 3, 5, 3, 2, 3, 1, 0,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 0, 1, 1,
	2, 3, 1, 1, 3, 1, 1, 1, 1, 1,
	2, 2, 3, 2, 1, 2, 2, 2, 0, 1,
	0, 1, 1, 1,
}
var asn1Chk = [...]int{

	-1000, -92, -1, -4, -2, 72, 73, -4, -93, -85,
	-86, 79, 15, -87, 80, -88, -3, 74, 75, -89,
	-90, -91, 20, 28, 7, 22, 80, -88, 81, 6,
	-91, 49, 49, 49, 29, 74, 8, 82, -5, -6,
	-15, -16, 21, 17, -7, -8, 30, -17, 4, 83,
	-18, -2, -3, -19, -20, -21, -36, -48, -2, -3,
	-9, 25, -10, -11, -12, -13, -2, -3, 83, 84,
	83, 79, -20, 6, -22, -23, -31, -32, -33, 85,
	-22, 83, -11, 25, 84, 79, -18, 80, -22, 6,
	-24, -26, -27, -28, -25, 12, 45, 46, -45, -29,
	32, 18, 72, 73, 10, 36, 43, 39, 16, 11,
	53, 26, 38, 44, 23, -66, -30, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 9, -34, 28, 20, -35, 52, 5, 42, 6,
	-2, -13, 80, -49, 79, -67, -68, -69, 81, 79,
	79, -81, -69, -78, 47, 79, -81, 79, 79, 87,
	48, 41, 48, 27, 48, 74, -37, -38, -39, 36,
	24, 50, -40, -41, -42, -43, -2, 74, 76, 79,
	77, 78, -44, -3, -14, -86, -50, 68, -51, -52,
	4, -53, -54, -55, -56, 81, -57, -76, -72, -78,
	-79, -37, 31, -77, -74, 25, -28, 35, 72, 73,
	-69, -70, -71, -50, -58, -59, -3, -60, -23, 68,
	-61, -62, -63, -3, -23, 13, -60, 37, -69, -61,
	37, -46, -47, -3, -82, -83, -84, -3, -40, 68,
	-2, 86, 87, 74, 80, 84, -94, 88, 51, 19,
	-95, 89, 33, 19, -51, -22, -73, 69, 91, -69,
	-67, 82, 80, 84, -23, 90, 80, 84, -23, -64,
	-65, 40, 14, 37, -80, -3, 80, -80, 80, 84,
	81, 80, 84, 81, -3, 84, 68, -53, -55, -54,
	-55, 82, -75, -57, 34, 91, 69, -59, -39, -40,
	-63, 70, -64, -37, -23, -31, -31, -47, -40, -39,
	-84, -40, -39, 74, 84, 91, -62, -26, -26, 82,
	82, 82, 82, 80, 84, -51, 71, 74, 84, 74,
	80,
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
	62, 197, 75, 76, 77, 0, 220, 220, 129, 83,
	99, 100, 130, 131, 86, 87, 88, 0, 0, 0,
	92, 93, 0, 95, 96, 97, 98, 183, 184, 185,
	186, 187, 188, 189, 190, 191, 192, 193, 194, 195,
	196, 0, 65, 72, 73, 0, 68, 69, 70, 0,
	37, 41, 43, 142, 0, 74, 198, 199, 0, 63,
	-2, 0, 221, 222, 0, -2, 0, 0, 0, 0,
	89, 90, 91, 94, 101, 0, 102, 103, 104, 105,
	106, 107, 108, 109, 110, 111, 0, 116, 113, 0,
	119, 120, 118, 117, 39, 38, 0, 143, 144, 147,
	0, 149, 151, 153, 155, 0, -2, 158, 159, 160,
	161, 223, 63, 214, 0, 0, 197, 206, 130, -2,
	200, 0, 202, 203, 0, 162, 63, 165, 166, 167,
	0, 171, 172, 63, 179, 0, 178, 218, 216, 0,
	218, 0, 133, 0, 0, 121, 122, 124, 127, 128,
	132, 66, 0, 0, 141, 0, 0, 137, 138, 0,
	0, 139, 140, 0, 0, 213, 0, 209, 0, 217,
	215, 201, 78, 63, 164, 0, 79, 63, 179, 176,
	180, 181, 0, 63, 63, 219, 80, 63, 84, 0,
	0, 85, 0, 0, 112, 0, 145, 150, 148, 152,
	154, 156, 204, 207, 208, 210, 211, 163, 168, 169,
	173, 63, 175, 182, 177, 0, 0, 134, 0, 0,
	123, 0, 0, 0, 0, 212, 0, 81, 82, 135,
	136, 125, 126, 114, 0, 146, 174, 0, 0, 0,
	115,
}
var asn1Tok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 90, 3, 3, 3, 3, 3, 3,
	81, 82, 3, 3, 84, 3, 87, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 83,
	91, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 85, 3, 86, 89, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 79, 88, 80,
}
var asn1Tok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78,
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
		//line parser/asn1p.y:263
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:270
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:274
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:281
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:285
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
		//line parser/asn1p.y:301
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:302
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:305
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:308
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:314
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:318
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:324
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:327
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:330
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:334
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:339
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:342
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:347
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:348
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:354
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:357
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:363
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:366
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:369
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:372
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:377
		{
			asn1VAL.module = nil
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:378
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:383
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
		//line parser/asn1p.y:402
		{
			asn1VAL.module = nil
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:406
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:412
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:419
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:423
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:427
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:434
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:435
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:438
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:448
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:452
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:459
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:464
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:468
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:476
		{
			asn1VAL.module = nil
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:477
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
		//line parser/asn1p.y:487
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:490
		{
			asn1VAL.xports = nil
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:493
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:499
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:504
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:512
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:517
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:522
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:531
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:534
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:544
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:548
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:552
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:559
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:569
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:574
		{
			asn1VAL.tag = nil
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:575
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:579
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:586
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:592
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:593
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:594
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:595
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:599
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:600
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:601
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 74:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:605
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:609
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 78:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:614
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrChoice
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:619
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSequence
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:624
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSet
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:630
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 82:
		asn1Dollar = asn1S[asn1pt-6 : asn1pt+1]
		//line parser/asn1p.y:636
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSetOf
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
			asn1Dollar[6].expr.Identifier = asn1Dollar[4].str
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:644
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 84:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:649
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:654
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:662
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:663
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:664
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:665
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:666
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:667
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:668
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:669
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:670
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 95:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:671
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 96:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:672
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 99:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:678
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 100:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:679
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 101:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:680
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 102:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:689
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 105:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:710
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 106:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:714
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 107:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:718
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 112:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:730
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 113:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:739
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 114:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:742
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
	case 115:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:754
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
	case 116:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:773
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 117:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:779
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 118:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:784
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 119:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:789
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:792
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:798
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:804
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 123:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:807
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 124:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:814
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 125:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:821
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 126:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:828
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:835
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 128:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:842
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 129:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:862
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 130:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:882
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 131:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:885
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:888
		{
			fmt.Println("XXX")
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:895
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:898
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 135:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:904
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 136:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:910
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 141:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:923
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 142:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:926
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 143:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:934
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 145:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:939
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 146:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:943
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 148:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:951
		{
		}
	case 150:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:957
		{
		}
	case 152:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:963
		{
		}
	case 154:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:970
		{
		}
	case 156:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:976
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSet
		}
	case 157:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:983
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeValue
		}
	case 158:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:987
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSubType
			/* FIXME: */
		}
	case 162:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:999
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 163:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1002
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 164:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1008
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 165:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1011
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 166:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1014
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 167:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1020
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 168:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1026
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 169:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1032
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 170:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1044
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 171:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1045
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 172:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1048
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 173:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1051
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 174:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1054
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 175:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1060
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 176:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1064
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 177:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1067
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			//$$.Meta = $3.Meta
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeComponentsOf
		}
	case 178:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1072
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 179:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1082
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagNoMark
			asn1VAL.marker.Value = nil
		}
	case 180:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1087
		{
			asn1VAL.marker = asn1Dollar[1].marker
		}
	case 181:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1091
		{
			fmt.Println("YYY")
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect
			asn1VAL.marker.Value = nil
		}
	case 182:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1097
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagDefault
			asn1VAL.marker.Value = asn1Dollar[2].value
		}
	case 183:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1105
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBMPString
		}
	case 184:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1106
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralString
			fmt.Println("WARNING: GeneralString is not fully supported")
		}
	case 185:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1110
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGraphicsString
			fmt.Println("WARNING: GraphicString is not fully supported")
		}
	case 186:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1114
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeIA5String
		}
	case 187:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1115
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeISO646String
		}
	case 188:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1116
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNumericString
		}
	case 189:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1117
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypePrintableString
		}
	case 190:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1118
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeT61String
			fmt.Println("WARNING: T61String is not fully supported")
		}
	case 191:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1122
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeTeletexString
		}
	case 192:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1123
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUniversalString
		}
	case 193:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1124
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtf8String
		}
	case 194:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1125
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVideoTexString
			fmt.Println("WARNING: VideotexString is not fully supported")
		}
	case 195:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1129
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVisibleString
		}
	case 196:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1130
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectDescriptor
		}
	case 197:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1136
		{
			asn1VAL.constraint = nil
		}
	case 200:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1141
		{
			// FIXME : Implementation incomplete
		}
	case 201:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1147
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
			// FIXME: Needs implementation
		}
	case 204:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1158
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1Dollar[2].constraint_type
		}
	case 206:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1165
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMin
		}
	case 208:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1172
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMax
		}
	case 209:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1178
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeSimpleRange
		}
	case 210:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1179
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeRightExcludedRange
		}
	case 211:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1180
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftExcludedRange
		}
	case 212:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1181
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftRightExcludedRange
		}
	case 213:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1185
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 214:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1189
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 215:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1195
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// FIXME : Implementation
		}
	case 216:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1202
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			/* FIXME : implementation not yet complete */
		}
	case 217:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1207
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
		}
	case 218:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1212
		{
			asn1VAL.str = ""
		}
	case 219:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1213
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 220:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1219
		{
			asn1VAL.constraint = nil
		}
	}
	goto asn1stack /* stack new state and value */
}
