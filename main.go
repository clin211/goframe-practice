package main

import (
	_ "corex/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"corex/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
