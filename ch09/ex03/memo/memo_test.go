// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo

import (
	"go_training/ch09/ex03/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestCloseFirst(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.SequentialCloseFirstCase(t, m)
}

func TestCloseSecond(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.SequentialCloseSecondCase(t, m)
}

func TestCloseThird(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	memotest.SequentialCloseThirdCase(t, m)
}
