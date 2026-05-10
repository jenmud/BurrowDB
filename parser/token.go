package parser

type TokenType int

const (
	// Special tokens
	EOF TokenType = iota
	ILLEGAL

	// Literals
	IDENT
	LABEL
	STRING
	INTEGER

	// Keywords
	KEYWORD // KEYWORD is a placeholder for all keywords, we will check the actual keyword in the parser.

	// Symbols
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LSQUARE   // [
	RSQUARE   // ]
	COMMA     // ,
	COLON     // :
	SEMICOLON // ;
	DOT       // .
)

// String returns the string representation of the TokenType.
func (t TokenType) String() string {
	return []string{
		"EOF",
		"ILLEGAL",
		"IDENT",
		"LABEL",
		"STRING",
		"INTEGER",
		"KEYWORD",
		"LPAREN",
		"RPAREN",
		"LBRACE",
		"RBRACE",
		"LSQUARE",
		"RSQUARE",
		"COMMA",
		"COLON",
		"SEMICOLON",
		"DOT",
	}[t]
}

// Token represents a lexical token.
type Token struct {
	Type  TokenType
	Value string
}
