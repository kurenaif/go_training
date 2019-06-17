package tempconv

import (
	"fmt"
	"math"
	"testing"
)

// 絶対誤差1e-6まで許容
const (
	EPS float64 = 0.000001
)

func TestCToF(t *testing.T) {
	var tests = []struct {
		celsius Celsius
		want    Fahrenheit
	}{
		{FreezingC, Fahrenheit(32)},
		{Celsius(-17.0), Fahrenheit(1.4)},
		{BoilingC, Fahrenheit(212)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CToF(%s)", test.celsius)
		got := CToF(test.celsius)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		fahrenheit Fahrenheit
		want       Celsius
	}{
		{Fahrenheit(32), FreezingC},
		{Fahrenheit(1.4), Celsius(-17.0)},
		{Fahrenheit(212), BoilingC},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("FToC(%s)", test.fahrenheit)
		got := FToC(test.fahrenheit)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		kelvin Kelvin
		want   Celsius
	}{
		{FreezingK, FreezingC},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToC(%s)", test.kelvin)
		got := KToC(test.kelvin)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		celsius Celsius
		want    Kelvin
	}{
		{FreezingC, FreezingK},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToC(%s)", test.celsius)
		got := CToK(test.celsius)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

// KToFとFToKはすでに実装されている関数を呼び出しているだけで、それぞれの関数はテスト済みなのでテスト不要
// 関数の呼び間違えのリスクはあるが、それは型で制限しているのでコンパイルエラーになってそもそもコンパイルは通らない
