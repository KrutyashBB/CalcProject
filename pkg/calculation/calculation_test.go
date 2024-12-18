package calculation_test

import (
	"testing"

	"github.com/KrutyashBB/CalcProject/pkg/calculation"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name string
		expression string
		expRes float64
	} {
		{
			name:           "simple",
			expression:     "1+1",
			expRes: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expRes: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expRes: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expRes: 0.5,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expRes{
				t.Fatalf("%f should be equal %f", val, testCase.expRes)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "simple",
			expression: "1+1*",
		},
		{
			name:       "priority",
			expression: "2+2**2",
		},
		{
			name:       "priority",
			expression: "((2+2-*(2",
		},
		{
			name:       "/",
			expression: "",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result %f was obtained", testCase.expression, val)
			}
		})
	}
}
