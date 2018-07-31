package main

import "fmt"

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}

func main() {
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(min(1, 2, 3, 4))

	fmt.Println(max())
	fmt.Println(min())
}
