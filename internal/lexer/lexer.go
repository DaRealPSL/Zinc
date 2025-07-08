package lexer

type Lexer struct {
	input        string
	position     int  //current char
	readPosition int  //next char
	ch           byte //current char
	line         int
	column       int
}

func New(input string) *Lexer {
	//optional: Add 'column: 1'
	l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCII null = EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
	l.position = l.readPosition
	l.readPosition++
}
