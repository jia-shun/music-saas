package main

import (
	"saas/orm"
	"saas/router"
)

func main()  {
	orm.InitDatabase()
	engine := router.Router()
	_ = engine.Run()
}
