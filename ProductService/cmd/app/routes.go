package main

import "github.com/gin-gonic/gin"

func (app *application) NewRouter() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	userRouter.Use(app.Controller.Authorize())
	{
		userRouter.POST("/products/post", app.Controller.Create)
		userRouter.PATCH("user/products", app.Controller.Update)
	}

	//r.GET("user/products", app.Controller.GetByUserID)

	//r.DELETE("user/products/:id", app.Controller.Delete)
	r.GET("/products/:id", app.Controller.GetByID)
	r.GET("user/products/search", app.Controller.Search)
	r.GET("user/products/search-by", app.Controller.SearchBy)
	return r
}
