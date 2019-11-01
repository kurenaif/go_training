package main

import (
	"go_training/ch02/ex03/popcount"
	"go_training/ch02/ex03/popcountloop"
	"go_training/ch02/ex04/popcountbitshift"
	"go_training/ch02/ex05/popcountlsb"
	"testing"
)

func makeNum(num int) uint64 {
	res := uint64(0)
	for i := 0; i < num; i++ {
		res = (res << 1) + 1
	}
	return res
}

var result int

func benchmark(b *testing.B, size int, target func(x uint64) int) {
	num := makeNum(size)
	b.ResetTimer()
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += target(num)
	}
	result = temp
}

// --------------------------------------------------------------------------------
// PopCount
// --------------------------------------------------------------------------------

// 1のケース後半の演習の比較用
func BenchmarkPopCount00(b *testing.B) { benchmark(b, 0, popcount.PopCount) }
func BenchmarkPopCount01(b *testing.B) { benchmark(b, 1, popcount.PopCount) }
func BenchmarkPopCount02(b *testing.B) { benchmark(b, 2, popcount.PopCount) }
func BenchmarkPopCount04(b *testing.B) { benchmark(b, 4, popcount.PopCount) }
func BenchmarkPopCount08(b *testing.B) { benchmark(b, 8, popcount.PopCount) }
func BenchmarkPopCount16(b *testing.B) { benchmark(b, 16, popcount.PopCount) }
func BenchmarkPopCount32(b *testing.B) { benchmark(b, 32, popcount.PopCount) }
func BenchmarkPopCount64(b *testing.B) { benchmark(b, 64, popcount.PopCount) }

// --------------------------------------------------------------------------------
// PopCountLoop
// --------------------------------------------------------------------------------

func BenchmarkPopCountLoop00(b *testing.B) { benchmark(b, 0, popcountloop.PopCount) }
func BenchmarkPopCountLoop01(b *testing.B) { benchmark(b, 1, popcountloop.PopCount) }
func BenchmarkPopCountLoop02(b *testing.B) { benchmark(b, 2, popcountloop.PopCount) }
func BenchmarkPopCountLoop04(b *testing.B) { benchmark(b, 4, popcountloop.PopCount) }
func BenchmarkPopCountLoop08(b *testing.B) { benchmark(b, 8, popcountloop.PopCount) }
func BenchmarkPopCountLoop16(b *testing.B) { benchmark(b, 16, popcountloop.PopCount) }
func BenchmarkPopCountLoop32(b *testing.B) { benchmark(b, 32, popcountloop.PopCount) }
func BenchmarkPopCountLoop64(b *testing.B) { benchmark(b, 64, popcountloop.PopCount) }

// --------------------------------------------------------------------------------
// PopCountBitshift
// --------------------------------------------------------------------------------

func BenchmarkPopCountBitShift00(b *testing.B) { benchmark(b, 0, popcountbitshift.PopCount) }
func BenchmarkPopCountBitShift01(b *testing.B) { benchmark(b, 1, popcountbitshift.PopCount) }
func BenchmarkPopCountBitShift02(b *testing.B) { benchmark(b, 2, popcountbitshift.PopCount) }
func BenchmarkPopCountBitShift04(b *testing.B) { benchmark(b, 4, popcountbitshift.PopCount) }
func BenchmarkPopCountBitShift08(b *testing.B) { benchmark(b, 8, popcountbitshift.PopCount) }
func BenchmarkPopCountBitShift16(b *testing.B) { benchmark(b, 16, popcountbitshift.PopCount) }
func BenchmarkPopCountBitShift32(b *testing.B) { benchmark(b, 32, popcountbitshift.PopCount) }
func BenchmarkPopCountBitShift64(b *testing.B) { benchmark(b, 64, popcountbitshift.PopCount) }

// --------------------------------------------------------------------------------
// PopCountLSB
// --------------------------------------------------------------------------------

func BenchmarkPopCountLSB00(b *testing.B) { benchmark(b, 0, popcountlsb.PopCount) }
func BenchmarkPopCountLSB01(b *testing.B) { benchmark(b, 1, popcountlsb.PopCount) }
func BenchmarkPopCountLSB02(b *testing.B) { benchmark(b, 2, popcountlsb.PopCount) }
func BenchmarkPopCountLSB04(b *testing.B) { benchmark(b, 4, popcountlsb.PopCount) }
func BenchmarkPopCountLSB08(b *testing.B) { benchmark(b, 8, popcountlsb.PopCount) }
func BenchmarkPopCountLSB16(b *testing.B) { benchmark(b, 16, popcountlsb.PopCount) }
func BenchmarkPopCountLSB32(b *testing.B) { benchmark(b, 32, popcountlsb.PopCount) }
func BenchmarkPopCountLSB64(b *testing.B) { benchmark(b, 64, popcountlsb.PopCount) }

// result
/*
BenchmarkPopCount00-8           	250021837	         4.66 ns/op
BenchmarkPopCount01-8           	260291910	         4.57 ns/op
BenchmarkPopCount02-8           	253419913	         4.72 ns/op
BenchmarkPopCount04-8           	257806824	         4.73 ns/op
BenchmarkPopCount08-8           	259122190	         4.62 ns/op
BenchmarkPopCount16-8           	255775983	         4.69 ns/op
BenchmarkPopCount32-8           	259811131	         4.59 ns/op
BenchmarkPopCount64-8           	261461794	         4.72 ns/op
BenchmarkPopCountLoop00-8       	59090173	        18.5 ns/op
BenchmarkPopCountLoop01-8       	64687146	        18.3 ns/op
BenchmarkPopCountLoop02-8       	64493756	        18.4 ns/op
BenchmarkPopCountLoop04-8       	63047233	        18.4 ns/op
BenchmarkPopCountLoop08-8       	65976120	        18.7 ns/op
BenchmarkPopCountLoop16-8       	61378778	        18.2 ns/op
BenchmarkPopCountLoop32-8       	59594438	        18.3 ns/op
BenchmarkPopCountLoop64-8       	64279242	        18.5 ns/op
BenchmarkPopCountBitShift00-8   	28188894	        39.9 ns/op
BenchmarkPopCountBitShift01-8   	28768699	        40.2 ns/op
BenchmarkPopCountBitShift02-8   	30271165	        40.3 ns/op
BenchmarkPopCountBitShift04-8   	29723136	        39.9 ns/op
BenchmarkPopCountBitShift08-8   	27618602	        39.6 ns/op
BenchmarkPopCountBitShift16-8   	30164505	        40.3 ns/op
BenchmarkPopCountBitShift32-8   	29581111	        40.3 ns/op
BenchmarkPopCountBitShift64-8   	29281387	        40.2 ns/op
BenchmarkPopCountLSB00-8        	560460021	         2.16 ns/op
BenchmarkPopCountLSB01-8        	412365390	         2.92 ns/op
BenchmarkPopCountLSB02-8        	303135662	         3.51 ns/op
BenchmarkPopCountLSB04-8        	266255257	         4.60 ns/op
BenchmarkPopCountLSB08-8        	181005154	         6.91 ns/op
BenchmarkPopCountLSB16-8        	100000000	        11.0 ns/op
BenchmarkPopCountLSB32-8        	59944405	        19.9 ns/op
BenchmarkPopCountLSB64-8        	26122514	        43.3 ns/op
*/
