package controller

import (
	repository "AuthService/internal/db/sqlc"
	"AuthService/internal/token"
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Queries repository.Querier
	Token   token.Maker
	Log     *log.Logger
}

type IController interface {
	Login(c *gin.Context)
	CreateNewUser(c *gin.Context)
}

func New(q repository.Querier, t token.Maker, l *log.Logger) *Controller {
	return &Controller{
		Queries: q,
		Log:     l,
		Token:   t,
	}
}

type input struct {
	Phone    string `json:"phone" binding:"required,numeric,gte=10"`
	Password string `json:"password" binding:"required"`
}

// Done
func (srv *Controller) Login(c *gin.Context) {
	var input input
	if err := c.ShouldBindJSON(&input); err != nil {
		srv.Log.Println("Package controller-CreateNewUser-ShouldBindJSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	user, err := srv.Queries.GetUserByPhone(ctx, input.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not exits"})
			return
		}
		srv.Log.Printf("Package controller-Login-GetUserByPhone: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if !comparePassword(user.Passwd, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}
	tk, _, err := srv.Token.CreateToken(int(user.ID), user.Phone, user.Isadmin, time.Minute*45)
	if err != nil {
		srv.Log.Println(err)
	}
	c.Header("Authorization", tk)
	c.JSON(http.StatusAccepted, gin.H{"Info": "Login successfully"})
}

// Done
func (srv *Controller) CreateNewUser(c *gin.Context) {
	var input input
	if err := c.ShouldBindJSON(&input); err != nil {
		srv.Log.Println("Package controller-CreateNewUser-ShouldBindJSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var CreateUser repository.CreateUserParams
	CreateUser.Username = "New User"
	CreateUser.Phone = input.Phone
	password, err := hashPassword(input.Password)
	if err != nil {
		srv.Log.Printf("Package controller-CreateNewUser-hashPassword: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't create account due to unexpected error"})
		return
	}
	CreateUser.Passwd = password

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	result, err := srv.Queries.CreateUser(ctx, CreateUser)
	if err != nil {
		srv.Log.Printf("Package controller-CreateNewUser: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Get the id from new created user to send to generate new token
	_, err = result.LastInsertId()
	if err != nil {
		srv.Log.Printf("Package controller-CreateNewUser: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "user created")
}
