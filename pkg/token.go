package pkg

import (
	"strings"
)

// TokenType types of tokens
type TokenType uint32

// predifined token types
const (
	Plus          TokenType = iota // +
	Minus                          // -
	Star                           // *
	Slash                          // /
	GE                             // >=
	GT                             // >
	EQ                             // ==
	LE                             // <=
	LT                             // <
	SemiColon                      // ;
	LeftParen                      // (
	RightParen                     // )
	Assignment                     // =
	If                             // If
	Else                           // Else
	Int                            // Int
	Identifier                     // 标识符
	IntLiteral                     // 整型字面量
	StringLiteral                  // 字符串字面量
)

func (t TokenType) String() string {
	return [...]string{
		"Plus",
		"Minus",
		"Star",
		"Slash",
		"GE",
		"GT",
		"EQ",
		"LE",
		"LT",
		"SemiColon",
		"LeftParen",
		"RightParen",
		"Assignment",
		"If",
		"Else",
		"Int",
		"Identifier",
		"IntLiteral",
		"StringLiteral",
	}[t]
}

// Token token structure
type Token struct {
	tokenType   TokenType
	tokenBuffer strings.Builder
}

// Type token type
func (token *Token) Type() TokenType {
	return token.tokenType
}

// String token string
func (token *Token) String() string {
	return token.tokenBuffer.String()
}
