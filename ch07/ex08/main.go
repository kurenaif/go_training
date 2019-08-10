package main

import (
	"fmt"
	"go_training/ch07/ex08/mulsort"
	"go_training/ch07/ex08/track"
	"sort"
)

var tracks = []*track.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, track.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, track.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, track.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, track.Length("4m24s")},
}

func main() {
	fmt.Println("byArtist:")
	sort.Sort(track.ByArtist(tracks))
	track.PrintTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(track.ByArtist(tracks)))
	track.PrintTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(track.ByYear(tracks))
	track.PrintTracks(tracks)

	fmt.Println("\nCustom:")
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *track.Track) bool {
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
	track.PrintTracks(tracks)

	// 練習問題7.8 sort.Interface ver.
	fmt.Println("\nMultipleSortIntarface:")
	var sorters mulsort.MultipleSortIntarface
	sorters.Add(track.ByTitle(tracks))
	sorters.Add(track.ByYear(tracks))
	sorters.Add(track.ByLength(tracks))
	sort.Sort(sorters)
	track.PrintTracks(tracks)

	// 練習問題7.8 sort.Stable ver.
	fmt.Println("\nStableSort:")
	sort.Sort(track.ByYear(tracks)) //一回崩す
	for sorter := sorters.Next(); sorter != nil; sorter = sorters.Next() {
		sort.Stable(*sorter)
	}
	track.PrintTracks(tracks)
}

type customSort struct {
	t    []*track.Track
	less func(x, y *track.Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
