package api

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine)  {
	api := r.Group("/api/api/")
	{
		order := OrderController{}
		api.POST("order/getItems",order.GetItems)
		api.POST("order/getItem",order.GetItem)

		product := ProductController{}
		api.POST("product/getItems",product.GetItems)
		api.POST("product/getItem",product.GetItem)

	}




}
