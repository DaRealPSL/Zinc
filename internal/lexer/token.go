package lexer

// TokenType represents the type of lexical tokens in Zinc
type TokenType string

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}

const (
	//special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"
	BOOL   = "BOOL"

	//operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	EQ  = "=="
	NEQ = "!="
	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="

	AND = "&&"
	OR  = "||"
	NOT = "!"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//keywords
	KEYWORD = "KEYWORD"
)

// keyword map
var keywords = map[string]TokenType{
	"let":      KEYWORD,
	"const":    KEYWORD,
	"var":      KEYWORD,
	"int":      KEYWORD,
	"float":    KEYWORD,
	"bool":     KEYWORD,
	"string":   KEYWORD,
	"char":     KEYWORD,
	"any":      KEYWORD,
	"void":     KEYWORD,
	"function": KEYWORD,
	"if":       KEYWORD,
	"else":     KEYWORD,
	"for":      KEYWORD,
	"while":    KEYWORD,
	"break":    KEYWORD,
	"continue": KEYWORD,
	"return":   KEYWORD,
	"struct":   KEYWORD,
	"obj":      KEYWORD,
	"new":      KEYWORD,
	"true":     BOOL,
	"false":    BOOL,
	"null":     KEYWORD,
	"import":   KEYWORD,
	"export":   KEYWORD,
	"from":     KEYWORD,
	"as":       KEYWORD,
	"switch":   KEYWORD,
	"case":     KEYWORD,
	"default":  KEYWORD,
	"in":       KEYWORD,
	"match":    KEYWORD,
	"typeof":   KEYWORD,
}

// LookupIdent returns IDENT unless it's a keyword or boolean
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
