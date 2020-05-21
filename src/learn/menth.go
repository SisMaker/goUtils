package main

import "fmt"

type N int

func (n N) value() {
	n++
	fmt.Println("v: %p %v \n ", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Println("v: %p %v \n ", n, *n)

}

func main() {
	var a N = 25
	p := &a
	p2 := &p
	(*p2).value()
	(*p2).pointer()

}
