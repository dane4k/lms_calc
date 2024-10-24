package main

import (
	"errors"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b
	"strings"
	"unicode"
)

func main() {
<<<<<<< HEAD
	expression := "1 + 2 * (3+4/2 - (1+2))*2+1"

	result, err := Calc(expression)
	fmt.Println(result, err)

}

var InvalidExpressionError = errors.New("invalid expression")
var StackIsEmptyError = errors.New("stack is empty")
var DivisionByZeroError = errors.New("division by zero")
=======
	expression := "(2 + 2) * 2 + ((13 - 1) + (2 - (3 + 4)))"

	_, _ = Calc(expression)

}

type AnyArr []interface{}

var InvalidExpressionError = errors.New("invalid expression")
var StackIsEmptyError = errors.New("stack is empty")
>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b

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
<<<<<<< HEAD
}

type Stack[T float64 | rune] struct {
=======
	'(': 999,
	')': 999,
}

type Stack[T int | rune] struct {
>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

<<<<<<< HEAD
func (s *Stack[T]) Pop() (T, error) {
=======
func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) Top() (T, error) {
>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b
	var defaultValue T
	if s.IsEmpty() {
		return defaultValue, StackIsEmptyError
	}
<<<<<<< HEAD
	top := s.Top()
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

func (s *Stack[T]) Top() T {
	return s.items[len(s.items)-1]
=======
	return s.items[len(s.items)-1], nil
>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func Calc(expression string) (float64, error) {
<<<<<<< HEAD
	var digitsStack Stack[float64]
=======
	var digitsStack Stack[int]
>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b
	var operatorsStack Stack[rune]

	exprSlice := []rune(strings.Replace(expression, " ", "", -1))

	for _, ch := range exprSlice {
		if unicode.IsDigit(ch) {
<<<<<<< HEAD
			digitsStack.Push(float64(ch - '0'))
		} else if strings.ContainsAny(string(ch), "+-/*()") {
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
=======
			digitsStack.Push(int(ch - '0'))
		} else if strings.ContainsAny(string(ch), "+-/*()") {
			topOperator, err := operatorsStack.Top()
			if err != nil {
				return 0, err
			}
			if OperatorsPriorities[topOperator] > OperatorsPriorities[ch] {
				operatorsStack.Push(ch)
			}

>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b
		} else {
			return 0, InvalidExpressionError
		}
	}
<<<<<<< HEAD
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
=======
	return 0, nil
}

func CalcTwoDigs(dig1, dig2 float64, operation rune) (float64, error) {
	return OperatorsMap[operation](dig1, dig2), nil
>>>>>>> 3abc6587ac52f6b7a3a71b1c5def01f28586747b
}
