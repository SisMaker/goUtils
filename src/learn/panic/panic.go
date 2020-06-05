package main

func test() {
	var a int
	defer func() {
		if p := recover(); p != nil {
			a = 1111
		}
	}()

	panic(2222)
	print(a)

}

func main() {
	test()

}
