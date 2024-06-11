package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gvalid"
)

type User struct {
	Name string `v:"required#请输入用户姓名"` // Name字段是必须的，错误消息为“请输入用户姓名”
	Type int    `v:"required#请选择用户类型"` // Type字段也是必须的，错误消息为“请选择用户类型”
}

// 定义全局变量
var (
	err  error        // 错误信息
	ctx  = gctx.New() // 创建新的上下文
	user = User{}     // 初始化一个User结构体变量
	data = g.Map{
		"name": "john", // 初始化一个映射，键为"name"，值为"john"
	}
)

func main() {
	// 将data的值扫描到user变量中，如果出错则抛出panic
	if err = gconv.Scan(data, &user); err != nil {
		panic(err)
	}

	// 创建一个新的验证器，关联的数据为data，需要验证的数据为user
	// 运行验证，如果出错，打印出错误信息
	err = g.Validator().Assoc(data).Data(user).Run(ctx)
	if err != nil {
		fmt.Println(err.(gvalid.Error).Items())
	}
}
