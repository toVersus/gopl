package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Printf("%08b %08b %08b\n", u, u+1, u*u)
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Printf("%08b %08b %08b\n", i, i+1, i*i)
	fmt.Println(i, i+1, i*i)
}
