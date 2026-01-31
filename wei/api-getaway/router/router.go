package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/handler/server"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/register", server.Register)
	r.POST("/login", server.Login)
	r.POST("/update", middleware.AuthToken(), server.Update)
	r.POST("/get/token", server.RefreshTokenHandler)
	return r
}
