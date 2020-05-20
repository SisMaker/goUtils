package main

import "fmt"

func worker(ch chan struct{}) {

	v, ok := <-ch
	if ok {
		fmt.Printf("IMY*********** get one null msg %T\n", v)
	}
	println("IMY*********** get one null msg \n")
	close(ch)
}

func main() {
	ch := make(chan struct{})
	go worker(ch)
	// send to a null msg to worker
	ch <- struct{}{}
	<-ch
	println("IMY***********  over\n")
}
