package lengthconv

import (
	"fmt"
	"math"
	"testing"
)

// 絶対誤差1e-6まで許容
const (
	EPS float64 = 1e-6
)

func TestMToF(t *testing.T) {
	var tests = []struct {
		meter Meter
		want  Feet
	}{
		{Meter(1), Feet(3.28084)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CToF(%s)", test.meter)
		got := MToF(test.meter)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}

func TestFToM(t *testing.T) {
	var tests = []struct {
		feet Feet
		want Meter
	}{
		{Feet(1), Meter(0.3048)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("FToC(%s)", test.feet)
		got := FToM(test.feet)
		diff := got - test.want
		if math.Abs(float64(diff)) > EPS {
			t.Errorf("%s = %s, want %s", descr, got, test.want)
		}
	}
}
