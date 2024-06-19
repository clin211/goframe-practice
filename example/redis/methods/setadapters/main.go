package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	c := gcache.New()
	adapter := gcache.New()
	c.SetAdapter(adapter)
	c.Set(ctx, "k1", g.Slice{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	fmt.Println(c.Get(ctx, "k1"))
}
