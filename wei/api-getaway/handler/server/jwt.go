package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/pkg"
)

func RefreshTokenHandler(c *gin.Context) {
	// 1. 从请求头获取旧token（和你现有接口传token的方式一致）
	oldToken := c.GetHeader("token")
	if oldToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "请传入旧token"})
		return
	}

	// 2. 调用刷新函数生成新token
	newToken, err := pkg.RefreshToken(oldToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	// 3. 返回新token
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "token刷新成功",
		"token": newToken,
	})
}
