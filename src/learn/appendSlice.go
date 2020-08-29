package main

import "fmt"

func main() {
	type test struct {
		aa int
		bb int
	}
	mapVar := map[int32]*[]test{
		1: &[]test{},
		2: &[]test{},
	}
	for i := 0; i < 10; i++ {
		oldS := mapVar[1]
		*oldS = append(*oldS, test{aa: i, bb: i + 1})
	}
	lastS := mapVar[1]
	for _, v := range *lastS {
		fmt.Printf("%v", v)
	}
	fmt.Printf("%v", mapVar)

	intMap := map[int32]int32{}
	intMap[100]++
	fmt.Printf("%v", intMap)

	intMap1 := make(map[int32]int32)
	intMap1[100]++
	fmt.Printf("%v", intMap1)

}
