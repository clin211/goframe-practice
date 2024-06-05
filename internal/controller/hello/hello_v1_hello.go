package hello

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"

	v1 "myapp/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	// 使用 g.Cfg() 获取配置
	name, err := g.Cfg().Get(ctx, "app.name")
	message := "success"
	if err != nil {
		message = err.Error()
	}

	// 使用 gcfg.Instance() 获取配置
	version, err := gcfg.Instance().Get(ctx, "app.version")
	if err != nil {
		message = err.Error()
	}

	// 读取自定义配置文件
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("local.yaml")

	// 后面读取的配置都是从 local.yaml 中读取
	mysqlConfig, err := g.Cfg().Get(ctx, "mysql")

	if err != nil {
		message = err.Error()
	}

	// 打印日志
	logger := glog.New()
	logger.Infof(ctx, "mysql config: %v", mysqlConfig)

	// 设置日志等级
	logger.SetLevel(glog.LEVEL_ALL | glog.LEVEL_INFO)
	logger.Infof(ctx, "version: %v", version)

	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code":    http.StatusOK,
		"message": message,
		"data": g.Map{
			"name":    name,
			"version": version,
			"mysql":   mysqlConfig,
		},
	})
	return
}
