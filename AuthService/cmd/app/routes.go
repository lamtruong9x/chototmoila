package main

import "github.com/gin-gonic/gin"

func (app *application) NewRouter() *gin.Engine {
	router := gin.Default()
	router.POST("auth/register", app.controller.CreateNewUser)
	router.POST("auth/login", app.controller.Login)
	return router
}
