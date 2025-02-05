package main

import (
	"main/config"
	"main/init/initRouter"
	"main/init/qmlog"
	"main/init/qmsql"
	"main/init/registTable"
	"net/http"
	"time"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	qmlog.InitLog()
	registTable.RegistTable(qmsql.InitMysql(config.Dbconfig.Admin))
	defer qmsql.DEFAULTDB.Close()
	Router := initRouter.InitRouter()
	//qmlog.QMLog.Info("服务器开启") // 日志测试代码
	s := &http.Server{
		Addr:           ":8888",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
