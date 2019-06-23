package popcountlsb

// 最下位bitを0にすることを繰り返すことで、PopCountを行う
func PopCount(x uint64) (cnt int) {
	for x > 0 {
		x = x & (x - 1)
		cnt++
	}
	return cnt
}
