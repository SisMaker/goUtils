package main

func consumer(data chan int, done chan bool) {
	//消费者
	for x := range data {
		println("recevie ", x)
	}
	done <- true
}

func product(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)

}

var X = 100

func main() {
	done := make(chan bool)
	data := make(chan int)
	go consumer(data, done)
	go product(data)
	v := <-done
	print("get done ", v)
	Y := 300
	print(X, Y)
	X, Y := 10, 20
	print(X, Y)

}
