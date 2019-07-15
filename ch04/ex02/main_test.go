package main

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	var tests = []struct {
		s             string
		hashAlgorithm string
		want          string
	}{
		{"hello", "sha256", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"hello", "sha384", "59e1748777448c69de6b800d7a33bbfb9ff1b463e44354c3553bcdb9c666fa90125a3c79f90397bdf5f6a13de828684f"},
		{"hello", "sha512", "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("hash(%s, %s)", test.s, test.hashAlgorithm)
		got, _ := hash(test.s, test.hashAlgorithm)
		if got != test.want {
			t.Errorf("%s = %s want %s\n", descr, got, test.want)
		}
	}
}
