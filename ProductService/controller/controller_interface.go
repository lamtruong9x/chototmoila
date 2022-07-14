package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Create(c *gin.Context)
	GetByUserID(c *gin.Context)
	//GetList(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
