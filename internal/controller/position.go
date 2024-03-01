package controller

import (
	"context"
	"goframe/api/backend"
	"goframe/internal/model"
	"goframe/internal/service"
)

// Position 内容管理
var Position = cPosition{}

type cPosition struct{}

// Create 创建
func (a *cPosition) Create(ctx context.Context, req *backend.PositionCreateReq) (res *backend.PositionCreateRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionCreateRes{Id: out.Id}, nil
}

// Delete 删除
func (a *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return &backend.PositionDeleteRes{}, err
}

// Update 更新
func (a *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
		Id: req.Id,
	})
	return &backend.PositionUpdateRes{req}, err
}

// GetList 获取列表
func (a *cPosition) GetList(ctx context.Context, req *backend.PositionGetListReq) (res *backend.PositionGetListRes, err error) {
	list, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	//fmt.Printf("list: %+v\n", list)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &backend.PositionGetListRes{List: list.List, Page: list.Page, Size: list.Size, Total: list.Total}, nil
}
