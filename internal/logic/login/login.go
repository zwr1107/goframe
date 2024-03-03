package login

import (
	"context"
	"goframe/internal/dao"
	"goframe/internal/model"
	"goframe/internal/model/entity"
	"goframe/internal/service"
	"goframe/utility"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// Login 登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	// 查询用户信息
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&adminInfo)

	//gutil.Dump(adminInfo, utility.EncryptPassword(in.Password, adminInfo.UserSalt))
	if err != nil {
		return err
	}
	// 验证密码
	if adminInfo.Password != utility.EncryptPassword(in.Password, adminInfo.UserSalt) {
		return gerror.New("账号或者密码不正确")
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 登录成功,更新用户信息
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}

// 注销
func (s *sLogin) Logout(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}
