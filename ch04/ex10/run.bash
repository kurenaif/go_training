set -eux
token=`cat token`
url="https://github.com/rust-lang/rust"
# echo $token $url
go run main.go $token $url
