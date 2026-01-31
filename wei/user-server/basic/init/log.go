package init

import (
	"context"
	"fmt"

	"github.com/gospacex/gospacex/core/logger"
	"github.com/gospacex/gospacex/core/storage/conf"
)

func LogInit() {
	err := logger.Init(context.Background(), conf.Cfg.Log)
	if err != nil {
		panic("logger初始化失败")
		return
	}
	fmt.Println("logger初始化成功")
}
