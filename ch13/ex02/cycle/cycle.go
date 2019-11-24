// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2019 kurenaif
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 359.

// Package equal provides a deep equivalence relation for arbitrary values.
package equal

import (
	"reflect"
	"unsafe"
)

type ptrType struct {
	x unsafe.Pointer
	t reflect.Type
}

func isCycle(x reflect.Value, seen map[ptrType]bool) bool {
	// 値を格納していない場合は流石に巡回していない
	if !x.IsValid() {
		return false
	}

	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := ptrType{xptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	//!-cyclecheck
	//!+
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCycle(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if !isCycle(x.Index(i), seen) {
				return false
			}
		}
		return true

	// ...struct and map cases omitted for brevity...
	//!-
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !isCycle(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
		//!+
	}
	panic("unreachable")
}

//!+comparison
// Equal reports whether x and y are deeply equal.
//!-comparison
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
//!+comparison
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

//!-comparison
