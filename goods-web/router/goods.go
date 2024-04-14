package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/middlewares"

	"mxshop-api/goods-web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("", goods.List)                                                   //商品列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New) //新建商品
		GoodsRouter.GET("/:id", goods.Detail)                                             //获取商品详情
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)
		GoodsRouter.GET("/:id/stocks", goods.Stocks)                                                    //获取商品库存
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)         //更新部分信息
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus) //

	}

}
