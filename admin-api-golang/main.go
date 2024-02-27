// 启动

package main

import (
	"admin-api-golang/common/config"
	"admin-api-golang/pkg/db"
	"admin-api-golang/pkg/log"
	"admin-api-golang/pkg/redis"
	"admin-api-golang/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 加载日志
	logInit := log.Log()
	// 设置启动模式
	gin.SetMode(config.Config.Server.Model)
	// 初始化路由
	routerInit := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: routerInit,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logInit.Info("Listen: %s \n", err)
		}
		logInit.Info("listen: %s \n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal)
	//监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	logInit.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logInit.Info("Server Shutdown:", err)
	}
	logInit.Info("Server exiting")
}

// 初始化连接
func init() {
	// mysql
	db.SetupDBLink()
	// redis
	redis.SetupRedisDb()
}
