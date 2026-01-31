package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/config"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/pkg"
)

var (
	logfile *os.File
	logmu   sync.Mutex
)

func CreateFile() {
	var err error
	logfile, err = os.OpenFile("logss.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("日式文件打开失败")
	}

}
func LogToken(userId, path, status string) {
	logmu.Lock()
	defer logmu.Unlock()
	data, _ := json.Marshal(map[string]string{
		"userid": userId,
		"path":   path,
		"status": status,
		"time":   time.Now().Format("2006010215"),
	})
	logfile.Write(append(data, '\n'))
}
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		userId := "5"
		path := c.Request.URL.Path

		if token == "" {
			LogToken(userId, path, "未登录") // 埋点
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "未登录",
			})
			c.Abort()
		}

		//todo 埋点：对jwt进行解析出用户信息，把用户/时间/路由/入参/返回等记录下来
		redistoken, err := config.Rdb.Exists(context.Background(), "key"+token).Result()
		if err != nil {
			LogToken(userId, path, "redis认证失败") // 埋点
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "redis认证失败",
			})

			c.Abort()
		}
		if redistoken > 0 {
			LogToken(userId, path, "被加入黑名单") // 埋
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "被加入黑名单",
			})
			c.Abort()
		}
		persontoken, err := pkg.PersonToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			c.Abort()
		}
		if c.Writer.Status() == http.StatusOK {
			c.Writer.Header().Set("Cache-Control", "public,max-age=600")
		}
		LogToken(userId, path, "完成") // 埋点
		RegisterCount := config.Rdb.Incr(context.Background(), "count").Val()
		if RegisterCount > 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "请勿重复操作",
			})
			c.Abort()
			return
		}
		c.Set("userId", persontoken["userId"])
		c.Next()

	}
}
