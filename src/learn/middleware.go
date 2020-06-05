package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", before, mainHandler, after)
	app.Run(iris.Addr(":8081"))

}

func before(ctx iris.Context) {
	shareInfo := "this is a shareInfo between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInfo)
	ctx.Next() //继续下一个handler
}

func after(ctx iris.Context) {
	println("after the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")
	info := ctx.Values().GetString("info")
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)
	ctx.Next()

}
