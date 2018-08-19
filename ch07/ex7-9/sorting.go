package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type LessFunc func(track1, track2 *Track) bool

var lessfuncs = map[string]LessFunc{
	"Title":  func(track1, track2 *Track) bool { return track1.Title < track2.Title },
	"Artist": func(track1, track2 *Track) bool { return track1.Artist < track2.Artist },
	"Album":  func(track1, track2 *Track) bool { return track1.Album < track2.Album },
	"Year":   func(track1, track2 *Track) bool { return track1.Year < track2.Year },
	"Length": func(track1, track2 *Track) bool { return track1.Length < track2.Length },
}

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

var trackTable = template.Must(template.New("Track Table").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
	  <meta charset="utf-8">
		<style>
		  table {
				border-collapse: collapse;
				border-spacing: 0px;
			}
		  table, th, td {
				padding: 5px;
				border: 1px solid black;
			}
		</style>
	</head>
	<body>
		<h1>Tracks</h1>
		<table>
		  <thead>
				<tr>
					<th><a href="/?sort=Title">Title</a></th>
					<th><a href="/?sort=Artist">Artist</a></th>
					<th><a href="/?sort=Album">Album</a></th>
					<th><a href="/?sort=Year">Year</a></th>
					<th><a href="/?sort=Length">Length</a></th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(w io.Writer, tracks []*Track) {
	if err := trackTable.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

type byRecentClick struct {
	tracks []*Track
	less   func(track1, track2 *Track) bool
}

func (x byRecentClick) Len() int           { return len(x.tracks) }
func (x byRecentClick) Less(i, j int) bool { return x.less(x.tracks[i], x.tracks[j]) }
func (x byRecentClick) Swap(i, j int)      { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sortBy := r.FormValue("sort")

		lessfn := func(track1, track2 *Track) bool {
			if sortBy == "" {
				return false
			}
			return lessfuncs[sortBy](track1, track2)
		}

		sort.Sort(byRecentClick{tracks, lessfn})
		printTracks(w, tracks)
	})
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
