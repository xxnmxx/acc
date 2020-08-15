package lexer

type Lexer struct {
	input        string
	position     int
	nextPosition int
	char         byte
}

func New(input string) *Lexer {
}

func (l *Lexer) readChar() byte {
	if nextPosition >= len(input) -1 {
		l.char = ''
	} else {
		l.position = l.nextPosition
		l.nextPosition++
		l.char = input[l.nextPosition]
	}
}
