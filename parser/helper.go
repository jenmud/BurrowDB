package parser

import "unicode"

// IsWhitespace returns true if the rune is a whitespace character.
func IsWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

// IsDigit returns true if the rune is a digit.
func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// IsAlpha returns true if the rune is a letter or an underscore.
func IsAlpha(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

// IsAlphaNumeric returns true if the rune is a letter, digit, or underscore.
func IsAlphaNumeric(r rune) bool {
	return IsAlpha(r) || IsDigit(r)
}

// IsSymbol returns true if the rune is a symbol character.
func IsSymbol(r rune) bool {
	switch r {
	case '(', ')', '{', '}', ':', ',', ';', '.':
		return true
	}

	return false
}
