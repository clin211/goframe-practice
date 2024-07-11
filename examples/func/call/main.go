// examples/func/call/main.go
package main

import (
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// formatDate 函数用于将时间格式化为指定的字符串格式
func formatDate(date time.Time) string {
	// 使用Format方法将时间格式化为"2006-01-02 15:04:05"的形式
	return date.Format("2006-01-02 15:04:05")
}

func main() {
	// 创建一个GoFrame的服务器实例
	s := g.Server()

	// 为服务器绑定一个处理根路径("/")的处理器
	s.BindHandler("/", func(r *ghttp.Request) {
		// 当请求根路径时，使用WriteTpl方法渲染模板并返回响应
		// 将当前时间作为参数"date"传递给模板，并提供formatDate函数用于时间格式化
		r.Response.WriteTpl("cell.html", g.Map{
			"date":       time.Now(), // 获取当前时间
			"formatDate": formatDate, // 将formatDate函数传递给模板
		})
	})

	// 设置服务器监听的端口号为 8199
	s.SetPort(8199)

	// 启动服务器，监听并处理HTTP请求
	s.Run()
}
