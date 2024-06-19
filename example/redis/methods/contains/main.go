package main

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	c := gcache.New()
	ctx := gctx.New()

	// 向缓存中设置一个键值对，键名为 "k"，值为 "v"，过期时间为 0 表示永不过期
	c.Set(ctx, "k", "v", 0)

	// 使用 Contains 方法检查缓存中是否存在键 "k"
	data, _ := c.Contains(ctx, "k")
	fmt.Println(data) // true，表示键 "k" 存在于缓存中

	// 使用 Contains 方法检查缓存中是否存在键 "k1"
	data1, _ := c.Contains(ctx, "k1")
	fmt.Println(data1) // false，表示键 "k1" 不存在于缓存中
}
