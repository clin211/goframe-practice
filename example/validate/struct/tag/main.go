package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type User struct {
	Uid   int    `v:"uid      @integer|min:1#|请输入用户ID"`
	Name  string `v:"name     @required|length:6,30#请输入用户名称|用户名称长度非法"`
	Pass1 string `v:"password1@required|password3"`
	Pass2 string `v:"password2@required|password3|same:Pass1#|密码格式不合法|两次密码不一致，请重新输入"`
}

func main() {
	var (
		ctx  = gctx.New()
		user = &User{
			Name:  "john",
			Pass1: "Abc123!@#",
			Pass2: "123",
		}
	)

	err := g.Validator().Data(user).Run(ctx)
	if err != nil {
		g.Dump(err.Items())
	}
}
