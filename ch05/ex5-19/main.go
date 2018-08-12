package main

import "fmt"

func main() {
	defer sayWorld()
	sayHello()
}

func sayWorld() {
	if p := recover(); p != nil {
		fmt.Println(p, "world!")
	}
}

func sayHello() {
	panic("Hello,")
}
