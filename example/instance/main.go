package main

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

func init() {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host: "localhost",
				Port: "3306",
				User: "root",
				Pass: "123456",
				Name: "user_center",
				Type: "mysql",
			},
		},
	})
}

func main() {
	db := g.DB()
	r, err := db.Model("users").Where("id", 1).One()
	if err != nil {
		panic(err)
	}
	g.Log().Println(r) // 2024-07-23 11:29:20.300 {"id":1,"name":"forest"}
}
