package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting + " " + strconv.Itoa(i))
		dur := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(dur) // 睡眠片刻
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	cnt, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < cnt; i++ {
		go SayGreetings("Hi"+strconv.Itoa(i), 1000)
	}
	log.Println("the main over!!!")
	time.Sleep(time.Second * 1000000)
}
