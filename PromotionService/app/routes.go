package main

import "github.com/gin-gonic/gin"

func (app *application) NewRouter() *gin.Engine {
	r := gin.Default()
	promotionRouter := r.Group("/purchase")
	promotionRouter.Use(app.Controller.Authorize())
	{
		promotionRouter.GET("/", app.Controller.Get)
		promotionRouter.POST("/", app.Controller.Create)
		promotionRouter.PATCH("/", app.Controller.Update)
		promotionRouter.DELETE("/:promotionid", app.Controller.Delete)
	}

	return r
}
