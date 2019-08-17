package eval

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"1", Env{}, "1.000000"},
		{"A", Env{"A": 12345}, "A"},
		{"+1", Env{}, "+1.000000"},
		{"(((((1+1)))))", Env{}, "1.000000+1.000000"},
		{"(((1+((2*3)))))", Env{}, "1.000000+(2.000000*3.000000)"},
		{"(A+B)*C", Env{"A": 345, "B": 456, "C": 567}, "(A+B)*C"},
		{"(((((A+B)))))*C", Env{"A": 345, "B": 456, "C": 567}, "(A+B)*C"},
		{"pow(x,3) + pow(y,3)", Env{"x": 12, "y": 1}, "pow(x,3.000000)+pow(y,3.000000)"},
		{"pow(x,3) + pow(y,3) * 4 + 3", Env{"x": 12, "y": 1}, "(pow(x,3.000000)+(pow(y,3.000000)*4.000000))+3.000000"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "(5.000000/9.000000)*(F-32.000000)"},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := expr.String()
		if got != test.want {
			t.Errorf("%s.String() got\n%q\n, want\n%q\n",
				test.expr, got, test.want)
		}

		expr2, err := Parse(got)
		if err != nil {
			t.Error(err)
			continue
		}

		gotEval := expr.Eval(test.env)
		gotEval2 := expr2.Eval(test.env)
		if gotEval != gotEval2 {
			t.Errorf("%s.Eval() != %s.Eval()", expr, expr2)
			t.Errorf("%s.Eval() => %v", expr, gotEval)
			t.Errorf("%s.Eval() => %v", expr2, gotEval2)
		}
	}
}
