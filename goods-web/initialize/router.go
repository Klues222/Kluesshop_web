package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/middlewares"
	router2 "mxshop-api/goods-web/router"
	"net/http"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/g/v1")
	router2.InitGoodsRouter(ApiGroup)
	router2.InitCategoryRouter(ApiGroup)
	router2.InitBannerRouter(ApiGroup)
	router2.InitBrandRouter(ApiGroup)
	return router
}
