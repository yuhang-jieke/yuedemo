package config

import (
	"github.com/redis/go-redis/v9"
	__ "github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/proto"
	"go.uber.org/zap"
)

var (
	UserClient __.UserClient
	Rdb        *redis.Client
	Logg       *zap.Logger
)
