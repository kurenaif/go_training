package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// MultipleSortInterface

type MultipleSortIntarface struct {
	sortInterface []sort.Interface
}

func (s *MultipleSortIntarface) Add(sorter sort.Interface) {
	s.sortInterface = append(s.sortInterface, sorter)
}

func (s MultipleSortIntarface) Len() int {
	if s.sortInterface == nil || len(s.sortInterface) == 0 {
		log.Printf("sort key is not setted. (may be sort is not enabled)")
		return 0
	}
	return s.sortInterface[0].Len()
}

func (s MultipleSortIntarface) Less(i, j int) bool {
	for _, sorter := range s.sortInterface {
		if sorter.Less(i, j) != sorter.Less(j, i) { //入れ替えると結果が変わる=>iとjは等しくない
			return sorter.Less(i, j)
		}
	}
	return false
}

func (s MultipleSortIntarface) Swap(i, j int) {
	if s.sortInterface == nil || len(s.sortInterface) == 0 {
		log.Printf("sort key is not setted. (may be sort is not enabled)")
		return
	}
	s.sortInterface[0].Swap(i, j)
}

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	//!-customcall
	printTracks(tracks)

	// 練習問題7.8 sort.Interface ver.
	fmt.Println("\nMultipleSortIntarface:")
	var sorter MultipleSortIntarface
	sorter.Add(byTitle(tracks))
	sorter.Add(byYear(tracks))
	sorter.Add(byLength(tracks))
	sort.Sort(sorter)
	printTracks(tracks)

	// 練習問題7.8 sort.Stable ver.
	fmt.Println("\nStableSort:")
	sort.Sort(byYear(tracks)) //一回崩す
	for _, sorter := range sorter.sortInterface {
		sort.Stable(sorter)
	}
	printTracks(tracks)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
