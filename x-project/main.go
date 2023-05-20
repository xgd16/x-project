package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"x-project/src/global"
	"x-project/src/lib"
	"x-project/src/service"
)

func main() {
	// 初始化系统配置
	global.InitSystemConfig()
	// 初始化系统服务
	service.InitService()
	// 维持
	lib.Maintain()
}
