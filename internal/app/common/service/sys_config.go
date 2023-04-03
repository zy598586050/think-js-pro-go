// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"tv3a/api/v1/system"
	"tv3a/internal/app/common/model/entity"
)

type ISysConfig interface {
	List(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error)
	Add(ctx context.Context, req *system.ConfigAddReq, userId uint64) (err error)
	CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int64) (err error)
	Get(ctx context.Context, id int) (res *system.ConfigGetRes, err error)
	Edit(ctx context.Context, req *system.ConfigEditReq, userId uint64) (err error)
	Delete(ctx context.Context, ids []int) (err error)
	GetConfigByKey(ctx context.Context, key string) (config *entity.SysConfig, err error)
	GetByKey(ctx context.Context, key string) (config *entity.SysConfig, err error)
}

var localSysConfig ISysConfig

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
