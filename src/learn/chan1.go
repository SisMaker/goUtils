package main

import (
	"fmt"
	"time"
)

func consumer(data chan int, ret chan bool) {
	//消费者
	tt := 1
	for {
		select {
		case <-data:
			fmt.Printf("aaaa %v \n", tt)
			tt += 2
			ret <- true
		}

	}
}

func product(data chan int, ret chan bool) {
	tt := 2
	for i := 0; i < 10; i++ {
		data <- i
		select {
		case <-ret:
			fmt.Printf("bbbb %v \n", tt)
			tt += 2
		}

	}
}

var X = 100

func main() {
	ret := make(chan bool, 1)
	data := make(chan int, 1)
	go consumer(data, ret)
	go product(data, ret)

	time.Sleep(50000000)
	print("get done ")

}
