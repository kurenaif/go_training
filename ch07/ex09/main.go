// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

//!+template
import (
	"fmt"
	"go_training/ch07/ex08/mulsort"
	"go_training/ch07/ex08/track"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type TrackList struct {
	Items  []*track.Track
	Titles []*TitleRow
}

// 表示用
type TitleRow struct {
	Title    string
	SortKind string //0: none, 1: asc, 2: desc
	Link     string //0: none, 1: asc, 2: desc
}

// 内部処理用
type TitleOrder struct {
	Title    string
	SortKind Order //0: none, 1: asc, 2: desc
}

var trackList = template.Must(template.New("trackList").Parse(`
<table>
<tr style='text-align: left'>
{{range .Titles}}
<th><a href={{.Link}}>{{.Title}} {{.SortKind}}</a> </th>
{{end}}
</tr>
{{range .Items}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

var titles []string = []string{
	"Title",
	"Artist",
	"Album",
	"Year",
	"Length",
}

var title2Num = map[string]int{
	"Title":  0,
	"Artist": 1,
	"Album":  2,
	"Year":   3,
	"Length": 4,
}

type Order int

const (
	None Order = iota
	Asc
	Desc
)

func (o Order) String() string {
	switch o {
	case None:
		return "ー"
	case Asc:
		return "▼"
	case Desc:
		return "▲"
	default:
		return "Unknown"
	}
}

func title2Sorter(title string) sort.Interface {
	switch title {
	case "Title":
		return track.ByTitle(tracks)
	case "Artist":
		return track.ByArtist(tracks)
	case "Album":
		return track.ByAlbum(tracks)
	case "Year":
		return track.ByYear(tracks)
	case "Length":
		return track.ByLength(tracks)
	}
	return nil
}

// sortOrder: 既存の並び順
// add: 追加する並び順
// 手前についかする！
func makeQueryParameter(sortOrder []TitleOrder, add TitleOrder) string {
	orders := []TitleOrder{}
	orders = append(orders, add)
	for _, order := range sortOrder {
		if order.Title == add.Title { //同じ場合、最初に持ってくる
			continue
		}
		orders = append(orders, order)
	}

	res := ""
	sep := ""
	for _, order := range orders {
		res += sep + order.Title + "=" + strconv.Itoa(int(order.SortKind))
		sep = "&"
	}
	return res
}

func printTracks(writer io.Writer, tracks []*track.Track, sortOrder []TitleOrder) {

	var sorters mulsort.MultipleSortIntarface
	for _, order := range sortOrder {
		sorter := title2Sorter(order.Title)
		if sorter == nil {
			log.Printf("title %q is not found. skip.", order.Title)
		}

		switch order.SortKind {
		case Asc:
			sorters.Add(sorter)
		case Desc:
			sorters.Add(sort.Reverse(sorter))
		}
	}
	sort.Sort(sorters)

	titleRow := []*TitleRow{}
	for _, title := range titles {
		titleRow = append(titleRow, &TitleRow{title, Order(0).String(), "http://localhost:8000?" + makeQueryParameter(sortOrder, TitleOrder{title, None})})
	}
	for _, order := range sortOrder {
		titleRow[title2Num[order.Title]].SortKind = Order(order.SortKind).String()
		nextKind := (order.SortKind + 1) % 3
		titleRow[title2Num[order.Title]].Link = "http://localhost:8000?" + makeQueryParameter(sortOrder, TitleOrder{order.Title, nextKind})
	}

	result := TrackList{tracks, titleRow}
	if err := trackList.Execute(writer, result); err != nil {
		log.Fatal(err)
	}
	if err := trackList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

var tracks = []*track.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, track.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, track.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, track.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, track.Length("4m24s")},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		sortOrder := []TitleOrder{}
		for k, v := range r.Form {
			for _, title := range titles {
				if k == title {
					num, err := strconv.Atoi(v[0])
					if err != nil {
						log.Printf("parameter %s skip. because: %v", k, err)
						continue
					}
					if num < 0 || 2 < num {
						log.Printf("parameter %s skip. because: num must be [0,2]", k)
						continue
					}
					sortOrder = append(sortOrder, TitleOrder{k, Order(num)})
				}
			}
		}
		printTracks(w, tracks, sortOrder)
	})
	fmt.Println("http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
