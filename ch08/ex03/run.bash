#!/bin/bash

set -x

go run ../reverb1/reverb.go &
sleep 1
go run ./netcat.go
