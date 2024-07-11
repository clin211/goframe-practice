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
		r.Response.WriteTpl("other.html", g.Map{
			"len": map[string]any{
				"string": "hello features",
				"array":  [5]string{"a", "b", "c", "d", "e"},
				"slice":  []string{"a", "b", "c", "d", "e"},
				"map":    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
				"chan":   make(chan int, 10),
			},
		})
	})

	// 设置服务器监听的端口号为 8199
	s.SetPort(8199)

	// 启动服务器，监听并处理HTTP请求
	s.Run()
}
