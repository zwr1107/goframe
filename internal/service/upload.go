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
	IUpload interface {
		Upload(ctx context.Context, in model.UploadCreateInput) (out model.UploadCreateOutput, err error)
		// Create 创建内容
		Create(ctx context.Context, in model.UploadCreateInput) (out model.UploadCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// GetList 获取列表
		GetList(ctx context.Context, in model.UploadGetListInput) (out *model.UploadGetListOutput, err error)
	}
)

var (
	localUpload IUpload
)

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}
