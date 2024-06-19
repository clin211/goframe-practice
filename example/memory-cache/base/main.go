package main

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 创建一个缓存对象，
	// 当然也可以便捷地直接使用gcache包方法
	var (
		ctx   = gctx.New()
		cache = gcache.New()
	)

	// 设置缓存，不过期
	err := cache.Set(ctx, "k1", "v1", 0)
	if err != nil {
		panic(err)
	}

	// 获取缓存值
	value, err := cache.Get(ctx, "k1")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	// 获取缓存大小
	size, err := cache.Size(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(size)

	// 缓存中是否存在指定键名
	b, err := cache.Contains(ctx, "k1")
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	// 删除并返回被删除的键值
	removedValue, err := cache.Remove(ctx, "k1")
	if err != nil {
		panic(err)
	}
	fmt.Println(removedValue)

	// 关闭缓存对象，让GC回收资源
	if err = cache.Close(ctx); err != nil {
		panic(err)
	}
}
