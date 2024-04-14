package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/order-web/api/shop_cart"
	"mxshop-api/order-web/middlewares"
)

func InitShopCartRouter(Router *gin.RouterGroup) {
	CartRouter := Router.Group("shopCart").Use(middlewares.JWTAuth())
	{
		CartRouter.GET("", shop_cart.List)          //购物车列表
		CartRouter.DELETE("/:id", shop_cart.Delete) //删除购物车信息
		CartRouter.POST("", shop_cart.New)          //添加商品
		CartRouter.PATCH("/:id", shop_cart.Update)  //更新条目

	}

}
