package main

import (
	"go_training/ch02/ex03/popcount"
	"go_training/ch02/ex03/popcountloop"
	"go_training/ch02/ex04/popcountbitshift"
	"go_training/ch02/ex05/popcountlsb"
	"math/rand"
	"strconv"
	"testing"
)

// --------------------------------------------------------------------------------
// PopCount
// --------------------------------------------------------------------------------

var result int

// 1のケース後半の演習の比較用
func BenchmarkPopCount11(b *testing.B) {
	// 二進数リテラルはないためParseUintで代用 0xFFFFでもよかった…？
	num, _ := strconv.ParseUint("1111111111111111111111111111111111111111111111111111111111111111", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCount(num)
	}
	result = temp
}

// 0のケース後半の演習の比較用
func BenchmarkPopCount00(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCount(num)
	}
	result = temp
}

// BenchmarkPopCountRand比較用
// 乱数のケースと比較するために特に代入せずに乱数だけ回すケース
func BenchmarkPopCount00rand(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCount(num)
		rand.Uint64() // 乱数のケースと比較するために乱数生成時間をあえてここで発生させる
	}
	result = temp
}

// 乱数のケース アクセスがバラバラになるので、ちょっと遅くなってほしい 完全な連続アクセスではないので意味はない…？
func BenchmarkPopCountRand(b *testing.B) {
	// 乱数は条件によって変わらないように固定する
	// 暗黙的に1で固定される(https://github.com/golang/go/blob/c8aec4095e089ff6ac50d18e97c3f46561f14f48/src/math/rand/rand.go#L236)が、ここではわかりやすくるために明示的に1で固定
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCount(rand.Uint64())
	}
	result = temp
}

// --------------------------------------------------------------------------------
// PopCountLoop
// --------------------------------------------------------------------------------

func BenchmarkPopCountLoop11(b *testing.B) {
	// 二進数リテラルはないためParseUintで代用 0xFFFFでもよかった…？
	num, _ := strconv.ParseUint("1111111111111111111111111111111111111111111111111111111111111111", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountloop.PopCount(num)
	}
	result = temp
}

func BenchmarkPopCountLoop00(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountloop.PopCount(num)
	}
	result = temp
}

func BenchmarkPopCountLoop00rand(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountloop.PopCount(num)
		rand.Uint64() // 乱数のケースと比較するために乱数生成時間をあえてここで発生させる
	}
	result = temp
}

func BenchmarkPopCountLoopRand(b *testing.B) {
	// 乱数は条件によって変わらないように固定する
	// 暗黙的に1で固定される(https://github.com/golang/go/blob/c8aec4095e089ff6ac50d18e97c3f46561f14f48/src/math/rand/rand.go#L236)が、ここではわかりやすくるために明示的に1で固定
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountloop.PopCount(rand.Uint64())
	}
	result = temp
}

// --------------------------------------------------------------------------------
// PopCountBitshift
// --------------------------------------------------------------------------------

func BenchmarkPopCountBitShift00(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountbitshift.PopCount(num)
	}
	result = temp
}

func BenchmarkPopCountBitShift11(b *testing.B) {
	// 二進数リテラルはないためParseUintで代用 0xFFFFでもよかった…？
	num, _ := strconv.ParseUint("1111111111111111111111111111111111111111111111111111111111111111", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountbitshift.PopCount(num)
	}
	result = temp
}

func BenchmarkPopCountBitShift00rand(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountbitshift.PopCount(num)
		rand.Uint64() // 乱数のケースと比較するために乱数生成時間をあえてここで発生させる
	}
	result = temp
}

func BenchmarkPopCountBitShiftRand(b *testing.B) {
	// 乱数は条件によって変わらないように固定する
	// 暗黙的に1で固定される(https://github.com/golang/go/blob/c8aec4095e089ff6ac50d18e97c3f46561f14f48/src/math/rand/rand.go#L236)が、ここではわかりやすくるために明示的に1で固定
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountbitshift.PopCount(rand.Uint64())
	}
	result = temp
}

// --------------------------------------------------------------------------------
// PopCountLSB
// --------------------------------------------------------------------------------

func BenchmarkPopCountLSB00(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountlsb.PopCount(num)
	}
	result = temp
}

func BenchmarkPopCountLSB11(b *testing.B) {
	// 二進数リテラルはないためParseUintで代用 0xFFFFでもよかった…？
	num, _ := strconv.ParseUint("1111111111111111111111111111111111111111111111111111111111111111", 2, 0)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountlsb.PopCount(num)
	}
	result = temp
}

func BenchmarkPopCountLSB00rand(b *testing.B) {
	num, _ := strconv.ParseUint("0000000000000000000000000000000000000000000000000000000000000000", 2, 0)
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountlsb.PopCount(num)
		rand.Uint64() // 乱数のケースと比較するために乱数生成時間をあえてここで発生させる
	}
	result = temp
}

func BenchmarkPopCountLSBRand(b *testing.B) {
	// 乱数は条件によって変わらないように固定する
	// 暗黙的に1で固定される(https://github.com/golang/go/blob/c8aec4095e089ff6ac50d18e97c3f46561f14f48/src/math/rand/rand.go#L236)が、ここではわかりやすくるために明示的に1で固定
	rand.Seed(1)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcountlsb.PopCount(rand.Uint64())
	}
	result = temp
}
