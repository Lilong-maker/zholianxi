package middleware

import (
	"context"
	"net/http"
	"zholianxi/bff/basic/config"
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
		ok, err := config.Rdb.SIsMember(context.Background(), "black:k1", getToken["userId"]).Result()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "登录过期",
			})
			c.Abort()
			return
		}
		if ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "已在黑名单",
			})
			c.Abort()
			return
		}
		c.Set("userId", getToken["userId"])
		c.Next()
	}
}
