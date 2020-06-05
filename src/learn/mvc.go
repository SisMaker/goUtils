package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

func main() {
	app := iris.New()
	mvc.Configure(app.Party("/root"), myMVC)
	app.Run(iris.Addr(":8082"))

}

func myMVC(app *mvc.Application) {
	app.Handle(new(MyController))

}

type MyController struct{}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler")

}

// GET: http://localhost:8080/root
func (m *MyController) Get() string { return "Hey" }

// GET: http://localhost:8080/root/something/{id:long}
func (m *MyController) MyCustomHandler(id int64) string {
	return "MyCustomHandler says Hey " + strconv.FormatInt(id, 10)
}
