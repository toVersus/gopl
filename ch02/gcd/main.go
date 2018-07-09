package main

func main() {
	println(gcd(12, 4))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
