package main

import (
	"os"
	"strconv"
)

func flib(index, max, l1, l2 int) int {
	if index == max {
		return l2
	}
	return flib(index+1, max, l2, l1+l2)
}

func f(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return f(n-1) + f(n-2)
	}
}

func main() {
	max, _ := strconv.Atoi(os.Args[1])

	if max == 1 {
		println("index1: \n", max)
		println("value1: \n", 1)
		return
	}
	if max == 2 {
		println("index1: \n", max)
		println("value1: \n", 1)
		return
	}

	sum := flib(3, max, 1, 2)
	println("index1: ", max)
	println("value1: ", sum)

	sum = f(max)
	println("index2: ", max)
	println("value2: ", sum)

}
