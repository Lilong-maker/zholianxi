package routers

import (
	"net/http"
	"time"
	"zholianxi/bff/handler/service"
	"zholianxi/bff/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	r.POST("GoodsAdd", middleware.Loggers(), middleware.Reg(), service.GoodsAdd)
	r.POST("Login", service.Login)
	r.POST("Uploads", service.Uploads)
	r.POST("GetGoods", service.GetGoods)
	r.POST("Ref", service.Ref)
	return r
}
