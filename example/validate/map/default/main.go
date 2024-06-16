package main

import (
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
		rules = map[string]string{
			"passport":  "required|length:6,16",
			"password":  "required|length:6,16|same:password2",
			"password2": "required|length:6,16",
		}
	)
	err := g.Validator().Rules(rules).Data(params).Run(ctx)
	if err != nil {
		g.Dump(err.Maps())
	}
}
