set -eux
go test -bench=. -test.benchmem
