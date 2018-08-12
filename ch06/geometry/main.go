package main

import (
	"fmt"
	"math"
	"reflect"
)

type Point struct{ X, Y float64 }

// Distance is implemented in legacy way
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance is same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(Point{1, 2}.Distance(q))

	(&p).ScaleBy(2.0)
	q.ScaleBy(2.0) // implicit (&q)
	fmt.Println(p.Distance(q))

	pptr := &Point{1, 2}
	fmt.Println(reflect.TypeOf(pptr))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	list := &IntList{}
	fmt.Println(list.Sum())
}
