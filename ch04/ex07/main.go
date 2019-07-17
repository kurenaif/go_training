// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	//!+array
	bs := []byte(string("😃😥😣→↑←↓DDR❢鬼人正邪アマノジャクa"))
	reverse(bs)
	fmt.Println(string(bs))
	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bytes := []byte(input.Text())
		reverse(bytes)
		fmt.Printf("%v\n", string(bytes))
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(bs []byte) {
	lIndex := 0 // 格納用Index
	rIndex := len(bs)
	lSpace := 0 // Sizeが違うものを交換したあとに生じる隙間
	rSpace := 0

	var lR, rR rune
	var lSize, rSize int

	// i: 移動するruneのbyteのはじめ
	// j: 移動するruneのbyteの終わり
	for i, j := 0, len(bs); i < (j - 1); { // 半開区間なので、[4:]と[:5]は同じものを指す。なので、j-1したものと比較する
		// 使用後ならば補充する(popするイメージ)
		// この代入はそのbyte列からそのruneを取り除くイメージ
		if lSize == 0 {
			lR, lSize = utf8.DecodeRune(bs[i:]) // 取得したサイズ
			lSpace += lSize
			i += lSize
		}
		if rSize == 0 {
			rR, rSize = utf8.DecodeLastRune(bs[:j])
			rSpace += rSize
			j -= rSize
		}
		// fmt.Fprintf(os.Stderr, "%c(%d): %d, %c(%d): %d\n", lR, lR, lSize, rR, rR, rSize)
		// fmt.Fprintf(os.Stderr, "lSpace: %d, rSpace: %d\n", lSpace, rSpace)
		// fmt.Fprintf(os.Stderr, "lIndex: %d, rIndex: %d\n", lIndex, rIndex)

		// 必ず下の2つのif文のどちらかは実行されることは証明可能。
		// 左側に代入するスペースが残っている
		if rSize != 0 && lSpace >= rSize {
			_ = utf8.EncodeRune(bs[lIndex:], rR) // => rSize
			lIndex += rSize
			lSpace -= rSize
			rSize = 0
			rR = 0
		}
		if lSize != 0 && rSpace >= lSize {
			_ = utf8.EncodeRune(bs[(rIndex-lSize):], lR) // => lSize
			rIndex -= lSize                              // 次に代入できるエリアの右端
			rSpace -= lSize                              // 現在合いているスペース
			lSize = 0
			lR = 0
		}
	}

	// lSpace= 2, rSpace = 4, lSize = 3, rSize = 3みたいなケースもあり
	// lSpace= 2, rSpace = 1, lSize = 0, rSize = 3になるケースが考えられる
	if rSize != 0 {
		_ = utf8.EncodeRune(bs[lIndex:], rR) // => rSize
		lIndex += rSize
		lSpace -= rSize
		rSize = 0
		rR = 0
	}
	if lSize != 0 {
		_ = utf8.EncodeRune(bs[(rIndex-lSize):], lR) // => lSize
		rIndex -= lSize                              // 次に代入できるエリアの右端
		rSpace -= lSize                              // 現在合いているスペース
		lSize = 0
		lR = 0
	}
}

//!-rev
/* 余りが出るケース: 😃😥😣→↑←↓DDR❢鬼人正邪アマノジャクa
😃(128515): 4, a(97): 1
lSpace: 4, rSpace: 1
lIndex: 0, rIndex: 61
😃(128515): 4, ク(12463): 3
lSpace: 3, rSpace: 4
lIndex: 1, rIndex: 61
😥(128549): 4, ャ(12515): 3
lSpace: 4, rSpace: 3
lIndex: 4, rIndex: 57
😥(128549): 4, ジ(12472): 3
lSpace: 1, rSpace: 6
lIndex: 7, rIndex: 57
😣(128547): 4, ジ(12472): 3
lSpace: 5, rSpace: 2
lIndex: 7, rIndex: 53
😣(128547): 4, ノ(12494): 3
lSpace: 2, rSpace: 5
lIndex: 10, rIndex: 53
→(8594): 3, ノ(12494): 3
lSpace: 5, rSpace: 1
lIndex: 10, rIndex: 49
→(8594): 3, マ(12510): 3
lSpace: 2, rSpace: 4
lIndex: 13, rIndex: 49
↑(8593): 3, マ(12510): 3
lSpace: 5, rSpace: 1
lIndex: 13, rIndex: 46
↑(8593): 3, ア(12450): 3
lSpace: 2, rSpace: 4
lIndex: 16, rIndex: 46
←(8592): 3, ア(12450): 3
lSpace: 5, rSpace: 1
lIndex: 16, rIndex: 43
←(8592): 3, 邪(37034): 3
lSpace: 2, rSpace: 4
lIndex: 19, rIndex: 43
↓(8595): 3, 邪(37034): 3
lSpace: 5, rSpace: 1
lIndex: 19, rIndex: 40
↓(8595): 3, 正(27491): 3
lSpace: 2, rSpace: 4
lIndex: 22, rIndex: 40
D(68): 1, 正(27491): 3
lSpace: 3, rSpace: 1
lIndex: 22, rIndex: 37
D(68): 1, 人(20154): 3
lSpace: 1, rSpace: 3
lIndex: 25, rIndex: 36
R(82): 1, 人(20154): 3
lSpace: 2, rSpace: 2
lIndex: 25, rIndex: 35
❢(10082): 3, 人(20154): 3
lSpace: 5, rSpace: 1
lIndex: 25, rIndex: 34
❢(10082): 3, 鬼(39740): 3
lSpace: 2, rSpace: 4
lIndex: 28, rIndex: 34
aクャジノマア邪正人���❢RDD↓←↑→😣😥😃

*/
