package main

func fa(a int) func(i int) int {
	return func(i int) int {
		println(&a, a)
		a = a + i
		return a
	}
}
func main() {
	f := fa(1) // f引用的外部闭包环境包括本次函数调用的形参a的值1
	g := fa(1) // g引用的外部闭包环境包括本次函数调用的形参a的值1
	// 此时 f, g 引用的闭包环境中的a的值 并不是同一个 而是两次函数调用产生的副本

	println(f(10))
	println(f(10))

	println(g(10))
	println(g(10))

}
