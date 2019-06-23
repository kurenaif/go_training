package popcountbitshift

// ポピュレーションカウントを最下位bitの検査を64回繰り返すことで求め、その結果を返す
func PopCount(x uint64) (cnt int) {
	for i := uint64(0); i < 64; i++ {
		cnt += int(x & 1)
		x >>= 1
	}
	return cnt
	// 以下のようなコードでも実現可能だが、p.50には64回繰り返すと書いてあるので、そちらに従う
	/*
		for ; x > 0; x >>= 1 {
			if (x & 1) == 1 {
				cnt++
			}
		}
	*/
}
