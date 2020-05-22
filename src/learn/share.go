package main

import (
	"os"
	"strconv"
	"time"
)

var share int = 0
var Cnt int

func GSet(index int, ch chan struct{}) {
	for i := 0; i < Cnt; i++ {
		//println("cur index ", index, "cur vale", share)
		share += 1
		ch <- struct{}{}
	}

}

func main() {
	Num, _ := strconv.Atoi(os.Args[1])
	Cnt, _ = strconv.Atoi(os.Args[2])
	println(Num, Cnt)
	ch := make(chan struct{})
	start := time.Now()
	for i := 0; i < Num; i++ {
		go GSet(i, ch)
	}
	GetNum := 0
	for {
		<-ch
		GetNum++
		//println("add  one num ", GetNum)
		if GetNum >= Cnt*Num {
			break
		}
	}
	use := time.Since(start)
	println("test over ", share, "should be ", Cnt*Num, "use time ", use)
}
