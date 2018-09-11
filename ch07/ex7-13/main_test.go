package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s: %v => %s, want %s", test.expr, test.env, got, test.want)
		}
	}
}

func TestString(t *testing.T) {
	tests := []string{
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"5 / 9 * (F - 32)",
	}

	for _, test := range tests {
		expr1, err := Parse(test)
		if err != nil {
			t.Error(err)
		}
		expr2, err := Parse(expr1.String())
		if err != nil {
			t.Error(err)
		}
		if expr1.String() != expr2.String() {
			t.Errorf("%s != %s", expr1.String(), expr2.String())
		}
	}
}
