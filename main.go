package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gview"
)

func main() {
	ctx := gctx.New()
	s := g.Server()

	// 设置 i18n 配置
	i18n := gi18n.New(gi18n.Options{
		Path: "i18n",
	})
	i18n.SetLanguage("en") // 默认语言

	// 自定义函数
	funcMap := map[string]interface{}{
		"i18n": func(key string) string {
			return i18n.Translate(ctx, key)
		},
	}

	// 添加自定义模板引擎并注册函数
	view := gview.New()
	view.BindFuncMap(funcMap)
	s.SetView(view)

	// 注册路由
	s.BindHandler("/", func(r *ghttp.Request) {
		lang := r.GetRequest("lang", "en").String()
		if lang != "" {
			i18n.SetLanguage(lang)
		}

		r.Response.WriteTpl("index.html", nil)
	})

	s.SetPort(8199)
	// 启动服务器
	s.Run()
}
