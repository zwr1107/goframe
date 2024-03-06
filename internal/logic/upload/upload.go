package Upload

import (
	"context"
	"goframe/internal/dao"
	"goframe/internal/model"
	"goframe/internal/service"
	"strings"
	"time"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/os/gfile"
)

type sUpload struct{}

func init() {
	service.RegisterUpload(New())
}

func New() *sUpload {
	return &sUpload{}
}

func (s *sUpload) Upload(ctx context.Context, in model.UploadCreateInput) (out model.UploadCreateOutput, err error) {
	//定义文件上传的目录
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		g.Log().Fatal(ctx, "文件上传配置路径不能为空")
	}
	uploadPath = uploadPath + "/" + time.Now().Format("20060102")
	//判断目录是否存在
	if !gfile.Exists(uploadPath) {
		//不存在则创建
		if err = gfile.Mkdir(uploadPath); err != nil {
			return out, err
		}
	}
	//定义文件名
	//fileName := gfile.Basename(in.File.Filename)
	////获取后缀
	//fileExt := gfile.Ext(fileName)
	////生成随机文件名
	//randomName := gtime.TimestampMilliStr() + fileExt
	//
	////本地目录
	//Src := uploadPath + "/" + randomName
	//Url := Src

	//文件名后缀转成小写
	in.File.Filename = gstr.ToLower(in.File.Filename)

	//保存文件
	fileName, err := in.File.Save(uploadPath, true)

	if err != nil {
		return out, err
	}

	//拼接路径
	Src := "/" + uploadPath + "/" + fileName
	//用strings join 拼接路径图片路径
	//join() 方法用于把数组中的所有元素放入一个字符串。元素是通过指定的分隔符进行分隔的。
	Url := strings.Join([]string{"", uploadPath, fileName}, "/")

	//添加数据
	lastInsertID, err := dao.FileInfo.Ctx(ctx).Data(g.Map{
		"name":    fileName,
		"src":     Src,
		"url":     Url,
		"user_id": 1,
	}).InsertAndGetId()
	if err != nil {
		return out, err
	}

	//返回数据
	return model.UploadCreateOutput{Id: int(lastInsertID), Name: fileName, Src: Src, Url: Url, UserId: 1}, err
}

// Create 创建内容
func (s *sUpload) Create(ctx context.Context, in model.UploadCreateInput) (out model.UploadCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.FileInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UploadCreateOutput{Id: int(lastInsertID)}, err
}

// Delete 删除
func (s *sUpload) Delete(ctx context.Context, id uint) error {
	//链式操作，Transaction() 用于事务处理
	return dao.FileInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//Unscoped() 用于真删除
		_, err := dao.FileInfo.Ctx(ctx).Where(dao.FileInfo.Columns().Id, id).Delete()
		return err
	})
}

// GetList 获取列表
func (s *sUpload) GetList(ctx context.Context, in model.UploadGetListInput) (out *model.UploadGetListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.FileInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.UploadGetListOutput{
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
	out.List = make([]model.UploadGetListOutputItem, 0, in.Size)
	//6.把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return

}
