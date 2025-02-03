package main

import (
	"ToDoList/conf"
	"ToDoList/routes"
)

func main() {
	// 1.加载配置文件
	conf.Init()
	// 2.初始化数据库
	// 3.初始化路由
	// 4.启动服务
	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}
