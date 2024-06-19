package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 创建一个新的缓存实例
	c := gcache.New()

	// 创建一个新的上下文实例
	ctx := gctx.New()

	// 设置缓存项，键名为 "k1"，对应的值是一个包含多个键值对的 Map，并且没有设置过期时间（0 表示永不过期）
	c.Set(ctx, "k1", g.Map{"k1": "v1", "k2": "v2"}, 0)

	// 使用 Values 方法获取缓存中所有的值，返回一个切片
	// data 包含缓存中的所有值，这里只有一个键 "k1"，它的值是一个 Map
	data, _ := c.Values(ctx)

	// 打印缓存中的所有值
	fmt.Println(data)
}
