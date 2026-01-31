package inits

import (
	"context"
	"fmt"

	"github.com/gospacex/gospacex/core/logger"
	cf "github.com/gospacex/gospacex/core/storage/conf"
	"github.com/yuhang-jieke/yuedemo/wei/api-getaway/basic/config"
)

func LogInit() {

	var logConf = cf.LogConfig{
		Level:         "info",
		MaxSize:       100, // megabytes
		MaxBackups:    5,
		MaxAge:        15, // 15 days
		Compress:      true,
		Path:          "C:\\Users\\ZhuanZ\\Desktop\\zuoye1.28\\yuedemo\\wei\\api-getaway\\basic\\cmd\\logs",
		ConsoleEnable: true,
	}
	err := logger.Init(context.Background(), &logConf)

	if err != nil {
		panic("logger初始化失败")
		return
	}
	config.Logg = logger.L
	fmt.Println("logger初始化成功")
}
