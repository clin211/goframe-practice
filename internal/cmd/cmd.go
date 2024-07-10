package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"corex/internal/controller/hello"
)

type User struct {
	Name  string
	Age   int
	Admin bool
}

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
			data := map[string]any{
				"nickname": "forest",
				"age":      "18",
			}
			s.BindHandler("/index", func(r *ghttp.Request) {
				r.Response.WriteTpl("index.html", data)
			})
			s.BindHandler("/template-tags", func(r *ghttp.Request) {
				// 创建用户数据列表
				var users = []struct {
					Name  string
					Age   int
					Admin bool
				}{
					{Name: "Alice", Age: 30, Admin: true},
					{Name: "Bob", Age: 25, Admin: false},
					{Name: "Charlie", Age: 28, Admin: false},
				}

				r.Response.WriteTpl("template-tags.html", g.Map{
					"data":  data,
					"users": users,
					"title": "GoFrame Template",
					"user":  User{Name: "forest", Age: 18, Admin: true},
					"footer": map[string]any{
						"copyright": "Copyright © 2019~2024",
						"links":     []string{"https://github.com/clin211", "https://github.com/gogf/gf"},
					},
				})
			})

			s.BindHandler("/template", func(r *ghttp.Request) {
				r.Response.WriteTpl("template.html", g.Map{
					"headerData": "GoFrame Template",
					"user":       User{Name: "forest", Age: 18, Admin: true},
					"footerData": map[string]any{
						"copyright": "Copyright © 2019~2024",
						"links":     []string{"https://github.com/clin211", "https://github.com/gogf/gf"},
					},
				})
			})
			s.Run()
			return nil
		},
	}
)
