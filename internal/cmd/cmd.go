package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"study/internal/controller/hello"
	"study/internal/logic/products"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})
			s.BindHandler("/template", func(r *ghttp.Request) {
				data, _ := products.List()
				r.Response.WriteTpl("index.tpl", g.Map{
					"data": data,
				})
			})
			s.Run()
			return nil
		},
	}
)
