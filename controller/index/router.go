package index

import (
	"ginApi/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine)  {
	//index 模块下 无需中间件过滤的
	indexNoAction :=r.Group("/api/index/")
	{
		test := TestController{}
		indexNoAction.POST("test/test",test.Test)
	}

	//index 模块下 需要校验 请求合法性的
	indexCheckApp := r.Group("/api/index/")
	{
		indexCheckApp.Use(middleware.CheckAppInstance())
		product := ProductController{}
		indexCheckApp.POST("product/getItems",product.GetItems)
		indexCheckApp.POST("product/getItem",product.GetItem)
	}

	//index 模块下 需要校验 请求合法性 以及需要登录才能访问的
	indexCheckUser := r.Group("/api/index/")
	{
		indexCheckUser.Use(middleware.CheckAppInstance(),middleware.CheckUserAuth())
		order := OrderController{}
		indexCheckUser.POST("order/getItems",order.GetItems)
		indexCheckUser.POST("order/getItem",order.GetItem)
	}




}
