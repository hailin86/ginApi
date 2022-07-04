package admin

import (
	"ginApi/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine)  {

	admin := r.Group("/api/admin/")
	{
		admin.Use(middleware.CheckAppInstance(),middleware.Permission())
		order := OrderController{}
		admin.POST("order/getItems",order.GetItems)
		admin.POST("order/getItem",order.GetItem)
		product := ProductController{}
		admin.POST("product/getItems",product.GetItems)
		admin.POST("product/getItem",product.GetItem)
		casbin := CasbinController{}
		admin.POST("casbin/add",casbin.Add)
	}
}
