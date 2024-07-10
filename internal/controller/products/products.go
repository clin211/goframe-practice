package products

import (
	"context"
	"study/api"

	"github.com/gogf/gf/net/ghttp"
)

// PostController 是 products 模块在 Controller 层的实现，用来处理 products 模块的请求.
type Products interface {
	List(ctx context.Context, req *ghttp.Request) (res []api.Product, err error)
}

type ProductsController struct {
}

// New 创建一个 post controller.
func New() *ProductsController {
	return &ProductsController{}
}
