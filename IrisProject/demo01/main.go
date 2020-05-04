package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	app.Get("/getrequest", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})

	app.Get("/userpath", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.WriteString("请求路径" + path)
	})

	app.Get("/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		userName := context.URLParam("username")
		app.Logger().Info(userName)

		pwd := context.URLParam("pwd")
		app.Logger().Info(pwd)
		context.HTML("<h1>" + userName + "," + pwd + "</h1>")
	})

	app.Post("/postlogin", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)

		name := context.PostValue("name")
		pwd := context.PostValue("pwd")
		app.Logger().Info(name, " ", pwd)
		context.HTML(name)
	})

	app.Post("/postjson", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("请求URL：", path)

		var person Person
		if err := context.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		context.Writef("Received: %#v\n", person)
	})

	app.Post("/postxml", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("请求URL：", path)

		var student Student
		if err := context.ReadXML(&student); err != nil {
			panic(err.Error())
		}
		context.Writef("Received: %#v\n", student)
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	StuName string `xml:"stu_name"`
	StuAge  int    `xml:"stu_age"`
}
