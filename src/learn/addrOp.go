package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("第一行 -a 变量类型为 %T\n", a)
	fmt.Printf("第二行 -b 变量类型为 %T\n", b)
	fmt.Printf("第三行 -c 变量类型为 %T\n", c)

	ptr = &a
	fmt.Printf("a的值为 %d \n", a)
	fmt.Printf("ptr为 %d\n", *ptr)
	fmt.Printf("ptr %d\n", ptr)
}
