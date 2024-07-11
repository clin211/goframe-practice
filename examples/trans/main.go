package main

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx  = gctx.New()
	i18n = gi18n.New()
)

func translate(key string) string {
	return i18n.Translate(ctx, key)
}

func main() {
	s := g.Server()
	i18n := g.I18n("i18n")
	i18n.SetLanguage("zh-CN")

	view := g.View()
	// 注册 i18n 函数到模板引擎
	view.BindFunc("_t", translate)

	s.BindHandler("/trans", func(r *ghttp.Request) {

		r.Response.WriteJson(g.Map{
			"lang":    r.Get("lang", "zh-CN"),
			"message": i18n.T(ctx, "hello"),
			"code":    http.StatusOK,
		})
	})

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			lang := r.GetRequest("lang", "zh-CN").String()
			r.SetCtx(gi18n.WithLanguage(r.Context(), lang))
			r.Middleware.Next()
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.WriteTpl("home.html", g.Map{
				"langCode": r.GetRequest("lang", "zh-CN").String(),
			})
		})
	})
	s.SetPort(8199)
	s.Run()
}
