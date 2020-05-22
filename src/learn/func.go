package main

func tets() (a int, s string, e error) {
	return 0, "", nil
}

func test(f func()) {
	f()
}

func makeFun() func(int, int) int {
	return func(x, y int) int {
		return x - y
	}

}

func main() {
	tets()

	// 直接调用匿名函数
	func(s string) {
		println("the string is ", s)
	}("hello world")

	// 将匿名函数复制给变量
	add := func(x, y int) int {
		return x + y
	}
	println("the add ret is ", add(1, 2))

	// 作为参数
	test(func() { println("this is one func") })

	// 作为返回值
	sub := makeFun()
	println("the sub fun retrun is ", sub(10, 1))

}
