package api

import (
	"fmt"
	"ginApi/common"
	"github.com/gin-gonic/gin"
)

type ProductController struct {

}

//商品列表
func (this *ProductController) GetItems (c *gin.Context)  {
	//


	common.Success(c,"open api product list")
}

//商品详情
func (this *ProductController) GetItem (c *gin.Context)  {
	apiParams , _ := c.Get("apiParams")
	fmt.Println("api params:",apiParams)
	params := apiParams.(map[string]interface{})
	productId ,ok := params["id"]
	if !ok {
		common.Failed(c,"param id required",1100,nil)
		return
	}
	//查找商品
	fmt.Println(productId)
	common.Success(c,params)
}
