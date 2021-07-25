package main

import (
	routerTool "errors_demo/router"
)

/*
function:程序入口-main方法
author:JZ
time:2021/07/09
*/
func main() {

	router := routerTool.InitRouter()                //获取初始化路由
	router.Run(":9090")                       //启动项目

}
