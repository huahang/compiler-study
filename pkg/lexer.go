package pkg

import (
	"fmt"
)

// LexerState state of this lexer
type LexerState int

const (
	INIT LexerState = iota
	INT1
	INT2
	INT3
	NUMBER
	ALPHA_NUMERIC
)

func isAlphabet(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func isAlphaNumeric(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func isNumber(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isBlank(c rune) bool {
	if c == ' ' {
		return true
	}
	if c == '\t' {
		return true
	}
	if isNewLine(c) {
		return true
	}
	return false
}

func isNewLine(c rune) bool {
	if c == '\n' {
		return true
	}
	if c == '\r' {
		return true
	}
	return false
}

// Tokenize tokenize
func Tokenize(script string, emit func(*Token)) (err error) {
	state := LexerState(INIT)
	token := &Token{}
	for pos, c := range script {
		if state == INIT {
			if c == 'i' {
				state = INT1
				token.tokenBuffer.WriteRune(c)
				continue
			} else if isNumber(c) {
				state = NUMBER
				token.tokenBuffer.WriteRune(c)
				continue
			} else if isAlphabet(c) {
				state = ALPHA_NUMERIC
				token.tokenBuffer.WriteRune(c)
				continue
			} else if c == '>' {

			} else if isBlank(c) {
				continue
			} else {
				err = fmt.Errorf("invalid rune %v at %v", c, pos)
				return err
			}
		} else if state == INT1 {
			if c == 'n' {
				state = INT2
				token.tokenBuffer.WriteRune(c)
				continue
			} else if isAlphaNumeric(c) {
				state = ALPHA_NUMERIC
				token.tokenBuffer.WriteRune(c)
				continue
			} else {
				emit(token)
				token = &Token{}
			}
		} else {
			err = fmt.Errorf("invalid rune %v at %v", c, pos)
		}
	}
	return nil
}
