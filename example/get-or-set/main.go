package main

import (
	"context" // 用于处理上下文，支持可取消的操作
	"fmt"     // 用于标准输入输出
	"time"    // 用于时间操作

	// 引入 gcache 和 gctx 是 GoFrame 框架的一部分，用于缓存和上下文处理
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 定义变量
	var (
		ch    = make(chan struct{}, 0) // 创建一个容量为0的无缓冲通道，用于同步控制
		ctx   = gctx.New()             // 创建一个新的 GoFrame 上下文
		key   = `key`                  // 定义缓存的键
		value = `value`                // 定义缓存的值
	)

	// 启动10个 goroutine
	for i := 0; i < 10; i++ {
		go func(index int) {
			<-ch // 等待通道信号，确保所有 goroutine 同时开始执行
			_, err := gcache.GetOrSetFuncLock(ctx, key, func(ctx context.Context) (interface{}, error) {
				// 这段代码块在缓存中不存在对应键的值时执行
				fmt.Println(index, "entered") // 打印当前 goroutine 的索引，显示哪个 goroutine 获取到了锁
				return value, nil             // 返回要缓存的值
			}, 0) // 第三个参数0表示该缓存项没有过期时间
			if err != nil {
				panic(err) // 如果发生错误，程序直接崩溃
			}
		}(i) // 将当前循环的索引传递给 goroutine
	}

	close(ch)               // 关闭通道，允许所有等待的 goroutine 开始执行
	time.Sleep(time.Second) // 主 goroutine 休眠一秒，等待其他 goroutine 执行完毕
}
