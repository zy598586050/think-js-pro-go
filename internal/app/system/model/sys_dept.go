/*
* @desc:部门model

* @Date:   2022/4/11 9:07
 */

package model

import "tv3a/internal/app/system/model/entity"

type SysDeptTreeRes struct {
	*entity.SysDept
	Children []*SysDeptTreeRes `json:"children"`
}
