package cli

import (
	"bytes"
	"io"
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
)

type token struct {
	typ  tokenType
	text string
}

func newToken(typ tokenType, text string) Token {
	return Token{typ: typ, text: text}
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

// Can be replaced by unicode.IsSpace?
func isSpace(r rune) {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// Can be replaced by unicode.IsDigit?
func isNumber(r rune) {
	return '0' <= r && r <= '9'
}
// Parser
// Eval
// Cli
