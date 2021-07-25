package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

/*
function:写入错误日志
author:JZ
time:2021-07-25
*/
func WriteErrorLog(error error, c *gin.Context) {

	logfile, err := os.OpenFile("./log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}

	defer logfile.Close()
	logger := log.New(logfile, "--------------------------------------------------------------------------------------\r\n", log.Ldate|log.Ltime)

	var logMsg string

	if c != nil{
		logMsg = fmt.Sprintf("\n请求地址:%v\n请求方式:%v\n请求ip:%v\n\n%+v",c.Request.RequestURI,c.Request.Method,c.ClientIP(),error)
	}else{
		logMsg = fmt.Sprintf("\n错误信息:\n%+v",error)
	}

	logger.Println(logMsg)

}
