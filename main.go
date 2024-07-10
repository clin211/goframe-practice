package main

import (
	_ "study/internal/packed"

	_ "study/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"study/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
