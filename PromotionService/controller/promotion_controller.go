package controller

import (
	entity "PromotionService/entities"
	"PromotionService/services"
	"PromotionService/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	productID   = "productid"
	promotionID = "promotionid"
)

type PromotionController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	//mid
	Authorize() gin.HandlerFunc
}

type promotionController struct {
	Maker   token.Maker
	service services.PromotionService
}

func NewPromotionController(service services.PromotionService, maker token.Maker) PromotionController {
	return &promotionController{service: service, Maker: maker}
}

func (ctrl *promotionController) Create(c *gin.Context) {
	var input entity.Promotion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	input.CreatedTime = time.Now()
	input.ExpiredTime = time.Now().AddDate(0, 0, 7) //add 7 days
	input.LastUpdate = time.Now()
	input.IsApprove = false

	if err := ctrl.service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})

}

func (ctrl *promotionController) Get(c *gin.Context) {
	// var results []*entity.Promotion
	results, err := ctrl.service.Get("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"results": results})

}

func (ctrl *promotionController) Update(c *gin.Context) {
	var input entity.Promotion
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	input.ExpiredTime = time.Now().AddDate(0, 0, 7) //add 7 days
	input.LastUpdate = time.Now()

	if err := ctrl.service.Update(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (ctrl *promotionController) Delete(c *gin.Context) {
	promotionID := c.Param(promotionID)
	//check isAdmin
	//todo thing
	//
	if err := ctrl.service.Delete(promotionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
