package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HEADER_KEY_AUTHORIZATION = "Authorization"
	NOT_CONTAIN_ACCESS_TOKEN = "request does not contain an access token"
	UserIDCtx                = "UserIDCtx"
	IsAdminCtx               = "isAdminCtx"
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
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": NOT_CONTAIN_ACCESS_TOKEN})
			context.Abort()
			return
		}
		//fmt.Printf("Payload: %+v\n", payload)
		context.Set(UserIDCtx, payload.UserID)
		context.Set(IsAdminCtx, payload.IsAdmin)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error2 ": err.Error()})
			context.Abort()
			return
		}

		context.Next()

	}
}
