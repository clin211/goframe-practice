package main

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	c := gcache.New()
	ctx := gctx.New()

	// 添加10个没有过期时间的元素
	for i := 0; i < 10; i++ {
		c.Set(ctx, i, i, 0)
	}

	// Size返回缓存中的项目数量
	n, _ := c.Size(ctx)
	fmt.Println(n)
}
