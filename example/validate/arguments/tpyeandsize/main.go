package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx      = gctx.New()
		rule     = "integer|between:6,16"
		messages = "请输入一个整数|参数大小不对啊老铁"
		value    = 5.66
	)

	if err := g.Validator().Rules(rule).Messages(messages).Data(value).Run(ctx); err != nil {
		g.Dump(err.Map())
	}
}
