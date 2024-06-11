package main

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"   // 引入GoFrame框架
	"github.com/gogf/gf/v2/net/ghttp" // 引入GoFrame的HTTP包
)

// BizReq 是我们的数据结构，它有三个字段：Account，Password和Password2
type BizReq struct {
	Account   string `v:"required"`                   // 必须有一个账户名
	Password  string `v:"required|ci|same:Password2"` // 必须有一个密码，并且它应该与Password2字段相同，ci表示大小写不敏感
	Password2 string `v:"required"`                   // 必须有一个Password2字段
}

func main() {
	s := g.Server()                             // 创建一个新的服务器实例
	s.BindHandler("/", func(r *ghttp.Request) { // 绑定一个处理器到根URL
		var data BizReq
		if err := r.Parse(&data); err != nil { // 尝试解析请求中的数据到我们的BizReq数据结构
			r.Response.WriteExit(err.Error()) // 如果有错误，我们将返回错误消息并退出
		}

		// 如果没有错误，我们将返回一个成功消息
		r.Response.WriteJson(g.Map{
			"code":    http.StatusOK,
			"message": "Account and Password validation passed",
			"data":    data,
		})
	})
	s.SetPort(8199) // 设置服务器端口为8199
	s.Run()         // 运行服务器
}
