/**
 * @Author: Sun
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/6/7 14:56
 */

package main

import (
	"DTCloud_iot/config"
	"DTCloud_iot/db"
	"DTCloud_iot/global"
	"DTCloud_iot/router"
	"fmt"
)

func init() {
	global.GlobalConfig = config.InitConfig()
	var err error
	global.DB, err = db.Conn2Postgresql(global.GlobalConfig.PGSql.Host,
		global.GlobalConfig.PGSql.User,
		global.GlobalConfig.PGSql.Pass,
		global.GlobalConfig.PGSql.DataBase,
		global.GlobalConfig.PGSql.Port)
	if err != nil {
		fmt.Println(err)
	}
	if global.DB == nil {
		fmt.Println("连接数据库失败！")
	}
}

func main() {

	r := router.NewRouter()

	r.Run(":2022")
}
