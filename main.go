package main

import (
	"errors"
	"strings"
	"unicode"
)

func main() {
	expression := "(2 + 2) * 2 + ((13 - 1) + (2 - (3 + 4)))"

	_, _ = Calc(expression)

}

type AnyArr []interface{}

var InvalidExpressionError = errors.New("invalid expression")
var StackIsEmptyError = errors.New("stack is empty")

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
	'(': 999,
	')': 999,
}

type Stack[T int | rune] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) Top() (T, error) {
	var defaultValue T
	if s.IsEmpty() {
		return defaultValue, StackIsEmptyError
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func Calc(expression string) (float64, error) {
	var digitsStack Stack[int]
	var operatorsStack Stack[rune]

	exprSlice := []rune(strings.Replace(expression, " ", "", -1))

	for _, ch := range exprSlice {
		if unicode.IsDigit(ch) {
			digitsStack.Push(int(ch - '0'))
		} else if strings.ContainsAny(string(ch), "+-/*()") {
			topOperator, err := operatorsStack.Top()
			if err != nil {
				return 0, err
			}
			if OperatorsPriorities[topOperator] > OperatorsPriorities[ch] {
				operatorsStack.Push(ch)
			}

		} else {
			return 0, InvalidExpressionError
		}
	}
	return 0, nil
}

func CalcTwoDigs(dig1, dig2 float64, operation rune) (float64, error) {
	return OperatorsMap[operation](dig1, dig2), nil
}
