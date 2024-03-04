package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminCreateReq  AdminAddReq 添加请求参数
type AdminCreateReq struct {
	//path 为请求路径, tags 为接口标签, method 为请求方法, summary 为接口描述
	g.Meta   `path:"/backend/admin/add" tags:"Admin" method:"post" summary:"添加"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色ids"`
	IsAdmin  int    `json:"is_admin"    dc:"是否超级管理员"`
}

// AdminCreateRes  AdminAddRes 添加返回参数
type AdminCreateRes struct {
	//新增后的id
	Id int `json:"id"`
}

// AdminDeleteReq  AdminDeleteReq 删除请求参数
type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" method:"post" tags:"Admin" summary:"删除"`
	Id     uint `v:"min:1#请选择需要删除的" json:"id" dc:"id"`
}

// AdminDeleteRes  AdminDeleteRes 删除返回参数
type AdminDeleteRes struct {
	//返回删除成功

}

// AdminUpdateReq 更新
type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update" method:"post" tags:"Admin" summary:"更新"`
	Id       uint   `json:"id"      v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色ids"`
	IsAdmin  int    `json:"is_admin"    dc:"是否超级管理员"`
}

// AdminUpdateRes 更新返回
type AdminUpdateRes struct {
	//返回成功
	*AdminUpdateReq //返回修改后的数据
}

// AdminGetListReq  AdminGetListReq 获取列表请求参数
type AdminGetListReq struct {
	g.Meta              `path:"/backend/admin/list" method:"get" tags:"Admin" summary:"获取列表"`
	Sort                int `json:"sort" in:"query" dc:"排序类型"`
	CommonPaginationReq     //分页参数
}

// AdminGetListRes  列表返回
type AdminGetListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// AdminGetInfoReq  获取信息请求
type AdminGetInfoReq struct {
	g.Meta `path:"/backend/admin/info" method:"get"`
}

// AdminGetInfoRes for gtoken
type AdminGetInfoRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	RoleIds string `json:"role_ids"`
	IsAdmin int    `json:"is_admin"`
}

// AdminGetInfoGtokenRes for jwt
type AdminGetInfoGtokenRes struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}
