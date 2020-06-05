package main

import (
	"fmt"
	eventWay2 "goUtils/learn/eventWay/eventWay"
)

func hand1(msg interface{}) error {
	println("IMY************ 111", msg)
	return nil
}

func hand2(msg interface{}) error {
	println("IMY************ 222", msg)
	return nil
}

func hand3(msg interface{}) error {
	fmt.Printf("IMY************ 333 %v\n", msg)
	return nil
}

func main() {
	eventWay2.Register(eventWay2.Event{Name: "str1"}, hand1)
	eventWay2.Register(eventWay2.Event{Name: "str1"}, hand2)
	eventWay2.Register(eventWay2.Event{Name: "str3"}, hand3)
	eventWay2.Publish(eventWay2.Event{Name: "str1"}, "teste1")
	eventWay2.Publish(eventWay2.Event{Name: "str3"}, "testt3")
	eventWay2.Publish(eventWay2.Event{Name: "str2"}, "fdfdfdfdf")

}
