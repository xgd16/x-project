package route

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xgd16/gf-x-tool/xTool"
	"x-project/src/global"
)

// Api 路由注册
func Api(r *ghttp.RouterGroup) {
	r.GET("/test", func(r *ghttp.Request) {
		xTool.FastResp(r).Resp(global.SystemConfig.Get("server.name").String() + " 服务启动成功")
	})
}
