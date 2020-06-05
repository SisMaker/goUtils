package main

import (
	"os"
	"strconv"
	"time"
)

func test() {}

func main() {
	start := time.Now()
	cnt, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < cnt; i++ {
		test()
	}
	println("IMY*****main*", time.Since(start).Microseconds(), "us")

}
