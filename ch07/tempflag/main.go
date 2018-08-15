package main

import (
	"flag"
	"fmt"

	"github.com/toversus/gopl/ch07/ex7-6"
)

var celsius = tempconv.CelsiusFlag("celsius", 20.0, "the celsius temperature")

func main() {
	flag.Parse()
	fmt.Println(*celsius)
}
