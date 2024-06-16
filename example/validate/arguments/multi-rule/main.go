package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx      = gctx.New()
		rule     = "url|min-length:11"
		value    = "goframe.org"
		messages = map[string]string{
			"url":        "请输入正确的URL地址",
			"min-length": "地址长度至少为{min}位",
		}
	)
	if err := g.Validator().Rules(rule).Messages(messages).Data(value).Run(ctx); err != nil {
		g.Dump(err.Map())
	}
}
