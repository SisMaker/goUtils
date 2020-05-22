package main

type Str interface {
	toStr()
	add()
}

type user struct {
	a int
	b string
}

type manger struct {
	*user
	aa int
	bb string
}

func (u user) toStr() {
	println("IMY********user", u.a, " ", u.b)
}

func (u user) add() {
	println("IMY********user222", u.a, " ", u.b)
}

func (m manger) toStr() {
	println("IMY********manger")
}

func overa(s Str) {
	println("TEST********", s)
	s.toStr()
}

func overb(s Str) {
	println("TEST********", s)
	s.add()
}

func main() {
	var m manger = manger{&user{22, "22uuuuuuuu"}, 1, "mmmmm"}
	var u user = user{2, "uuuuuuuuu"}

	u.toStr()
	m.toStr()

	overa(m)
	overa(u)

	overb(u)
	overb(m)
}
