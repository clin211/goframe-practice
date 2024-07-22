package main

import (
	_ "orm/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"orm/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
