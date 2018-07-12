package popcount

import (
	"testing"

	"github.com/toversus/gopl/ch02/popcount"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountTableLookup returns the population count (number of set bits) of x.
func PopCountTableLookup(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>uint(i))])
	}
	return count
}

var result int

func BenchmarkPopCount(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = popcount.PopCount(uint64(i))
	}
	result = n
}

func BenchmarkPopCountTableLookup(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = PopCountTableLookup(uint64(i))
	}
	result = n
}
