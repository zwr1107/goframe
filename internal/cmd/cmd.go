package cmd

import (
	"context"
	"goframe/internal/controller"
	"goframe/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goframe/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				//
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					hello.NewV1(),
					controller.Rotation,     //轮播图接口
					controller.Position,     //广告位接口
					controller.Login,        //登录接口
					controller.Admin.Create, // 管理员
					controller.Admin.Update, // 管理员
				)

				// 后台管理 ,需要token验证
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})

				})

			})
			s.Run()
			return nil
		},
	}
)
