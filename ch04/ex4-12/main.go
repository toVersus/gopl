package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/toversus/gopl/ch04/ex4-12/xkcd"
)

const usage = `xkcd get INDEX`

func main() {
	if len(os.Args) < 2 {
		log.Fatalln(usage)
	}
	action := os.Args[1]
	if action != "get" {
		return
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("INDEX (%s) must be an integer", os.Args[1])
	}
	comic, err := xkcd.GetComic(n)
	if err != nil {
		log.Fatal("Error getting comic", err)
	}
	fmt.Println(comic)
}
