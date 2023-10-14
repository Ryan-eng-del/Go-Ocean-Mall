package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_jwt "ocean_mall/account/jwt"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" || len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "invalid token",
			})
			c.Abort()
			return
		}

		j := _jwt.NewJWT()

		parseToken, err := j.ParseJWT(token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("claims", parseToken)
		c.Next()
	}
}
