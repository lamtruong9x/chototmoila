package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HEADER_KEY_AUTHORIZATION = "Authorization"
	NOT_CONTAIN_ACCESS_TOKEN = "request does not contain an access token"
	userIDCtx                = "userIDCtx"
	isAdminCtx               = "isAdminCtx"
)

func (c controller) Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {

		tokenString := context.GetHeader(HEADER_KEY_AUTHORIZATION)

		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": NOT_CONTAIN_ACCESS_TOKEN})
			context.Abort()
			return
		}

		payload, err := c.Maker.VerifyToken(tokenString)
		fmt.Printf("Payload: %+v\n", payload)
		fmt.Println("payload_id", payload.UserID)
		context.Set(userIDCtx, payload.UserID)
		context.Set(isAdminCtx, payload.IsAdmin)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error2 ": err.Error()})
			context.Abort()
			return
		}

		context.Next()

	}
}