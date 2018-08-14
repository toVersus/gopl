package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{[]int{3, 4, 5, 2, 1, 7, 7, 9, 10}, "{1 2 3 4 5 7 7 9 10}"},
	}
	for _, test := range tests {
		result := Sort(test.input)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%+v, want=%+v\n", result, test.want)
		}
	}
}
