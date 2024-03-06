package controller

import (
	"context"
	"goframe/api/backend"
	"goframe/internal/model"
	"goframe/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// Upload 内容管理
var Upload = cUpload{}

type cUpload struct{}

// Delete 删除
func (a *cUpload) Delete(ctx context.Context, req *backend.UploadDeleteReq) (res *backend.UploadDeleteRes, err error) {
	err = service.Upload().Delete(ctx, uint(req.Id))
	return &backend.UploadDeleteRes{}, err
}

// GetList 获取列表
func (a *cUpload) GetList(ctx context.Context, req *backend.UploadGetListReq) (res *backend.UploadGetListRes, err error) {
	list, err := service.Upload().GetList(ctx, model.UploadGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	//fmt.Printf("list: %+v\n", list)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &backend.UploadGetListRes{List: list.List, Page: list.Page, Size: list.Size, Total: list.Total}, nil
}

// upload file 上传文件
func (a *cUpload) Upload(ctx context.Context, req *backend.UploadCreateReq) (res *backend.UploadCreateRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}
	upload, err := service.Upload().Upload(ctx, model.UploadCreateInput{
		UploadCreateUpdateBase: model.UploadCreateUpdateBase{
			File: req.File,
		},
	})

	if err != nil {
		return nil, err
	}
	return &backend.UploadCreateRes{
		Id:     upload.Id,
		Name:   upload.Name,
		Src:    upload.Src,
		Url:    upload.Url,
		UserId: upload.UserId,
	}, nil
}
