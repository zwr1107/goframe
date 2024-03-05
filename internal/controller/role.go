package controller

import (
	"context"
	"goframe/api/backend"
	"goframe/internal/model"
	"goframe/internal/service"
)

// Role 内容管理
var Role = cRole{}

type cRole struct{}

// Create 创建
func (a *cRole) Create(ctx context.Context, req *backend.RoleCreateReq) (res *backend.RoleCreateRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleCreateRes{Id: out.Id}, nil
}

// Delete 删除
func (a *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return &backend.RoleDeleteRes{}, err
}

// Update 更新
func (a *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Desc: req.Desc,
			Name: req.Name,
		},
		Id: req.Id,
	})
	return &backend.RoleUpdateRes{req}, err
}

// GetList 获取列表
func (a *cRole) GetList(ctx context.Context, req *backend.RoleGetListReq) (res *backend.RoleGetListRes, err error) {
	list, err := service.Role().GetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	//fmt.Printf("list: %+v\n", list)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &backend.RoleGetListRes{List: list.List, Page: list.Page, Size: list.Size, Total: list.Total}, nil
}

// AddPermission 添加权限
func (a *cRole) AddPermission(ctx context.Context, req *backend.AddPermissionReq) (res *backend.AddPermissionRes, err error) {
	out, err := service.Role().AddPermission(ctx, model.RoleAddPermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AddPermissionRes{Id: out.Id}, nil
}

// DeletePermission 删除权限
func (a *cRole) DeletePermission(ctx context.Context, req *backend.DeletePermissionReq) (res *backend.DeletePermissionRes, err error) {
	err = service.Role().DeletePermission(ctx, model.RoleDeletePermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	return &backend.DeletePermissionRes{}, err
}
