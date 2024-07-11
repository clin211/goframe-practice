// examples/func/index/main.go
package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	// 创建一个GoFrame的服务器实例
	s := g.Server()

	// 为服务器绑定一个处理根路径("/")的处理器
	s.BindHandler("/", func(r *ghttp.Request) {
		// 当请求根路径时，使用WriteTpl方法渲染模板并返回响应
		// 将当前时间作为参数"date"传递给模板，并提供formatDate函数用于时间格式化
		r.Response.WriteTpl("index.html", g.Map{
			"Map":    map[string]any{"key": "this is map"},
			"Slice":  []string{"this is slice 1", "this is slice 2"},
			"Array":  [2]string{"this is array 1", "this is array 2"},
			"String": "hello features",
		})
	})

	// 设置服务器监听的端口号为 8199
	s.SetPort(8199)

	// 启动服务器，监听并处理HTTP请求
	s.Run()
}
