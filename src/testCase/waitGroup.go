package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting + " " + strconv.Itoa(i))
		dur := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(dur) // 睡眠片刻
	}
	wg.Done()
	log.Println("green over !!!")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	cnt, _ := strconv.Atoi(os.Args[1])
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go SayGreetings("Hi"+strconv.Itoa(i), 10)
	}

	wg.Wait()
	wg.Wait()
	log.Println("the main over!!!")
}
