# go_training

## メモ

```
$ go run filename.go
```
以外で実行する必要がある場合、`run.bash` に書く

```
$ go test -v # (ファイル名なし)
```
ではない場合、`test.bash` に書く

## 第1章 チュートリアル

* p.6 `i++` は式であり文ではない
* p.12 ネストしたmapは扱いがしんどい https://stackoverflow.com/questions/44305617/nested-maps-in-golang

### p.23 server2の挙動についてのメモ

firefoxとcurlでは `/count` は正常に動作したが、chromeでのみ、 `/count` にアクセスするたびにcountがインクリメントされていった。

環境: ubuntu18.10, chrome Version 74.0.3729.131 (Official Build) (64-bit)

原因: chromeはアクセスするたびに `favicon.ico` を取りに行くため

### p.23 server2

スラッシュで終わっているハンドラのパターンはそのパターンを接頭辞として持つすべてのURLに一致する。: `count/` とかくと、`count/xxx` などもすべて `counter()`が呼ばれる


## 質問queue

* p.9 配列の要素数1で差が出た要因がわからない
* 標準入力のテストのベストプラクティスを知りたい
	* 関数に分ける？ ex04みたいに頑張る？
