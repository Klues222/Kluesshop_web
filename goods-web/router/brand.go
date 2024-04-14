package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/api/brands"
)

func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("brands")
	{

		BrandRouter.GET("", brands.List)          // 品牌列表页
		BrandRouter.DELETE("/:id", brands.Delete) // 删除品牌
		BrandRouter.POST("", brands.New)          //新建品牌
		BrandRouter.PUT("/:id", brands.Update)    //修改品牌信息
	}

	CategoryBrandRouter := Router.Group("categorybrands")
	{
		CategoryBrandRouter.GET("", brands.CategoryBrandList)          // 类别品牌列表页
		CategoryBrandRouter.DELETE("/:id", brands.DeleteCategoryBrand) // 删除类别品牌
		CategoryBrandRouter.POST("", brands.NewCategoryBrand)          //新建类别品牌
		CategoryBrandRouter.PUT("/:id", brands.UpdateCategoryBrand)    //修改类别品牌
		CategoryBrandRouter.GET("/:id", brands.GetCategoryBrandList)   //获取分类的品牌
	}
}
