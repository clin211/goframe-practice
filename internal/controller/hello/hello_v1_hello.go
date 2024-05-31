package hello

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"

	v1 "myapp/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	name := g.Cfg("app.name")
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"code":    http.StatusOK,
		"message": "success",
		"data": g.Map{
			"name": name,
		},
	})
	return
}
