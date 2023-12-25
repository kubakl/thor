package lex

type resolverFn func() *Token

// Lexer converts text into tokens.
type Lexer interface {
	// Generate generates a list of tokens.
	Generate() []*Token
}

type lexer struct {
	idx int

	line    int
	linePos int

	input string
}

// NewLexer returns a new instance of lexer.
func NewLexer(input string) Lexer {
	return &lexer{input: input, linePos: 1, line: 1}
}

func (l *lexer) Generate() []*Token {
	var tokens []*Token
	var resolversByChar = map[byte]resolverFn{
		'\n': l.newLineResolver,
		' ':  l.emptyResolver,
		'\r': l.emptyResolver,
		'\t': l.emptyResolver,
		'!':  l.exclamResolver,
	}

	for l.idx < len(l.input) {
		char := l.input[l.idx]

		token := resolversByChar[char]()
		if token != nil {
			tokens = append(tokens, token)
		}

		l.advance()
	}

	tokens = append(tokens, NewToken(EOF, nil, l.line, l.linePos, l.linePos))
	return tokens
}

func (l *lexer) exclamResolver() *Token {
	if l.peak() == '=' {
		l.advance()
		return NewToken(NEQ, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(EXCLAM, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) simpleResolver(t TokenType) resolverFn {
	return func() *Token {
		return NewToken(t, nil, l.line, l.linePos, l.linePos)
	}
}

func (l *lexer) emptyResolver() *Token {
	return nil
}

func (l *lexer) newLineResolver() *Token {
	l.line++
	return nil
}

func (l *lexer) peak() (value byte) {
	idx := l.idx + 1

	if len(l.input) > idx {
		value = l.input[idx]
	}

	return
}

func (l *lexer) advance() {
	l.idx++
	l.linePos++
}
