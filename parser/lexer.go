package parser

import "unicode/utf8"

// stateFunc is a function that represents a state in the lexer.
type stateFunc func(*Lexer) stateFunc

// Lexer is a struct that holds the state of the lexer and provides methods for lexing the query input.
type Lexer struct {
	input  string
	start  int
	pos    int
	width  int
	state  stateFunc
	Tokens chan Token // output channel for tokens, make this buffered to avoid blocking
}

// NewLexer creates a new Lexer with the given input string.
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  input,
		Tokens: make(chan Token, 2),
	}

	go l.run()

	return l
}

func (l *Lexer) run() {
	// start running with the initial state and run until the state is nil (end of input)
	for state := lexStatement; state != nil; {
		state = state(l)
	}

	// make sure that you close the channel so that things shutdown.
	close(l.Tokens)
}

// next returns the next rune in the input string and advances the lexer's position.
func (l *Lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return -1
	}

	// extract the char as a rune value, and advance the position.
	r := rune(l.input[l.pos])
	r, w := utf8.DecodeRuneInString(l.input[l.pos:]) // this will set the width of the rune
	l.width = w
	l.pos += w

	return r
}

// ignore ignores the current token and resets the start position to the current position.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// peek returns the next rune in the input string without advancing the lexer's position.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup backs up the lexer's position by one rune.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// acceptRun takes a validation function and will continue advancing till the validation returns False.
func (l *Lexer) acceptRun(valid func(rune) bool) {
	for valid(l.next()) {
	}
	l.backup()
}

// emit emits a token of the given type to the output channel.
func (l *Lexer) emit(t TokenType) {
	token := Token{
		Type:  t,
		Value: l.input[l.start:l.pos],
	}

	// send the token over the channel
	l.Tokens <- token

	// advance the start position to the current position.
	l.start = l.pos
}

// errorf emits an illegal token with the given error message.
func (l *Lexer) errorf(format string, args ...interface{}) stateFunc {
	token := Token{
		Type:  ILLEGAL,
		Value: l.input[l.start:l.pos],
	}

	// send the token over the channel
	l.Tokens <- token

	// retuning a nil state function will end the lexing process.
	return nil
}
