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


# module関連

* `go mod init` 使ってディレクトリ内で完結している
* あるいはgithub.comなんちゃら

# ベンチマーク書くときの注意

副作用なしのコードを書くと消滅する可能性がある

# めも

* 演習問題3.8: すごい拡大したものを出力しないと違いがわからん
* 演習問題3.8: big.Ratはめちゃめちゃでかいのでiteration数を減らさないといけない
* 演習問題4.3: 配列のサイズは決め打ちで
* 演習問題4.4: gcd
* 演習問題4.13: お金払わないとできないのでやらなくていい
* 演習問題4.14: 一回のアクセスでデータを持っておいてなんかする

## 質問queue
