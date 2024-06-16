package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	type User struct {
		Age  int
		Name string
	}
	var (
		ctx   = gctx.New()
		user  = User{Name: "john"}
		rules = map[string]string{
			"Name": "required|length:6,16",
			"Age":  "between:18,30",
		}
		messages = map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为{min}到{max}个字符",
			},
			"Age": "年龄为18到30周岁",
		}
	)

	err := g.Validator().Rules(rules).Messages(messages).Data(user).Run(ctx)
	if err != nil {
		g.Dump(err.Maps())
	}
}
