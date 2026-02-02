package middleware

import (
	"net/http"
	"zholianxi/bff/pkg"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "未登录",
			})
		}
		getToken, err := pkg.GenToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "token未解析",
			})
		}
		c.Set("userId", getToken["userId"].(string))
		c.Next()
	}
}
