#!/bin/bash

set -eux

cat out.gif | go run main.go > /dev/null
cat out.jpg | go run main.go > /dev/null
cat out.png | go run main.go > /dev/null

cat out.gif | go run main.go -e png > out.gif.png
file out.gif.png
cat out.gif | go run main.go -e jpg > out.gif.jpg
file out.gif.jpg
cat out.gif | go run main.go -e gif > out.gif.gif
file out.gif.gif

cat out.jpg | go run main.go -e png > out.jpg.png
file out.jpg.png
cat out.jpg | go run main.go -e jpg > out.jpg.jpg
file out.jpg.jpg
cat out.jpg | go run main.go -e gif > out.jpg.gif
file out.jpg.gif

cat out.png | go run main.go -e png > out.png.png
file out.png.png
cat out.png | go run main.go -e jpg > out.png.jpg
file out.png.jpg
cat out.png | go run main.go -e gif > out.png.gif
file out.png.gif

rm out.gif.gif out.gif.jpg out.gif.png out.jpg.gif out.jpg.jpg out.jpg.png out.png.gif out.png.jpg out.png.png
