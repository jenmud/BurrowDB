package main

import (
	"fmt"

	"burrow/parser"
)

func main() {
	inputs := []string{
		`
			CREATE (n:Person:Dog {"name": "bob", "age": 30, "meta": {"hobbies": ["coding", "hiking"]}})
		 	RETURN n.name, n.age. n.meta.hobbies;
		`,
		`MATCH (n) RETURN n;`,
	}

	for _, input := range inputs {
		fmt.Println("INPUT:", input)

		lexer := parser.NewLexer(input)

		for tok := range lexer.Tokens {
			fmt.Printf("%-12s %q\n", tok.Type, tok.Value)
		}

		fmt.Println()
	}
}
