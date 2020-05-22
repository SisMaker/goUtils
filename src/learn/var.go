package main

import "fmt"

func main() {

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var n int
	fmt.Println(n)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var a1 *int
	var a2 []int
	var a3 map[string]int
	var a4 chan int
	var a5 func(string) int
	var a6 error // error 是接口
	fmt.Printf("%v %v %v %v %v %v\n", a1, a2, a3, a4, a5, a6)
}
