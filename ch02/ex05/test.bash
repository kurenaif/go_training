set -eux
go test -v ./...
go test -bench=.
