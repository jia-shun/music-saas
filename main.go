package main

import (
	"music/core"
	"music/global"
	"music/initialize"
)

func main() {
	global.VIPER = core.Viper()
	global.LOG = core.Zap()
	global.DB = initialize.Gorm()
	initialize.MysqlTables(global.DB)
	db, _ := global.DB.DB()
	defer db.Close()
	core.RunServer()
}
