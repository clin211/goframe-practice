package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx    = gctx.New()
		params = map[string]interface{}{
			"passport":  "",
			"password":  "123456",
			"password2": "1234567",
		}
		rules = []string{
			"passport@required|length:6,16#账号不能为空|账号长度应当在{min}到{max}之间",
			"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在{min}到{max}之间|两次密码输入不相等",
			"password2@required|length:6,16#",
		}
	)
	err := g.Validator().Rules(rules).Data(params).Run(ctx)
	if err != nil {
		fmt.Println(err.Map())
		fmt.Println(err.FirstItem())
		fmt.Println(err.FirstError())
	}
}
