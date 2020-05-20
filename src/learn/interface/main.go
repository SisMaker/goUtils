package main

import (
	"goUtils/learn/interface/multiInter"
)

type T struct{}

func (T) M1() { println("IMY******* M1") }
func (T) M2() { println("IMY******* M2") }

func f1(i multiInter.I1) { i.M1() }
func f2(i multiInter.I2) { i.M2() }
func main() {
	t := T{}
	f1(t)
	f2(t)
}
