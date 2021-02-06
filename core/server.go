package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"music-saas/global"
	"music-saas/initialize"
	"strconv"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.CONFIG.System.UseCache {
		initialize.Redis()
	}
	Router := initialize.Routers()
	port := global.CONFIG.System.Addr
	address := fmt.Sprintf(":%d", port)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on: " + strconv.Itoa(port))
	global.LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 30000 * time.Millisecond
	s.WriteTimeout = 30000 * time.Millisecond
	s.MaxHeaderBytes = 1 << 20
	return s
}
