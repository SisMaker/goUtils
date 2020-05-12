package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("第 1111  c 的值 %d \n", c)
	c = a - b
	fmt.Printf("第 2222  c 的值 %d \n", c)
	c = a * b
	fmt.Printf("第 3333  c 的值 %d \n", c)
	c = a*1.0 / b
	fmt.Printf("第 4444  c 的值 %d \n", c)
	c = a % b
	fmt.Printf("第 5555  c 的值 %d \n", c)
	a++
	fmt.Printf("第 6666  c 的值 %d \n", a)
	a--
	fmt.Printf("第 7777  c 的值 %d \n", a)
}
