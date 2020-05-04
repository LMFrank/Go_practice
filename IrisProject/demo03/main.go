package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	userParty := app.Party("/users", func(context context.Context) {
		context.Next()
	})

	userParty.Get("/register", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("用户注册功能", path)
		context.HTML("<h1>用户注册功能</h1>")
	})

	usersRouter := app.Party("/admin", userMiddleware)

	usersRouter.Done(func(context context.Context) {
		context.Application().Logger().Infof("response send to " + context.Path())
	})

	usersRouter.Get("/info", func(context context.Context) {
		context.HTML("<h1>用户信息</h1>")
		context.Next()
	})

	usersRouter.Get("/query", func(context context.Context) {
		context.HTML("<h1>查询信息</h1>")
		context.Next()
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

func userMiddleware(context iris.Context) {
	context.Next()
}
