package init

import (
	"fmt"

	"github.com/gospacex/gospacex/core/storage/conf"
	"github.com/gospacex/gospacex/core/storage/db/mysql"
	"github.com/yuhang-jieke/yuedemo/wei/user-server/basic/config"
	"github.com/yuhang-jieke/yuedemo/wei/user-server/model"
)

func MysqlInit() {

	var err error
	config.DB, err = mysql.Init(true, "debug", conf.Cfg.Mysql)

	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	fmt.Printf("数据库连接成功", err)
	err = config.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("数据表迁移失败")
	}
	fmt.Printf("数据表迁移成功")
}
