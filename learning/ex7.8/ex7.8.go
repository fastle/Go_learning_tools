package main

import (
	"fmt"
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

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout,0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 计算各列宽度， 输出表格
}

func length(s string) time.Duration {
    d, err := time.ParseDuration(s)
    if err != nil {
        panic(s)
    }
    return d
}


type table struct {
	t    []*Track
	keys []string // keys to sort by
}

func (t table) Len() int {
	return len(t.t)
}

func (t table) Less(i, j int) bool {
	for p := len(t.keys) - 1; p >= 0; p--{
		fmt.Println(t.keys[p])
		switch t.keys[p] {
		case "Title":
			if t.t[i].Title != t.t[j].Title {
				return t.t[i].Title < t.t[j].Title
			}
		case "Artist":
			if t.t[i].Artist != t.t[j].Artist {
				return t.t[i].Artist < t.t[j].Artist
			}
		case "Album":
			if t.t[i].Album != t.t[j].Album {
				return t.t[i].Album < t.t[j].Album
			}
		case "Year":
			if t.t[i].Year != t.t[j].Year {
				return t.t[i].Year < t.t[j].Year
			}
		}
	}
	return false // all keys are equal
}

func (t table) Swap(i, j int) {
	t.t[i], t.t[j] = t.t[j], t.t[i]
}

func setPrime(t *table, key string) {
	t.keys = append(t.keys, key)
}
func main() {
	table := table{tracks, []string{}}
	setPrime(&table, "Year")
	setPrime(&table, "Title")
	sort.Sort(table)
	PrintTracks(table.t)
}