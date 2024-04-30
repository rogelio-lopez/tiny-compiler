package main

import (
	"fmt"
	"log"
	"strconv"
)

type Tokens struct {
	typeOf string
	value  string
	inInt  int
}
type Node struct {
	typeOf string
	name   string
	value  int

	body   []Node
	params []Node
}

func main() {
	const test string = "(add 2 (subtract 41 2))"
	tokens := tokenizer(test)

	var current = 0
	var ast = Node{
		typeOf: "program",
		body:   parser(tokens, &current),
	}
	fmt.Println(ast)
}

func traverser() {

}

/*
 *   {
 *     type: 'Program',
 *     body: [{
 *       type: 'CallExpression',
 *       name: 'add',
 *       params: [{
 *         type: 'NumberLiteral',
 *         value: '2'
 *       }, {
 *         type: 'CallExpression',
 *         name: 'subtract',
 *         params: [{
 *           type: 'NumberLiteral',
 *           value: '4'
 *         }, {
 *           type: 'NumberLiteral',
 *           value: '2'
 *         }]
 *       }]
 *     }]
 *   }
 */
// [{paren ( 0} ->{opp add 0} {num 2 2} {paren ( 0} {opp subtract 0} {num 41 41} {num 2 2} {paren ) 0} {paren ) 0}]

func parser(tokens []Tokens, currentToken *int) []Node {
	var returnNode = []Node{}

	for *currentToken < len(tokens)-1 {
		*currentToken++

		//if opening ( then call this funciton again
		if tokens[*currentToken].typeOf == "opp" {
			//append to returnNode []Node to then return to body
			var val = tokens[*currentToken].value
			returnNode = append(returnNode, Node{
				typeOf: "CallExpression",
				name:   val,
				params: parser(tokens, currentToken),
			})

		} else if tokens[*currentToken].typeOf == "num" {
			num, err := strconv.Atoi(tokens[*currentToken].value)
			if err != nil {
				log.Fatalln("strconv error")
			}

			returnNode = append(returnNode, Node{
				typeOf: "NumberLiteral",
				value:  num,
			})

		} else if tokens[*currentToken].typeOf == "paren" && tokens[*currentToken].value == ")" {
			return returnNode
		}
	}

	return returnNode
}

func tokenizer(input string) []Tokens {
	var current int = 0
	var tokenArr []Tokens = []Tokens{}
	var tokenStr string = ""

	for current < len(input) {
		char := input[current]

		//space = 32 ascii
		if char >= 97 && char <= 122 { // letters
			tokenStr = tokenStr + string(char)

			// need to keep looking until end of space
			if current < len(input)-1 {
				if input[current+1] < 97 || input[current+1] > 122 {

					tokenArr = append(tokenArr, Tokens{typeOf: "opp", value: tokenStr})
					tokenStr = ""
				}
			}
		} else if char >= 48 && char <= 57 { // numbers
			tokenStr = tokenStr + string(char)

			// need to keep looking until end of space
			if current < len(input)-1 {
				if input[current+1] < 48 || input[current+1] > 57 {
					val, err := strconv.Atoi(tokenStr)
					if err != nil {
						log.Fatal("Error trying to strconv")
					}
					tokenArr = append(tokenArr, Tokens{typeOf: "num", value: tokenStr, inInt: val})
					tokenStr = ""
				}
			}
		} else if char == 40 { // paren (
			tokenArr = append(tokenArr, Tokens{typeOf: "paren", value: "("})
		} else if char == 41 { // paren )
			tokenArr = append(tokenArr, Tokens{typeOf: "paren", value: ")"})
		}
		current++
	}
	return tokenArr
}
