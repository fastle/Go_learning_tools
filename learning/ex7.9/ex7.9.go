//使用html/template包（§4.6）替代printTracks将tracks展示成一个HTML表格。将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
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

var trackTable = template.Must(template.New("Track").Parse(`
<h1> Tracks </h1>
<table>
<tr style='text-align: left'>
    <th onclick="submitform('Title')">Title
        <form action="" name="Title" method="post">
            <input type="hidden" name="orderby" value="Title"/>
        </form>
    </th>
    <th>Artist
        <form action="" name="Artist" method="post">
            <input type="hidden" name="orderby" value="Artist"/>
        </form>
    </th>
    <th>Album
        <form action="" name="Album" method="post">
            <input type="hidden" name="orderby" value="Album"/>
        </form>
    </th>
    <th onclick="submitform('Year')">Year
        <form action="" name="Year" method="post">
            <input type="hidden" name="orderby" value="Year"/>
        </form>
    </th>
    <th onclick="submitform('Length')">Length
        <form action="" name="Length" method="post">
            <input type="hidden" name="orderby" value="Length"/>
        </form>
    </th>
</tr>
{{range .T}}
<tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
    <td>{{.Album}}</td>
    <td>{{.Year}}</td>
    <td>{{.Length}}</td>
</tr>
{{end}}
</table>

<script>
function submitform(formname) {
    document[formname].submit();
}
</script>
`))

func length(s string) time.Duration {
    d, err := time.ParseDuration(s)
    if err != nil {
        panic(s)
    }
    return d
}


type table struct {
	T    []*Track
	keys []string // keys to sort by
}

func (t table) Len() int {
	return len(t.T)
}

func (t table) Less(i, j int) bool {
	for p := len(t.keys) - 1; p >= 0; p--{
		fmt.Println(t.keys[p])
		switch t.keys[p] {
		case "Title":
			if t.T[i].Title != t.T[j].Title {
				return t.T[i].Title < t.T[j].Title
			}
		case "Artist":
			if t.T[i].Artist != t.T[j].Artist {
				return t.T[i].Artist < t.T[j].Artist
			}
		case "Album":
			if t.T[i].Album != t.T[j].Album {
				return t.T[i].Album < t.T[j].Album
			}
		case "Year":
			if t.T[i].Year != t.T[j].Year {
				return t.T[i].Year < t.T[j].Year
			}
		}
	}
	return false // all keys are equal
}

func (t table) Swap(i, j int) {
	t.T[i], t.T[j] = t.T[j], t.T[i]
}

func setPrime(t *table, key string) {
	t.keys = append(t.keys, key)
}


func printTracks(w io.Writer, x *table) {
        trackTable.Execute(w, x)
}
func main() {
	table := table{tracks, []string{}}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if err := r.ParseForm(); err != nil {
            fmt.Printf("ParseForm: %v\n", err)
        }
		for k, v := range r.Form {
			if k == "orderby" {
				setPrime(&table, v[0])
			}
		}
		sort.Sort(table)
		printTracks(w, &table)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}