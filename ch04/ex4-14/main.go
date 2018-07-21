package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/toversus/gopl/ch04/ex4-14/github"
)

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a><td>
</tr>
{{end}}
</table>
`))

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

type IssueCache []github.Issue

func NewIssueCache(owner, repo string) (ic IssueCache, err error) {
	return github.GetIssues(owner, repo)
}

func (ic IssueCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.SplitN(r.URL.Path, "/", -1)
	if len(path) < 3 || path[2] == "" {
		log.Fatalf("invalid url: %s\n", r.URL.Path)
	}
	numStr := path[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Issue number isn't a number: '%s'", numStr)))
		if err != nil {
			log.Fatalf("Error writing response for %s: %s", r.Body, err)
		}
		return
	}
	issueTemplate.Execute(w, ic[num])
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("usage: githubserver OWNER REPO")
	}
	owner := os.Args[1]
	repo := os.Args[2]

	issueCache, err := NewIssueCache(owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", issueCache)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
