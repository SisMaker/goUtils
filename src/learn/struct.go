package main

import (
	"time"
)

type Emplyee struct {
	Id        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerId int
}

var dilbert Emplyee

func test() {
	dilbert.Salary -= 5000
	posPtr := &dilbert.Position
	*posPtr = "fdfdsfdfd" + *posPtr

	// 点号同样可以用在结构体指针上
	var emPtr *Emplyee = &dilbert
	emPtr.Position += "fddddddd"

	// 等价于下面这句
	(*emPtr).Position += "fddddddd"

	//没有任何成员变量的结构体成为空结构体
	set := make(map[string]struct{})
	if _, ok := set[str]; !ok {
		set[str] = struct{}{}

	}

}
