package main

func fa(base int) (func(int) int, func(int) int) {
	println(&base, base)
	add := func(i int) int {
		println("adddddddd")
		base -= i
		println(&base, base)
		return base
	}

	sub := func(i int) int {
		println("subbbbbbbb")
		base -= i
		println(&base, base)
		return base
	}
	return add, sub

}

func main() {
	// f, g闭包引用的base是同一个， 是fa函数调用传递过来的实参值
	f, g := fa(0)

	println(f(1), g(2))

	s, k := fa(0)
	println(s(1), k(2))
}
