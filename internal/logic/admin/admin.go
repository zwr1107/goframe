package Admin

import (
	"context"
	"goframe/internal/dao"
	"goframe/internal/model"
	"goframe/internal/model/entity"
	"goframe/internal/service"
	"goframe/utility"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/go-errors/errors"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/util/grand"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

// Create 创建内容
func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	////判断用户名是否存在
	adminInfo := entity.AdminInfo{}
	err = dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&adminInfo)
	if err == nil {
		return out, errors.New("用户名已存在")
	}

	//密码加盐,随机生成一个盐值
	Salt := grand.S(10)
	//密码加密
	in.Password = utility.EncryptPassword(in.Password, Salt)
	//盐值
	in.UserSalt = Salt

	//插入数据返回id
	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{Id: int(lastInsertID)}, err
}

// Delete 删除
func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	//链式操作，Transaction() 用于事务处理
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//Unscoped() 用于真删除
		_, err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Id, id).Delete()
		return err
	})
}

// Update 修改
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		//判断密码是否为空
		if in.Password != "" {
			//密码加盐,随机生成一个盐值
			Salt := grand.S(10)
			//密码加密
			in.Password = utility.EncryptPassword(in.Password, Salt)
			//盐值
			in.UserSalt = Salt
		} else {
			//不修改密码的情况下，不修改盐值
			in.UserSalt = ""
		}

		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id). //它用于在更新或插入操作中排除某些字段。这个方法接受一个或多个字段名作为参数，并在执行数据库操作时忽略这些字段。
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 获取列表
func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.AdminInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.AdminGetListOutputItem, 0, in.Size)
	//6.把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return

}

func (s *sAdmin) GetUserByUserNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	} else {
		return g.Map{
			"id":       adminInfo.Id,
			"username": adminInfo.Name,
		}
	}
}
