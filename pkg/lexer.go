package pkg

import "fmt"

// LexerState state of this lexer
type lexerState int

const (
	sInit lexerState = iota
	sInt1
	sInt2
	sInt3
	sNumberic
	sAlphaNumeric
	sGT
	sLT
	sEQ
	sGE
	sLE
	sSemi
	sError
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

func isNumberic(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isAlphaNumeric(c rune) bool {
	if isAlphabet(c) {
		return true
	}
	if isNumberic(c) {
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

func initToken(c rune) (token *Token, state lexerState, err error) {
	if isBlank(c) {
		token = nil
		state = sInit
	} else {
		token = &Token{}
		if c == 'i' {
			token.tokenType = Int
			token.tokenBuffer.WriteRune(c)
			state = sInt1
		} else if isNumberic(c) {
			token.tokenType = IntLiteral
			token.tokenBuffer.WriteRune(c)
			state = sNumberic
		} else if isAlphabet(c) {
			token.tokenType = Identifier
			token.tokenBuffer.WriteRune(c)
			state = sAlphaNumeric
		} else if c == '=' {
			token.tokenType = EQ
			token.tokenBuffer.WriteRune(c)
			state = sEQ
		} else if c == '>' {
			token.tokenType = GT
			token.tokenBuffer.WriteRune(c)
			state = sGT
		} else if c == '<' {
			token.tokenType = LT
			token.tokenBuffer.WriteRune(c)
			state = sLT
		} else if c == ';' {
			token.tokenType = SemiColon
			token.tokenBuffer.WriteRune(c)
			state = sSemi
		} else {
			token = nil
			state = sError
			err = fmt.Errorf("Invalid rune: %v", c)
		}
	}
	return token, state, err
}

// Tokenize tokenize
func Tokenize(script string, emit func(*Token)) (pos int, err error) {
	var (
		state        = lexerState(sInit)
		token *Token = nil
	)
	for pos, c := range script {
		switch state {
		case sInit:
			{
				token, state, err = initToken(c)
				if err != nil {
					return pos, err
				}
				continue
			}
		case sInt1:
			{
				if c == 'n' {
					state = sInt2
					token.tokenBuffer.WriteRune(c)
				} else if isAlphaNumeric(c) {
					state = sAlphaNumeric
					token.tokenType = Identifier
					token.tokenBuffer.WriteRune(c)
				} else {
					token.tokenType = Identifier
					emit(token)
					token, state, err = initToken(c)
					if err != nil {
						return pos, err
					}
				}
				continue
			}
		case sInt2:
			{
				if c == 't' {
					state = sInt3
					token.tokenBuffer.WriteRune(c)
				} else if isAlphaNumeric(c) {
					state = sAlphaNumeric
					token.tokenType = Identifier
					token.tokenBuffer.WriteRune(c)
				} else {
					token.tokenType = Identifier
					emit(token)
					token, state, err = initToken(c)
					if err != nil {
						return pos, err
					}
				}
				continue
			}
		case sInt3:
			{
				if isAlphaNumeric(c) {
					state = sAlphaNumeric
					token.tokenType = Identifier
					token.tokenBuffer.WriteRune(c)
				} else {
					emit(token)
					token, state, err = initToken(c)
					if err != nil {
						return pos, err
					}
				}
				continue
			}
		case sAlphaNumeric:
			{
				if isAlphaNumeric(c) {
					token.tokenBuffer.WriteRune(c)
				} else {
					emit(token)
					token, state, err = initToken(c)
					if err != nil {
						return pos, err
					}
				}
				continue
			}
		case sNumberic:
			{
				if isNumberic(c) {
					token.tokenBuffer.WriteRune(c)
				} else {
					emit(token)
					token, state, err = initToken(c)
					if err != nil {
						return pos, err
					}
				}
				continue
			}
		case sGT, sLT:
			{
				if c != '=' {
					emit(token)
					token, state, err = initToken(c)
					if err != nil {
						return pos, err
					}
					continue
				}
				if state == sGT {
					token.tokenType = GE
					state = sGE
				}
				if state == sLT {
					token.tokenType = LE
					state = sLE
				}
				token.tokenBuffer.WriteRune(c)
			}
		case sEQ, sGE, sLE, sSemi:
			{
				emit(token)
				token, state, err = initToken(c)
				if err != nil {
					return pos, err
				}
			}
		}
	}
	if token != nil {
		emit(token)
	}
	return -1, nil
}
