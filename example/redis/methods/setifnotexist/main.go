package main

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	c := gcache.New()
	// 当键名不存在时写入，并设置过期时间为1000毫秒
	k1, err := c.SetIfNotExist(ctx, "k1", "v1", 1000*time.Millisecond)
	fmt.Println("不存在则写入，过期时间为1000毫秒：", k1, err)

	// 当键名已存在时返回false
	k2, err := c.SetIfNotExist(ctx, "k1", "v2", 1000*time.Millisecond)
	fmt.Println("已存在则返回false：", k2, err)

	// 打印当前的键值列表
	keys1, _ := c.Keys(ctx)
	fmt.Println("当前的键值列表：", keys1)

	// 如果`duration` == 0，则不会过期。如果`duration` < 0 或 给定`value`为nil，则删除`key`。
	c.SetIfNotExist(ctx, "k1", 0, -10000)

	// 等待1秒，K1: V1会自动过期
	time.Sleep(1200 * time.Millisecond)

	// 再次打印当前的键值对，发现K1: V1已经过期
	keys2, _ := c.Keys(ctx)
	fmt.Println("过期则返回 nil：", keys2)
}
