package controller

import (
	"context"
	"goframe/api/backend"
	"goframe/internal/model"
	"goframe/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

// Create 创建
func (a *cAdmin) Create(ctx context.Context, req *backend.AdminCreateReq) (res *backend.AdminCreateRes, err error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminCreateRes{Id: out.Id}, nil
}

// Delete 删除
func (a *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	return &backend.AdminDeleteRes{}, err
}

// Update 更新
func (a *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
		Id: req.Id,
	})
	return &backend.AdminUpdateRes{req}, err
}

// GetList 获取列表
func (a *cAdmin) GetList(ctx context.Context, req *backend.AdminGetListReq) (res *backend.AdminGetListRes, err error) {
	list, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	//fmt.Printf("list: %+v\n", list)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &backend.AdminGetListRes{List: list.List, Page: list.Page, Size: list.Size, Total: list.Total}, nil
}

func (c *cAdmin) Info(ctx context.Context, req *backend.AdminGetInfoReq) (res *backend.AdminGetInfoRes, err error) {
	return &backend.AdminGetInfoRes{
		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
		IdentityKey: service.Auth().IdentityKey,
		Payload:     service.Auth().GetPayload(ctx),
	}, nil
}
