package middleware

import (
	"encoding/json"
	"log"
	"os"
	"time"
	"zholianxi/bff/pkg"

	"gitee.com/zuyanlongnb666/crawl/reg"
	"github.com/gin-gonic/gin"
)

func Reg() gin.HandlerFunc {
	middleware := reg.NegotiationCacheMiddleware("v0.0.1", "Etag", "no-cache")
	return middleware
}
func Loggers() gin.HandlerFunc {
	return func(c *gin.Context) {
		trackData := map[string]interface{}{
			"path":      c.Request.URL.Path,
			"method":    c.Request.Method,
			"timestamp": time.Now().UnixMilli(),
			"status":    "success",
			"msg":       "验证通过",
		}
		token := c.Request.Header.Get("token")
		if token == "" {
			trackData["status"] = "fail"
			trackData["msg"] = "请求头中无token"
			trackData["user_id"] = ""
			// 记录埋点日志
			trackEvent(trackData)
			//marshal, err := json.Marshal(trackData)
			//if err != nil {
			//	return
			//}

			c.JSON(400, gin.H{
				"msg":  "请登录",
				"data": 500,
			})
			c.Abort()
		}
		getToken, err := pkg.GenToken(token)
		if err != nil {
			trackData["status"] = "fail"
			trackData["msg"] = "token解析失败: " + err.Error()
			trackData["user_id"] = ""
			trackEvent(trackData)
			marshal, err := json.Marshal(trackData)
			if err != nil {
				return
			}
			logFile.Write(append(marshal, '\n'))
			c.JSON(400, gin.H{
				"msg":  "解析失败",
				"data": 500,
			})
			c.Abort()
		}
		userId := getToken["userId"].(string)
		trackData["user_id"] = userId
		// 记录埋点日志
		marshal, err := json.Marshal(trackData)
		if err != nil {
			return
		}
		logFile.Write(append(marshal, '\n'))

		trackEvent(trackData)
		c.Set("userId", userId)
		c.Next()
	}
}
func trackEvent(data map[string]interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("埋点日志序列化失败: %v", err)
		return
	}
	// 输出到控制台，生产环境可以写入日志文件或发送到日志平台
	log.Println("track:", string(jsonData))
}

var logFile *os.File

func CreateFile() {
	var err error
	logFile, err = os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
}
