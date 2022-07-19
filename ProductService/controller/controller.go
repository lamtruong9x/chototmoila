package controller

import (
	"chotot_product_ltruong/dto"
	"chotot_product_ltruong/entity"
	"chotot_product_ltruong/service"
	"chotot_product_ltruong/token"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	Maker   token.Maker
	Service service.Service
}

func New(svc service.Service, maker token.Maker) *controller {
	return &controller{Service: svc, Maker: maker}
}

// Hard coded for now
const (
	limitPerPage = 10

	PRODUCT_NAME_FIELD = "product_name"
	ADDRESS_FIELD      = "address"
	KEYWORD            = "k"
	ADDRESS            = "a"
	PARAM              = "param"
	VALUE              = "value"
	EXPIRED_TIME       = 24 * 30
)

func (ctrl *controller) Create(c *gin.Context) {
	var input dto.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Controller - Create: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt(UserIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	input.UserId = userID
	input.CreatedTime = time.Now()
	input.ExpiredTime = time.Now().Add(time.Hour * EXPIRED_TIME)
	input.Priority = false

	if err := ctrl.Service.Create(&input); err != nil {
		log.Printf("Controller - Create: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "product created"})
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

	userID := c.GetInt(UserIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

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

func (ctr *controller) GetByID(c *gin.Context) {
	idString := c.Param("id")

	productID, err := strconv.Atoi(idString)
	if err != nil || productID < 1 {
		c.JSON(http.StatusBadRequest, err)
	}

	product, err := ctr.Service.GetByID(productID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, product)
}
func (ctrl *controller) Purchase(c *gin.Context) {
	if !c.GetBool(IsAdminCtx) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allow for this function"})
		return
	}
	var input dto.ProductUpdate
	input.Priority = true
	ctrl.superUpdate(&input, c)
}

func (ctrl *controller) Update(c *gin.Context) {
	var input dto.ProductUpdate
	if input.Priority != false {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allow for this function"})
		return
	}
	ctrl.superUpdate(&input, c)

}

func (ctrl *controller) superUpdate(input *dto.ProductUpdate, c *gin.Context) {
	if err := c.ShouldBindJSON(input); err != nil {
		log.Printf("Controller - Update: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// hand matching
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	input.UserId = c.GetInt(UserIDCtx)
	input.Id = id

	product, err := ctrl.Service.Update(input)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	// Get userID from ctx
	userID := c.GetInt(UserIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := ctrl.Service.Delete(id, userID); err != nil {
		log.Printf("Controller - Delete: %v\n", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (ctrl *controller) Search(c *gin.Context) {
	keyword := "'%" + c.Query(KEYWORD) + "%'"
	address := "'%" + c.Query(ADDRESS) + "%'"
	if keyword == "'%%'" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	query := fmt.Sprintf("%s like %s", PRODUCT_NAME_FIELD, keyword)
	if address != "'%%'" {
		query += fmt.Sprintf(" and %s like %s", ADDRESS_FIELD, address)
	}

	products, _ := ctrl.Service.Search(query)
	fmt.Println(query)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (ctrl *controller) SearchBy(c *gin.Context) {
	param := c.Query(PARAM)
	value := c.Query(VALUE)

	if param == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	query := fmt.Sprintf("%s = %s", param, value)

	products, _ := ctrl.Service.Search(query)
	//fmt.Println(query)

	c.JSON(http.StatusOK, gin.H{"data": products})
}
