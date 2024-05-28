package main

import (
	_ "myapp/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"myapp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
