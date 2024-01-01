package lex

type resolverFn func() *Token

// Lexer converts text into tokens.
type Lexer interface {
	// Generate generates a list of tokens.
	Generate() Iterator[*Token]
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

func (l *lexer) Generate() Iterator[*Token] {
	var tokens []*Token
	for l.idx < len(l.input) {
		char := l.input[l.idx]

		token := l.resolvers()[char]()
		if token != nil {
			tokens = append(tokens, token)
		}

		l.advance()
	}

	tokens = append(tokens, NewToken(EOF, nil, l.line, l.linePos, l.linePos))
	return NewIterator(tokens)
}

func (l *lexer) resolvers() map[byte]resolverFn {
	return map[byte]resolverFn{
		'\n': l.newLineResolver,
		' ':  l.emptyResolver,
		'\r': l.emptyResolver,
		'\t': l.emptyResolver,
		'!':  l.exclamResolver,
		'%':  l.simpleResolver(MOD),
		'^':  l.simpleResolver(POW),
		'&':  l.ampResolver,
		'*':  l.astResolver,
		'(':  l.simpleResolver(LPAREN),
		')':  l.simpleResolver(RPAREN),
		'-':  l.minusResolver,
		'+':  l.plusResolver,
		'=':  l.eqResolver,
		':':  l.simpleResolver(COLON),
		';':  l.simpleResolver(SEMICOL),
		',':  l.simpleResolver(COMMA),
		'.':  l.simpleResolver(DOT),
		'/':  l.divResolver,
		'{':  l.simpleResolver(LBRAC),
		'}':  l.simpleResolver(RBRAC),
		'[':  l.simpleResolver(LSQUARE),
		']':  l.simpleResolver(RSQUARE),
		'>':  l.gtltResolver(GT),
		'<':  l.gtltResolver(LT),
		'|':  l.orResolver,
	}
}

func (l *lexer) exclamResolver() *Token {
	if l.peak() == '=' {
		l.advance()
		return NewToken(NEQ, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(EXCLAM, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) ampResolver() *Token {
	if l.peak() == '&' {
		l.advance()
		return NewToken(AND, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(AMP, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) astResolver() *Token {
	if l.peak() == '=' {
		l.advance()
		return NewToken(TIMESEQ, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(AST, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) divResolver() *Token {
	if l.peak() == '=' {
		l.advance()
		return NewToken(DIVEQ, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(DIV, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) minusResolver() *Token {
	if l.peak() == '-' {
		l.advance()
		return NewToken(MINMIN, nil, l.line, l.linePos-1, l.linePos)
	}

	if l.peak() == '=' {
		l.advance()
		return NewToken(MINUSEQ, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(MINUS, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) plusResolver() *Token {
	if l.peak() == '+' {
		l.advance()
		return NewToken(PLUSPLUS, nil, l.line, l.linePos-1, l.linePos)
	}

	if l.peak() == '=' {
		l.advance()
		return NewToken(PLUSEQ, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(PLUS, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) eqResolver() *Token {
	if l.peak() == '=' {
		l.advance()
		return NewToken(EQ, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(ASSIGN, nil, l.line, l.linePos, l.linePos)
}

func (l *lexer) gtltResolver(tok TokenType) resolverFn {
	return func() *Token {
		if l.peak() == '=' {
			l.advance()
			return NewToken(tok+1, nil, l.line, l.linePos-1, l.linePos)
		}

		return NewToken(tok, nil, l.line, l.linePos, l.linePos)
	}
}

func (l *lexer) orResolver() *Token {
	if l.peak() == '|' {
		l.advance()
		return NewToken(OR, nil, l.line, l.linePos-1, l.linePos)
	}

	return NewToken(ILLEGAL, nil, l.line, l.linePos, l.linePos)
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
