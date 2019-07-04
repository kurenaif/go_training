# ベンチマーク結果

* `BenchmarkEchoFor`: for文を使った文字列生成
* `BenchmarkEchoJoin`: Joinを使った文字列生成
* `xxx_yyy`: xxxが配列の要素数、yが配列の1要素の文字列の長さ
    * e.g.) 2_3ならば ["aaa", "aaa"]

```
+ go test -bench=. -test.benchmem
goos: linux
goarch: amd64
pkg: go_training/ch01/ex03
BenchmarkEchoFor100_100-8    	   10000	    120372 ns/op	  540720 B/op	      99 allocs/op
BenchmarkEchoJoin100_100-8   	  300000	      5944 ns/op	   20480 B/op	       2 allocs/op
BenchmarkEchoFor1_100-8      	100000000	        12.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkEchoJoin1_100-8     	300000000	         4.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkEchoFor100_1-8      	  200000	      8744 ns/op	   10768 B/op	      99 allocs/op
BenchmarkEchoJoin100_1-8     	 2000000	       951 ns/op	     416 B/op	       2 allocs/op
PASS
ok  	go_training/ch01/ex03	10.897s
```

## 所感

* 配列の長さ、文字列の長さ両方に依存？
* Joinに比べ、Forのほうがメモリ割り当て回数が圧倒的に少ないので早い？
* 配列の長さ1,文字列の長さ100のケースで差がでた要因が謎 sepにスペースを代入しているせい？
* 参考実装: join https://golang.org/src/strings/strings.go?s=10789:10829#L415
* 参考実装: concatstring: https://github.com/golang/go/blob/master/src/runtime/string.go#L23
* joinは配列の長さが1ならそのままreturnする仕様になっていた
* +はそうではないので、要素数1はその違い？
