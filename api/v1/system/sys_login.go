/*
* @desc:登录

* @Date:   2022/4/27 21:51
 */

package system

import (
	commonApi "tv3a/api/v1/common"
	"tv3a/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta     `path:"/login" tags:"登录" method:"post" summary:"用户登录"`
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}

type UserLoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	MenuList    []*model.UserMenus  `json:"menuList"`
	Permissions []string            `json:"permissions"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/logout" tags:"登录" method:"get" summary:"退出登录"`
	commonApi.Author
}

type UserLoginOutRes struct {
}

type PvReq struct {
	g.Meta    `path:"/pv" tags:"上报" method:"post" summary:"数据上报"`
	Url       string `p:"url" v:"required"`
	Event     string `p:"event" v:"required"`
	UtmSource string `p:"utm_source"`
	UtmId     string `p:"utm_id"`
}

type PvRes struct {
}
