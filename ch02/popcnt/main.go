package main

import (
	"fmt"

	"github.com/toversus/gopl/ch02/popcount"
)

var pc [256]byte

func main() {
	fmt.Printf("%b\n", 1000)
	fmt.Println(popcount.PopCount(1000))
}
