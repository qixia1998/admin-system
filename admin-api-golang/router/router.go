// 访问接口路由配置

package router

import (
	"admin-api-golang/common/config"
	"admin-api-golang/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 跌机恢复
	router.Use(gin.Recovery())
	// 跨域中间件
	router.Use(middleware.Cors())
	// 图片访问路径静态文件夹可直接访问
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	// 日志log中间件
	router.Use(middleware.Logger())
	register(router)
	return router
}

// 路由接口
func register(router *gin.Engine) {
	// todo 后续接口url
}
