package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	expression := ""
	result, err := Calc(expression)
	fmt.Println(result, err)

}

var InvalidExpressionError = errors.New("invalid expression")
var StackIsEmptyError = errors.New("stack is empty")
var DivisionByZeroError = errors.New("division by zero")

var OperatorsMap = map[rune]func(float64, float64) float64{
	'*': func(a, b float64) float64 { return a * b },
	'/': func(a, b float64) float64 { return a / b },
	'+': func(a, b float64) float64 { return a + b },
	'-': func(a, b float64) float64 { return a - b },
}

var OperatorsPriorities = map[rune]int{
	'*': 2,
	'/': 2,
	'+': 1,
	'-': 1,
}

type Stack[T float64 | rune] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var defaultValue T
	if s.IsEmpty() {
		return defaultValue, StackIsEmptyError
	}
	top := s.Top()
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

func (s *Stack[T]) Top() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func Calc(expression string) (float64, error) {
	var digitsStack Stack[float64]
	var operatorsStack Stack[rune]

	if len(expression) == 0 {
		return 0, InvalidExpressionError
	}
	if _, ok := OperatorsMap[rune(expression[len(expression)-1])]; ok {
		return 0, InvalidExpressionError
	}
	exprSlice := []rune(strings.Replace(expression, " ", "", -1))

	lastIsOperator := true

	for _, ch := range exprSlice {
		if unicode.IsDigit(ch) {
			digitsStack.Push(float64(ch - '0'))
			lastIsOperator = false
		} else if strings.ContainsAny(string(ch), "+-/*()") {
			if _, ok := OperatorsMap[ch]; ok {
				if lastIsOperator {
					return 0, InvalidExpressionError
				}
				lastIsOperator = true
			}

			if ch == '(' {
				operatorsStack.Push(ch)
			} else if ch == ')' {
				for !operatorsStack.IsEmpty() {
					topOperator, _ := operatorsStack.Pop()
					if topOperator == '(' {
						break
					}
					if err := CalcTwoDigs(&digitsStack, topOperator); err != nil {
						return 0, err
					}
				}
			} else {
				for !operatorsStack.IsEmpty() && OperatorsPriorities[operatorsStack.Top()] >= OperatorsPriorities[ch] {
					topOperator, _ := operatorsStack.Pop()
					if err := CalcTwoDigs(&digitsStack, topOperator); err != nil {
						return 0, err
					}
				}
				operatorsStack.Push(ch)
			}
		} else {
			return 0, InvalidExpressionError
		}
	}
	for !operatorsStack.IsEmpty() {
		topOperator, _ := operatorsStack.Pop()
		if err := CalcTwoDigs(&digitsStack, topOperator); err != nil {
			return 0, err
		}
	}
	return digitsStack.Top(), nil

}

func CalcTwoDigs(digitsStack *Stack[float64], topOperator rune) error {
	dig2, _ := digitsStack.Pop()
	dig1, _ := digitsStack.Pop()
	if dig2 == 0 && topOperator == '/' {
		return DivisionByZeroError
	}
	res := OperatorsMap[topOperator](dig1, dig2)
	digitsStack.Push(res)
	return nil
}
