package main

import (
	"fmt"
	"goUtils/learn/interface/multiInter"
)

type T struct {
	aa int
	bb string
}

type T2 struct {
	aa2 int
	bb2 string
}

func (t T) M1() { println("IMY******* M1", t.bb) }
func (T) M2()   { println("IMY******* M2") }

func (T2) M2() { println("IMY******* M2") }

func f1(i multiInter.I1) { i.M1() }
func f2(i multiInter.I2) { i.M2() }
func main() {
	t := T{}
	f1(t)
	f2(t)
	var i interface{} = T{11, "fdfdfdf"}
	fmt.Printf("%v %[1]T\n", i.(T))
	ii := i.(T)
	f1(ii)
	var t2 interface{} = T2{11, "aaaaaaaa"}
	fmt.Printf("%v %[1]T\n", t2)
}
