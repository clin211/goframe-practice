package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx  = gctx.New()
		rule = `regex:\d{6,}|\D{6,}|max-length:16`
	)

	if err := g.Validator().Rules(rule).Data(`123456`).Run(ctx); err != nil {
		fmt.Println(err)
	}

	if err := g.Validator().Rules(rule).Data(`abcde6`).Run(ctx); err != nil {
		fmt.Println(err)
	}
}
