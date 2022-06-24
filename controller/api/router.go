package api

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine)  {
	order := OrderController{}
	r.POST("/api/api/order/getItems",order.GetItems)
	r.POST("/api/api/order/getItem",order.GetItem)

	product := ProductController{}
	r.POST("/api/api/product/getItems",product.GetItems)
	r.POST("/api/api/product/getItem",product.GetItem)

}
