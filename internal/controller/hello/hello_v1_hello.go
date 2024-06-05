package hello

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"

	v1 "myapp/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	// 使用 g.Cfg() 获取配置
	name, err := g.Cfg().Get(ctx, "app.name")
	fmt.Println("name: ", name)
	message := "success"
	if err != nil {
		message = err.Error()
	}

	// 使用 gcfg.Instance() 获取配置
	version, err := gcfg.Instance().Get(ctx, "app.version")
	if err != nil {
		message = err.Error()
	}

	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code":    http.StatusOK,
		"message": message,
		"data": g.Map{
			"name":    name,
			"version": version,
		},
	})
	return
}
