package api

import (
	"github.com/kataras/iris"
)

func Routes (app *iris.Application) {
	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
			"message": "Works!",
			"status": 200,
        })
	})
}