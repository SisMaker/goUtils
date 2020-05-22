package main

import (
	"fmt"
	"math"
)

var XX = 200

const YY = 300 //常量不会分配地址

func main() {
	const (
		x uint16 = 120
		y
		s
		s1 = "abc"
		z
	)
	println(x, " ", y, " ", s, " ", s1, " ", z)

	const (
		a = iota
		a1
		b float32 = iota
		c
	)
	println(a, " ", a1, " ", b, " ", c)

	// println(&XX, &YY) 			// error
	fmt.Printf("%v %v \n", math.MaxUint8, math.MaxInt64)

}
