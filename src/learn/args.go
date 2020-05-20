package main

import "fmt"

func test(a int, b int, c ...int) {
	fmt.Printf("IMY***********  %v %v %v\n", a, b, c)
}

func main() {
	// test(1,2,3, 4, 5, 6, 7)      // it is ok
	//c := []int{2, 3, 4}
	// test(3,4, c...)
	test(3, 4)
	//test(3,4,[]int{3,4,5})
}
