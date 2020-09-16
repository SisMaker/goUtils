package main

import (
	"os"
	"strconv"
	"time"
)

func test(i int) {
	println("IMY***********", i)
	defer test(i*100000000000000000/137 + 11 + 253)
}

func main() {
	cnt, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < cnt; i++ {
		go test(1)
	}
	time.Sleep(120 * time.Second)
}
