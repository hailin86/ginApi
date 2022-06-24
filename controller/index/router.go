package index

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine)  {
	order := OrderController{}
	r.POST("/api/index/order/getItems",order.GetItems)
	r.POST("/api/index/order/getItem",order.GetItem)

	product := ProductController{}
	r.POST("/api/index/product/getItems",product.GetItems)
	r.POST("/api/index/product/getItem",product.GetItem)

}
