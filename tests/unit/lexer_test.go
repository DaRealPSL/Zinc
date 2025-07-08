package unit

import (
	"testing"

	"github.com/DaRealPSL/zinc/internal/lexer"
)

func TestLexerBasicTokens(t *testing.T) {
	input := `
	let age = 42;
	string name = "Pepijn";
	bool happy = true;
	if (age > 10) {
		print("You're awesome!");
	}
	`

	expectedTypes := []lexer.TokenType{
		lexer.KEYWORD, lexer.IDENT, lexer.ASSIGN, lexer.INT, lexer.SEMICOLON,
		lexer.KEYWORD, lexer.IDENT, lexer.ASSIGN, lexer.STRING, lexer.SEMICOLON,
		lexer.KEYWORD, lexer.IDENT, lexer.ASSIGN, lexer.BOOL, lexer.SEMICOLON,
		lexer.KEYWORD, lexer.LPAREN, lexer.IDENT, lexer.GT, lexer.INT, lexer.RPAREN,
		lexer.LBRACE,
		lexer.IDENT, lexer.LPAREN, lexer.STRING, lexer.RPAREN, lexer.SEMICOLON,
		lexer.RBRACE,
		lexer.EOF,
	}

	l := lexer.New(input)

	for i, expected := range expectedTypes {
		tok := l.NextToken()
		t.Logf("token[%d] = %+v\n", i, tok)
		if tok.Type != expected {
			t.Fatalf("test[%d] - token type wrong. expected=%q, got=%q", i, expected, tok.Type)
		}
	}
}

func TestLexerIgnoresComments(t *testing.T) {
	input := `
		// this is a single-line comment
		let x = 10; /* and this is
		a multiline comment */
		let y = 20; // another comment
	`

	expectedTypes := []lexer.TokenType{
		lexer.KEYWORD, lexer.IDENT, lexer.ASSIGN, lexer.INT, lexer.SEMICOLON,
		lexer.KEYWORD, lexer.IDENT, lexer.ASSIGN, lexer.INT, lexer.SEMICOLON,
		lexer.EOF,
	}

	l := lexer.New(input)

	for i, expected := range expectedTypes {
		tok := l.NextToken()
		t.Logf("token[%d] = %+v\n", i, tok)
		if tok.Type != expected {
			t.Fatalf("test[%d] - expected=%q, got=%q", i, expected, tok.Type)
		}
	}
}
