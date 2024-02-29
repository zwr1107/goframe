package controller

import (
	"context"
	"goframe/api/backend"
	"goframe/internal/model"
	"goframe/internal/service"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

// Create 创建
func (a *cRotation) Create(ctx context.Context, req *backend.RotationCreateReq) (res *backend.RotationCreateRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationCreateRes{RotationId: out.RotationId}, nil
}

// Delete 删除
func (a *cRotation) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.Id)
	return &backend.RotationDeleteRes{}, err
}

// Update 更新
func (a *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
		Id: req.Id,
	})
	return &backend.RotationUpdateRes{req}, err
}

// GetList 获取列表
func (a *cRotation) GetList(ctx context.Context, req *backend.RotationGetListReq) (res *backend.RotationGetListRes, err error) {
	list, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	//fmt.Printf("list: %+v\n", list)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &backend.RotationGetListRes{List: list.List, Page: list.Page, Size: list.Size, Total: list.Total}, nil
}
