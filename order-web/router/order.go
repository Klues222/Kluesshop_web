package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/order-web/api/pay"
	"mxshop-api/order-web/middlewares"

	"mxshop-api/order-web/api/order"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("order").Use(middlewares.JWTAuth())
	{
		OrderRouter.GET("", order.List)        //订单列表
		OrderRouter.POST("", order.New)        //新建订单
		OrderRouter.GET("/:id/", order.Detail) //订单详情

	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}

}
