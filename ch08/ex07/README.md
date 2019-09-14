depth指定しないと実行時間かかるのでdepth指定できるようにしました。
depthのex06と意味は同じです。
depthのぶんで取れるだけ取って、取れなかった分はミラーライングしないようになっています。

*```
go run main.go -depth 2 https://golang.org/
```
