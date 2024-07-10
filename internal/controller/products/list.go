package products

import (
	"study/api"
	"study/internal/logic/products"
)

func List(ctrl *ProductsController, req interface{}) ([]api.Product, error) {
	return products.List()
}
