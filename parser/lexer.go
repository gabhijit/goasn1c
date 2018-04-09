// Lexer using ideas from
// https://golang.org/src/text/template/parse/lex.go
// Copyright Go Authors

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

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type itemType int

type Pos int

type item struct {
	typ  itemType
	pos  Pos
	val  string
	line int
}

const (
	dash      = "-"
	slashStar = "/*"
	starSlash = "*/"

	capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters   = "abcdefghijklmnopqrstuvwxyz"
	digits         = "0123456789"
	capitalWords   = capitalLetters + digits + dash

	eof = -1
)

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	case i.typ >= itemReservedABSENT && i.typ <= itemReservedWITH:
		return fmt.Sprintf("<%s>", i.val)
	case len(i.val) > 20:
		return fmt.Sprintf("%.20q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

type stateFn func(*lexer) stateFn

type lexer struct {
	name   string
	input  string
	pos    Pos
	start  Pos
	width  Pos
	line   int
	items  chan item
	states []stateFn
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
	// Correct newline count.
	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos], l.line}
	// Some items contain text internally. If so, count their newlines.
	switch t {
	case itemWSpace, itemComment:
		l.line += strings.Count(l.input[l.start:l.pos], "\n")
	}
	l.start = l.pos
}

// consume is just like emit, but doesn't actually emit
func (l *lexer) consume(t itemType) {
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.line += strings.Count(l.input[l.start:l.pos], "\n")
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...), l.line}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() item {
	return <-l.items
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	for range l.items {
	}
}

// lex creates a new scanner for the input string.
func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
		line:  1,
	}
	go l.run()
	return l
}

// states are -
// 1. lexStart
// 2. lexStartModuleHeader
// 3. lexStartModuleBody
// 4. lexAfterModule

// runs the state machine for the lexer
func (l *lexer) run() {

	for state := lexStart; state != nil; {
		state = state(l)
	}
	close(l.items)
	fmt.Println("closed")
}

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func (l *lexer) consumeWhitespace() {
	for isWhiteSpace(l.peek()) {
		l.next()
	}
	l.ignore()
}

func (l *lexer) consumeComment() bool {

	x := l.accept("-") && l.accept("-")
	if x == false {
		return false
	}
	// consumed --
	for {
		fmt.Println(l.start, l.pos)
		switch r := l.next(); {
		case r == '-':
			if l.peek() == '-' {
				l.next()
				return true
			}
		case r == '\n':
			return true
		default:
			fmt.Println(r)
		}
	}
	return false
}

func (l *lexer) consumeCstyleComment() bool {

	x := l.accept("/") && l.accept("*")
	if x == false {
		return false
	}
	for {
		switch l.next() {
		case '*':
			if l.peek() == '/' {
				l.next()
				return true
			}
		}
	}
	return false
}

func lexStart(l *lexer) stateFn {

	l.consumeWhitespace()

	if l.peek() == '-' {
		if x := l.consumeComment(); x == false {
			// FIXME : return error here
			fmt.Println(x)
			return nil
		} else {
			l.ignore()
		}
		//fmt.Println(l.start, l.pos)
		return lexStart
	}

	if l.peek() == '/' {
		if l.consumeCstyleComment() == false {
			// FIXME: return error here
			return nil
		} else {
			l.ignore()
		}
		return lexStart
	}

	// Now we expect typereference like module identifier
	if l.accept(capitalLetters) {
		return stateModuleHeader
	} else {
		return nil
	}

	return nil

}

func (l *lexer) processOid() error {

	for {
		r := l.next()
		if r == '}' {
			l.backup()
			break
		}
	}

	r := l.next()
	if r == '}' {
		l.emit(itemOID)
		l.start = l.pos
		return nil
	} else {
		return errors.New("expected OID not terminated.")
	}

}

func (l *lexer) processWord(word string, state lexerState) error {

	last := word[len(word)-1]

	// Word can be
	// 1. A word starting with capital letter
	// 2. A keyword
	// 3. A word starting with small letter

	// depending upon state we are in, we emit different tokens

	if strings.IndexAny(word, digits) == 0 {
		return errors.New("A word cannot start with a digit")
	}

	if last == '-' {
		return errors.New("A Word cannot end with a -.")
	}

	if strings.Index(word, "--") >= 0 {
		return errors.New("A word cannot have a dash followed by a dash (--).")
	}

	if val, ok := reservedWordsMap[word]; ok {
		l.emit(val)
		l.start = l.pos
		return nil
	}

	if strings.IndexAny(word, capitalLetters) == 0 {
		if state == lexerStateModuleHeader {
			l.emit(itemModuleReference)
		} else {
			l.emit(itemTypeReference)
		}
		l.start = l.pos
		return nil
	} else {
		// FIXME: verify
		l.emit(itemValueReference)
		l.start = l.pos
		return nil
	}

	// should never come here.
	return errors.New("Unknown error occurred.")
}

func (l *lexer) processAssignment() error {

	l.acceptRun(":=")

	t := l.input[l.start:l.pos]

	if t == "::=" {
		l.emit(itemAssignment)
		l.start = l.pos
		return nil
	}

	return errors.New("invalid assignment found.")
}

func stateModuleHeader(l *lexer) stateFn {

	state := lexerStateModuleHeader

	l.backup()
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// keep accumulating
		default:
			switch {
			case r == '{':
				l.backup()
				err := l.processOid()
				if err != nil {
					return l.errorf("lex error")
				} else {
					l.consumeWhitespace()
				}
			case r == ':':
				l.backup()
				err := l.processAssignment()
				if err != nil {
					return l.errorf("lex error")
				} else {
					l.consumeWhitespace()
				}
			case isWhiteSpace(r):
				l.backup()
				word := l.input[l.start:l.pos]
				// Word can be a keyword
				fmt.Println("before: %d:%d", l.start, l.pos)
				err := l.processWord(word, state)
				fmt.Println("after: %d:%d", l.start, l.pos)
				if err != nil {
					return l.errorf("lex error")
				} else {
					l.consumeWhitespace()
				}
			case r == eof:
				break
			}
		}

	}
	return nil

}

type lexerState int

const (
	lexerStateStart lexerState = iota
	lexerStateModuleHeader
	lexerStateModuleBody
)

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

	itemOID
	itemWSpace
	itemEOF
)

var reservedWordsMap = map[string]itemType{
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
