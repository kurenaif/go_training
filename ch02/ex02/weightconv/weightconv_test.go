package weightconv

import (
	"fmt"
	"math"
	"testing"
)

// 絶対誤差1e-6まで許容
const (
	EPS float64 = 1e-6
)

func TestKToP(t *testing.T) {
	var tests = []struct {
		kilogram Kilogram
		want     Pounds
	}{
		// 1/0.45359237で計算
		{Kilogram(1.0), Pounds(2.20462262185)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToP(%s)", test.kilogram)
		got := KToP(test.kilogram)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestPToK(t *testing.T) {
	var tests = []struct {
		pounds Pounds
		want   Kilogram
	}{
		{Pounds(1), Kilogram(0.45359237)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("PToK(%s)", test.pounds)
		got := PToK(test.pounds)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}
