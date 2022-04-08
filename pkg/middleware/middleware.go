package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/NwokoyeChigozie/quik_task/internal/model"
	"github.com/NwokoyeChigozie/quik_task/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var MyIdentity *model.UserIdentity

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearerToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenStr = strArr[1]
		}

		if tokenStr == "" {
			r := utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token could not be found!", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, r)
			return
		}

		token, err := utility.TokenValid(tokenStr)
		if err != nil {
			r := utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token is invalid!", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, r)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID, ok := claims["user_id"].(string) //convert the interface to string
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token is invalid!", "Unauthorized", nil))
			return
		}

		authoriseStatus, ok := claims["authorised"].(bool) //check if token is authorised for middleware
		if !ok && !authoriseStatus {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utility.BuildErrorResponse(http.StatusUnauthorized, "error", "Token is invalid!", "Unauthorized", nil))
			return
		}

		userId, _ := strconv.Atoi(userID)
		MyIdentity = &model.UserIdentity{
			ID: userId,
		}

	}
}

func ValidateRequestUser(requestedDataUserID int) bool {
	currenctUserID := MyIdentity.ID
	return currenctUserID == requestedDataUserID
}
