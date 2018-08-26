package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var listTmpl = template.Must(template.New("list").Parse(`
<html>
<body>
<table>
	<tr>
		<th>item</th>
		<th>price</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td>{{$k}}</td>
		<td>{{$v}}</td>
	</tr>
{{end}}
</table>
</body>
</html>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	listTmpl.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
