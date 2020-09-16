package main

import "time"

func main() {
	println("hello world")

	go func() {
		time.Sleep(time.Second)
		//panic(123)
	}()

	defer func() {
		recover()
	}()
	println("IMY************")
	for {
		_ = new(int64)
		//time.Sleep(time.Second)
	}
}
