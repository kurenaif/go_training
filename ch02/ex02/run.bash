set -ux
# コマンドライン入力
go run ./main.go 3 5.4 0.0 -2.0
# コマンドライン入力(不正)
go run ./main.go test
# コマンドライン入力(一部不正)
go run ./main.go 3 test 5
# 標準入力
echo 3 5.4 0.0 -2.0 | go run ./main.go
# 標準入力(不正)
echo test | go run ./main.go
# 標準入力(一部不正)
echo 3 test 5 | go run ./main.go
# 標準入力(ユーザー入力)
go run ./main.go
