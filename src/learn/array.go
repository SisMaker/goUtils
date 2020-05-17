package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Println("%d %d \n", i, v)

	}
	//var tet , bb int = 1, 2
	var (tet int; bb string)
	print(tet, bb)
	months := [...]string{1:"January", 12:"December"}
	Q2 := months[4:7]
	fmt.Println("%T %T", months, Q2)

}
