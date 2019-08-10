package track

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type ByTitle []*Track

func (x ByTitle) Len() int           { return len(x) }
func (x ByTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x ByTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByArtist []*Track

func (x ByArtist) Len() int           { return len(x) }
func (x ByArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x ByArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByAlbum []*Track

func (x ByAlbum) Len() int           { return len(x) }
func (x ByAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x ByAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByYear []*Track

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ByLength []*Track

func (x ByLength) Len() int           { return len(x) }
func (x ByLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x ByLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
