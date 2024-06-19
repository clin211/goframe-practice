package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	c := gcache.New()
	ctx := gctx.New()

	// 使用 SetMap 方法向缓存中设置多个键值对，键 "k1" 对应的值为 "v1"，过期时间为 0 表示永不过期
	c.SetMap(ctx, g.MapAnyAny{"k1": "v1"}, 0)

	// 使用 Data 方法获取整个缓存的数据
	data, _ := c.Data(ctx)
	fmt.Println(data) // map[k1:v1]
}
