package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadCreateUpdateBase 创建/修改内容基类
type UploadCreateUpdateBase struct {
	File *ghttp.UploadFile //上传文件对象
}

// UploadCreateInput 创建内容
type UploadCreateInput struct {
	UploadCreateUpdateBase

	//File *ghttp.UploadFile
}

// UploadCreateOutput 创建内容返回结果
type UploadCreateOutput struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Src    string `json:"src"`
	Url    string `json:"url"`
	UserId uint   `json:"user_id"`
}

// UploadUpdateInput 修改内容
type UploadUpdateInput struct {
	UploadCreateUpdateBase
	Id uint `json:"id"`
}

// UploadGetListInput 获取内容列表
type UploadGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// UploadGetListOutput 查询列表结果
type UploadGetListOutput struct {
	List  []UploadGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

type UploadGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	Desc      string      `json:"desc"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
