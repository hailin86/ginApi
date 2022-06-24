package admin

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine)  {
	order := OrderController{}
	r.POST("/api/admin/order/getItems",order.GetItems)
	r.POST("/api/admin/order/getItem",order.GetItem)

	product := ProductController{}
	r.POST("/api/admin/product/getItems",product.GetItems)
	r.POST("/api/admin/product/getItem",product.GetItem)

}
