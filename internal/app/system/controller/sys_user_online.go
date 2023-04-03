/*
* @desc:在线用户管理

* @Date:   2023/1/10 17:23
 */

package controller

import (
	"context"
	"tv3a/api/v1/system"
	"tv3a/internal/app/system/service"
)

var UserOnline = new(SysUserOnlineController)

type SysUserOnlineController struct{}

func (c *SysUserOnlineController) List(ctx context.Context, req *system.SysUserOnlineSearchReq) (res *system.SysUserOnlineSearchRes, err error) {
	res, err = service.SysUserOnline().GetOnlineListPage(ctx, req)
	return
}

func (c *SysUserOnlineController) ForceLogout(ctx context.Context, req *system.SysUserOnlineForceLogoutReq) (res *system.SysUserOnlineForceLogoutRes, err error) {
	err = service.SysUserOnline().ForceLogout(ctx, req.Ids)
	return
}
