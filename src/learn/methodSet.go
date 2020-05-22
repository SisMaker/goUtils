package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type T struct {
	S
}

func (S) sVal()  {}
func (*S) sPtr() {}
func (T) tVal()  {}
func (*T) tPtr() {}

func methodSet(I interface{}) {
	t := reflect.TypeOf(I)
	fmt.Println(t)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var t T
	methodSet(t)
	println("-----------------")
	methodSet(&t)
}