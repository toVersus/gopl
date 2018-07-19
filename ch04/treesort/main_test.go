package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{4, 1, 3, 2, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, test := range tests {
		result := test.input
		if Sort(result); !reflect.DeepEqual(result, test.want) {
			t.Errorf("got=%+v, want=%+v\n", result, test.want)
		}
	}
}
