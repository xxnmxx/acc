package token

const (
	// Set token type here.
	NEW   = "NEW"
	EQUAL = "EQUAL"
	DOT   = "DOT"
)

type Token struct {
	tokenType tokenType
	literal   string
}

func (t *Token) Type() tokenType {
	return t.tokenType
}
