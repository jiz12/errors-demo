package login

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) {

	router.GET("/user/find", userFind) //查找用户信息

}
