package lexer

import "github.com/pippokairos/pizza/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var nextToken token.Token

	switch l.ch {
	case '=':
		nextToken = newToken(token.ASSIGN, l.ch)
	case '+':
		nextToken = newToken(token.PLUS, l.ch)
	case '(':
		nextToken = newToken(token.LPAREN, l.ch)
	case ')':
		nextToken = newToken(token.RPAREN, l.ch)
	case '{':
		nextToken = newToken(token.LBRACE, l.ch)
	case '}':
		nextToken = newToken(token.RBRACE, l.ch)
	case ',':
		nextToken = newToken(token.COMMA, l.ch)
	case ';':
		nextToken = newToken(token.SEMICOLON, l.ch)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		if isLetter(l.ch) {
			nextToken.Literal = l.readIdentifierOrLiteral(isLetter)
			nextToken.Type = token.LookupIdent(nextToken.Literal)

			return nextToken // Needed to avoid reading the next character again
		} else if isDigit(l.ch) {
			nextToken.Literal = l.readIdentifierOrLiteral(isDigit)
			nextToken.Type = token.INT

			return nextToken
		} else {
			nextToken = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return nextToken
}

func newToken(tokenType token.TokenType, tokenLiteral byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(tokenLiteral)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	startPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) readIdentifierOrLiteral(identifyingFunc func(byte) bool) string {
	startPosition := l.position
	for identifyingFunc(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}
