package main

import "fmt"
const boilingF = 212.0

func main(){
	var F = boilingF
	var c = (F - 32) * 5 /9
	fmt.Printf("boiling point = %g F or %g C\n", F, c)

}