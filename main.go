package main

import (
	"music-saas/core"
	"music-saas/global"
	"music-saas/initialize"
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
