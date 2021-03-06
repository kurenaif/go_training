// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2019 kurenaif
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	var movie2 Movie
	err = NewDecoder(bytes.NewReader(data)).Decode(&movie2)
	if err != nil {
		t.Fatalf("err expect nil, got %v", err)
	}
	if !reflect.DeepEqual(movie2, strangelove) {
		t.Fatalf("\n-----------------------------------------------------------\nexpect:\n%v\ngot:\n%v\n--------------------------------------------------\n", movie, strangelove)
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func Test2(t *testing.T) {
	type S struct {
		Title string
		Flag  bool
		C64   complex64
		C128  complex128
		F32   float32
		F64   float64
		// interface は無理っぽい
	}

	type MyType string

	target := S{"Hello", true, 64 + 1i, 128 + 2i, 32.0, 64.0}
	data, err := Marshal(target)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	var temp S
	err = NewDecoder(bytes.NewReader(data)).Decode(&temp)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(temp, target) {
		t.Fatalf("\n-----------------------------------------------------------\nexpect:\n%v\ngot:\n%v\n--------------------------------------------------\n", target, temp)
	}
}

func Test3(t *testing.T) {
	// c64 := 1 + 2i
	i := 3

	data, err := Marshal(i)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
}
