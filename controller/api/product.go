package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {

}

//商品列表
func (this *ProductController) GetItems (c *gin.Context)  {

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":"我是 api product list",
	})
}

//商品详情
func (this *ProductController) GetItem (c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":"我是 api product detail",
	})
}
