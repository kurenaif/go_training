package popcountloop

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) (cnt int) {
	// 注意！！！ 演算順序はbitshiftのほうが先！
	for i := uint(0); i < 8; i++ {
		cnt += int(pc[byte(x>>(i*8))])
	}
	return
}
