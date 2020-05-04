package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	app.Handle("GET", "/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		app.Logger().Error("request path:", path)
	})

	app.Handle("POST", "/postcommit", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("post request, the url is:", path)
		context.HTML(path)
	})

	app.Get("/hello/{name}", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		name := context.Params().Get("name")
		context.HTML("<h1>" + name + "</h1>")
	})

	app.Get("/api/users/{userid:uint64}", func(context context.Context) {
		userID, err := context.Params().GetUint("userid")

		if err != nil {
			context.JSON(map[string]interface{}{
				"requestcode": 201,
				"message":     "bad request",
			})
			return
		}

		context.JSON(map[string]interface{}{
			"requestcode": 200,
			"user_id":     userID,
		})
	})

	app.Get("/api/users/{isLogin:bool}", func(context context.Context) {
		isLogin, err := context.Params().GetBool("isLogin")
		if err != nil {
			context.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			context.WriteString("已登录")
		} else {
			context.WriteString("未登录")
		}

		context.Params()
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
