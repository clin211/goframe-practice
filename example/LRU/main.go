package main

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {

	var (
		ctx   = gctx.New()
		cache = gcache.New(2) // 设置LRU淘汰数量
	)

	// 添加10个元素，不过期
	for i := 0; i < 10; i++ {
		if err := cache.Set(ctx, i, i, 0); err != nil {
			panic(err)
		}
	}
	size, err := cache.Size(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	keys, err := cache.Keys(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

	// 读取键名1，保证该键名是优先保留
	value, err := cache.Get(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	// 等待一定时间后(默认1秒检查一次)，
	// 元素会被按照从旧到新的顺序进行淘汰
	time.Sleep(3 * time.Second)
	size, err = cache.Size(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	keys, err = cache.Keys(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}
