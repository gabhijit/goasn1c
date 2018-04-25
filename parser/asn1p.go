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

//line parser/asn1p.y:1447

//line yacctab:1
var asn1Exca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	17, 26,
	-2, 45,
	-1, 150,
	80, 168,
	-2, 63,
	-1, 151,
	80, 168,
	-2, 63,
	-1, 191,
	69, 203,
	91, 203,
	-2, 155,
	-1, 205,
	87, 17,
	-2, 129,
}

const asn1Private = 57344

const asn1Last = 424

var asn1Act = [...]int{

	178, 183, 217, 229, 167, 75, 263, 211, 163, 191,
	225, 187, 16, 218, 16, 147, 186, 145, 216, 196,
	181, 74, 10, 188, 222, 304, 171, 4, 250, 4,
	286, 259, 245, 154, 215, 221, 235, 241, 234, 313,
	310, 18, 79, 52, 311, 59, 83, 67, 65, 269,
	251, 79, 261, 270, 59, 185, 73, 256, 315, 221,
	303, 257, 276, 67, 68, 69, 273, 261, 238, 51,
	52, 58, 185, 66, 240, 165, 201, 81, 70, 309,
	58, 80, 197, 308, 307, 67, 203, 164, 244, 66,
	215, 306, 165, 201, 282, 88, 51, 18, 200, 197,
	255, 166, 37, 203, 164, 84, 274, 79, 271, 148,
	140, 66, 28, 317, 215, 200, 292, 272, 166, 182,
	268, 18, 260, 204, 205, 172, 18, 173, 175, 176,
	174, 79, 190, 141, 237, 79, 142, 87, 11, 153,
	204, 205, 172, 18, 173, 175, 176, 174, 152, 190,
	212, 219, 219, 226, 230, 214, 220, 220, 231, 161,
	165, 151, 206, 180, 150, 17, 18, 17, 18, 209,
	223, 26, 164, 14, 213, 149, 144, 85, 165, 201,
	71, 233, 18, 202, 316, 197, 166, 232, 314, 203,
	164, 61, 247, 172, 18, 165, 5, 6, 172, 18,
	302, 200, 236, 160, 166, 285, 35, 164, 5, 6,
	172, 18, 173, 175, 176, 174, 252, 253, 258, 248,
	254, 166, 5, 6, 287, 262, 204, 205, 172, 18,
	173, 175, 176, 174, 63, 190, 275, 50, 5, 6,
	48, 18, 277, 5, 6, 172, 18, 173, 175, 176,
	174, 5, 6, 33, 18, 280, 278, 32, 212, 284,
	93, 54, 219, 214, 290, 288, 279, 220, 289, 293,
	281, 226, 31, 295, 230, 291, 297, 299, 231, 300,
	298, 296, 213, 301, 267, 159, 294, 131, 104, 109,
	95, 157, 155, 219, 108, 305, 101, 82, 220, 266,
	156, 114, 46, 34, 111, 312, 158, 86, 5, 6,
	100, 18, 137, 42, 105, 72, 112, 107, 134, 49,
	246, 106, 113, 96, 97, 265, 133, 21, 242, 43,
	12, 110, 117, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 129, 130, 36, 24, 30, 138,
	102, 103, 15, 139, 89, 29, 243, 239, 3, 136,
	22, 7, 25, 8, 1, 20, 27, 19, 23, 13,
	9, 228, 227, 195, 194, 198, 192, 283, 199, 249,
	193, 208, 207, 146, 115, 264, 210, 189, 184, 143,
	57, 224, 98, 177, 170, 169, 168, 162, 56, 135,
	132, 78, 77, 76, 116, 99, 92, 91, 94, 90,
	55, 53, 47, 41, 40, 179, 64, 62, 60, 45,
	44, 39, 38, 2,
}
var asn1Pact = [...]int{

	150, -1000, 150, -1000, -1000, -1000, -1000, -1000, 59, 315,
	-1000, 93, 340, 91, -1000, -1000, 31, -1000, -1000, 349,
	340, -1000, 223, 208, 204, 274, -1000, -1000, 132, 338,
	-1000, -1000, -1000, -1000, -1000, 20, 292, -1000, 312, -1000,
	272, -1000, 236, -1000, 179, -1000, 166, -19, -5, -1000,
	-1000, 101, -1000, 179, -1000, -1000, -1000, -1000, 50, -43,
	-6, -1000, 179, -1000, 21, -1000, 98, -1000, -1000, 179,
	-1000, 57, -1000, -43, 348, -1000, 278, -1000, 298, 307,
	347, -1000, -1000, 150, 179, 56, -1000, -1000, -1000, 97,
	-1000, 28, -1000, -1000, -1000, 96, 85, 82, -1000, -1000,
	69, 60, -54, -1000, -1000, -1000, -1000, 244, 259, 243,
	-1000, -1000, 279, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 237, -1000, -1000, -1000, 129, -1000, -1000, -1000, 136,
	59, -1000, -1000, -1000, 51, -1000, 28, -1000, 51, -34,
	22, 22, 107, 119, 150, -1000, -1000, -1000, -1000, -1000,
	-48, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -51, -1000, -1000, 128, -1000, -1000, -1000, -1000, -1000,
	-1000, 54, -1000, -16, -14, 309, -1, -1000, 301, -1000,
	68, -1000, -1000, -1000, -1000, -1000, -1000, -43, -1000, -41,
	28, 28, 28, -1000, -54, -1000, -1000, 18, -1000, -1000,
	-23, -1000, -43, -1000, -1000, -59, 42, -17, -1000, -43,
	285, 247, -1000, 40, -31, -1000, 27, 37, -18, -1000,
	25, -1000, -1000, -1000, -1000, 107, -22, -1000, 174, 154,
	-1000, -1000, 154, 154, -1000, -1000, 154, 12, -1000, 171,
	-61, 155, -1000, -1000, -1000, -1000, -1000, -34, -1000, 124,
	-1000, 46, 285, -1000, -1000, -1000, 136, -43, -1000, -1000,
	107, 124, -1000, 119, 124, -1000, 126, -24, -1, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -66, -1000, -1000,
	-1000, -1000, 22, -1000, -1000, -1000, -1000, 9, 2, -1000,
	1, -3, -40, 68, -1000, -32, -1000, -1000, -1000, -1000,
	-1000, 114, -1000, -1000, -26, 110, 33, -1000,
}
var asn1Pgo = [...]int{

	0, 423, 26, 0, 358, 422, 421, 420, 419, 418,
	417, 234, 416, 48, 415, 414, 413, 412, 237, 411,
	261, 410, 21, 5, 409, 408, 407, 406, 183, 405,
	404, 403, 402, 401, 400, 399, 398, 19, 397, 8,
	4, 396, 395, 394, 393, 392, 391, 10, 390, 389,
	20, 1, 388, 16, 11, 23, 387, 9, 386, 7,
	24, 18, 2, 13, 6, 385, 384, 17, 383, 15,
	382, 381, 380, 379, 378, 377, 376, 375, 374, 373,
	372, 371, 3, 370, 22, 369, 352, 367, 365, 327,
	364, 363, 357, 356,
}
var asn1R1 = [...]int{

	0, 90, 1, 1, 91, 4, 83, 83, 84, 84,
	85, 85, 86, 86, 86, 3, 2, 2, 87, 87,
	88, 88, 89, 89, 89, 89, 5, 5, 6, 7,
	7, 8, 8, 9, 9, 10, 10, 14, 14, 11,
	12, 12, 13, 13, 13, 15, 15, 16, 16, 16,
	17, 17, 18, 18, 18, 19, 19, 20, 20, 20,
	21, 22, 23, 31, 31, 32, 33, 35, 35, 35,
	35, 34, 34, 34, 24, 26, 26, 27, 27, 27,
	27, 25, 25, 25, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 30, 30, 30,
	36, 37, 37, 38, 38, 38, 38, 38, 38, 39,
	39, 41, 41, 41, 40, 44, 43, 42, 42, 80,
	81, 81, 82, 82, 82, 82, 82, 28, 45, 45,
	45, 46, 46, 47, 47, 92, 92, 93, 93, 49,
	48, 50, 50, 50, 50, 51, 51, 52, 52, 53,
	53, 54, 54, 55, 55, 56, 56, 56, 56, 56,
	58, 58, 59, 59, 59, 60, 60, 60, 61, 61,
	62, 62, 62, 63, 63, 63, 63, 64, 64, 65,
	65, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 66, 67, 67, 68, 68, 69,
	70, 71, 72, 74, 74, 75, 75, 73, 73, 73,
	73, 76, 76, 77, 78, 79, 57,
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
	4, 1, 4, 4, 1, 1, 1, 2, 2, 2,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 2,
	4, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 5, 9, 1, 1, 1, 1, 1, 1,
	1, 3, 1, 4, 4, 1, 1, 1, 1, 1,
	3, 1, 3, 4, 4, 1, 1, 1, 1, 3,
	4, 1, 1, 3, 5, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 3, 1, 1, 1, 1, 1,
	1, 3, 2, 1, 1, 1, 3, 3, 0, 1,
	1, 3, 5, 3, 2, 3, 1, 0, 1, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 1, 2, 3,
	1, 1, 3, 1, 1, 1, 1, 1, 2, 2,
	3, 2, 1, 2, 2, 2, 1,
}
var asn1Chk = [...]int{

	-1000, -90, -1, -4, -2, 72, 73, -4, -91, -83,
	-84, 79, 15, -85, 80, -86, -3, 74, 75, -87,
	-88, -89, 20, 28, 7, 22, 80, -86, 81, 6,
	-89, 49, 49, 49, 29, 74, 8, 82, -5, -6,
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
	79, 79, 79, 79, 87, 48, 41, 48, 27, 48,
	74, -37, -38, -39, 36, 24, 50, -40, -41, -42,
	-43, -2, 74, 76, 79, 77, 78, -44, -3, -14,
	-84, -50, 68, -51, -52, 4, -53, -54, -55, -56,
	81, -57, -76, -72, -78, -79, -37, 31, -77, -74,
	47, 25, -28, 35, 72, 73, -69, -70, -71, -50,
	-58, -59, -3, -60, -23, 68, -61, -62, -63, -3,
	-23, 13, -60, -61, -46, -47, -3, -80, -81, -82,
	-3, -40, 68, -2, 86, 87, 74, 80, 84, -92,
	88, 51, 19, -93, 89, 33, 19, -51, -22, -73,
	69, 91, -69, -69, -67, 82, 80, 84, -23, 90,
	80, 84, -23, -64, -65, 40, 14, 37, 80, 80,
	84, 81, 80, 84, 81, -3, 84, 68, -53, -55,
	-54, -55, 82, -75, -57, 34, 91, 69, -59, -39,
	-40, -63, 70, -64, -37, -23, -47, -40, -39, -82,
	-40, -39, 74, 84, 91, -62, 82, 82, 82, 82,
	80, 84, -51, 71, 74, 84, 74, 80,
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
	62, 195, 75, 76, 77, 0, 0, 0, 127, 81,
	97, 98, 128, 129, 84, 85, 86, 0, 0, 0,
	90, 91, 0, 93, 94, 95, 96, 181, 182, 183,
	184, 185, 186, 187, 188, 189, 190, 191, 192, 193,
	194, 0, 65, 72, 73, 0, 68, 69, 70, 0,
	37, 41, 43, 140, 0, 74, 196, 197, 0, 63,
	-2, -2, 0, 0, 0, 87, 88, 89, 92, 99,
	0, 100, 101, 102, 103, 104, 105, 106, 107, 108,
	109, 0, 114, 111, 0, 117, 118, 116, 115, 39,
	38, 0, 141, 142, 145, 0, 147, 149, 151, 153,
	0, -2, 156, 157, 158, 159, 216, 63, 212, 0,
	0, 0, 195, 204, 128, -2, 198, 0, 200, 201,
	0, 160, 63, 163, 164, 165, 0, 169, 170, 63,
	177, 0, 176, 0, 0, 131, 0, 0, 119, 120,
	122, 125, 126, 130, 66, 0, 0, 139, 0, 0,
	135, 136, 0, 0, 137, 138, 0, 0, 211, 0,
	207, 0, 214, 215, 213, 199, 78, 63, 162, 0,
	79, 63, 177, 174, 178, 179, 0, 63, 80, 82,
	0, 0, 83, 0, 0, 110, 0, 143, 148, 146,
	150, 152, 154, 202, 205, 206, 208, 209, 161, 166,
	167, 171, 63, 173, 180, 175, 132, 0, 0, 121,
	0, 0, 0, 0, 210, 0, 133, 134, 123, 124,
	112, 0, 144, 172, 0, 0, 0, 113,
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
		//line parser/asn1p.y:260
		{
			AllModules = asn1Dollar[1].grammar
			fmt.Println(AllModules)
		}
	case 2:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:267
		{
			asn1VAL.grammar = asn1types.NewAsn1Grammar()
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[1].module)
		}
	case 3:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:271
		{
			asn1VAL.grammar = asn1Dollar[1].grammar
			asn1VAL.grammar.Modules = append(asn1VAL.grammar.Modules, asn1Dollar[2].module)
		}
	case 4:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:278
		{
			currentModule = asn1types.NewAsn1Module()
		}
	case 5:
		asn1Dollar = asn1S[asn1pt-9 : asn1pt+1]
		//line parser/asn1p.y:282
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
		//line parser/asn1p.y:298
		{
			asn1VAL.oid = nil
		}
	case 7:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:299
		{
			asn1VAL.oid = asn1Dollar[1].oid
		}
	case 8:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:302
		{
			asn1VAL.oid = asn1Dollar[2].oid
		}
	case 9:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:305
		{
			asn1VAL.oid = nil
		}
	case 10:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:311
		{
			asn1VAL.oid = asn1types.NewAsn1Oid()
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[1].oid_arc)
		}
	case 11:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:315
		{
			asn1VAL.oid = asn1Dollar[1].oid
			asn1VAL.oid.Arcs = append(asn1VAL.oid.Arcs, asn1Dollar[2].oid_arc)
		}
	case 12:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:321
		{ /* iso */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = -1
		}
	case 13:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:324
		{ /* iso(1) */
			asn1VAL.oid_arc.Name = asn1Dollar[1].str
			asn1VAL.oid_arc.Num = asn1Dollar[3].num
		}
	case 14:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:327
		{ /* 1 */
			asn1VAL.oid_arc.Name = ""
			asn1VAL.oid_arc.Num = asn1Dollar[1].num
		}
	case 15:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:331
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 16:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:336
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 17:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:339
		{
			asn1VAL.str = asn1Dollar[1].str
		}
	case 18:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:344
		{
			asn1VAL.module_flags = asn1types.ModuleFlagNoFlags
		}
	case 19:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:345
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 20:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:351
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags
		}
	case 21:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:354
		{
			asn1VAL.module_flags = asn1Dollar[1].module_flags | asn1Dollar[2].module_flags
		}
	case 22:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:360
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExplicitTags
		}
	case 23:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:363
		{
			asn1VAL.module_flags = asn1types.ModuleFlagImplicitTags
		}
	case 24:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:366
		{
			asn1VAL.module_flags = asn1types.ModuleFlagAutomaticTags
		}
	case 25:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:369
		{
			asn1VAL.module_flags = asn1types.ModuleFlagExtensibilityImplied
		}
	case 26:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:374
		{
			asn1VAL.module = nil
		}
	case 27:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:375
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 28:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:380
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
		//line parser/asn1p.y:399
		{
			asn1VAL.module = nil
		}
	case 31:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:403
		{
			asn1VAL.module = asn1Dollar[2].module
		}
	case 32:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:409
		{
			// FIXME: Need to figure out how to call the Lexer's Error()
			return -1
		}
	case 33:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:416
		{
			asn1VAL.module = asn1types.NewAsn1Module()
		}
	case 35:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:420
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[1].xports)
		}
	case 36:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:424
		{
			asn1VAL.module = asn1Dollar[1].module
			asn1VAL.module.Imports = append(asn1VAL.module.Imports, asn1Dollar[2].xports)
		}
	case 37:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:431
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = nil
			asn1VAL.aid.Value = nil
		}
	case 38:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:432
		{
			asn1VAL.aid = asn1types.NewAsn1AssignedIdentifier()
			asn1VAL.aid.Oid = asn1Dollar[1].oid
			asn1VAL.aid.Value = nil
		}
	case 39:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:435
		{
			asn1VAL.xports = asn1Dollar[1].xports

			asn1VAL.xports.Type = asn1types.Asn1XportsTypeImport
			asn1VAL.xports.FromModule = asn1Dollar[3].str
			asn1VAL.xports.Identifier = asn1Dollar[4].aid
		}
	case 40:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:445
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 41:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:449
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 42:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:456
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 43:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:461
		{ /* Completely equivalent to above */
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 44:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:465
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
		}
	case 45:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:473
		{
			asn1VAL.module = nil
		}
	case 46:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:474
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
		//line parser/asn1p.y:484
		{
			asn1VAL.xports = asn1Dollar[2].xports
		}
	case 48:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:487
		{
			asn1VAL.xports = nil
		}
	case 49:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:490
		{
			/* Empty EXPORTS clause effectively prohibits export. */
			asn1VAL.xports = asn1types.NewAsn1Xports()
		}
	case 50:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:496
		{
			asn1VAL.xports = asn1types.NewAsn1Xports()
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[1].expr)
		}
	case 51:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:501
		{
			asn1VAL.xports = asn1Dollar[1].xports
			asn1VAL.xports.Type = asn1types.Asn1XportsTypeExport
			asn1VAL.xports.Members = append(asn1VAL.xports.Members, asn1Dollar[3].expr)
		}
	case 52:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:509
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 53:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:514
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 54:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:519
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExportVariable
		}
	case 55:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:528
		{
			asn1VAL.module = asn1Dollar[1].module
		}
	case 56:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:531
		{
			asn1VAL.module = asn1Dollar[1].module
			for _, m := range asn1Dollar[2].module.Members {
				asn1VAL.module.Members = append(asn1VAL.module.Members, m)
			}
		}
	case 57:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:541
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 58:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:545
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 59:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:549
		{
			asn1VAL.module = asn1types.NewAsn1Module()
			asn1VAL.module.Members = append(asn1VAL.module.Members, asn1Dollar[1].expr)
		}
	case 60:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:556
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			// FIXME : Need to add code for type of expression
		}
	case 62:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:566
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 63:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:571
		{
			asn1VAL.tag = nil
		}
	case 64:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:572
		{
			asn1VAL.tag = asn1Dollar[1].tag
		}
	case 65:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:576
		{
			asn1VAL.tag = asn1Dollar[1].tag
			asn1VAL.tag.Mode = asn1Dollar[2].tag.Mode
		}
	case 66:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:583
		{
			asn1VAL.tag = asn1Dollar[2].tag
			asn1VAL.tag.Val = asn1Dollar[3].num
		}
	case 67:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:589
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassContextSpec
		}
	case 68:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:590
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassUniversal
		}
	case 69:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:591
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassApplication
		}
	case 70:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:592
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Class = asn1types.Asn1TagClassPrivate
		}
	case 71:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:596
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeDefault
		}
	case 72:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:597
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 73:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:598
		{
			asn1VAL.tag = asn1types.NewAsn1Tag()
			asn1VAL.tag.Mode = asn1types.Asn1TagModeImplicit
		}
	case 74:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:602
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 75:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:606
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 78:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:611
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrChoice
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 79:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:616
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSequence
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 80:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:621
		{
			fmt.Println("AAA")
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeConstrSet
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 81:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:629
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1Dollar[1].expr_type

		}
	case 82:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:634
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeInteger
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 83:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:639
		{
			asn1VAL.expr = asn1Dollar[3].expr
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeEnumerated
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 84:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:647
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBoolean
		}
	case 85:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:648
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNull
		}
	case 86:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:649
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeReal
		}
	case 87:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:650
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeOctetString
		}
	case 88:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:651
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEmbeddedPdv
		}
	case 89:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:652
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeCharacterString
		}
	case 90:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:653
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtcTime
		}
	case 91:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:654
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralizedTime
		}
	case 92:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:655
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectIdentifier
		}
	case 93:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:656
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeRelativeOid
		}
	case 94:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:657
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeExternal
		}
	case 97:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:663
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeInteger
		}
	case 98:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:664
		{
			fmt.Println("DDD")
			asn1VAL.expr_type = asn1types.Asn1ExprTypeEnumerated
		}
	case 99:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:665
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBitString
		}
	case 100:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:674
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
		}
	case 103:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:695
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeNull
		}
	case 104:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:699
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeFalse
		}
	case 105:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:703
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeTrue
		}
	case 110:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:715
		{
			ref := asn1types.NewAsn1Reference()
			_ = ref
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 111:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:724
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 112:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:727
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
		//line parser/asn1p.y:739
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
		//line parser/asn1p.y:758
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 115:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:764
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 116:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:769
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 117:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:774
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 118:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:777
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 119:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:783
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// TODO handle Enumeration validation
		}
	case 120:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:789
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 121:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:792
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// $$.Members = append($$.Members, $3)
		}
	case 122:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:799
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 123:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:806
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 124:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:813
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 125:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:820
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			//$$.Identifier = $1;
		}
	case 126:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:827
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible

			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = "..."
		}
	case 127:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:847
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeReference
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeTypeRef
		}
	case 128:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:867
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 129:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:870
		{
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 130:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:873
		{
			fmt.Println("XXX")
			asn1VAL.ref = asn1types.NewAsn1Reference()
		}
	case 131:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:880
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 132:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:883
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 133:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:889
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 134:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:895
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeUniversal
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValue
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 139:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:908
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
		}
	case 140:
		asn1Dollar = asn1S[asn1pt-4 : asn1pt+1]
		//line parser/asn1p.y:911
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeValueSet
		}
	case 141:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:919
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 143:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:924
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 144:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:928
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeExtensibilityMark
		}
	case 146:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:936
		{
		}
	case 148:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:942
		{
		}
	case 150:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:948
		{
		}
	case 152:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:955
		{
		}
	case 154:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:961
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSet
		}
	case 155:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:968
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeValue
		}
	case 156:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:972
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1types.Asn1ConstraintTypeSubType
			/* FIXME: */
		}
	case 160:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:984
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 161:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:987
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 162:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:993
		{
			asn1VAL.expr = asn1Dollar[2].expr
		}
	case 163:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:996
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 164:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:999
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 165:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1005
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 166:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1011
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 167:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1017
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			asn1VAL.expr.Identifier = "..."
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeExtensible
			asn1VAL.expr.Meta = asn1types.Asn1ExprMetaTypeType
		}
	case 168:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1029
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 169:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1030
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 170:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1033
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
		}
	case 171:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1036
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 172:
		asn1Dollar = asn1S[asn1pt-5 : asn1pt+1]
		//line parser/asn1p.y:1039
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 173:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1045
		{
			asn1VAL.expr = asn1Dollar[2].expr
			asn1VAL.expr.Identifier = asn1Dollar[1].str
		}
	case 174:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1049
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 175:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1052
		{
			asn1VAL.expr = asn1types.NewAsn1Expression()
			//$$.Meta = $3.Meta
			asn1VAL.expr.Type = asn1types.Asn1ExprTypeComponentsOf
		}
	case 176:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1057
		{
			asn1VAL.expr = asn1Dollar[1].expr
		}
	case 177:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1067
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagNoMark
			asn1VAL.marker.Value = nil
		}
	case 178:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1072
		{
			asn1VAL.marker = asn1Dollar[1].marker
		}
	case 179:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1076
		{
			fmt.Println("YYY")
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagOptional | asn1types.Asn1MarkerFlagIndirect
			asn1VAL.marker.Value = nil
		}
	case 180:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1082
		{
			asn1VAL.marker = asn1types.NewAsn1Marker()
			asn1VAL.marker.Flags = asn1types.Asn1MarkerFlagDefault
			asn1VAL.marker.Value = asn1Dollar[2].value
		}
	case 181:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1090
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeBMPString
		}
	case 182:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1091
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGeneralString
			fmt.Println("WARNING: GeneralString is not fully supported")
		}
	case 183:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1095
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeGraphicsString
			fmt.Println("WARNING: GraphicString is not fully supported")
		}
	case 184:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1099
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeIA5String
		}
	case 185:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1100
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeISO646String
		}
	case 186:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1101
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeNumericString
		}
	case 187:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1102
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypePrintableString
		}
	case 188:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1103
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeT61String
			fmt.Println("WARNING: T61String is not fully supported")
		}
	case 189:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1107
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeTeletexString
		}
	case 190:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1108
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUniversalString
		}
	case 191:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1109
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeUtf8String
		}
	case 192:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1110
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVideoTexString
			fmt.Println("WARNING: VideotexString is not fully supported")
		}
	case 193:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1114
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeVisibleString
		}
	case 194:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1115
		{
			asn1VAL.expr_type = asn1types.Asn1ExprTypeObjectDescriptor
		}
	case 195:
		asn1Dollar = asn1S[asn1pt-0 : asn1pt+1]
		//line parser/asn1p.y:1121
		{
			asn1VAL.constraint = nil
		}
	case 198:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1126
		{
			// FIXME : Implementation incomplete
		}
	case 199:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1132
		{
			asn1VAL.constraint = asn1Dollar[2].constraint
			// FIXME: Needs implementation
		}
	case 202:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1143
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			asn1VAL.constraint.Type = asn1Dollar[2].constraint_type
		}
	case 204:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1150
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMin
		}
	case 206:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1157
		{
			asn1VAL.value = asn1types.NewAsn1Value()
			asn1VAL.value.Type = asn1types.Asn1ValueTypeMax
		}
	case 207:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1163
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeSimpleRange
		}
	case 208:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1164
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeRightExcludedRange
		}
	case 209:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1165
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftExcludedRange
		}
	case 210:
		asn1Dollar = asn1S[asn1pt-3 : asn1pt+1]
		//line parser/asn1p.y:1166
		{
			asn1VAL.constraint_type = asn1types.Asn1ConstraintTypeLeftRightExcludedRange
		}
	case 211:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1170
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 212:
		asn1Dollar = asn1S[asn1pt-1 : asn1pt+1]
		//line parser/asn1p.y:1174
		{
			asn1VAL.value = asn1types.NewAsn1Value()
		}
	case 213:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1180
		{
			asn1VAL.expr = asn1Dollar[1].expr
			// FIXME : Implementation
		}
	case 214:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1187
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
			/* FIXME : implementation not yet complete */
		}
	case 215:
		asn1Dollar = asn1S[asn1pt-2 : asn1pt+1]
		//line parser/asn1p.y:1192
		{
			asn1VAL.constraint = asn1types.NewAsn1Constraint()
		}
	}
	goto asn1stack /* stack new state and value */
}
