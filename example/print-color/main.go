package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.TODO()
	g.Log().Debug(ctx, "Debug")
	g.Log().Info(ctx, "Info")
	g.Log().Notice(ctx, "Notice")
	g.Log().Warning(ctx, "Warning")
	g.Log().Error(ctx, "Error")
}
