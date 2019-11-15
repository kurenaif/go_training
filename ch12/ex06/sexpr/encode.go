// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2019 kurenaif
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid: // ignore

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.Complex128, reflect.Complex64:
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(c), imag(c))

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "t")
		}

	case reflect.Interface:
		// type output
		t := v.Elem().Type()

		leftBuffer := new(bytes.Buffer)
		rightBuffer := new(bytes.Buffer)

		if t.Name() == "" { // 名前がつけられてないtypeはそのまま表示する
			fmt.Fprintf(leftBuffer, "%q", t)
		} else {
			fmt.Fprintf(leftBuffer, "\"%s.%s\" ", t.PkgPath(), t.Name()) //一意ではないとはこういうことか？
		}

		// value output
		if err := encode(rightBuffer, v.Elem()); err != nil {
			return err
		}

		if len(rightBuffer.Bytes()) != 0 {
			buf.WriteByte('(')
			buf.Write(leftBuffer.Bytes())
			buf.WriteByte(' ')
			buf.Write(rightBuffer.Bytes())
			buf.WriteByte(')')
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)

		content := new(bytes.Buffer)

		isFirst := true

		for i := 0; i < v.Len(); i++ {
			if isFirst {
				isFirst = false
				content.WriteByte(' ')
			}
			if err := encode(content, v.Index(i)); err != nil {
				return err
			}
		}

		if len(content.Bytes()) != 0 {
			buf.WriteByte('(')
			buf.Write(content.Bytes())
			buf.WriteByte(')')
		}

	case reflect.Struct: // ((name value) ...)

		content := new(bytes.Buffer)

		isFirst := true

		for i := 0; i < v.NumField(); i++ {
			rightBuffer := new(bytes.Buffer)
			if err := encode(rightBuffer, v.Field(i)); err != nil {
				return err
			}

			if len(rightBuffer.Bytes()) != 0 {
				if isFirst {
					content.WriteByte(' ')
					isFirst = false
				}
				content.WriteByte('(')
				fmt.Fprintf(content, "%s", v.Type().Field(i).Name)
				content.WriteByte(' ')
				content.Write(rightBuffer.Bytes())
				content.WriteByte(')')
			}
		}

		if len(content.Bytes()) != 0 {
			buf.WriteByte('(')
			buf.Write(content.Bytes())
			buf.WriteByte(')')
		}

	case reflect.Map: // ((key value) ...)
		isFirst := true
		content := new(bytes.Buffer)

		for _, key := range v.MapKeys() {
			if isFirst {
				content.WriteByte(' ')
				isFirst = false
			}

			leftBuffer := new(bytes.Buffer)
			rightBuffer := new(bytes.Buffer)

			if err := encode(leftBuffer, key); err != nil {
				return err
			}
			if err := encode(rightBuffer, v.MapIndex(key)); err != nil {
				return err
			}

			if len(rightBuffer.Bytes()) != 0 {
				content.WriteByte('(')
				content.Write(leftBuffer.Bytes())
				content.WriteByte(' ')
				content.Write(rightBuffer.Bytes())
				content.WriteByte(')')
			}
		}

		if len(content.Bytes()) != 0 {
			buf.WriteByte('(')
			buf.Write(content.Bytes())
			buf.WriteByte(')')
		}

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// !-encode
