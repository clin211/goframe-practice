package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g" // 导入GoFrame框架
)

func main() {

	BailRule() // 调用我们的BailRule函数
}

func BailRule() {
	// 定义一个结构体BizReq，包含四个字段：Account，QQ，Password，Password2
	type BizReq struct {
		Account   string `v:"bail|required|length:6,16|same:QQ"` //Account字段必须存在，长度在6-16之间，并且必须与QQ字段相同
		QQ        string // QQ字段没有特殊要求
		Password  string `v:"required|same:Password2"` // Password字段必须存在，并且必须与Password2字段相同
		Password2 string `v:"required"`                // Password2字段必须存在
	}
	var (
		ctx = context.Background() // 创建一个context背景
		req = BizReq{
			Account:   "clin",           // 设置Account为"clin"
			QQ:        "123456",         // 设置QQ为"123456"
			Password:  "password123456", // 设置Password为"password123456"
			Password2: "password123456", // 设置Password2为"password123456"
		}
	)
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		fmt.Println(err) // 如果验证失败，打印错误信息
	}
}
