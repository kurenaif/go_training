package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif" // register GIF decoder
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	format := flag.String("e", "png", "image extension(jpg, png or gif) default: png")
	flag.Parse()
	switch *format {
	case "png":
		if err := toPNG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	case "jpg":
		if err := toJPEG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	case "gif":
		if err := toGIF(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, nil)
}
