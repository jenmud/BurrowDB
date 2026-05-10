package parser

import (
	"strings"
)

// lexStatement is the state function for parsing a statement.
func lexStatement(l *Lexer) stateFunc {
	r := l.next()

	switch {

	// nothing to do here
	case r == -1:
		l.emit(EOF)
		return nil

	// check that we have a space, and advance to the next token ignoring the space.
	case IsWhitespace(r):
		l.ignore()

	// check that we have a letter, and advance the next action to lexIdentifier.
	case IsAlpha(r):
		l.backup() // backup so that we can include the first letter in the identifier
		return lexIdentifier

		// check that we have a number, and advance the next action to lexIdentifier.
	case IsDigit(r):
		l.backup()
		return lexInteger

	// check that we have a double quote, and advance the next action to lexString.
	case r == '"':
		l.ignore() // ignore the opening quote
		return lexString

	// check that we have a left parenthesis, and emit it.
	case r == '(':
		l.emit(LPAREN)

	// check that we have a right parenthesis, and emit it.
	case r == ')':
		l.emit(RPAREN)

	// check that we have a left brace, and emit it.
	case r == '{':
		l.emit(LBRACE)

	// check that we have a right brace, and emit it.
	case r == '}':
		l.emit(RBRACE)

	case r == '[':
		l.emit(LSQUARE)

	case r == ']':
		l.emit(RSQUARE)

	// check that we have a colon, and emit it.
	case r == ':':
		l.emit(COLON)

	// check that we have a comma, and emit it.
	case r == ',':
		l.emit(COMMA)

	// check that we have a comma, and emit it.
	case r == '.':
		l.emit(DOT)
	}

	// continue lexing the statement.
	return lexStatement
}

// lexIdentifier emit an identifier token.
func lexIdentifier(l *Lexer) stateFunc {

	l.acceptRun(func(r rune) bool {
		return IsAlphaNumeric(r)
	})

	word := l.input[l.start:l.pos]

	switch strings.ToUpper(word) {

	case "CREATE", "MATCH", "RETURN":
		l.emit(KEYWORD)

	default:
		l.emit(IDENT)
	}

	return lexStatement
}

// lexString emit an string token.
func lexString(l *Lexer) stateFunc {

	// walk through the string until we find a closing double quote.
	for {
		r := l.next()

		switch r {
		case -1:
			return l.errorf("unexpected end of input in string literal")

		case '"':
			l.backup()
			l.emit(STRING)
			l.next()
			l.ignore()
			return lexStatement

		case '\\':
			l.next() // skip the next character after a backslash
		}

	}

}

func lexInteger(l *Lexer) stateFunc {

	for {
		r := l.next()

		if !IsDigit(r) {
			l.backup()
			l.emit(INTEGER)
			return lexStatement
		}

	}

}
