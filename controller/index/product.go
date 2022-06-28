package index

import (
	"fmt"
	"ginApi/cache"
	"ginApi/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ProductController struct {

}

//商品列表
func (this *ProductController) GetItems (c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":nil,
	})
}

//商品详情
func (this *ProductController) GetItem (c *gin.Context)  {
	id := 1
	pc := &cache.ProductCache{}
	res := pc.Get(id) // nil
	fmt.Println("res:",res)
	product := &model.Product{
		Id:       1,
		Name:     "测试商品",
		Price:    100.55,
		Status:   100,
		Created:  time.Now(),
		Modified: time.Now(),
	}
	flag := pc.Set(1,product,120)
	fmt.Println("set flag:",flag)

	res1 := pc.Get(id) // product
	fmt.Println("res1:",res1)

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":res1,
	})
}