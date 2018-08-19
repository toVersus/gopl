package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

type byRecentClick struct {
	tracks []*Track
	less   func(track1, track2 *Track) bool
}

func (x byRecentClick) Len() int           { return len(x.tracks) }
func (x byRecentClick) Less(i, j int) bool { return x.less(x.tracks[i], x.tracks[j]) }
func (x byRecentClick) Swap(i, j int)      { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

func main() {
	clicked := []string{"Artist", "Title", "Title", "Year", "Artist"}

	lessfn := func(track1, track2 *Track) bool {
		if len(clicked) == 0 {
			return track1.Title < track2.Title
		}

		return lessfuncs[clicked[0]](track1, track2)
	}

	sort.Sort(byRecentClick{tracks, lessfn})
	printTracks(tracks)
}
