package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadCreateReq  UploadAddReq 添加请求参数
type UploadCreateReq struct {
	//path 为请求路径, tags 为接口标签, method 为请求方法, summary 为接口描述
	g.Meta `path:"/backend/upload/add" tags:"upload" method:"post" summary:"添加"`
	File   *ghttp.UploadFile `json:"file" v:"required#文件不能为空" dc:"文件"`
}

// UploadCreateRes  UploadAddRes 添加返回参数
type UploadCreateRes struct {
	//新增后的id
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Src    string `json:"src"`
	Url    string `json:"url"`
	UserId uint   `json:"user_id"`
}

type UploadDeleteReq struct {
	g.Meta `path:"/backend/upload/delete" tags:"upload" method:"post" summary:"删除"`
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
}

// UploadDeleteRes  UploadDeleteRes 删除返回参数
type UploadDeleteRes struct {
	//返回删除成功

}

// UploadGetListReq  UploadGetListReq 获取列表请求参数
type UploadGetListReq struct {
	g.Meta              `path:"/backend/upload/list" method:"get" tags:"upload" summary:"获取列表"`
	Sort                int `json:"sort" in:"query" dc:"排序类型"`
	CommonPaginationReq     //分页参数
}

// UploadGetListRes  列表返回
type UploadGetListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
