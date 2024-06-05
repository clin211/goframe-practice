package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.TODO()
	g.Log().Line().Print(ctx, "this is the short file name with its line number")
	g.Log().Line(true).Print(ctx, "lone file name with line number")
}
