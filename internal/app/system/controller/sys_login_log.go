/*
* @desc:登录日志管理

* @Date:   2022/4/24 22:14
 */

package controller

import (
	"context"
	"tv3a/api/v1/system"
	"tv3a/internal/app/system/service"
)

var LoginLog = loginLogController{}

type loginLogController struct {
	BaseController
}

func (c *loginLogController) List(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error) {
	res, err = service.SysLoginLog().List(ctx, req)
	return
}

func (c *loginLogController) Delete(ctx context.Context, req *system.LoginLogDelReq) (res *system.LoginLogDelRes, err error) {
	err = service.SysLoginLog().DeleteLoginLogByIds(ctx, req.Ids)
	return
}

func (c *loginLogController) Clear(ctx context.Context, req *system.LoginLogClearReq) (res *system.LoginLogClearRes, err error) {
	err = service.SysLoginLog().ClearLoginLog(ctx)
	return
}
