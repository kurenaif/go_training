set -eux
go test -bench=.
go test -v ./...
