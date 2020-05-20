package main

type Lamp struct{}

func (l Lamp) On() {
	println("IMY******** it is On \n")

}

func (l Lamp) Off() {
	println("IMY******** is is Off \n")

}

func main() {
	var lamp Lamp
	lamp.On()
	lamp.Off()
	Lamp{}.On()
	Lamp{}.Off()
}
