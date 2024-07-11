package main

import (
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func lowercase(s string) string {
	return strings.ToLower(s)
}

func main() {
	// 创建一个GoFrame的服务器实例
	s := g.Server()

	// 为模板绑定自定义函数
	view := g.View()
	view.BindFunc("uppercase", uppercase)
	view.BindFunc("lowercase", lowercase)

	// 为服务器绑定一个处理根路径("/")的处理器
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("custom.html", g.Map{})
	})

	// 设置服务器监听的端口号为 8199
	s.SetPort(8199)

	// 启动服务器，监听并处理HTTP请求
	s.Run()
}
