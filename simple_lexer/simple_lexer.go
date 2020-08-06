package main

import (
	"fmt"

	"github.com/huahang/compiler-study/pkg"
)

func main() {
	scripts := []string{
		"int age = 45;",
		"inta age = 45;",
		"in age = 45;",
		"age >= 45;",
		"age > 45;",
	}
	for _, script := range scripts {
		pkg.Tokenize(script, func(token *pkg.Token) {
			fmt.Printf("Token: %v %v\n", token.Type().String(), token.String())
		})
		fmt.Println("")
	}
}
