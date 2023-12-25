package lex

type TokenType int8

const (
	EOF TokenType = iota
	IDENT
	KEYWORD
	ILLEGAL

	INT
	FLOAT
	STRING
	CHAR
	BOOL

	EXCLAM
	MOD
	POW
	AMP
	AST
	LPAREN
	RPAREN
	MINUS
	PLUS
	ASSIGN
	COLON
	SEMICOL
	COMMA
	DOT
	DIV
	LBRAC
	RBRAC
	LSQUARE
	RSQUARE

	EQ
	NEQ
	GT
	GTEQ
	LT
	LTEQ
	AND
	OR
)

// Token represents a token.
type Token struct {
	Type  TokenType
	Value any

	Line  int
	Start int
	End   int
}

// NewToken returns a new instance of a token.
func NewToken(t TokenType, val any, line, start, end int) *Token {
	return &Token{
		Type:  t,
		Value: val,
		Line:  line,
		Start: start,
		End:   end,
	}
}
