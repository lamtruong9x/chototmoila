package controller

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
	"chotot_product_ltruong/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type controller struct {
	Service service.Service
}

func New(svc service.Service) *controller {
	return &controller{Service: svc}
}

// Hard coded for now
const (
	userID       = 1
	limitPerPage = 10
)

func (ctrl *controller) Create(c *gin.Context) {
	var input dto.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Controller - Create: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Hard coded
	input.UserId = userID
	input.CreatedTime = time.Now()
	input.ExpiredTime = time.Now().Add(time.Hour * 24)

	if err := ctrl.Service.Create(&input); err != nil {
		log.Printf("Controller - Create: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// Return 10 latest products each page
func (ctrl *controller) GetByUserID(c *gin.Context) {
	pageNum := 0
	if s := c.Query("page"); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
			return
		}
		pageNum = n
	}
	products := make([]*entity.Product, 0, 10)
	products, err := ctrl.Service.GetByUserID(userID, limitPerPage, pageNum)
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "you don't have any products to browse"})
		return
	}
	if err != nil {
		log.Printf("Controller - GetByUserID: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "you don't have any products to browse"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (ctrl *controller) GetByName(c *gin.Context) {

}

func (ctrl *controller) Update(c *gin.Context) {
	var input dto.ProductUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Controller - Update: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	product, err := ctrl.Service.Update(&input)
	if err != nil {
		log.Printf("Controller - Update: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ctrl *controller) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := ctrl.Service.Delete(id); err != nil {
		log.Printf("Controller - Delete: %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
