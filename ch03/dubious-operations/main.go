package main

import (
	"fmt"
	"math"
)

func main() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Printf("nan == nan: %t\nnan < nan: %t\nnan > nan: %t\n",
		nan == nan, nan < nan, nan > nan)
}
