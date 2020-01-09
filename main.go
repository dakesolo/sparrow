package main

import (
	"main/route"
	"main/share"
	"net/http"
)

func main() {

	//初始化路径
	share.InitPath()

	//初始化日志
	share.InitLog()

	//初始化数据库
	share.InitDB()

	//初始化OTS
	share.InitOTS()

	//初始化路由
	mux := http.NewServeMux()
	route.InitRoute(mux)

	//运行server
	share.Run(mux)
}
