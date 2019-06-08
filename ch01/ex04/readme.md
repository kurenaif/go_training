# 入力と出力

```
$go run ex04.go a.txt b.txt
(sum: 2 ):
	a.txt	1
	b.txt	1
hello(sum: 3 ):
hello	a.txt	3
hollow(sum: 2 ):
hollow	a.txt	1
hollow	b.txt	1
text(sum: 4 ):
text	a.txt	2
text	b.txt	2
```

## 所感

* mapは順序が変わるhashtable
	* テストはちゃんとmap同士で比較しなければならない
* `.ex4_test.go`(うまく行かないので隠しファイル) にやろうとした痕跡を残す
	* 標準入力のテストはソースコードの通りのやり方でできる
	* lineごとに昇順ソートしたまでは良いが、中身までソートしなければ単純な文字列比較では不可能
