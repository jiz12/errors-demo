package database

import (
	"database/sql"
	"errors_demo/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql驱动
	"github.com/pkg/errors"
)

var Db *sql.DB

/*
function:初始化数据库配置
author:JZ
time:2020-08-19
*/
func init() {
	initPlatformDB() //初始化DqDb
}

func initPlatformDB() {


	/*声明，以下数据库为真实数据库信息，仅供demo使用。若有恶意行为，将追究法律责任。*/

	//配置MySQL连接参数
	username := "root"  //账号
	password := "root2020" //密码
	host := "47.115.18.17" //数据库地址，可以是Ip或者域名
	port := 3306 //数据库端口
	Dbname := "zkzj_platform_db" //数据库名
	timeout := "10s" //连接超时，10秒
	var err error
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.WriteErrorLog(errors.Wrap(err,"堆栈信息"),nil)
		panic(err)
	}

}