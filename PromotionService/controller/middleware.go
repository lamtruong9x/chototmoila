package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HEADER_KEY_AUTHORIZATION = "Authorization"
	NOT_CONTAIN_ACCESS_TOKEN = "request does not contain an access token"
	FORBIDDEN                = "you are not allow for this function"
	UserIDCtx                = "UserIDCtx"
	IsAdminCtx               = "isAdminCtx"
)

func (c *promotionController) Authorize() gin.HandlerFunc {
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

		if !payload.IsAdmin {
			context.JSON(http.StatusForbidden, gin.H{"error": FORBIDDEN})
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
