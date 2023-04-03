/*
* @desc:缓存处理

* @Date:   2023/2/1 18:12
 */

package system

import (
	commonApi "tv3a/api/v1/common"

	"github.com/gogf/gf/v2/frame/g"
)

type CacheRemoveReq struct {
	g.Meta `path:"/cache/remove" tags:"缓存管理" method:"delete" summary:"清除缓存"`
	commonApi.Author
}

type CacheRemoveRes struct {
	commonApi.EmptyRes
}
