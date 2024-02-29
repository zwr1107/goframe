package Rotation

import (
	"context"
	"goframe/internal/dao"
	"goframe/internal/model"
	"goframe/internal/service"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建内容
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	//链式操作，Transaction() 用于事务处理
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//Unscoped() 用于真删除
		_, err := dao.RotationInfo.Ctx(ctx).Where(dao.RotationInfo.Columns().Id, id).Unscoped().Delete()
		return err
	})
}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id). //它用于在更新或插入操作中排除某些字段。这个方法接受一个或多个字段名作为参数，并在执行数据库操作时忽略这些字段。
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}
