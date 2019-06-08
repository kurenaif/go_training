# ベンチマーク結果

* `BenchmarkEchoFor`: for文を使った文字列生成
* `BenchmarkEchoJoin`: Joinを使った文字列生成
* `xxx_yyy`: xxxが配列の要素数、yが配列の1要素の文字列の長さ
    * e.g.) 2_3ならば ["aaa", "aaa"]

```
go test -bench=. -test.benchmem
goos: linux
goarch: amd64
BenchmarkEchoFor100_100-8    	   20000	     88725 ns/op	  540720 B/op	      99 allocs/op
BenchmarkEchoJoin100_100-8   	  300000	      4203 ns/op	   20480 B/op	       2 allocs/op
BenchmarkEchoFor1_100-8      	100000000	        12.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkEchoJoin1_100-8     	300000000	         4.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkEchoFor100_1-8      	  200000	      7895 ns/op	   10768 B/op	      99 allocs/op
BenchmarkEchoJoin100_1-8     	 2000000	       959 ns/op	     416 B/op	       2 allocs/op
```

## 所感

* 配列の長さ、文字列の長さ両方に依存？
* Joinに比べ、Forのほうがメモリ割り当て回数が圧倒的に少ないので早い？
* 配列の長さ1,文字列の長さ100のケースで差がでた要因が謎 sepにスペースを代入しているせい？