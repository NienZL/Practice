package model

// import (
// 	"errors"
// 	"fmt"
// )

// const (
// 	Literal = iota // Number = 0 , operator = 1
// 	Operator
// )

// type Token struct {
// 	Char  string
// 	Type  int
// 	index int
// }

// type Scanner struct {
// 	Source string //user input
// 	char   string //current character
// 	index  int    // current index
// }

// func Scan(input string) []*Token {
// 	s := &Scanner{
// 		Source: input,
// 		char:   string(input[0]),
// 		index:  0,
// 	}

// 	token := s.ToToken()
// 	return token
// }

// func (s *Scanner) ToToken() []*Token {
// 	token := make([]*Token, 0)
// 	for {
// 		tok := s.nextToken()
// 		// read next token  until no more strings
// 		if tok == nil {
// 			break
// 		}

// 		token = append(token, tok)
// 	}

// 	return token
// }

// // handle all character issue
// func (s *Scanner) nextToken() *Token {
// 	// if still have unhandled char
// 	if s.index >= len(s.Source) {
// 		return nil
// 	}
// 	// skip whitespace
// 	var err error
// 	for s.isWhiteSpace(s.char) && err == nil {
// 		err = s.nextChar()
// 	}

// 	//開始處理有意義的字符
// 	startIndex := s.index
// 	var token *Token
// 	switch s.char {
// 	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
// 		for s.isDigitNum(s.char) && s.nextChar() == nil {
// 		}
// 		token = &Token{
// 			Char:  s.Source[startIndex:s.index],
// 			Type:  Literal,
// 			index: startIndex,
// 		}
// 	case "+", "-", "*", "/", "(", ")":
// 		token = &Token{
// 			Char:  s.char,
// 			Type:  Operator,
// 			index: startIndex,
// 		}
// 		err = s.nextChar()
// 	default:
// 		fmt.Println("輸入的字元錯誤")
// 	}
// 	return token
// }

// func (s *Scanner) isWhiteSpace(c string) bool {
// 	return c == " "
// }

// func (s *Scanner) isDigitNum(c string) bool {
// 	return "0" <= c && c <= "9" || c == "."
// }

// func (s *Scanner) nextChar() error {
// 	s.index++
// 	if s.index < len(s.Source) {
// 		s.char = string(s.Source[s.index])
// 		return nil
// 	} else {
// 		// no more next char
// 		return errors.New("EOF")
// 	}
// }

// // for i:=0;i<i:= len(input),i++{  // scan every char of input string
// // 	char:=input[i] // current char
// // 	switch char{
// // 	case '0','1','2','3','4','5','6','7','8','9':

// // 	}
// // }
