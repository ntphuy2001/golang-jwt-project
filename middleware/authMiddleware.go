package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helper "github.com/ntphuy2001/golang-jwt-project/helpers"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("Email", claims.Email)
		c.Set("First_name", claims.FirstName)
		c.Set("Last_name", claims.LastName)
		c.Set("Uid", claims.Uid)
		c.Set("User_type", claims.UserType)
		c.Next()
	}
}
