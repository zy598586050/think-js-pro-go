/*
* @desc:后台路由

* @Date:   2022/2/18 17:34
 */

package router

import (
	"context"
	"tv3a/internal/app/system/controller"
	"tv3a/internal/app/system/service"
	"tv3a/library/libRouter"

	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Bind(
			//登录
			controller.Login,
		)
		//登录验证拦截
		service.GfToken().Middleware(group)
		//context拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		group.Bind(
			controller.User,
			controller.Menu,
			controller.Role,
			controller.Dept,
			controller.Post,
			controller.DictType,
			controller.DictData,
			controller.Config,
			controller.Monitor,
			controller.LoginLog,
			controller.OperLog,
			controller.Personal,
			controller.UserOnline,
			controller.Cache, // 缓存处理
		)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
