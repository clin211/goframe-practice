package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx  = gctx.New()
		rule = "length:6,16"
	)

	if err := g.Validator().Rules(rule).Data("123456").Run(ctx); err != nil {
		fmt.Println(err.String())
	}
	if err := g.Validator().Rules(rule).Data("12345").Run(ctx); err != nil {
		fmt.Println(err.String())
	}
}
