// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2019 kurenaif
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 359.

// Package equal provides a deep equivalence relation for arbitrary values.
package cycle

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
			if isCycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	// ...struct and map cases omitted for brevity...
	//!-
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCycle(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
		//!+
	}

	// atomicな要素
	return false
}

func IsCycle(x interface{}) bool {
	seen := make(map[ptrType]bool)
	return isCycle(reflect.ValueOf(x), seen)
}

//!-comparison
