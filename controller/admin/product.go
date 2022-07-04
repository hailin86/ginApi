package admin

import (
	"ginApi/common"
	"github.com/gin-gonic/gin"
)

type ProductController struct {

}

//商品列表
func (this *ProductController) GetItems (c *gin.Context)  {


	common.ResultSuccess(c,"admin product list")
}

//商品详情
func (this *ProductController) GetItem (c *gin.Context)  {

	common.ResultSuccess(c,"admin product detail")
}