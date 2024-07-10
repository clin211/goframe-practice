package products

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"study/api"
)

func List() ([]api.Product, error) {
	url := "https://fakestoreapi.com/products"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// 将数据解析到结构体中
	var products []api.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return products, err
}
