package main

func main() {
	f := func() interface{} {
		println("inner 00000000")
		return func() {
			println("inner 11111111111")

		}
		//println("inner 222222222")
	}()
	d := f.(func())
	d()
	test(test("111111111", "2222222222"), test("333333333", "444444444444"))
}

func test(a string, b string) string {

	println("testttt", a, b)
	return a
}
