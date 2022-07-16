package main

import "github.com/gin-gonic/gin"

func (app *applicatiton) NewRouter() *gin.Engine {
	r := gin.Default()
	//r.POST("user/products/post", app.Controller.Create)
	//r.GET("user/products", app.Controller.GetByUserID)
	//r.PATCH("user/products", app.Controller.Update)
	//r.DELETE("user/products/:id", app.Controller.Delete)
	r.GET("/products/:id", app.Controller.GetByID)
	r.GET("user/products/search", app.Controller.Search)
	r.GET("user/products/search-by", app.Controller.SearchBy)
	return r
}
