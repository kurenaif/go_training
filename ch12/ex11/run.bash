set -eux

go run main.go &
sleep 1
curl 'http://localhost:12345/search'
curl 'http://localhost:12345/search?l=golang&l=programming'
curl 'http://localhost:12345/search?l=golang&l=programming&max=100'
curl 'http://localhost:12345/search?x=true&l=golang&l=programming'
curl 'http://localhost:12345/search?q=hello&x=123'
curl 'http://localhost:12345/search?q=hello&max=lots'
