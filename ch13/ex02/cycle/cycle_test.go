// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package cycle

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	one := 1

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1 := make(chan int)
	var ch1ro <-chan int = ch1

	type mystring string

	var iface1 interface{} = &one

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false},
		{"foo", false},
		// slices
		{[]string{"foo"}, false},
		// slice cycles
		{cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		{
			map[string][]int{},
			false,
		},
		// pointers
		{&one, false},
		{new(bytes.Buffer), false},
		// pointer cycles
		{cyclePtr1, true},
		{cyclePtr2, true},
		// functions
		{(func())(nil), false},
		{func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, false},
		// channels
		{ch1, false},
		{ch1ro, false}, // NOTE: not equal
		// interfaces
		{&iface1, false},
	} {
		if IsCycle(test.x) != test.want {
			t.Errorf("IsCycle(%v) = %t",
				test.x, !test.want)
		}
	}
}

func Example_equal() {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	//!+
	fmt.Println(IsCycle([]int{1, 2, 3}))      // "false"
	fmt.Println(IsCycle([]string{"foo"}))     // "false"
	fmt.Println(IsCycle([]string(nil)))       // "false"
	fmt.Println(IsCycle(map[string]int(nil))) // "false"
	fmt.Println(IsCycle(a))                   // "true"
	fmt.Println(IsCycle(b))                   // "true"
	fmt.Println(IsCycle(c))                   // "true"
	//!-

	// Output:
	// false
	// false
	// false
	// false
	// true
	// true
	// true
}

func Example2_equal() {
	type Node struct {
		left  *Node
		right *Node
	}
	var n2 Node
	var n Node
	n.left = &n2
	n.right = &n2

	//!+
	fmt.Println(IsCycle(n))
	//!-

	// Output:
	// false
}
