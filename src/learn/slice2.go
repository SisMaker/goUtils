package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	Num, _ := strconv.Atoi(os.Args[1])

	news := append(s[:Num-1], s[Num:]...)
	fmt.Printf("%v \n", news)
}
