package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slogv2/src/main/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "token验证失败",
				"data": err.Error(),
			})
			c.Abort()
		}
		c.Next()
	}
}
