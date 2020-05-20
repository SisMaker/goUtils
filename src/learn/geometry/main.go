package main

import (
	"fmt"
	"goUtils/learn/geometry/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance())

	pptr := &geometry.Point{7, 7}
	pptr.ScaleBy(2.0)
	fmt.Println(*pptr)
	(*pptr).ScaleBy(2.0)
	fmt.Println(*pptr)

	fmt.Println((*pptr).Distance(q))

	p.ScaleBy(2.0)
	fmt.Println(p)
}
