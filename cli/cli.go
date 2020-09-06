package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// AST
type Node interface {
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type newStatement struct {
}

func (ns *newStatement) String() string {}
func (ns *newStatement) statementNode() {}

// Token
type tokenType int

const (
	_ tokenType = iota
	tokenNew
	tokenInput
	tokenDot
	tokenAssign
	tokenIdentifier
)

type token struct {
	typ  tokenType
	text string
}

func newToken(typ tokenType, text string) Token {
	return Token{typ: typ, text: text}
}

// Can be generalize like regKeyword(s string,typ tokenType, fn())?
var keywords map[string]tokenType = map[string]tokenType{
	"new": tokenNew,
	"in":  tokenInput,
}

func lookupKeywords(kw string) tokenType {
	typ, ok := keywords[kw]
	if ok {
		return typ
	}
	return tokenIdentifier
}

// Lexer
const EofRune rune = -1

type lexer struct {
	rd       io.RuneReader
	peeking  bool
	peekRune rune
	last     rune
	buf      bytes.Buffer
}

func newLexer(rd io.RuneReader) *lexer {
	return &lexer{
		rd: rd,
	}
}

func (l *lexer) skipSpace() rune {
	comment = false
	for {
		r := l.read()
		if r == '\n' || r == EofRune {
			return r
		}
		if r == ';' {
			comment = true
			continue
		}
		if !comment && !isSpace(r) {
			l.back(r)
			return r
		}
	}
}

func (l *lexer) skipToNewline() {
	for l.last != '\n' && l.last != EofRune {
		l.nextRune()
	}
	l.peeking = false
}

// wip
func (l *lexer) next() *token {
	for {
		r := l.read()
		switch {
		case unicode.IsLetter(r):

		default:
		}
	}
}

func (l *lexer) read() rune {
	if l.peeking {
		l.peeking = false
		return l.peekRune
	}
	return l.nextRune()
}

func (l *lexer) nextRune() rune {
	r, _, err := l.rd.ReadRune()
	if err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr)
		}
		r = EofRune
	}
	l.last = r
	return r
}

func (l *lexer) peek() rune {
	if l.peeking {
		return l.peekRune
	}
	r := l.read()
	l.peeking = true
	l.peekRune = r
	return r
}

func (l *lexer) back(r rune) {
	l.peeking = true
	l.peekRune = r
}

func (l *lexer) accum(r rune, valid func(rune) bool) {
	l.buf.Reset()
	for {
		l.buf.WriteRune(r)
		r := l.read()
		if r == EofRune {
			return
		}
		if !valid(r) {
			l.back(r)
			return
		}
	}
}

//func (l *lexer) endToken() {
//	if r := l.peek(); isAlphanum(r) || !isSpace(r) && r != '(' && r != ')' && r != '.' && r != EofRune {
//		errorf("invalid token after %s", &l.buf)
//	}
//}

// Can be replaced by unicode.IsSpace?
func isSpace(r rune) {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// Can be replaced by unicode.IsDigit?
func isNumber(r rune) {
	return '0' <= r && r <= '9'
}

func isFloat(r rune) {
}

func isAlphanum(r rune) bool {
}
