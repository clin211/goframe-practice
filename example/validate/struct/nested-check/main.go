package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	type Pass struct {
		Pass1 string `v:"password1@required|same:password2#请输入您的密码|您两次输入的密码不一致"`
		Pass2 string `v:"password2@required|same:password1#请再次输入您的密码|您两次输入的密码不一致"`
	}
	type User struct {
		Pass
		Id   int
		Name string `valid:"name@required#请输入您的姓名"`
	}
	var (
		ctx  = gctx.New()
		user = &User{
			Name: "john",
			Pass: Pass{
				Pass1: "1",
				Pass2: "2",
			},
		}
	)
	err := g.Validator().Data(user).Run(ctx)
	g.Dump(err.Maps())
}
