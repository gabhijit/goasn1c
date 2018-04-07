// Copyright (c) 2018 Abhijit Gadgil <gabhijit@iitbombay.org>. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Implements Lexer using ideas from - Rob Pike's talk about 'Lexical Scanning in Go'

// Lexical Specification from ITU-T X.680 (072002)
// http://www.itu.int/rec/T-REC-X.680-200207-S/en

package parser

type itemType int

type item struct {
	typ itemType
	val string
}

const (
	itemError itemType = iota // Error occurred

	itemTypeReference   // 'typereference' from section 11.2
	itemIdentifier      // 'identifier' from section 11.3
	itemValueReference  // 'valuereference' from section 11.4
	itemModuleReference // 'modulereference' from section 11.5
	itemComment         // 'comment from section 11.6
	itemEmpty           // 'empty' from section 11.7
	itemNumber          // 'number from section 11.8
	itemRealNumber      //'realnumber from section 11.9
	itemBstring         // 'bstring' from section 11.10
	itemXmlBstring      // 'xmlbstring' from section 11.11
	itemHstring         // 'hstring' from section 11.12
	itemXmlHstring      // 'xmlhstring' from section 11.13
	itemCstring         // 'cstring' from section 11.14
	itemXmlCstring      // 'xmlcstring' from section 11.15
	itemAssignment      // section 11.16 ':=='
	itemRangeSeparator  // section 11.17 '..'
	itemEllipsis        // section 11.18 '...'
	itemTwoLeftBracket  // section 11.19 '[['
	itemTwoRightBracket //section 11.20 ']]'
	itemXmlEndTagStart  // section 11.21 '</'
	itemXmlSingleTagEnd // section 11.22 '/>'
	itemXmlBooleanTrue  // section 11.23
	itemXmlBooleanFalse // section 11.24
	// TODO: How should we treat 11.25?
	itemSymbol // section 11.26

	// Reserved Keywords
	itemReservedABSENT
	itemReservedABSTRACT_SYNTAX
	itemReservedALL
	itemReservedAPPLICATION
	itemReservedAUTOMATIC
	itemReservedBEGIN
	itemReservedBMPString
	itemReservedBOOLEAN
	itemReservedBY
	itemReservedCHARACTER
	itemReservedCHOICE
	itemReservedCLASS
	itemReservedCOMPONENT
	itemReservedCOMPONENTS
	itemReservedCONSTRAINED
	itemReservedCONTAINING
	itemReservedDEFAULT
	itemReservedDEFINITIONS
	itemReservedEMBEDDED
	itemReservedENCODED
	itemReservedEND
	itemReservedENUMERATED
	itemReservedEXCEPT
	itemReservedEXPLICIT
	itemReservedEXPORTS
	itemReservedEXTENSIBILITY
	itemReservedEXTERNAL
	itemReservedFALSE
	itemReservedFROM
	itemReservedGeneralString
	itemReservedGeneralizedTime
	itemReservedGraphicString
	itemReservedIA5String
	itemReservedIDENTIFIER
	itemReservedIMPLICIT
	itemReservedIMPLIED
	itemReservedIMPORTS
	itemReservedINCLUDES
	itemReservedINSTANCE
	itemReservedINTEGER
	itemReservedINTERSECTION
	itemReservedISO646String
	itemReservedMAX
	itemReservedMIN
	itemReservedMINUS_INFINITY
	itemReservedNULL
	itemReservedNumericString
	itemReservedOBJECT
	itemReservedOCTET
	itemReservedOF
	itemReservedOPTIONAL
	itemReservedObjectDescriptor
	itemReservedPATTERN
	itemReservedPDV
	itemReservedPLUS_INFINITY
	itemReservedPRESENT
	itemReservedPRIVATE
	itemReservedPrintableString
	itemReservedREAL
	itemReservedRELATIVE_OID
	itemReservedSEQUENCE
	itemReservedSET
	itemReservedSIZE
	itemReservedSTRING
	itemReservedSYNTAX
	itemReservedT61String
	itemReservedTAGS
	itemReservedTRUE
	itemReservedTYPE_IDENTIFIER
	itemReservedTeletexString
	itemReservedUNION
	itemReservedUNIQUE
	itemReservedUNIVERSAL
	itemReservedUTCTime
	itemReservedUTF8String
	itemReservedUniversalString
	itemReservedVideotexString
	itemReservedVisibleString
	itemReservedWITH
	itemReservedlBIT
)

const reservedWordsMap = map[string]itemType{
	"ABSENT":          itemReservedABSENT,
	"ABSTRACT-SYNTAX": itemReservedABSTRACT_SYNTAX,
	"ALL":             itemReservedALL,
	"APPLICATION":     itemReservedAPPLICATION,
	"AUTOMATIC":       itemReservedAUTOMATIC,
	"BEGIN":           itemReservedBEGIN,
	"BIT":             itemReservedlBIT,
	"BMPString":       itemReservedBMPString,
	"BOOLEAN":         itemReservedBOOLEAN,
	"BY":              itemReservedBY,
	"CHARACTER":       itemReservedCHARACTER,
	"CHOICE":          itemReservedCHOICE,
	"CLASS":           itemReservedCLASS,
	"COMPONENT":       itemReservedCOMPONENT,
	"COMPONENTS":      itemReservedCOMPONENTS,
	"CONSTRAINED":     itemReservedCONSTRAINED,
	"CONTAINING":      itemReservedCONTAINING,
	"DEFAULT":         itemReservedDEFAULT,
	"DEFINITIONS":     itemReservedDEFINITIONS,
	"EMBEDDED":        itemReservedEMBEDDED,
	"ENCODED":         itemReservedENCODED,
	"END":             itemReservedEND,
	"ENUMERATED":      itemReservedENUMERATED,
	"EXCEPT":          itemReservedEXCEPT,
	"EXPLICIT":        itemReservedEXPLICIT,
	"EXPORTS":         itemReservedEXPORTS,
	"EXTENSIBILITY":   itemReservedEXTENSIBILITY,
	"EXTERNAL":        itemReservedEXTERNAL,
	"FALSE":           itemReservedFALSE,
	"FROM":            itemReservedFROM,
	"GeneralString":   itemReservedGeneralString,
	"GeneralizedTime": itemReservedGeneralizedTime,
	"GraphicString":   itemReservedGraphicString,
	"IA5String":       itemReservedIA5String,
	"IDENTIFIER":      itemReservedIDENTIFIER,
	"IMPLICIT":        itemReservedIMPLICIT,
	"IMPLIED":         itemReservedIMPLIED,
	"IMPORTS":         itemReservedIMPORTS,
	"INCLUDES":        itemReservedINCLUDES,
	"INSTANCE":        itemReservedINSTANCE,
	"INTEGER":         itemReservedINTEGER,
	"INTERSECTION":    itemReservedINTERSECTION,
	"ISO646String":    itemReservedISO646String,
	"MAX":             itemReservedMAX,
	"MIN":             itemReservedMIN,
	"MINUS-INFINITY":  itemReservedMINUS_INFINITY,
	"NULL":            itemReservedNULL,
	"NumericString":   itemReservedNumericString,
	"OBJECT":          itemReservedOBJECT,
	"OCTET":           itemReservedOCTET,
	"OF":              itemReservedOF,
	"OPTIONAL":        itemReservedOPTIONAL,
	"Descriptor":      itemReservedObjectDescriptor,
	"PATTERN":         itemReservedPATTERN,
	"PDV":             itemReservedPDV,
	"PLUS-INFINITY":   itemReservedPLUS_INFINITY,
	"PRESETN":         itemReservedPRESENT,
	"PRIVATE":         itemReservedPRIVATE,
	"PrintableString": itemReservedPrintableString,
	"REAL":            itemReservedREAL,
	"RELATIVE-OID":    itemReservedRELATIVE_OID,
	"SEQUENCE":        itemReservedSEQUENCE,
	"SET":             itemReservedSET,
	"SIZE":            itemReservedSIZE,
	"STRING":          itemReservedSTRING,
	"SYNTAX":          itemReservedSYNTAX,
	"T61String":       itemReservedT61String,
	"TAGS":            itemReservedTAGS,
	"TRUE":            itemReservedTRUE,
	"TYPE-IDENTIFIER": itemReservedTYPE_IDENTIFIER,
	"TeletexString":   itemReservedTeletexString,
	"UNION":           itemReservedUNION,
	"UNIQUE":          itemReservedUNIQUE,
	"UNIVERSAL":       itemReservedUNIVERSAL,
	"UTCTime":         itemReservedUTCTime,
	"UTF8String":      itemReservedUTF8String,
	"UniversalString": itemReservedUniversalString,
	"VideotexString":  itemReservedVideotexString,
	"VisibleString":   itemReservedVisibleString,
	"WITH":            itemReservedWITH,
}
