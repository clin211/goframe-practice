package main

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func main() {
	_, err := gjson.Encode(func() {})
	fmt.Printf("normal error line: %v\n\n", err)

	_, err = gjson.Encode(func() {})
	fmt.Printf("stack error line: %+v\n", err)
}
