package lexer

import "github.com/kidzrockboom/Galain/token"

type Lexer struct {
    input string
    position int  // current position in input ( points to current char)
    readPosition int // current reading position in input (after current char)
    char byte // current char under examination
}

func New(input string) *Lexer {
    lexer := &Lexer {input: input}
    lexer.readChar()
    return lexer
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.char = 0
    } else {
        l.char = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token 

    switch l.char {
    case '=':
        tok = newToken(token.ASSIGN, l.char)
    case ';':
        tok = newToken(token.SEMICOLON, l.char)
    case '(':
        tok = newToken(token.LEFTPAREN, l.char)
    case ')': 
        tok = newToken(token.RIGHTPAREN, l.char)
    case ',':
        tok = newToken(token.COMMA, l.char)
    case '+':
        tok = newToken(token.PLUS, l.char)
    case '{':
        tok = newToken(token.LEFTBRACE, l.char)
    case '}':
        tok = newToken(token.RIGHTBRACE, l.char)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }

    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(char)}
}
