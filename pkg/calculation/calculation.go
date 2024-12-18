package calculation

import (
	"strconv"
	"strings"
)

func Calc(expr string) (float64, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	priority := map[rune]int{'*': 2, '/': 2, '+': 1, '-': 1, '(': 0, ')': 0}
	var err error

	var nums []float64
	var ops []rune

	for i := 0; i < len(expr); i++ {
		chr := rune(expr[i])

		switch {
		case chr >= '0' && chr <= '9':
			startInd := i
			for i < len(expr)-1 && (expr[i+1] >= '0' && expr[i+1] <= '9' || expr[i+1] == '.') {
				i++
			}

			num, err := strconv.ParseFloat(expr[startInd:i+1], 64)
			if err != nil {
				return 0, ErrNumParsing
			}
			nums = append(nums, num)

		case chr == '+' || chr == '*' || chr == '-' || chr == '/':
			if len(ops) > 0 && priority[ops[len(ops)-1]] >= priority[chr] {
				if nums, ops, err = execExpression(nums, ops); err != nil {
					return 0, err
				}
			}
			ops = append(ops, chr)

		case chr == '(':
			ops = append(ops, chr)

		case chr == ')':
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if nums, ops, err = execExpression(nums, ops); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, ErrParenthesisSequence
			}
			ops = ops[:len(ops)-1]
		default:
			return 0, ErrExpression
		}
	}

	for len(ops) > 0 {
		if nums, ops, err = execExpression(nums, ops); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 || len(ops) != 0 {
		return 0, ErrExpression
	}

	return nums[0], nil
}

func execExpression(nums []float64, ops []rune) ([]float64, []rune, error) {
	if len(nums) < 2 {
		return nil, nil, ErrExpression
	}
	a := nums[len(nums)-2]
	b := nums[len(nums)-1]
	op := ops[len(ops)-1]

	nums = nums[:len(nums)-2]
	ops = ops[:len(ops)-1]

	var res float64
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		if b == 0 {
			return nil, nil, ErrDivisionByZero
		}
		res = a / b
	}
	return append(nums, res), ops, nil
}
