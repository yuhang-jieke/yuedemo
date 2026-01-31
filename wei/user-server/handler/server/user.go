package server

import (
	"context"

	"github.com/yuhang-jieke/yuedemo/wei/user-server/basic/config"
	__ "github.com/yuhang-jieke/yuedemo/wei/user-server/handler/proto"
	"github.com/yuhang-jieke/yuedemo/wei/user-server/model"
)

type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {
	user := model.User{
		Name:    in.Name,
		Age:     int(in.Age),
		Address: in.Address,
	}
	if err := user.Registers(config.DB); err != nil {
		panic("注册失败")
	}
	return &__.RegisterResp{
		Greet: "注册成功",
	}, nil
}
func (s *Server) Update(_ context.Context, in *__.UpdateReq) (*__.UpdateResp, error) {

	var user model.User
	if err := user.Update(config.DB, in); err != nil {
		panic("修改失败")
	}
	if user.ID != uint(in.UserId) {
		panic("未登录")
	}
	return &__.UpdateResp{
		Greet: "修改成功",
	}, nil
}
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {

	var user model.User
	if err := user.FindName(config.DB, in.Name); err != nil {
		panic("查询失败")
	}
	if in.Age != int64(user.Age) {
		panic("年龄不正确")
	}
	return &__.LoginResp{
		Greet:  "登录成功",
		UserId: int64(user.ID),
	}, nil
}
