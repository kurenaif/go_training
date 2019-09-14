#!/bin/bash
set -eux

go run main.go -depth 2 https://golang.org
cd golang.org
python3 -m http.server 8000 
