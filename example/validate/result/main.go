package main

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gvalid"
)

func main() {
	type User struct {
		Name string `v:"required#请输入用户姓名"`
		Type int    `v:"required|min:1#|请选择用户类型"`
	}
	var (
		err  error
		ctx  = gctx.New()
		user = User{}
	)
	if err = g.Validator().Data(user).Run(ctx); err != nil {
		g.Dump(err.(gvalid.Error).Maps())
		g.Dump(gerror.Current(err))
	}
}
