package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	tackCh := make(chan int, 10)
	go worker(tackCh)
	timeStart := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		tackCh <- i
	}
	timeEnd := time.Now().UnixNano()
	fmt.Printf("IMY******over %v %v \n", timeEnd-timeStart, (timeEnd-timeStart)/100000000)
	os.Exit(1)
}

func worker(taskCh chan int) {
	const N = 5
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				<-taskCh
				// fmt.Printf("finis task:%d by worker %d\n",task, id)
				// time.Sleep(time.Second)
			}
		}(i)
	}
}
