package initialize

import (
	"fmt"
	"gin_web_server/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func SqlDB() *sqlx.DB {
	//fmt.Println("222")
	dbUserName := global.GvaConfig.App.DbUserName
	dbPassWord := global.GvaConfig.App.DbPassWord
	dbName := global.GvaConfig.App.DbName
	dbUrl := global.GvaConfig.App.DbUrl
	dbPort := global.GvaConfig.App.DbPort
	//sqlStr := fmt.Sprintf("%s", dbUserName+":"+dbPassWord+"@tcp("+dbUrl+":"+dbPort+")"+"/"+dbName)
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUserName, dbPassWord, dbUrl, dbPort, dbName)

	//global.GvaLog.Info(sqlStr)
	//db, err := sqlx.Open(global.GvaConfig.App.DbType, "root:admin@tcp(localhost:3306)/demo")
	db, err := sqlx.Open(global.GvaConfig.App.DbType, sqlStr)

	if err != nil {
		panic(err.Error())
	}
	var Db = db
	err = Db.Ping()

	if err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
		global.GvaLog.Error("数据连接失败", zap.Error(err))
	} else {
		//fmt.Println("connect to db success")
		global.GvaLog.Info("mysql连接成功")
	}
	return Db
}
