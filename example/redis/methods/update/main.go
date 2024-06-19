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

	// 使用SetMap方法将多个键值对添加到缓存中，设置的过期时间为0，表示这些键值对不会过期
	c.SetMap(ctx, g.MapAnyAny{"k1": "v1", "k2": "v2", "k3": "v3"}, 0)

	// 从缓存中获取键 "k1" 对应的值，并打印出来
	k1, _ := c.Get(ctx, "k1")
	fmt.Println(k1)

	// 从缓存中获取键 "k2" 对应的值，并打印出来
	k2, _ := c.Get(ctx, "k2")
	fmt.Println(k2)

	// 从缓存中获取键 "k3" 对应的值，并打印出来
	k3, _ := c.Get(ctx, "k3")
	fmt.Println(k3)

	// 更新缓存中键 "k1" 对应的值为 "v11"
	// re 是更新之前的值
	// exist 表示键 "k1" 是否存在于缓存中
	re, exist, _ := c.Update(ctx, "k1", "v11")
	fmt.Println(re, exist)

	// 尝试更新缓存中键 "k4" 对应的值为 "v44"
	// 由于 "k4" 不存在于缓存中，所以 exist 为 false，re 无法提供旧值
	re1, exist1, _ := c.Update(ctx, "k4", "v44")
	fmt.Println(re1, exist1)

	// 再次从缓存中获取键 "k1"、"k2" 和 "k3" 对应的值，打印出来以确认更新后的状态
	kup1, _ := c.Get(ctx, "k1")
	fmt.Println(kup1)

	kup2, _ := c.Get(ctx, "k2")
	fmt.Println(kup2)

	kup3, _ := c.Get(ctx, "k3")
	fmt.Println(kup3)
}
