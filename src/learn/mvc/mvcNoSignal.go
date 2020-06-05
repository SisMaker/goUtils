package main

import (
	"context"
	"github.com/kataras/iris"
	"time"
)

func main() {
	app := iris.New()
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// 关闭所有主机
		println("IMY*********** todo stop")
		app.Shutdown(ctx)
	})
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>hi I just exist in order to see if the server is closed</h1>")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)
}
