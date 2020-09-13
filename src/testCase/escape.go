package main

type Val struct {
	Name string
	Age  int
}

func NewValue(name string, age int) *Val {
	c := new(Val)
	c.Name = name
	c.Age = age
	return c
}

func main() {
	NewValue("Test", 10)
}
