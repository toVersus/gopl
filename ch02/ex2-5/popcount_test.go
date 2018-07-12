package popcount

import (
	"reflect"
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

func PopCountShifting(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCountClearRightmostBit(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func testPopCount(t *testing.T, f func(uint64) int) {
	tests := []struct {
		input uint64
		want  int
	}{
		{1000, 6},
	}

	for _, test := range tests {
		if result := f(test.input); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got = %d, want = %d", result, test.want)
		}
	}
}

func TestPopCountClearRightmostBit(t *testing.T) {
	testPopCount(t, PopCountClearRightmostBit)
}

var result int

func benchmark(b *testing.B, f func(uint64) int) {
	var n int
	for i := 0; i < b.N; i++ {
		n = f(uint64(i))
	}
	result = n
}

func BenchmarkPopCount(b *testing.B) {
	benchmark(b, popcount.PopCount)
}

func BenchmarkPopCountTableLookup(b *testing.B) {
	benchmark(b, PopCountTableLookup)
}

func BenchmarkPopCountShifting(b *testing.B) {
	benchmark(b, PopCountShifting)
}

func BenchmarkPopCountClearRightmostBit(b *testing.B) {
	benchmark(b, PopCountClearRightmostBit)
}
