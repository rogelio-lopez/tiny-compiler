package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	const test string = "(add 2 (subtract 4 2))"
	tokenizer(test)
}

type tokens struct {
	typeOf string
	value  int
}

func tokenizer(input string) []tokens {
	var tokenArr []tokens = []tokens{}

	for _, v := range input {
		//space = 32 ascii
		if v != 32 {
			fmt.Println(string(v))
		}
	}
	return tokenArr
}
