GOARCH=amd64 go test -v
GOARCH=386 go test -v

GOARCH=amd64 go test -bench .
GOARCH=386 go test -bench .

