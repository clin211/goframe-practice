package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gview"
)

func main() {
	s := g.Server()

	// 设置 i18n 配置
	i18n := gi18n.New(gi18n.Options{
		Path: "i18n",
	})
	i18n.SetLanguage("en") // 默认语言

	// 添加自定义模板引擎
	s.SetView(gview.New())

	// 注册路由
	s.BindHandler("/", func(r *ghttp.Request) {
		lang := r.Get("lang").String()
		if lang != "" {
			i18n.SetLanguage(lang)
		}

		r.Response.WriteTpl("index.html", g.Map{
			"hello": i18n.Translate(r.Context(), "hello"),
		})
	})

	// 启动服务器
	s.Run()
}
