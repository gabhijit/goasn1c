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
// http://www.itu.int/rec/T-REC-X.681-200207-S/en

package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type itemType int

type Pos int

// an item is a token that is 'emit'ted by the Lexer
type item struct {
	typ  itemType
	pos  Pos
	val  string
	line int
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	case i.typ >= itemReservedABSENT && i.typ <= itemReservedWITH:
		return fmt.Sprintf("<%d:%s:%d>", i.line, i.val, i.typ)
	case len(i.val) > 80:
		return fmt.Sprintf("%d:%.80q...:%d", i.line, i.val, i.typ)
	}
	return fmt.Sprintf("%d:%q:%d", i.line, i.val, i.typ)
}

const (
	dash            = "-"
	capitalLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters    = "abcdefghijklmnopqrstuvwxyz"
	digits          = "0123456789"
	terminalLetters = "{}();,|^! \r\n\t"
	binaryAlphabet  = "01"
	hexAlphabet     = "0123456789abcdefABCDEF"
	quote           = "'"
	whiteSpaces     = " \t\r\n"
	hexString       = hexAlphabet + whiteSpaces + quote
	binaryString    = binaryAlphabet + whiteSpaces + quote
	idLetters       = smallLetters + capitalLetters + digits + dash

	eof = -1
)

// lexer : internal state of tokenizer
type lexer struct {
	name  string
	input string
	pos   Pos
	start Pos
	width Pos
	line  int
	items chan item
	debug bool
	asn1Lexer
}

func (l *lexer) Lex(lval *asn1SymType) int {

	i := l.nextItem()

	if l.debug {
		fmt.Println(i)
	}

	switch t := i.typ; {
	case t >= itemReservedABSENT && t <= itemReservedWITH:
		off := i.typ - itemReservedABSENT
		return Tok_ABSENT + int(off)

	case t == itemModuleReference || t == itemTypeReference:
		lval.str = i.val
		return Tok_TypeReference

	case t == itemCapitalReference:
		lval.str = i.val
		return Tok_CAPITALREFERENCE

	case t == itemValueFieldReference:
		lval.str = i.val
		return Tok_ValueFieldReference

	case t == itemTypeFieldReference:
		lval.str = i.val
		return Tok_TypeFieldReference

	case t == itemAssignment:
		return Tok_Assignment

	case t == itemEllipsis:
		return Tok_Ellipsis

	case t == itemRangeSeparator:
		return Tok_TwoDots

	case t == itemNumber:
		lval.num, _ = strconv.ParseInt(i.val, 10, 64)
		return Tok_Number

	case t == itemIdentifier || t == itemValueReference:
		lval.str = i.val
		return Tok_Identifier

	case t == itemSymbol:
		runes := []rune(i.val)
		return int(runes[0])

	case t == itemCstring:
		lval.str = i.val
		return Tok_CString

	case t == itemBstring:
		lval.str = i.val
		return Tok_BString

	case t == itemHstring:
		lval.str = i.val
		return Tok_HString

	case t == itemEOF:
		return 0
	}

	return 0
}

func (l *lexer) Error(s string) {
	fmt.Println("Error:", s)
}

type stateFn func(*lexer) stateFn

// All the following functions are taken from go/src/text/template/parse/lex.go
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
func lex(name, input string, debug bool) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
		line:  1,
		debug: debug,
	}
	go l.run()
	return l
}

// runs the state machine for the lexer
func (l *lexer) run() {

	for state := lexStart; state != nil; {
		state = state(l)
	}
	close(l.items)
}

// eat up all white spaces
func (l *lexer) consumeWS() {
	for isWhiteSpace(l.peek()) {
		l.next()
	}
	l.start = l.pos
	l.ignore()
}

// consume comment -- followed by '\n' or another '--' on same line
func (l *lexer) consumeComment() bool {

	x := l.accept("-") && l.accept("-")
	if x == false {
		return false
	}
Loop:
	// consumed --
	for {
		switch r := l.next(); {
		case r == '-':
			if l.peek() == '-' {
				l.next()
				return true
			}
		case r == '\n':
			return true
		case r == eof:
			break Loop
		default:
			// pass
		}
	}
	return false
}

// consume c-style comments
func (l *lexer) consumeCstyleComment() bool {

	x := l.accept("/") && l.accept("*")
	if x == false {
		return false
	}
Loop:
	for {
		switch l.next() {
		case '*':
			if l.peek() == '/' {
				l.next()
				return true
			}
		case eof:
			break Loop
		}
	}
	return false
}

// processing of a 'word' in the context of Module Header
// usually, a keyword modulename or oid references
func (l *lexer) processWordModuleHeader(word string) error {

	if err := validateWord(word); err != nil {
		return err
	}

	if val, ok := reservedWordsMap[word]; ok {
		l.emit(val)
		l.start = l.pos
		return nil
	}

	if strings.IndexAny(word, capitalLetters) == 0 {
		l.emit(itemModuleReference)
		l.start = l.pos
		return nil
	}

	// should never come here.
	return errors.New("Unknown error occurred.")
}

// processing of a 'word' in the context of module body
// pretty much everything
func (l *lexer) processWordModuleBody(word string) error {

	if err := validateWord(word); err != nil {
		return err
	}

	if val, ok := reservedWordsMap[word]; ok {
		l.emit(val)
		return nil
	}

	if strings.IndexAny(word, capitalLetters) == 0 {
		if strings.ToUpper(word) == word {
			// All Caps is a valid Object Class or Type, Parser will decide what it is
			l.emit(itemCapitalReference)
		} else {
			l.emit(itemTypeReference)
		}
		return nil
	}

	if strings.IndexAny(word, smallLetters) == 0 {
		l.emit(itemValueReference)
		return nil
	}

	if strings.IndexAny(word, "&@") == 0 {
		if strings.IndexAny(word, capitalLetters) == 1 {
			l.emit(itemTypeFieldReference)
		} else if strings.IndexAny(word, smallLetters) == 1 {
			l.emit(itemValueFieldReference)
		}
		return nil
	}

	// try a number
	_, err := strconv.ParseInt(word, 10, 64)

	if err != nil {
		return err
	} else {
		l.emit(itemNumber)
		return nil
	}

}

// process word in the oid context
// FIXME: may be this is not required (it's up to the parser to decide whether this is right oid
func (l *lexer) processWordOid(word string) error {

	if err := validateWord(word); err != nil {
		return err
	}

	if strings.IndexAny(word, smallLetters) == 0 {
		l.emit(itemIdentifier)
		l.start = l.pos
		return nil
	}

	// should never come here.
	return errors.New("Unknown error occurred.")
}

// process word in the oid context
// FIXME: may be this is not required (it's up to the parser to decide whether this is right oid)
func (l *lexer) processOid() error {

	for {
		switch r := l.next(); {
		case r == '{':
			l.emit(itemSymbol)
		case unicode.IsLower(r):
			l.acceptRun(idLetters)
			word := l.input[l.start:l.pos]
			err := l.processWordOid(word)
			if err != nil {
				return errors.New("lex error")
			}
		case unicode.IsDigit(r):
			l.acceptRun(digits)
			l.emit(itemNumber)
			s := l.peek()
			if isWhiteSpace(s) || s == ')' {
				l.start = l.pos
			} else {
				return errors.New("number in oid should be followed by whitespace or ')'")
			}
		case r == '(':
			l.emit(itemSymbol)
		case r == ')':
			l.emit(itemSymbol)
		case r == '}':
			l.emit(itemSymbol)
			return nil
		default:
			return errors.New("Unknown char found.")
		}
		l.consumeWS()
	}
	// commenting below thanks to `go vet`, but not sure
	//return errors.New("Unknown error occurred.")
}

// special handling of ":" character ('cos it could mean assignment or a standalone token or a separator (parametarized type)
func (l *lexer) handleColon() error {

	l.acceptRun(":=")
	w := l.input[l.start:l.pos]

	switch len(w) {
	case 1:
		l.emit(itemSymbol)
		l.start = l.pos
		return nil
	case 2:
		return errors.New("unaccepted lexical item " + w)
	case 3:
		if w == "::=" {
			l.emit(itemAssignment)
			l.start = l.pos
			return nil
		} else {
			return errors.New("unaccepted lexical item " + w)
		}
	default:
		return errors.New("unaccepted lexical item " + w)
	}

}

// special handling of "." could be a separator(single .) a range separator (..) or ellipsis (...)
func (l *lexer) handleDot() error {

	l.acceptRun(".")
	w := l.input[l.start:l.pos]

	switch len(w) {
	case 1:
		l.emit(itemSymbol)
	case 2:
		l.emit(itemRangeSeparator)
	case 3:
		l.emit(itemEllipsis)
	default:
		return errors.New("too many dots.")
	}
	l.start = l.pos
	l.consumeWS()

	return nil
}

// left/right brackets two consecutive brackets have special meaning '[[' and ']]'
func (l *lexer) handleLeftRightBracket(b rune) error {

	s := string(b)
	l.acceptRun(s)
	w := l.input[l.start:l.pos]

	switch len(w) {
	case 1:
		l.emit(itemSymbol)
	case 2:
		if b == '[' {
			l.emit(itemTwoLeftBracket)
		} else {
			l.emit(itemTwoRightBracket)
		}
	default:
		return errors.New("wrong number of brakcets " + w)
	}
	l.start = l.pos
	l.consumeWS()

	return nil
}

// handing of "-". Don't think this is 100% right yet, but most of the test cases pass so far pass
func (l *lexer) handleDash() error {

	l.acceptRun("-")
	w := l.input[l.start:l.pos]

	if len(w) == 2 {
		l.start = l.pos
	Loop:
		for {
			switch r := l.next(); {
			case r == '\n':
				l.start = l.pos
				return nil
			case r == '-':
				if l.peek() == '-' {
					l.next()
					l.start = l.pos
					return nil
				}
			case r == eof:
				break Loop
			}
		}
		return errors.New("unterminated comment")
	} else if len(w) == 1 {
		l.emit(itemSymbol)
		l.start = l.pos
		l.consumeWS()
		return nil
	}

	return errors.New("error processing dash")
}

// handling of "/".
func (l *lexer) handleSlash() error {

	var w string

	count := 1
	if l.accept("/") {
		if l.accept("*") {
			w = l.input[l.start:l.pos]
		} else {
			l.emit(itemSymbol)
			return nil
		}
	}

	if len(w) == 2 {
		l.start = l.pos
	Loop:
		for count > 0 {
			switch r := l.next(); {
			case r == '*':
				if l.peek() == '/' {
					count--
					l.next()
					l.start = l.pos
				}

			case r == '/':
				if l.peek() == '*' {
					l.next()
					count++
				}
			case r == eof:
				break Loop
			}
		}
	}

	if count == 0 {
		return nil
	}

	return errors.New("Unterminated C++ style comment")
}

// handle '\'' single quote (usually means handling binarystring and hexString
func (l *lexer) handleSingleQuote() error {

	save := l.pos
	l.accept("'")

	l.acceptRun(binaryString)

	if l.peek() == 'B' {
		l.next()
		l.emit(itemBstring)
		return nil
	} else {
		l.pos = save
		l.start = save
		l.acceptRun(hexString)

		if l.peek() == 'H' {
			l.next()
			l.emit(itemHstring)
			return nil

		}
		return errors.New("invalid bstring or hstring")
	}
}

// handle the double quote "\"". Note doesn't handle "" yet
func (l *lexer) handleDoubleQuote() error {
	l.accept("\"")
Loop:
	for {
		switch r := l.next(); {
		case r == '"':
			s := l.peek()
			if s == '"' {
				l.next()
				continue
			} else {
				l.emit(itemCstring)
				l.start = l.pos
				return nil
			}
		case r == eof:
			break Loop
		}
	}
	return errors.New("unterminated double quote.")
}

// make sure it's a valid word (starts with a letter or number, does not end in a dash and no two consequtive dashes
func validateWord(word string) error {

	last := word[len(word)-1]

	// Interpretation of the word is responsibility of the parser. We make sure that the word is what constitutes a
	// 1. valid identifier, typereference, valuereference, objectreference, number and so on

	// FIXME: We need to implement scan number (or reuse) but this should not fail word validation
	/* if strings.IndexAny(word, digits) == 0 {
		return errors.New("A word cannot start with a digit")
	} */

	if last == '-' {
		return errors.New("A Word cannot end with a -.")
	}

	if strings.Index(word, "--") >= 0 {
		return errors.New("A word cannot have a dash followed by a dash (--).")
	}

	if strings.Index(word, "&") > 0 {
		return errors.New("& not allowed anywhere other than beginning of word.")
	}

	return nil
}

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isTerminator(r rune) bool {
	return strings.ContainsRune(terminalLetters, r)
}

// start of a module(ish) thing
// FIXME: This should be folded into handleModuleHeader and handleModuleBody
func lexStart(l *lexer) stateFn {

	l.consumeWS()

	if l.next() == eof {
		l.emit(itemEOF)
		return nil
	}
	l.backup()

	if l.peek() == '-' {
		if x := l.consumeComment(); x == false {
			// FIXME : return error here
			return nil
		} else {
			l.start = l.pos
			l.ignore()
		}
		return lexStart
	}

	if l.peek() == '/' {
		if l.consumeCstyleComment() == false {
			// FIXME: return error here
			return nil
		} else {
			l.start = l.pos
			l.ignore()
		}
		return lexStart
	}

	// Now we expect typereference like module identifier
	if l.accept(capitalLetters) {
		return lexModuleHeader
	} else {
		return nil
	}

}

// module header state
// FIXME : this needs to be in similar structure as module body
func lexModuleHeader(l *lexer) stateFn {

	l.backup()
Loop:
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
					l.consumeWS()
				}
			case r == ':':
				l.backup()
				err := l.handleColon()
				if err != nil {
					return l.errorf("lex error")
				} else {
					l.consumeWS()
					// We are gone past the ::=
					return lexModuleBody
				}
			case isWhiteSpace(r):
				l.backup()
				word := l.input[l.start:l.pos]
				err := l.processWordModuleHeader(word)
				if err != nil {
					return l.errorf("lex error")
				} else {
					l.consumeWS()
				}
			case r == eof:
				l.emit(itemEOF)
				break Loop
			}
		}

	}
	return nil

}

// module body state - pretty much everything happens here.
func lexModuleBody(l *lexer) stateFn {

	accum := false
	l.consumeWS()

Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			accum = true
		case r == '&':
			accum = true
		case r == '@':
			accum = true
		case r == '-':
			if l.peek() == '-' {
				l.backup()
				err := l.handleDash()
				if err != nil {
					return l.errorf("lexer error 1")
				}
			} else {
				accum = true
			}
		case r == '\'':
			l.backup()
			err := l.handleSingleQuote()
			if err != nil {
				return l.errorf("lexer error")
			}
		case r == '/':
			l.backup()
			err := l.handleSlash()
			if err != nil {
				return l.errorf("lexer error")
			}
		case r == '"':
			l.backup()
			err := l.handleDoubleQuote()
			if err != nil {
				return l.errorf("lexer error")
			}
		case r == eof:
			break Loop
		default:
			l.backup()
			if accum {
				word := l.input[l.start:l.pos]
				err := l.processWordModuleBody(word)
				if err != nil {
					return l.errorf("lexer error")
				}
				if word == "END" {
					return lexStart
				}

				l.start = l.pos
			}
			switch {
			case r == '.':
				err := l.handleDot()
				if err != nil {
					return l.errorf("lexer error")
				}
			case r == ':':
				err := l.handleColon()
				if err != nil {
					return l.errorf("lex error")
				}
			case r == '[' || r == ']':
				err := l.handleLeftRightBracket(r)
				if err != nil {
					return l.errorf("lex error")
				}
			/* case r == '-':
			err := l.handleDash()
			if err != nil {
				return l.errorf("lex error")
			} */
			case isTerminator(r):
				l.next()
				if !isWhiteSpace(r) {
					l.emit(itemSymbol)
					l.start = l.pos
				}
			case r == eof:
				l.emit(itemEOF)
				break
			}
			l.consumeWS()
			accum = false
		}
	}
	return nil
}

// items from X.680 specification - TODO: review this
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
	itemReservedBIT
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

	itemCapitalReference
	// X.681 Section 7
	itemObjectClassReference
	itemObjectReference
	itemObjectSetReference
	itemTypeFieldReference
	itemValueFieldReference
	itemValueSetFieldReference
	itemObjectFieldReference
	itemObjectSetFieldReference

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
	"BIT":             itemReservedBIT,
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
