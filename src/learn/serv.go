package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/valyala/tcplisten"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	// 输出html
	// 请求方式：GET
	// 访问地址：http://localhost:8080/welcome
	app.Handle("GET", "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")

	})
	// 输出字符串
	// 类似于app.Handle("GET", "/ping", [...])
	// 请求方式GET
	// 请求地址：http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("Pong")
	})

	// 输出JSON
	// 请求方式：GET
	// 请求地址： http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "hello Iris!"})
	})

	listenerCfg := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
	}
	l, err := listenerCfg.NewListener("tcp4", ":8080")
	if err != nil {
		app.Logger().Fatal(err)
	}
	app.Run(iris.Listener(l)) //8080 监听端口

}
