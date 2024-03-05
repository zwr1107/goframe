package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoleCreateReq  RoleAddReq 添加请求参数
type RoleCreateReq struct {
	//path 为请求路径, tags 为接口标签, method 为请求方法, summary 为接口描述
	g.Meta `path:"/backend/role/add" tags:"role" method:"post" summary:"添加"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

// RoleCreateRes  RoleAddRes 添加返回参数
type RoleCreateRes struct {
	//新增后的id
	Id int `json:"id"`
}

// RoleDeleteReq  RoleDeleteReq 删除请求参数
type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"post" tags:"role" summary:"删除"`
	Id     uint `v:"min:1#请选择需要删除的" json:"id" dc:"id"`
}

// RoleDeleteRes  RoleDeleteRes 删除返回参数
type RoleDeleteRes struct {
	//返回删除成功

}

// RoleUpdateReq 更新
type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"post" tags:"role" summary:"更新"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
	Id     uint   `json:"id"  v:"required#请选择需要更新的数据" dc:"id"`
}

// RoleUpdateRes 更新返回
type RoleUpdateRes struct {
	//返回成功
	*RoleUpdateReq //返回修改后的数据
}

// RoleGetListReq  RoleGetListReq 获取列表请求参数
type RoleGetListReq struct {
	g.Meta              `path:"/backend/role/list" method:"get" tags:"role" summary:"获取列表"`
	CommonPaginationReq //分页参数
}

// RoleGetListRes  列表返回
type RoleGetListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AddPermissionReq struct {
	g.Meta       `path:"/backend/role/add/permission" method:"post" tags:"角色" summary:"角色添加权限接口"`
	RoleId       uint `json:"role_id" desc:"角色id"`
	PermissionId uint `json:"permission_id" desc:"权限id"`
}

type AddPermissionRes struct {
	Id uint `json:"id"`
}

type DeletePermissionReq struct {
	g.Meta       `path:"/backend/role/delete/permission" method:"post" tags:"角色" summary:"角色删除权限接口"`
	RoleId       uint `json:"role_id" desc:"角色id"`
	PermissionId uint `json:"permission_id" desc:"权限id"`
}

type DeletePermissionRes struct {
}
