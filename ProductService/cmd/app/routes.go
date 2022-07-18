package main

import "github.com/gin-gonic/gin"

func (app *application) NewRouter() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	userRouter.Use(app.Controller.Authorize())
	{
		userRouter.POST("/products", app.Controller.Create)
		userRouter.PATCH("/products/:id", app.Controller.Update)
		userRouter.DELETE("/products/:id", app.Controller.Delete)
	}

	//r.GET("user/products", app.Controller.GetByUserID)
	//r.DELETE("user/products/:id", app.Controller.Delete)
	r.GET("/products/:id", app.Controller.GetByID)
	r.GET("/products/search", app.Controller.Search)
	r.GET("/products/search-by", app.Controller.SearchBy)
	return r
}
