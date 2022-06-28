package index

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine)  {
	index := r.Group("/api/index/")
	{
		order := OrderController{}
		index.POST("order/getItems",order.GetItems)
		index.POST("order/getItem",order.GetItem)

		product := ProductController{}
		index.POST("product/getItems",product.GetItems)
		index.POST("product/getItem",product.GetItem)

		test := TestController{}
		index.POST("test/test",test.Test)
	}


}
