package service

import (
	"fmt"
	"x-project/src/service/queue"
	"x-project/src/service/web"
)

// 填写注册服务
var register = map[string]func(){
	"WEB":   web.Service,
	"Queue": queue.Service,
}

// InitService 初始化系统服务
func InitService() {
	if len(register) <= 0 {
		fmt.Println("没有启动任何基础服务")
		return
	}

	for s, f := range register {
		go f()
		fmt.Printf("%s 服务已启动\n", s)
	}
}
