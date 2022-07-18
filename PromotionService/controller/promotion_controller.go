package controller

import (
	entity "PromotionService/entities"
	"PromotionService/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PromotionController interface {
	Create(c *gin.Context)
	// Get(c *gin.Context)
	// Update(c *gin.Context)
	// Delete(c *gin.Context)
}

type promotionController struct {
	service services.PromotionService
}

func NewPromotionController(service services.PromotionService) PromotionController {
	return &promotionController{service: service}
}

func (ctrl *promotionController) Create(c *gin.Context) {
	var input entity.PromotionEntity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	input.CreatedTime = time.Now()
	input.LastUpdate = time.Now()
	input.IsApprove = false

	if err := ctrl.service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})

}

// func (ctrl *promotionController) Get(c *gin.Context) ([]*entity.PromotionEntity, error) {

// }

// func (ctrl *promotionController) Update(c *gin.Context) error {

// }

// func (ctrl *promotionController) Delete(c *gin.Context) error {

// }
