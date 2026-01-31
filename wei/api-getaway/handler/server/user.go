package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/config"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/inits"
	__ "github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/proto"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/handler/request"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/pkg"
)

func Register(c *gin.Context) {
	var form request.User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	result, _ := config.Rdb.Get(context.Background(), "key"+form.Name).Result()
	if form.Name == result {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "已经注册过",
		})
		return
	}
	_, err := config.UserClient.Register(c, &__.RegisterReq{
		Name:    form.Name,
		Age:     int64(form.Age),
		Address: form.Address,
	})
	if err != nil {
		panic("注册失败")
	}
	err = config.Rdb.Set(context.Background(), "key"+form.Name, form.Name, 0).Err()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "加入缓存失败",
		})
		return
	}
	go func() {
		usermap := map[string]string{
			"name":    form.Name,
			"age":     strconv.Itoa(form.Age),
			"address": form.Address,
		}
		_, err = inits.ElasticClient.Index().Index("usermap").BodyJson(usermap).Do(context.Background())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "es同步失败",
			})
			return
		}
	}()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
	return
}
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	r, err := config.UserClient.Login(c, &__.LoginReq{
		Name: form.Name,
		Age:  int64(form.Age),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "登录失败" + err.Error(),
		})
		return
	}
	handler, err := pkg.TokenHandler(r.Greet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "token生成失败" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": handler,
	})
	return
}
func Update(c *gin.Context) {
	var form request.Update

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不正确",
		})
		return
	}
	_, err := config.UserClient.Update(c, &__.UpdateReq{
		Id:      int64(form.Id),
		Address: form.Address,
	})
	if err != nil {
		panic("修改失败")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
	return
}
