package main

import (
	"bytes"
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var right, left bytes.Buffer
	for l := t; l != nil; l = l.left {
		right.WriteString(fmt.Sprintf(" %d", l.value))
	}
	for r := t.right; r != nil; r = r.right {
		left.WriteString(fmt.Sprintf("%d ", r.value))
	}
	return "{" + reverse(right.String()) + strings.TrimRight(left.String(), " ") + "}"
}

// Sort sorts values in place.
func Sort(values []int) string {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root.String()
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func reverse(s string) string {
	t := strings.Split(s, "")
	l := len(t)
	for i := 0; i < l/2; i++ {
		t[i], t[l-i-1] = t[l-i-1], t[i]
	}
	return strings.Join(t, "")
}

func main() {
	n := []int{3, 4, 5, 2, 1, 7, 7, 9, 10}
	Sort(n)
}
