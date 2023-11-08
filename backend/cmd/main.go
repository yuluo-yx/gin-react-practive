package main

import (
	"backend/conf"
	"backend/routes"
)

func main() {

	conf.Init()

	// 路由
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)

}
