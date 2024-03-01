package Position

import (
	"context"
	"goframe/internal/dao"
	"goframe/internal/model"
	"goframe/internal/service"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

// Create 创建内容
func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PositionCreateOutput{Id: int(lastInsertID)}, err
}

// Delete 删除
func (s *sPosition) Delete(ctx context.Context, id uint) error {
	//链式操作，Transaction() 用于事务处理
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//Unscoped() 用于真删除
		_, err := dao.PositionInfo.Ctx(ctx).Where(dao.PositionInfo.Columns().Id, id).Delete()
		return err
	})
}

// Update 修改
func (s *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.PositionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PositionInfo.Columns().Id). //它用于在更新或插入操作中排除某些字段。这个方法接受一个或多个字段名作为参数，并在执行数据库操作时忽略这些字段。
			Where(dao.PositionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 获取列表
func (s *sPosition) GetList(ctx context.Context, in model.PositionGetListInput) (out *model.PositionGetListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.PositionInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.PositionGetListOutput{
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
	out.List = make([]model.PositionGetListOutputItem, 0, in.Size)
	//6.把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return

}
