#!/bin/bash

set -x

go run ./reverb.go &
sleep 1
go run ../ex03/netcat.go
