package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g" // 导入GoFrame框架
)

func main() {
	// 定义一个名为BizReq的结构体，有两个字段：Value1和Value2，都是整数数组
	type BizReq struct {
		Value1 []int `v:"foreach|in:1,2,3"` // Value1字段的每个元素都必须在1,2,3之中
		Value2 []int `v:"foreach|in:1,2,3"` // Value2字段的每个元素都必须在1,2,3之中
	}
	var (
		ctx = context.Background() // 创建一个上下文环境
		req = BizReq{
			Value1: []int{1, 2, 3}, // 设置Value1为1,2,3
			Value2: []int{3, 4, 5}, // 设置Value2为3,4,5
		}
	)
	// 验证数据，如果数据不满足验证规则，那么验证器会返回一个错误
	if err := g.Validator().Bail().Data(req).Run(ctx); err != nil {
		fmt.Println(err.String()) // 打印错误信息
	}
}
