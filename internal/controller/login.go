package controller

import (
	"context"
	"goframe/api/backend"
	"goframe/internal/service"
)

// Login 内容管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	//res = &backend.LoginDoRes{}
	//err = service.Login().Login(ctx, model.UserLoginInput{
	//	Name:     req.Name,
	//	Password: req.Password,
	//})
	//if err != nil {
	//	return
	//}
	//gutil.Dump(res)
	//res.Info = service.Session().GetUser(ctx)
	//
	//return
	res = &backend.LoginDoRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}

func (c *cLogin) RefreshToken(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
	res = &backend.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
