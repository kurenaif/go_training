package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", float64(l))
}

func (u unary) String() string {
	res := fmt.Sprintf("%c%s", u.op, u.x)
	switch u.x.(type) {
	case unary, binary:
		res = "(" + res + ")"
	}
	return res
}

func (b binary) String() string {
	lhs := b.x.String()
	rhs := b.y.String()

	switch b.x.(type) {
	case unary, binary:
		lhs = "(" + lhs + ")"
	}
	switch b.y.(type) {
	case unary, binary:
		rhs = "(" + rhs + ")"
	}
	return fmt.Sprintf("%s%c%s", lhs, b.op, rhs)
}

func (c call) String() string {
	res := c.fn + "("
	sep := ""
	for _, arg := range c.args {
		res += sep + arg.String()
		sep = ","
	}
	res += ")"
	return res
}

func (b bracket) String() string {
	res := string(b.op) + "("
	sep := ""
	for _, arg := range b.args {
		res += sep + arg.String()
		sep = ","
	}
	res += ")" + string(b.op)
	return res
}
