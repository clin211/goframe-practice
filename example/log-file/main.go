package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

// 设置日志等级
func main() {
	ctx := context.TODO()
	path := "./glog"
	glog.SetPath(path)
	glog.SetStdoutPrint(false)
	// 使用默认文件名称格式
	glog.Print(ctx, "标准文件名称格式，使用当前时间时期")
	// 通过SetFile设置文件名称格式
	glog.SetFile("stdout.log")
	glog.Print(ctx, "设置日志输出文件名称格式为同一个文件")
	// 链式操作设置文件名称格式
	glog.File("stderr.log").Print(ctx, "支持链式操作")
	glog.File("error-{Ymd}.log").Print(ctx, "文件名称支持带gtime日期格式")
	glog.File("access-{Ymd}.log").Print(ctx, "文件名称支持带gtime日期格式")

	list, err := gfile.ScanDir(path, "*")
	g.Dump(err)
	g.Dump(list)
}
