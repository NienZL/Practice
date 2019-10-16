package model

import (
	"unicode"
)

func toPostfix(input string) []string {
	operatorStack := new(Stack)
	var postfix []string
	var digitNumber []string
	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if char == "." || unicode.IsDigit(rune(input[i])) {
			digitNumber = append(digitNumber, string(input[i]))
		} else {
			if len(digitNumber) > 0 {
				postfix = append(postfix, saveNumber(digitNumber))
				digitNumber = digitNumber[:0] //flush digitNumber temp
			}

			switch char {
			case " ":
				continue
			case "(":
				operatorStack.Push("(")
			case ")":
				for n := operatorStack.Len(); n > 0; n-- {
					if operatorStack.Peek() == "(" {
						operatorStack.Remove()
						break
					}
					postfix = append(postfix, operatorStack.Pop())
				}
			default: // + - * /
				if operatorStack.Len() > 0 {
					if operatorStack.Peek() == "(" {
						operatorStack.Push(char)
					} else {
						for n := operatorStack.Len(); n > 0 && isHigher(operatorStack.Peek(), char); n-- {
							//  true:top>=char ; false:top<char ;  when top>=char pop()
							postfix = append(postfix, operatorStack.Pop())
						}
						operatorStack.Push(char)
					}
				} else {
					//nothing in stack, directly push operator
					operatorStack.Push(char)
				}
			}
		}
	}

	//pop all elements in the stack and integer temp
	if len(digitNumber) > 0 {
		postfix = append(postfix, saveNumber(digitNumber))
	}

	for operatorStack.Len() > 0 {
		postfix = append(postfix, operatorStack.Pop())
	}
	return postfix
}

func saveNumber(input []string) string {
	var number string
	for i := 0; i < len(input); i++ {
		number += input[i]
	}
	// fmt.Printf("check the number reader : %s\n", number)
	return number
}

func isHigher(top string, char string) bool { // true:top<char ; false: top>=char
	switch top {
	case "*", "/":
		return true
	case "+", "-":
		if char == "*" || char == "/" {
			return false
		} else {
			return true
		}
	}
	return false
}

// import (
// 	"fmt"
// 	"strconv"
// )

// type Tree struct {
// 	Tokens    []*Token
// 	currTok   *Token
// 	currIndex int
// 	Err 	error
// }

// type Node interface {
// 	toString() string // handle diffrent input type
// }

// type NumberNode struct {
// 	Value float64
// }

// type OperatorNode struct {
// 	Operator string
// 	Left,
// 	Right Node
// }

// func (n NumberNode) toString() string {
// 	return fmt.Sprintf("NumberNode:%s", strconv.FormatFloat(n.Value, 'f', 0, 64))
// }
// func (o OperatorNode) toString() string {
// 	return fmt.Sprintf("OperatorNode:(%s %s %s)", o.Operator, o.Left.toString(), o.Right.toString())
// }

// var priority = map[string]int{"+": 10, "-": 10, "*": 20, "/": 20}

// func (t *Tree) getNextToken() *Token {
// 	t.currIndex++
// 	if t.currIndex < len(t.Tokens) {
// 		t.currTok = t.Tokens[t.currIndex]
// 		return t.currTok
// 	}
// 	return nil
// }

// func (t *Tree) getTokenPriority() int {
// 	if p, ok := priority[t.currTok.Char]; ok {
// 		return p
// 	}
// 	return -1
// }

// // generate Number Node
// func (t *Tree) parseNumber() NumberNode{
// 	f64, err := strconv.ParseFloat(t.currTok.Char,64)
// 	if err!= nil {
// 		t.Err = errors.New(
// 			fmt.Sprintf("%v \n want '(' or '0-9' but get wrong character", err.Error(), t.currTok.Char)
// 		)
// 		return NumberNode{}
// 	}
// 	n := NumberNode{
// 		Value: f64,
// 	}
// 	t.getNextToken()
// 	return n
// }

// //generate other nodes
// func (t *Tree) parsePrimary() Node{
// 	switch t.currTok.Type{
// 	case Literal:
// 		return t.parseNumber()
// 	case Operator:
// 		if t.currTok.Char == "(" {
// 			t.getNextToken()

// 		}
// 	}
// }

// func (t *Tree) ParseExpression() Node {
// 	left := t.parsePrimary()
// 	return t.parseBTree(0,left)
// }

// func (t *Tree) parseBTree(currPriority int, left Node) Node{

// }

// func (t *Tree) getNextToken() *Token{
// 	t.currIndex++
// 	if t.currIndex < len(t.Tokens) {
// 		t.currTok = t.Tokens[]
// 	}
// }

// func (t *Tree) ParsePrimary() Node{ //handle all token and return a node
// 	switch t.currTok.Type{
// 	case Literal:
// 		return t.parseNumber()
// 	case Operator:
// 		// handle "(" ")"
// 		if t.currTok.Char == "(" {
// 			t.getNextToken()
// 		}
// 	}
// }
