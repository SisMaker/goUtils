package main

import "fmt"

type test struct {
	aa int
	bb int
	cc string
}

func main() {
	var t interface{} = test{
		1,
		2,
		"fdfss"}

	switch p := t.(type) {
	case test:
		fmt.Printf("tetttttt %T %v \n", t, p)
	default:
		fmt.Printf("%v", p)
	}

	//fmt.Printf("***** %v \n", t)
}
