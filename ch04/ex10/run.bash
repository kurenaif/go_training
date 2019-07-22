set -eux
token=`cat token`
go run main.go $token  https://github.com/rust-lang/rust
