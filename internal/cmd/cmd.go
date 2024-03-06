package cmd

import (
	"context"
	"goframe/api/backend"
	"goframe/internal/consts"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/service"
	"goframe/utility"
	"goframe/utility/response"
	"strconv"

	"github.com/gogf/gf/v2/util/gmode"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 静态目录设置
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			//fmt.Println("uploadPath:", uploadPath)

			//g.Dump("uploadPath:", uploadPath)
			if uploadPath == "" {
				g.Log().Fatal(ctx, "文件上传配置路径不能为空")
				//panic("文件上传配置路径不能为空")
			}
			s.AddStaticPath("/upload", uploadPath)

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}

			// 启动gtoken
			gfAdminToken := &gtoken.GfToken{
				CacheMode:        2,
				ServerName:       "goframe",                                                 //服务名称
				LoginPath:        "/backend/login",                                          //登录路径
				LoginBeforeFunc:  loginFunc,                                                 //登录前的函数
				LoginAfterFunc:   loginAfterFunc,                                            //登录后的函数
				LogoutPath:       "/backend/user/logout",                                    //退出路径
				AuthPaths:        g.SliceStr{"/backend/admin/info"},                         //拦截路径
				AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				AuthAfterFunc:    authAfterFunc,
				MultiLogin:       true, //是否允许多点登录
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//gtoken中间件绑定
				//err := gfAdminToken.Middleware(ctx, group)
				//if err != nil {
				//	panic(err)
				//}
				group.Bind(
					controller.Rotation,      // 轮播图
					controller.Position,      // 手工位
					controller.Admin.Create,  // 管理员
					controller.Admin.Update,  // 管理员
					controller.Admin.Delete,  // 管理员
					controller.Admin.GetList, // 管理员
					controller.Login,         // 登录
					controller.Role,          //角色管理
					//controller.Permission,    //权限管理
					controller.Upload, //上传
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth) //for jwt
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})
					//
					//group.Bind(
					//	controller.Role,
					//)
				})
			})
			s.Run()
			return nil
		},
	}
)

// todo 迁移到合适的位置
func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// todo 迁移到合适的位置
// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		//去掉前缀
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		g.Dump("adminId:", adminId)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		//获得权限id，定义一个切片
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		//通过权限id查询权限
		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    10 * 24 * 60 * 60, //单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	//g.Dump("respData:", respData)
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)

	if err != nil {
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}
