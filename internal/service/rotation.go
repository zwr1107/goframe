// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe/internal/model"
)

type (
	IRotation interface {
		// Create 创建内容
		Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.RotationUpdateInput) error
		// GetList 获取列表
		GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error)
	}
)

var (
	localRotation IRotation
)

func Rotation() IRotation {
	if localRotation == nil {
		panic("implement not found for interface IRotation, forgot register?")
	}
	return localRotation
}

func RegisterRotation(i IRotation) {
	localRotation = i
}
