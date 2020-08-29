package main

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	a := 5
	t := &a
	defer func() {
		*t = *t + 5
	}()
	return *t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	println(f1())
	println(f2())
	println(f3())
}
