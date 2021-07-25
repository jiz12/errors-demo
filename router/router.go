package router

import (
	"errors_demo/login"
	"github.com/gin-gonic/gin"
)

/*
function:初始化路由主方法
author:JZ
time:2021-07-09
*/
func InitRouter() *gin.Engine {

	router := gin.Default() //初始化路由
	/*登录路由*/
	loginRouter := router.Group("api/v1/login")
	login.InitRouter(loginRouter)
	/*登录路由*/
	return router
}
