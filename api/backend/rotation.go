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
type RotationCreateRes struct {
	//新增后的id
	RotationId int `json:"rotation_id"`
}
