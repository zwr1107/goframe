package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PositionCreateReq  PositionAddReq 添加请求参数
type PositionCreateReq struct {
	//path 为请求路径, tags 为接口标签, method 为请求方法, summary 为接口描述
	g.Meta    `path:"/backend/position/add" tags:"Position" method:"post" summary:"添加"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
	Sort      int    `json:"sort"    dc:"排序"`
}

// PositionCreateRes  PositionAddRes 添加返回参数
type PositionCreateRes struct {
	//新增后的id
	Id int `json:"id"`
}

// PositionDeleteReq  PositionDeleteReq 删除请求参数
type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete" method:"post" tags:"Position" summary:"删除"`
	Id     uint `v:"min:1#请选择需要删除的" json:"id" dc:"id"`
}

// PositionDeleteRes  PositionDeleteRes 删除返回参数
type PositionDeleteRes struct {
	//返回删除成功

}

// PositionUpdateReq 更新
type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update" method:"post" tags:"Position" summary:"更新"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
	Sort      int    `json:"sort"    dc:"排序"`
	Id        uint   `json:"id"  v:"required#请选择需要更新的数据" dc:"id"`
}

// PositionUpdateRes 更新返回
type PositionUpdateRes struct {
	//返回成功
	*PositionUpdateReq //返回修改后的数据
}

// PositionGetListReq  PositionGetListReq 获取列表请求参数
type PositionGetListReq struct {
	g.Meta              `path:"/backend/position/list" method:"get" tags:"Position" summary:"获取列表"`
	Sort                int `json:"sort" in:"query" dc:"排序类型"`
	CommonPaginationReq     //分页参数
}

// PositionGetListRes  列表返回
type PositionGetListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
