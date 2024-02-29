package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RotationCreateReq  RotationAddReq 添加轮播图请求参数
type RotationCreateReq struct {
	//path 为请求路径, tags 为接口标签, method 为请求方法, summary 为接口描述
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"添加轮播图"`
	//图片链接 json 标签为请求参数, v 为验证规则, dc 为参数描述
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	//跳转链接
	Link string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	//排序
	Sort int `json:"sort" dc:"排序"`
}

// RotationCreateRes  RotationAddRes 添加轮播图返回参数
type RotationCreateRes struct {
	//新增后的id
	RotationId int `json:"rotation_id"`
}

// RotationDeleteReq  RotationDeleteReq 删除轮播图请求参数
type RotationDeleteReq struct {
	g.Meta `path:"/backend/rotation/delete" method:"post" tags:"Rotation" summary:"删除轮播图"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" json:"id" dc:"轮播图id"`
}

// RotationDeleteRes  RotationDeleteRes 删除轮播图返回参数
type RotationDeleteRes struct {
	//返回删除成功

}

// 更新
type RotationUpdateReq struct {
	g.Meta `path:"/backend/rotation/update" method:"post" tags:"Rotation" summary:"更新轮播图"`
	Id     uint   `json:"id" v:"min:1#请选择需要修改的轮播图" dc:"轮播图id"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort" dc:"排序"`
}

// RotationUpdateRes 更新返回
type RotationUpdateRes struct {
	//返回成功
	*RotationUpdateReq
}

// RotationGetListReq  RotationGetListReq 获取轮播图列表请求参数
type RotationGetListReq struct {
	g.Meta `path:"/backend/rotation/list" method:"get" tags:"Rotation" summary:"获取轮播图列表"`
	Page   int `json:"page" v:"min:1#页码必须大于0" dc:"页码"`
	Size   int `json:"size" v:"min:1#每页数量必须大于0" dc:"每页数量"`
}
