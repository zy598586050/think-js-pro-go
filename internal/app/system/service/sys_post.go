// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"tv3a/api/v1/system"
	"tv3a/internal/app/system/model/entity"
)

type ISysPost interface {
	List(ctx context.Context, req *system.PostSearchReq) (res *system.PostSearchRes, err error)
	Add(ctx context.Context, req *system.PostAddReq) (err error)
	Edit(ctx context.Context, req *system.PostEditReq) (err error)
	Delete(ctx context.Context, ids []int) (err error)
	GetUsedPost(ctx context.Context) (list []*entity.SysPost, err error)
}

var localSysPost ISysPost

func SysPost() ISysPost {
	if localSysPost == nil {
		panic("implement not found for interface ISysPost, forgot register?")
	}
	return localSysPost
}

func RegisterSysPost(i ISysPost) {
	localSysPost = i
}
