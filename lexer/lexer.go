package lexer

import "github.com/OPC-16/Monkey/token"

type Lexer struct {
    input        string
    position     int    // current position in input (points to current char)
    readPosition int    // current reading position in input (after current char)
    ch           byte   // current char under examination
}

// returns a new lexer
func New(input string) *Lexer {
    l := &Lexer{ input: input}
    l.readChar()
    return l
}

// gives the next character and advances position in input string
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0        // 0 is ASCII code for "NUL"
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhiteSpace()

    switch l.ch {
        case '=':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)
                tok = token.Token{ Type: token.EQ, Literal: literal}
            } else {
                tok = newToken(token.ASSIGN, l.ch)
            }
        case '+':
            tok = newToken(token.PLUS, l.ch)
        case '-':
            tok = newToken(token.MINUS, l.ch)
        case '!':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)
                tok = token.Token{ Type: token.NOT_EQ, Literal: literal}
            } else {
                tok = newToken(token.BANG, l.ch)
            }
        case '/':
            tok = newToken(token.SLASH, l.ch)
        case '*':
            tok = newToken(token.ASTERISK, l.ch)
        case '<':
            tok = newToken(token.LT, l.ch)
        case '>':
            tok = newToken(token.GT, l.ch)
        case ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case '(':
            tok = newToken(token.LPAREN, l.ch)
        case ')':
            tok = newToken(token.RPAREN, l.ch)
        case ',':
            tok = newToken(token.COMMA, l.ch)
        case '{':
            tok = newToken(token.LBRACE, l.ch)
        case '}':
            tok = newToken(token.RBRACE, l.ch)
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        default:
            if isLetter(l.ch) {
                tok.Literal = l.readIdentifier()
                tok.Type = token.LookupIdent(tok.Literal)
                return tok

            } else if isDigit(l.ch) {
                tok.Type = token.INT
                tok.Literal = l.readNumber()
                return tok

            } else {
                tok = newToken(token.ILLEGAL, l.ch)
            }
    }

    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{ Type: tokenType, Literal: string(ch)}
}

// it reads in an identifier and advances our lexer’s positions until it encounters a non-letter-character.
func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

// reads only integer number.
// for sake of simplicity we don't support floats, or numbers in hex notation or in octal notation
func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

// checks whether the given argument is a letter.
// this function is important because changing it determines the language our interpreter will be able to parse.
// because of `ch == '_'` check, we can use variables with names like foo_bar
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhiteSpace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    } 
}

// it is really similar to readChar(), except that it doesn’t increment l.position and l.readPosition.
func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}
