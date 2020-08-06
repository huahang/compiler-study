package main

import (
	"fmt"

	"github.com/huahang/compiler-study/pkg"
)

func main() {
	var script string
	script = "int i = 8"
	pkg.Tokenize(script, func(token *pkg.Token) {
		fmt.Printf("Token type: %v, value: %v\n", token.Type(), token.String())
	})
}
