package admin

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine)  {

	admin := r.Group("/api/admin/")
	{
		order := OrderController{}
		admin.POST("order/getItems",order.GetItems)
		admin.POST("order/getItem",order.GetItem)
		product := ProductController{}
		admin.POST("product/getItems",product.GetItems)
		admin.POST("product/getItem",product.GetItem)
	}

	//order := OrderController{}
	//r.POST("/api/admin/order/getItems",order.GetItems)
	//r.POST("/api/admin/order/getItem",order.GetItem)

	//product := ProductController{}
	//r.POST("/api/admin/product/getItems",product.GetItems)
	//r.POST("/api/admin/product/getItem",product.GetItem)

}
