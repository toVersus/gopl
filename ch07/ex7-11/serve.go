package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := product{db: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type product struct {
	sync.Mutex
	db map[string]dollars
}

func (p *product) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	newPrice, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request for price: %+v\n", priceStr)
		return
	}

	p.Lock()
	p.db[item] = dollars(newPrice)
	p.Unlock()
}

func (p *product) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := p.db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	p.Lock()
	delete(p.db, item)
	p.Unlock()
}

func (p *product) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := p.db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	p.Lock()
	fmt.Fprintf(w, "%s\n", price)
	p.Unlock()
}

func (p *product) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := p.db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	priceStr := req.URL.Query().Get("price")
	newPrice, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request for price: %+v\n", priceStr)
		return
	}

	p.Lock()
	p.db[item] = dollars(newPrice)
	p.Unlock()
}
