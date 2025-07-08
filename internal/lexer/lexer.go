package lexer

import "unicode"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position (after current char)
	ch           byte // current char under examination
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0 //null byte for EOF
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token
	tok.Line = l.line
	tok.Column = l.column

	switch l.ch {
	case '=':
		tok = l.newToken(ASSIGN, string(l.ch))
	case '+':
		tok = l.newToken(PLUS, string(l.ch))
	case '-':
		tok = l.newToken(MINUS, string(l.ch))
	case '*':
		tok = l.newToken(ASTERISK, string(l.ch))
	case '/':
		tok = l.newToken(SLASH, string(l.ch))
	case '>':
		tok = l.newToken(GT, string(l.ch))
	case '<':
		tok = l.newToken(LT, string(l.ch))
	case ';':
		tok = l.newToken(SEMICOLON, string(l.ch))
	case ',':
		tok = l.newToken(COMMA, string(l.ch))
	case '(':
		tok = l.newToken(LPAREN, string(l.ch))
	case ')':
		tok = l.newToken(RPAREN, string(l.ch))
	case '{':
		tok = l.newToken(LBRACE, string(l.ch))
	case '}':
		tok = l.newToken(RBRACE, string(l.ch))
	case '"':
		tok.Type = STRING
		tok.Lexeme = l.readString()
		return tok // keep this: fixed double-skip
	case 0:
		tok.Type = EOF
		tok.Lexeme = ""
	default:
		if isLetter(l.ch) {
			tok.Lexeme = l.readIdentifier()
			tok.Type = LookupIdent(tok.Lexeme)
			return tok
		} else if isDigit(l.ch) {
			tok.Type, tok.Lexeme = l.readNumber()
			return tok
		} else {
			tok = l.newToken(ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) newToken(tokenType TokenType, ch string) Token {
	return Token{
		Type:   tokenType,
		Lexeme: ch,
		Line:   l.line,
		Column: l.column,
	}
}

func (l *Lexer) skipWhitespace() {
	for {
		if l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
			l.readChar()
		} else if l.ch == '/' && l.peekChar() == '/' {
			//skip single-line comments
			l.readChar() //consume first '/'
			l.readChar() //consume second '/'
			for l.ch != '\n' && l.ch != 0 {
				l.readChar()
			}
		} else if l.ch == '/' && l.peekChar() == '*' {
			//skip multi-line comments
			l.readChar() //consume first '/'
			l.readChar() //consume '*'
			for {
				if l.ch == 0 {
					break //we've hit an unexpected EOF
				}
				if l.ch == '*' && l.peekChar() == '/' {
					l.readChar() //consume '*'
					l.readChar() //consume '/'
					break
				}
				l.readChar()
			}
		} else {
			break
		}
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() (TokenType, string) {
	pos := l.position
	var tokType TokenType = INT

	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' {
		tokType = FLOAT
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return tokType, l.input[pos:l.position]
}

func (l *Lexer) readString() string {
	start := l.position + 1 // skip opening quote
	l.readChar()

	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}

	str := l.input[start:l.position]
	l.readChar() // skip closing quote
	return str
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
