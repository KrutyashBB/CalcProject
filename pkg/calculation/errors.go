package calculation

import "errors"

var (
	ErrDivisionByZero      = errors.New("division by zero")
	ErrNumParsing          = errors.New("number parsing error")
	ErrParenthesisSequence = errors.New("incorrect bracket sequence")
	ErrExpression          = errors.New("Expression is not valid")
)
