package init

import (
	"fmt"

	redis2 "github.com/gospacex/gospacex/core/storage/cache/redis"
	"github.com/gospacex/gospacex/core/storage/conf"
	"github.com/yuhang-jieke/yuedemo/wei/user-server/basic/config"
)

func RedisInit() {
	var err error
	r := conf.Cfg.Redis
	fmt.Println(r)
	err = redis2.Init(true, conf.Cfg.Redis)
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}
	config.RDB = redis2.RC
	fmt.Println("redis连接成功")
}
