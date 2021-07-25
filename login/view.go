package login

import (
	"errors_demo/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)
import 	orm "errors_demo/database"

/*
function:查找用户信息
author:JZ
time:2021-07-25
*/
func userFind(c *gin.Context){

	/*初始化用户id*/
	userId := -1

	/*查询用户信息*/
	userName,err := findUserInfoById(userId)

	/*记录日志*/
	if err != nil{
		log.WriteErrorLog(err,c)
		c.JSON(http.StatusOK,gin.H{"code":10001,"error":"未查询到用户信息。"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"code":200,"userName":userName,"msg":"发送成功！"})

}

/*
function:根据id查找用户信息
author:JZ
time:2021-07-25
*/
func findUserInfoById(userId int)(string,error){

	var userName string

	err := orm.Db.QueryRow(" SELECT user_name FROM account_base_t WHERE id = ?",userId).Scan(&userName)

	/*处理异常*/
	if err != nil{
		return userName,errors.Wrapf(err,"用户id:%v",userId)
	}

	return userName,nil
}