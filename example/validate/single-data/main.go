package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// 定义全局变量
var (
	err  error        // 用于存储错误信息
	ctx  = gctx.New() // 创建新的上下文
	data = g.Map{
		"password": "123", // 初始化一个映射，键为"password"，值为"123"
	}
)

func main() {
	// 创建一个新的验证器，规则是"gte:18"，表示值必须大于或等于18；数据为16，消息为"未成年人不允许注册哟"；最后运行验证
	err = g.Validator().
		Rules("gte:18").
		Data(16).
		Messages("未成年人不允许注册哟").
		Run(ctx)
	fmt.Println(err.Error())

	// 创建另一个新的验证器，规则是"required-with:password"，表示如果"password"字段存在，那么这个字段就是必须的；数据为空字符串，关联的数据为data，消息为"请输入确认密码"；最后运行验证
	err = g.Validator().Data("").Assoc(data).
		Rules("required-with:password").
		Messages("请输入确认密码").
		Run(ctx)

	fmt.Println(err.Error())
}
