package main

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	var tests = []struct {
		args []int
		want []int
	}{
		{[]int{0, 1, 2, 3}, []int{0}},
		{[]int{-3, -2, -1, 0}, []int{-3}},
		{[]int{-3}, []int{-3}},
		{[]int{}, nil},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("min(%v)", test.args)
		got := min(test.args...)
		if got == nil && test.want == nil {
			// success
			continue
		}
		if got == nil || test.want == nil {
			// 両方nilではないが、gotかtest.wantのどちらかがnil => error
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
		want := test.want[0]
		if *got != want {
			t.Errorf("%s = %d, want %d", descr, got, want)
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		args []int
		want []int
	}{
		{[]int{0, 1, 2, 3}, []int{3}},
		{[]int{-3, -2, -1, 0}, []int{0}},
		{[]int{-3}, []int{-3}},
		{[]int{}, nil},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("max(%v)", test.args)
		got := max(test.args...)
		if got == nil && test.want == nil {
			// success
			continue
		}
		if got == nil || test.want == nil {
			// 両方nilではないが、gotかtest.wantのどちらかがnil => error
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
		want := test.want[0]
		if *got != want {
			t.Errorf("%s = %d, want %d", descr, got, want)
		}
	}
}

func TestMin2(t *testing.T) {
	var tests = []struct {
		args []int
		want []int
	}{
		{[]int{0, 1, 2, 3}, []int{0}},
		{[]int{-3, -2, -1, 0}, []int{-3}},
		{[]int{-3}, []int{-3}},
		{[]int{}, nil},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("max(%v)", test.args)
		got, err := min2(test.args...)

		// errorを想定するケース
		if test.want == nil {
			if err == nil {
				t.Errorf("%s = (%d, %v), want (0, one argument is requi....)", descr, got, err)
			}
		} else { // そうでないケース
			want := test.want[0]
			if err != nil {
				t.Errorf("%s = (%d, %v), want (%d, nil)", descr, got, err, want)
			}
			if got != want {
				t.Errorf("%s = %d, want %d", descr, got, want)
			}
		}
	}
}

func TestMax2(t *testing.T) {
	var tests = []struct {
		args []int
		want []int
	}{
		{[]int{0, 1, 2, 3}, []int{0}},
		{[]int{-3, -2, -1, 0}, []int{-3}},
		{[]int{-3}, []int{-3}},
		{[]int{}, nil},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("max(%v)", test.args)
		got, err := min2(test.args...)

		// errorを想定するケース
		if test.want == nil {
			if err == nil {
				t.Errorf("%s = (%d, %v), want (0, one argument is requi....)", descr, got, err)
			}
		} else { // そうでないケース
			want := test.want[0]
			if err != nil {
				t.Errorf("%s = (%d, %v), want (%d, nil)", descr, got, err, want)
			}
			if got != want {
				t.Errorf("%s = %d, want %d", descr, got, want)
			}
		}
	}
}
