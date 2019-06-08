package main

import (
	"testing"
)

func MakeCase(array_size int, str_length int) (res []string) {
	str := ""
	for i := 0; i < str_length; i++ {
		str += "a"
	}
	for i := 0; i < array_size; i++ {
		res = append(res, str)
	}
	return
}

func BenchmarkEchoFor100_100(b *testing.B) {
	target := MakeCase(100, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoFor(target)
	}
}

func BenchmarkEchoJoin100_100(b *testing.B) {
	target := MakeCase(100, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoJoin(target)
	}
}

func BenchmarkEchoFor1_100(b *testing.B) {
	target := MakeCase(1, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoFor(target)
	}
}

func BenchmarkEchoJoin1_100(b *testing.B) {
	target := MakeCase(1, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoJoin(target)
	}
}

func BenchmarkEchoFor100_1(b *testing.B) {
	target := MakeCase(100, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoFor(target)
	}
}

func BenchmarkEchoJoin100_1(b *testing.B) {
	target := MakeCase(100, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EchoJoin(target)
	}
}
